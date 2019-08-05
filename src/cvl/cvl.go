package cvl
import (
	"fmt"
	"os"
	"strings"
	"regexp"
	"time"
	 log "github.com/golang/glog"
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/antchfx/xmlquery"
	"github.com/antchfx/jsonquery"
	"cvl/internal/yparser"
	. "cvl/internal/util"
	"sync"
	"flag"
)

//DB number 
const (
	APPL_DB uint8 = 0 + iota
	ASIC_DB
	COUNTERS_DB
	LOGLEVEL_DB
	CONFIG_DB
	PFC_WD_DB
	FLEX_COUNTER_DB = PFC_WD_DB
	STATE_DB
	SNMP_OVERLAY_DB
	INVALID_DB
)

const DEFAULT_CACHE_DURATION uint16 = 300 /* 300 sec */
const MAX_BULK_ENTRIES_IN_PIPELINE int = 50

var reLeafRef *regexp.Regexp = nil
var reHashRef *regexp.Regexp = nil
var cvlInitialized bool
var dbNameToDbNum map[string]uint8

//map of lua script loaded
var luaScripts map[string]*redis.Script

//var tmpDbCache map[string]interface{} //map of table storing map of key-value pair
					//m["PORT_TABLE] = {"key" : {"f1": "v1"}}
//Important schema information to be loaded at bootup time
type modelTableInfo struct {
	dbNum uint8
	modelName string
	module *yparser.YParserModule
	keys []string
	redisKeyDelim string
	redisKeyPattern string
	mapLeaf []string //for 'mapping  list'
	leafRef map[string][]string //for storing all leafrefs for a leaf in a table, 
				//multiple leafref possible for union 
	mustExp map[string]string
	tablesForMustExp map[string]CVLOperation
}


/* CVL Error Structure. */
type CVLErrorInfo struct {
	TableName string      /* Table having error */
	ErrCode  CVLRetCode   /* CVL Error return Code. */
	CVLErrDetails string  /* CVL Error Message details. */ 
	Keys    []string      /* Keys of the Table having error. */
        Value    string        /* Field Value throwing error */
	Field	 string        /* Field Name throwing error . */
	Msg     string        /* Detailed error message. */
	ConstraintErrMsg  string  /* Constraint error message. */
	ErrAppTag string
}

type CVL struct {
	redisClient *redis.Client
	yp *yparser.YParser
	tmpDbCache map[string]interface{} //map of table storing map of key-value pair
	batchLeaf string
	chkLeafRefWithOthCache bool
}

type modelNamespace struct {
	prefix string
	ns string
}

type modelDataInfo struct {
	modelNs map[string]modelNamespace//model namespace 
	tableInfo map[string]modelTableInfo //redis table to model name and keys
}

//Struct for storing global DB cache to store DB which are needed frequently like PORT
type dbCachedData struct {
	root *yparser.YParserNode //Root of the cached data
	startTime time.Time  //When cache started
	expiry uint16    //How long cache should be maintained in sec
}

//Global data cache for redis table
type cvlGlobalSessionType struct {
	db map[string]dbCachedData
	pubsub *redis.PubSub
	stopChan chan int //stop channel to stop notification listener
	cv *CVL
	mutex *sync.Mutex
}
var cvg cvlGlobalSessionType

//Single redis client for validation
var redisClient *redis.Client

//Stores important model info
var modelInfo modelDataInfo

type keyValuePairStruct struct {
	key string
	values []string
}

func TRACE_LOG(level log.Level, tracelevel CVLTraceLevel, fmtStr string, args ...interface{}) {
	TRACE_LEVEL_LOG(level, tracelevel, fmtStr, args...)
}

func CVL_LOG(level CVLLogLevel, fmtStr string, args ...interface{}) {
	CVL_LEVEL_LOG(level, fmtStr, args...)
}

//package init function 
func init() {
	if (os.Getenv("CVL_SCHEMA_PATH") != "") {
		CVL_SCHEMA = os.Getenv("CVL_SCHEMA_PATH") + "/"
	}

	if (os.Getenv("CVL_DEBUG") != "") {
		SetTrace(true)
	}

	ConfigFileSyncHandler()
	cvlCfgMap := ReadConfFile()

	if (strings.Compare(cvlCfgMap["LOGTOSTDERR"], "true") == 0) {
		flag.Set("logtostderr", "true")
		flag.Set("stderrthreshold", cvlCfgMap["STDERRTHRESHOLD"])
		flag.Set("v", cvlCfgMap["VERBOSITY"])
	}

	CVL_LOG(INFO ,"Current Values of CVL Configuration File %v", cvlCfgMap)

	//regular expression for leafref and hashref finding
	reLeafRef = regexp.MustCompile(`.*[/]([a-zA-Z]*:)?(.*)[/]([a-zA-Z]*:)?(.*)`)
	reHashRef = regexp.MustCompile(`\[(.*)\|(.*)\]`)

	Initialize()

	cvg.db = make(map[string]dbCachedData)

	//Global session keeps the global cache
	cvg.cv, _ = ValidationSessOpen()
	//Create buffer channel of length 1
	cvg.stopChan = make(chan int, 1)
	//Initialize mutex
	cvg.mutex = &sync.Mutex{}

	_, err := redisClient.ConfigSet("notify-keyspace-events", "AKE").Result()
	if err != nil {
		CVL_LOG(ERROR ,"Could not enable notification error %s", err)
	}

	dbCacheSet(false, "PORT", 0)
}

func Debug(on bool) {
	yparser.Debug(on)
}

//Get attribute value of xml node
func getXmlNodeAttr(node *xmlquery.Node, attrName string) string {
	for _, attr := range node.Attr {
		if (attrName == attr.Name.Local) {
			return attr.Value
		}
	}

	return ""
}

//Store useful schema data during initialization
func storeModelInfo(modelFile string, module *yparser.YParserModule) { //such model info can be maintained in C code and fetched from there 
	f, err := os.Open(CVL_SCHEMA + modelFile)
	root, err := xmlquery.Parse(f)

	if  err != nil {
		return
	}
	f.Close()

	//model is derived from file name
	tokens := strings.Split(modelFile, ".")
	modelName := tokens[0]

	//Store namespace
	modelNs := modelNamespace{}

	nodes := xmlquery.Find(root, "//module/namespace")
	if (nodes != nil) {
		modelNs.ns = nodes[0].Attr[0].Value
	}

	nodes = xmlquery.Find(root, "//module/prefix")
	if (nodes != nil) {
		modelNs.prefix = nodes[0].Attr[0].Value
	}

	modelInfo.modelNs[modelName] = modelNs

	//Store metadata present in each list
	nodes = xmlquery.Find(root, "//module/container/list")
	if (nodes == nil) {
		return
	}

	for  _, node := range nodes {
		//for each list
		tableName :=  node.Attr[0].Value
		tableInfo := modelTableInfo{modelName: modelName}
		//Store the reference for list node to be used later
		listNode := node
		node = node.FirstChild
		//Default database is CONFIG_DB since CVL works with config db mainly
		tableInfo.module = module
		tableInfo.dbNum = CONFIG_DB
		//default delim '|'
		tableInfo.redisKeyDelim = "|"

		fieldCount := 0

		//Check for meta data in schema
		for node !=  nil {
			switch node.Data {
			case "db-name":
				tableInfo.dbNum = dbNameToDbNum[node.Attr[0].Value]
				fieldCount++
			case "key":
				tableInfo.keys = strings.Split(node.Attr[0].Value," ")
				fieldCount++
			case "key-delim":
				tableInfo.redisKeyDelim = node.Attr[0].Value
				fieldCount++
			case "key-pattern":
				tableInfo.redisKeyPattern = node.Attr[0].Value
				fieldCount++
			case "map-leaf":
				tableInfo.mapLeaf = strings.Split(node.Attr[0].Value," ")
				fieldCount++
			}
			node = node.NextSibling
		}

		//Find and store all leafref under each table
		/*
		if (listNode == nil) {
			//Store the tableInfo in global data
			modelInfo.tableInfo[tableName] = tableInfo

			continue
		}
		*/

		leafRefNodes := xmlquery.Find(listNode, "//type[@name='leafref']")
		if (leafRefNodes == nil) {
			//Store the tableInfo in global data
			modelInfo.tableInfo[tableName] = tableInfo

			continue
		}

		tableInfo.leafRef = make(map[string][]string)
		for _, leafRefNode := range leafRefNodes {
			if (leafRefNode.Parent == nil || leafRefNode.FirstChild == nil) {
				continue
			}

			//Get the leaf/leaf-list name holding this leafref
			//Note that leaf can have union of leafrefs
			leafName := ""
			for node := leafRefNode.Parent; node != nil; node = node.Parent {
				if  (node.Data == "leaf" || node.Data == "leaf-list") {
					leafName = getXmlNodeAttr(node, "name")
					break
				}
			}

			//Store the leafref path
			if (leafName != "") {
				tableInfo.leafRef[leafName] = append(tableInfo.leafRef[leafName],
				getXmlNodeAttr(leafRefNode.FirstChild, "value"))
			}
		}

		//Find all 'must' expression and store the agains its parent node
		mustExps := xmlquery.Find(listNode, "//must")
		if (mustExps == nil) {
			//Update the tableInfo in global data
			modelInfo.tableInfo[tableName] = tableInfo
			continue
		}

		tableInfo.mustExp = make(map[string]string)
		for _, mustExp := range mustExps {
			if (mustExp.Parent == nil) {
				continue
			}
			parentName := ""
			for node := mustExp.Parent; node != nil; node = node.Parent {
				//assuming must exp is at leaf or list level
				if  (node.Data == "leaf" || node.Data == "leaf-list" ||
				node.Data == "list") {
					parentName = getXmlNodeAttr(node, "name")
					break
				}
			}
			if (parentName != "") {
				tableInfo.mustExp[parentName] = getXmlNodeAttr(mustExp, "condition")
			}
		}

		//Update the tableInfo in global data
		modelInfo.tableInfo[tableName] = tableInfo
	}
}

//Find the tables names in must expression, these tables data need to be fetched 
//during semantic validation
func addTableNamesForMustExp() {

	for tblName, tblInfo := range  modelInfo.tableInfo {
		if (tblInfo.mustExp == nil) {
			continue
		}

		for _, mustExp := range tblInfo.mustExp {
			tblInfo.tablesForMustExp = make(map[string]CVLOperation)
			var op CVLOperation = OP_NONE
			//Check if 'must' expression should be executed for a particular operation
			if (strings.Contains(mustExp,
			"/scommon:operation/scommon:operation != 'CREATE'") == true) {
				op = op | OP_CREATE
			} else if (strings.Contains(mustExp,
			"/scommon:operation/scommon:operation != 'UPDATE'") == true) {
				op = op | OP_UPDATE
			} else if (strings.Contains(mustExp,
			"/scommon:operation/scommon:operation != 'DELETE'") == true) {
				op = op | OP_DELETE
			}
			//store the current table always so that expression check with other row
			tblInfo.tablesForMustExp[tblName] = op

			//check which table name is present in the must expression
			for tblNameSrch, _ := range modelInfo.tableInfo {
				if (tblNameSrch == tblName) {
					continue
				}
				//Table name should appear like "../VLAN_MEMBER/tagging_mode' or '
				// "/prt:PORT/prt:ifname"
				re := regexp.MustCompile(fmt.Sprintf(".*[/]([a-zA-Z]*:)?%s[\\[/]", tblNameSrch))
				matches := re.FindStringSubmatch(mustExp)
				if (len(matches) > 0) {
					//stores the table name 
					tblInfo.tablesForMustExp[tblNameSrch] = op
				}
			}
		}

		//update map
		modelInfo.tableInfo[tblName] = tblInfo
	}
}

//Split key into table prefix and key
func splitRedisKey(key string) (string, string) {
	for tblName,_ := range modelInfo.tableInfo {
		//Check if table prefix matches to any schema table
		//i.e. has 'VLAN|' or 'PORT|' etc.
		if (strings.HasPrefix(key, tblName + modelInfo.tableInfo[tblName].redisKeyDelim)) {
			prefixLen := len(tblName) + 1
			return tblName, key[prefixLen:]
		}
	}

	return "",""
}

//Convert Redis key to Yang keys, if multiple key components are there,
//they are separated based on Yang schema
func getRedisToYangKeys(tableName string, redisKey string)[]keyValuePairStruct{
	keyNames := modelInfo.tableInfo[tableName].keys
	//First split all the keys components
	keyVals := strings.Split(redisKey, modelInfo.tableInfo[tableName].redisKeyDelim) //split by DB separator
	//Store patterns for each key components by splitting using key delim
	keyPatterns := strings.Split(modelInfo.tableInfo[tableName].redisKeyPattern,
			modelInfo.tableInfo[tableName].redisKeyDelim) //split by DB separator

	if (len(keyNames) != len(keyVals)) {
		return nil //number key names and values does not match
	}

	mkeys := []keyValuePairStruct{}
	//For each key check the pattern and store key/value pair accordingly
	for  idx, keyName := range keyNames {

		//check if key-pattern contains specific key pattern
		if (keyPatterns[idx+1] == fmt.Sprintf("({%s},)*", keyName)) {   // key pattern is "({key},)*" i.e. repeating keys seperated by ','   
			repeatedKeys := strings.Split(keyVals[idx], ",")
			mkeys = append(mkeys, keyValuePairStruct{keyName, repeatedKeys})

		} else if (keyPatterns[idx+1] == fmt.Sprintf("{%s}", keyName)) { //no specific key pattern - just "{key}"

			//Store key/value mapping     
			mkeys = append(mkeys, keyValuePairStruct{keyName,  []string{keyVals[idx]}})
		}
	}

	return mkeys
}


//Add child node to a parent node
func(c *CVL) addChildNode(tableName string, parent *yparser.YParserNode, name string) *yparser.YParserNode {

	//return C.lyd_new(parent, modelInfo.tableInfo[tableName].module, C.CString(name))
	return c.yp.AddChildNode(modelInfo.tableInfo[tableName].module, parent, name)
}

//Add specific entries by looking at must expression
//Must expression may need single or multiple entries
//It can be within same table or across multiple tables
//Aggregate function such count() can be quite expensive and 
//should be avoid through this function
func (c *CVL) addTableEntryForMustExp(op CVLOperation, tableName string) CVLRetCode {
	if (modelInfo.tableInfo[tableName].mustExp == nil) {
		return CVL_SUCCESS
	}
	return CVL_SUCCESS
}

//Add all other table data for validating all 'must' exp for tableName
func (c *CVL) addTableDataForMustExp(op CVLOperation, tableName string) CVLRetCode {

	if (modelInfo.tableInfo[tableName].mustExp == nil) {
		return CVL_SUCCESS
	}

	for mustTblName, mustOp := range modelInfo.tableInfo[tableName].tablesForMustExp {
		//First check if must expression should be executed for the ven operation
		if (mustOp != OP_NONE) && ((mustOp & op) == OP_NONE) {
			//must to be excuted for particular operation, but current operation 
			//is not the same one
			continue
		}

		//Check in global cache first and merge to session cache
		if topNode, _ := dbCacheGet(mustTblName); topNode != nil {
			var errObj yparser.YParserError
			//If global cache has the table, add to the session validation
			TRACE_LOG(INFO_API, TRACE_CACHE, "Adding global cache to session cache for table %s", tableName)
			if errObj = c.yp.CacheSubtree(true, topNode); errObj.ErrCode != yparser.YP_SUCCESS {
				return CVL_SYNTAX_ERROR
			}
		} else { //Put the must table in global table and add to session cache
			cvg.cv.chkLeafRefWithOthCache = true
			dbCacheSet(false, mustTblName, 100*DEFAULT_CACHE_DURATION) //Keep the cache for default duration
			cvg.cv.chkLeafRefWithOthCache = false

			if topNode, ret := dbCacheGet(mustTblName); topNode != nil {
				var errObj yparser.YParserError
				//If global cache has the table, add to the session validation
				TRACE_LOG(INFO_API, TRACE_CACHE, "Global cache created, add the data to session cache for table %s", tableName)
				if errObj = c.yp.CacheSubtree(true, topNode); errObj.ErrCode != yparser.YP_SUCCESS {
					return CVL_SYNTAX_ERROR
				}
			} else if (ret == CVL_SUCCESS) {
				TRACE_LOG(INFO_API, TRACE_CACHE, "Global cache empty, no data in Redis for table %s", tableName)
				return CVL_SUCCESS
			} else {
				CVL_LOG(ERROR ,"Could not create global cache for table %s", mustTblName)
				return CVL_ERROR
			}


			/*
			tableKeys, err:= redisClient.Keys(mustTblName +
			modelInfo.tableInfo[mustTblName].redisKeyDelim + "*").Result()

			if (err != nil) {
				continue
			}

			for _, tableKey := range tableKeys {
				tableKey = tableKey[len(mustTblName+ modelInfo.tableInfo[mustTblName].redisKeyDelim):] //remove table prefix
				if (c.tmpDbCache[mustTblName] == nil) {
					c.tmpDbCache[mustTblName] = map[string]interface{}{tableKey: nil}
				} else {
					tblMap := c.tmpDbCache[mustTblName]
					tblMap.(map[string]interface{})[tableKey] =nil
					c.tmpDbCache[mustTblName] = tblMap
				}
			}
			*/
		}
	}

	return CVL_SUCCESS
}

func (c *CVL) addUpdateDataToCache(tableName string, redisKey string) {
	if (c.tmpDbCache[tableName] == nil) {
		c.tmpDbCache[tableName] = map[string]interface{}{redisKey: nil}
	} else {
		tblMap := c.tmpDbCache[tableName]
		tblMap.(map[string]interface{})[redisKey] =nil
		c.tmpDbCache[tableName] = tblMap
	}
}

//Check delete constraint for leafref if key/field is deleted
func (c *CVL) checkDeleteConstraint(cfgData []CVLEditConfigData,
			tableName, keyVal, field string) CVLRetCode {
	var leafRefs []tblFieldPair
	if (field != "") {
		//Leaf or field is getting deleted
		leafRefs = c.findUsedAsLeafRef(tableName, field)
	} else {
		//Entire entry is getting deleted
		leafRefs = c.findUsedAsLeafRef(tableName, modelInfo.tableInfo[tableName].keys[0])
	}

	//The entry getting deleted might have been referred from multiple tables
	//Return failure if at-least one table is using this entry
	for _, leafRef := range leafRefs {
		TRACE_LOG(INFO_API, (TRACE_DELETE | TRACE_SEMANTIC), "Checking delete constraint for leafRef %s/%s", leafRef.tableName, leafRef.field)
		//Check in dependent data first, if the referred entry is already deleted
		leafRefDeleted := false
		for _, cfgDataItem := range cfgData {
			if (cfgDataItem.VType == VALIDATE_NONE) &&
			(cfgDataItem.VOp == OP_DELETE ) &&
			(strings.HasPrefix(cfgDataItem.Key, (leafRef.tableName + modelInfo.tableInfo[leafRef.tableName].redisKeyDelim + keyVal + modelInfo.tableInfo[leafRef.tableName].redisKeyDelim))) {
				//Currently, checking for one entry is being deleted in same session
				//We should check for all entries
				leafRefDeleted = true
				break
			}
		}

		if (leafRefDeleted == true) {
			continue //check next leafref
		}

		//Else, check if any referred enrty is present in DB
		var nokey []string
		refKeyVal, err := luaScripts["find_key"].Run(redisClient, nokey, leafRef.tableName,
		modelInfo.tableInfo[leafRef.tableName].redisKeyDelim, leafRef.field, keyVal).Result()
		if (err == nil &&  refKeyVal != "") {
			CVL_LOG(ERROR, "Delete will violate the constraint as entry %s is referred in %s", tableName, refKeyVal)

			return CVL_SEMANTIC_ERROR
		}
	}


	return CVL_SUCCESS
}

//Add the data which are referring this key
func (c *CVL) updateDeleteDataToCache(tableName string, redisKey string) {
	if _, existing := c.tmpDbCache[tableName]; existing == false {
		return
	} else {
		tblMap := c.tmpDbCache[tableName]
		if _, existing := tblMap.(map[string]interface{})[redisKey]; existing == true {
			delete(tblMap.(map[string]interface{}), redisKey)
		}
		c.tmpDbCache[tableName] = tblMap
	}
}

//Find which all tables (and which field) is using given (tableName/field)
// as leafref
//Use LUA script to find if table has any entry for this leafref

type tblFieldPair struct {
	tableName string
	field string
}

func (c *CVL) findUsedAsLeafRef(tableName, field string) []tblFieldPair {

	var tblFieldPairArr []tblFieldPair

	for tblName, tblInfo := range  modelInfo.tableInfo {
		if (tableName == tblName) {
			continue
		}
		if (len(tblInfo.leafRef) == 0) {
			continue
		}

		for fieldName, leafRefs  := range tblInfo.leafRef {
			found := false
			//Find leafref by searching table and field name
			for _, leafRef := range leafRefs {
				if ((strings.Contains(leafRef, tableName) == true) &&
				(strings.Contains(leafRef, field) == true)) {
					tblFieldPairArr = append(tblFieldPairArr,
					tblFieldPair{tblName, fieldName})
					//Found as leafref, no need to search further
					found = true
					break
				}
			}

			if (found == true) {
				break
			}
		}
	}

	return tblFieldPairArr
}

//Add leafref entry for caching
//It has to be recursive in nature, as there can be chained leafref
func (c *CVL) addLeafRef(config bool, tableName string, name string, value string) {

	if (config == false) {
		return
	}

	//Check if leafRef entry is there for this field
	if (len(modelInfo.tableInfo[tableName].leafRef[name]) > 0) { //array of leafrefs for a leaf
		for _, leafRef  := range modelInfo.tableInfo[tableName].leafRef[name] {

			//Get reference table name from the path and the leaf name
			matches := reLeafRef.FindStringSubmatch(leafRef)

			//We have the leafref table name and the leaf name as well
			if (matches != nil && len(matches) == 5) { //whole + 4 sub matches
				refTableName := matches[2]
				redisKey := value

				//Check if leafref dependency can also be met from 'must' table
				if (c.chkLeafRefWithOthCache == true) {
					found := false
					for mustTbl, _ := range modelInfo.tableInfo[tableName].tablesForMustExp {
						if mustTbl == refTableName {
							found = true
							break
						}
					}
					if (found == true) {
						//Leafref data will be available from must table dep data, skip this leafref entry
						continue
					}
				}

				//only key is there, value wil be fetched and stored here, 
				//if value can't fetched this entry will be deleted that time
				if (c.tmpDbCache[refTableName] == nil) {
					c.tmpDbCache[refTableName] = map[string]interface{}{redisKey: nil}
				} else {
					tblMap := c.tmpDbCache[refTableName]
					_, exist := tblMap.(map[string]interface{})[redisKey]
					if (exist == false) {
						tblMap.(map[string]interface{})[redisKey] = nil
						c.tmpDbCache[refTableName] = tblMap
					}
				}
			}
		}
	}
}


func (c *CVL) addChildLeaf(config bool, tableName string, parent *yparser.YParserNode, name string, value string) {
	//Batch leaf creation
	c.batchLeaf = c.batchLeaf + name + "#" + value + "#"
	//Check if this leaf has leafref,
	//If so add the add redis key to its table so that those 
	// details can be fetched for dependency validation

	c.addLeafRef(config, tableName, name, value)
}


func (c *CVL) checkFieldMap(fieldMap *map[string]string) map[string]interface{} {
	fieldMapNew := map[string]interface{}{}

	for field, value := range *fieldMap {
		if (field == "NULL") {
			continue
		} else if (field[len(field)-1:] == "@") {
			//last char @ means it is a leaf-list/array of fields
			field = field[:len(field)-1] //strip @
			//split the values seprated using ','
			strArr := strings.Split(value, ",")
			//fieldMapNew[field] = strings.Split(value, ",")
			arrMap := make([]interface{}, 0)//len(strArr))
			for _, strArrItem := range strArr {
				arrMap = append(arrMap, strArrItem)
			}
			fieldMapNew[field] = arrMap//make([]interface{}, len(strArr))
		} else {
			fieldMapNew[field] = value
		}
	}

	return fieldMapNew
}

//Fetch given table entries using pipeline
func (c *CVL) fetchTableDataToTmpCache(tableName string, dbKeys map[string]interface{}) int {

	TRACE_LOG(INFO_API, TRACE_CACHE, "\n%v, Entered fetchTableDataToTmpCache", time.Now())

	totalCount := len(dbKeys)
	if (totalCount == 0) {
		//No entry to be fetched
		return 0
	}

	entryFetched := 0
	bulkCount := 0
	bulkKeys := []string{}
	for dbKey, val := range dbKeys { //for all keys
		if (val != nil) { //skip entry already fetched
			mapTable := c.tmpDbCache[tableName]
			delete(mapTable.(map[string]interface{}), dbKey) //delete entry already fetched
			continue
		}

		bulkKeys = append(bulkKeys, dbKey)
		bulkCount = bulkCount + 1

		if(bulkCount != totalCount) && ((bulkCount % MAX_BULK_ENTRIES_IN_PIPELINE) != 0) {
			//Accumulate entries to be fetched
			continue
		}

		mCmd := map[string]*redis.StringStringMapCmd{}

		pipe := redisClient.Pipeline()

		for _, dbKey := range bulkKeys {

			redisKey := tableName + modelInfo.tableInfo[tableName].redisKeyDelim + dbKey
			//Check in global cache first and merge to session cache

			//Otherwise fetch it from Redis
			mCmd[dbKey] = pipe.HGetAll(redisKey) //write into pipeline
			if mCmd[dbKey] == nil {
				CVL_LOG(ERROR, "Failed pipe.HGetAll('%s')", redisKey)
			}
		}

		_, err := pipe.Exec()
		if err != nil {
			CVL_LOG(ERROR, "Failed to fetch details for table %s", tableName)
			return 0
		}
		pipe.Close()
		bulkKeys = nil

		mapTable := c.tmpDbCache[tableName]

		for key, val := range mCmd {
			res, err := val.Result()
			if (err != nil || len(res) == 0) {
				//no data found, don't keep blank entry
				delete(mapTable.(map[string]interface{}), key)
				continue
			}
			//exclude table name and delim
			keyOnly := key
			fieldMap := c.checkFieldMap(&res)
			mapTable.(map[string]interface{})[keyOnly] = fieldMap
			entryFetched = entryFetched + 1
		}
	}

	TRACE_LOG(INFO_API, TRACE_CACHE,"\n%v, Exiting fetchTableDataToTmpCache", time.Now())

	return entryFetched
}

//populate redis data to cache
func (c *CVL) fetchDataToTmpCache() *yparser.YParserNode {
	TRACE_LOG(INFO_API, TRACE_CACHE, "\n%v, Entered fetchToTmpCache", time.Now())

	entryToFetch := 0
	var root *yparser.YParserNode = nil
	var errObj yparser.YParserError

	for entryToFetch = 1; entryToFetch > 0; { //Force to enter the loop for first time
		//Repeat until all entries are fetched 
		entryToFetch = 0
		for tableName, dbKeys := range c.tmpDbCache { //for each table
			entryToFetch = entryToFetch + c.fetchTableDataToTmpCache(tableName, dbKeys.(map[string]interface{}))
		} //for each table

		//If no table entry delete the table  itself
		for tableName, dbKeys := range c.tmpDbCache { //for each table
			if (len(dbKeys.(map[string]interface{}))  == 0) {
				 delete(c.tmpDbCache, tableName)
				 continue
			}
		}

		if (entryToFetch == 0) {
			//No more entry to fetch
			break
		}

		if (Tracing == true) {
			jsonDataBytes, _ := json.Marshal(c.tmpDbCache)
			jsonData := string(jsonDataBytes)
			TRACE_LOG(INFO_API, TRACE_CACHE, "Top Node=%v\n", jsonData)
		}

		data, err := jsonquery.ParseJsonMap(&c.tmpDbCache)

		if (err != nil) {
			return nil
		}

		//Build yang tree for each table and cache it
		for jsonNode := data.FirstChild; jsonNode != nil; jsonNode=jsonNode.NextSibling {
			TRACE_LOG(INFO_API, TRACE_CACHE, "Top Node=%v\n", jsonNode.Data)
			//Visit each top level list in a loop for creating table data
			topNode, _ := c.generateTableData(true, jsonNode)
			if (root == nil) {
				root = topNode
			} else {
				if root, errObj = c.yp.MergeSubtree(root, topNode); errObj.ErrCode != yparser.YP_SUCCESS {
					return nil
				}
			}
		}
	} // until all dependent data is fetched

	if root != nil && Tracing == true {
		dumpStr := c.yp.NodeDump(root)
		TRACE_LOG(INFO_DETAIL, TRACE_CACHE, "Dependent Data = %v\n", dumpStr)
	}

	TRACE_LOG(INFO_API, TRACE_CACHE, "\n%v, Exiting fetchToTmpCache", time.Now())
	return root
}


func (c *CVL) clearTmpDbCache() {
	for key, _ := range c.tmpDbCache {
		delete(c.tmpDbCache, key)
	}
}

func (c *CVL) generateTableFieldsData(config bool, tableName string, jsonNode *jsonquery.Node,
parent *yparser.YParserNode) CVLRetCode {

	//Traverse fields
	for jsonFieldNode := jsonNode.FirstChild; jsonFieldNode!= nil;
	jsonFieldNode = jsonFieldNode.NextSibling {
		//Add fields as leaf to the list
		if (jsonFieldNode.Type == jsonquery.ElementNode &&
		jsonFieldNode.FirstChild != nil &&
		jsonFieldNode.FirstChild.Type == jsonquery.TextNode) {

			if (len(modelInfo.tableInfo[tableName].mapLeaf) == 2) {//mapping should have two leaf always
				//Values should be stored inside another list as map table
				listNode := c.addChildNode(tableName, parent, tableName) //Add the list to the top node
				c.addChildLeaf(config, tableName,
				listNode, modelInfo.tableInfo[tableName].mapLeaf[0],
				jsonFieldNode.Data)

				c.addChildLeaf(config, tableName,
				listNode, modelInfo.tableInfo[tableName].mapLeaf[1],
				jsonFieldNode.FirstChild.Data)

			} else {
				//check if it is hash-ref, then need to add only key from "TABLE|k1"
				hashRefMatch := reHashRef.FindStringSubmatch(jsonFieldNode.FirstChild.Data)

				if (hashRefMatch != nil && len(hashRefMatch) == 3) {
				/*if (strings.HasPrefix(jsonFieldNode.FirstChild.Data, "[")) &&
				(strings.HasSuffix(jsonFieldNode.FirstChild.Data, "]")) &&
				(strings.Index(jsonFieldNode.FirstChild.Data, "|") > 0) {*/

					c.addChildLeaf(config, tableName,
					parent, jsonFieldNode.Data,
					hashRefMatch[2]) //take hashref key value
				} else {
					c.addChildLeaf(config, tableName,
					parent, jsonFieldNode.Data,
					jsonFieldNode.FirstChild.Data)
				}
			}

		} else if (jsonFieldNode.Type == jsonquery.ElementNode &&
		jsonFieldNode.FirstChild != nil &&
		jsonFieldNode.FirstChild.Type == jsonquery.ElementNode) {
			//Array data e.g. VLAN members
			for  arrayNode:=jsonFieldNode.FirstChild; arrayNode != nil;

			arrayNode = arrayNode.NextSibling {
				c.addChildLeaf(config, tableName,
				parent, jsonFieldNode.Data,
				arrayNode.FirstChild.Data)
			}
		}
	}

	return CVL_SUCCESS
}

func (c *CVL) generateTableData(config bool, jsonNode *jsonquery.Node)(*yparser.YParserNode, CVLErrorInfo) {
	var cvlErrObj CVLErrorInfo
	tableName := fmt.Sprintf("%s",jsonNode.Data)
	c.batchLeaf = ""

	var topNode *yparser.YParserNode
	topNode = c.yp.AddChildNode(modelInfo.tableInfo[tableName].module,
	nil, modelInfo.tableInfo[tableName].modelName)


	//Traverse each key instance
	for jsonNode = jsonNode.FirstChild; jsonNode != nil; jsonNode = jsonNode.NextSibling {

		//For each field check if is key 
		//If it is key, create list as child of top container
		// Get all key name/value pairs
		keyValuePair := getRedisToYangKeys(tableName, jsonNode.Data)
		keyCompCount := len(keyValuePair)
		totalKeyComb := 1
		var keyIndices []int

		//Find number of all key combinations
		//Each key can have one or more key values, which results in nk1 * nk2 * nk2 combinations
		idx := 0
		for i,_ := range keyValuePair {
			totalKeyComb = totalKeyComb * len(keyValuePair[i].values)
			keyIndices = append(keyIndices, 0)
		}

		for  ; totalKeyComb > 0 ; totalKeyComb-- {

			//Add table i.e. create list element
			listNode := c.addChildNode(tableName, topNode, tableName) //Add the list to the top node

			//For each key combination
			//Add keys as leaf to the list
			for idx = 0; idx < keyCompCount; idx++ {
				c.addChildLeaf(config, tableName,
				listNode, keyValuePair[idx].key,
				keyValuePair[idx].values[keyIndices[idx]])
			}

			//Get all fields under the key field and add them as children of the list
			c.generateTableFieldsData(config, tableName, jsonNode, listNode)

			//Check which key elements left after current key element
			var next int = keyCompCount - 1
			for  ((next > 0) && ((keyIndices[next] +1) >=  len(keyValuePair[next].values))) {
				next--
			}
			//No more combination possible
			if (next < 0) {
				break
			}

			keyIndices[next]++

			//Reset indices for all other key elements
			for idx = next+1;  idx < keyCompCount; idx++ {
				keyIndices[idx] = 0
			}

			TRACE_LOG(INFO_API, TRACE_CACHE, "Starting batch leaf creation - %s\n", c.batchLeaf)
			//process batch leaf creation
			if errObj := c.yp.AddMultiLeafNodes(modelInfo.tableInfo[tableName].module, listNode, c.batchLeaf); errObj.ErrCode != yparser.YP_SUCCESS {
				cvlErrObj = CreateCVLErrObj(errObj)
				return nil, cvlErrObj 
			}
			c.batchLeaf = ""
		}
	}

	return topNode, cvlErrObj
}

func (c *CVL) translateToYang(jsonMap *map[string]interface{}) (*yparser.YParserNode, CVLErrorInfo) {

	var  cvlErrObj CVLErrorInfo
	//Parse the map data to json tree
	data, _ := jsonquery.ParseJsonMap(jsonMap)
	var root *yparser.YParserNode
	root = nil
	var errObj yparser.YParserError

	for jsonNode := data.FirstChild; jsonNode != nil; jsonNode=jsonNode.NextSibling {
		TRACE_LOG(INFO_API, TRACE_LIBYANG, "Top Node=%v\n", jsonNode.Data)
		//Visit each top level list in a loop for creating table data
		topNode, cvlErrObj  := c.generateTableData(true, jsonNode)

		if  topNode == nil {
			cvlErrObj.ErrCode = CVL_SYNTAX_ERROR
			return nil, cvlErrObj
		}

		if (root == nil) {
			root = topNode
		} else {
			if root, errObj = c.yp.MergeSubtree(root, topNode); errObj.ErrCode != yparser.YP_SUCCESS {
				return nil, cvlErrObj
			}
		}
	}

	return root, cvlErrObj
}

//Validate config - syntax and semantics
func (c *CVL) validate (data *yparser.YParserNode) CVLRetCode {

	depData := c.fetchDataToTmpCache()
	/*
	if (depData != nil) {
		if (0 != C.lyd_merge_to_ctx(&data, depData, C.LYD_OPT_DESTRUCT, ctx)) {
			TRACE_LOG(1, "Failed to merge status data\n")
		}
	}

	if (0 != C.lyd_data_validate(&data, C.LYD_OPT_CONFIG, ctx)) {
		fmt.Println("Validation failed\n")
		return CVL_SYNTAX_ERROR
	}*/

	TRACE_LOG(INFO_DATA, TRACE_LIBYANG, "\nValidate1 data=%v\n", c.yp.NodeDump(data))
	errObj := c.yp.ValidateData(data, depData)
	if yparser.YP_SUCCESS != errObj.ErrCode {
		return CVL_FAILURE
	}

	return CVL_SUCCESS
}

func  CreateCVLErrObj(errObj yparser.YParserError) CVLErrorInfo {

	cvlErrObj :=  CVLErrorInfo {
		TableName : errObj.TableName,
		ErrCode   : CVLRetCode(errObj.ErrCode),
		CVLErrDetails : cvlErrorMap[CVLRetCode(errObj.ErrCode)],
		Keys      : errObj.Keys,
		Value     : errObj.Value,
		Field     : errObj.Field,
		Msg       : errObj.Msg,
		ConstraintErrMsg : errObj.ErrTxt,
		ErrAppTag  : errObj.ErrAppTag,
	}


	return cvlErrObj

}

//Perform syntax checks
func (c *CVL) validateSyntax(data *yparser.YParserNode) (CVLErrorInfo, CVLRetCode) {
	var cvlErrObj CVLErrorInfo
	TRACE_LOG(INFO_DATA, TRACE_LIBYANG, "Validating syntax \n....")

	if errObj  := c.yp.ValidateSyntax(data); errObj.ErrCode != yparser.YP_SUCCESS {

		retCode := CVLRetCode(errObj.ErrCode)

			cvlErrObj =  CVLErrorInfo {
		             TableName : errObj.TableName,
		             ErrCode   : CVLRetCode(errObj.ErrCode),		
			     CVLErrDetails : cvlErrorMap[retCode], 
			     Keys      : errObj.Keys,
			     Value     : errObj.Value,
			     Field     : errObj.Field,
			     Msg       : errObj.Msg,
			     ConstraintErrMsg : errObj.ErrTxt,
			     ErrAppTag	: errObj.ErrAppTag,
			}



		return  cvlErrObj, retCode
	}

	return cvlErrObj, CVL_SUCCESS
}

//Perform semantic checks 
func (c *CVL) validateSemantics(data *yparser.YParserNode, appDepData *yparser.YParserNode) (CVLErrorInfo, CVLRetCode) {
	var cvlErrObj CVLErrorInfo
	
	if (SkipSemanticValidation() == true) {
		return cvlErrObj, CVL_SUCCESS
	}

	//Get dependent data from 
	depData := c.fetchDataToTmpCache() //fetch data to temp cache for temporary validation

	if (Tracing == true) {
		TRACE_LOG(INFO_API, TRACE_SEMANTIC, "Validating semantics data=%s\n depData =%s\n, appDepData=%s\n....", c.yp.NodeDump(data), c.yp.NodeDump(depData), c.yp.NodeDump(appDepData))
	}

	if errObj := c.yp.ValidateSemantics(data, depData, appDepData); errObj.ErrCode != yparser.YP_SUCCESS {

		retCode := CVLRetCode(errObj.ErrCode)

		cvlErrObj =  CVLErrorInfo {
			TableName : errObj.TableName,
		        ErrCode   : CVLRetCode(errObj.ErrCode),		
			CVLErrDetails : cvlErrorMap[retCode], 
			Keys      : errObj.Keys,
			Value     : errObj.Value,
			Field     : errObj.Field,
			Msg       : errObj.Msg,
			ConstraintErrMsg : errObj.ErrTxt,
			ErrAppTag	: errObj.ErrAppTag,
		}



		return  cvlErrObj, retCode
	}

	return cvlErrObj ,CVL_SUCCESS
}

//Add config data item to accumulate per table
func (c *CVL) addCfgDataItem(configData *map[string]interface{}, cfgDataItem CVLEditConfigData) (string, string){
	var cfgData map[string]interface{}//:= map[string]interface{}
	cfgData = *configData

	for tblName,_ := range modelInfo.tableInfo {
		//Check if table prefix matches to any schema table
		//i.e. has 'VLAN|' or 'PORT|' etc.
		if (strings.HasPrefix(cfgDataItem.Key, tblName + modelInfo.tableInfo[tblName].redisKeyDelim)) {
			prefixLen := len(tblName) + 1
			if (cfgDataItem.VOp == OP_DELETE) {
				//Don't add data it is delete operation
				return tblName, cfgDataItem.Key[prefixLen:]
			}
			if _, existing := cfgData[tblName]; existing {
				fieldsMap := cfgData[tblName].(map[string]interface{})
				fieldsMap[cfgDataItem.Key[prefixLen:]] = c.checkFieldMap(&cfgDataItem.Data)
			} else {
				fieldsMap := make(map[string]interface{})
				fieldsMap[cfgDataItem.Key[prefixLen:]] = c.checkFieldMap(&cfgDataItem.Data)
				cfgData[tblName] = fieldsMap
			}

			return tblName, cfgDataItem.Key[prefixLen:]
		}
	}

	return "",""
}

//Get table entry from cache for redis key
func dbCacheEntryGet(tableName, key string) (*yparser.YParserNode, CVLRetCode) {
	//First check if the table is cached
	topNode, _ := dbCacheGet(tableName)


	if (topNode != nil) {
		//Convert to Yang keys
		keyValuePair := getRedisToYangKeys(tableName, key)

		//Find if the entry is cached
		keyCompStr := ""
		for _, keyValItem := range keyValuePair {
			keyCompStr = keyCompStr + fmt.Sprintf("[%s='%s']",
			keyValItem.key, keyValItem.values[0])
		}

		entryNode := yparser.FindNode(topNode, fmt.Sprintf("//%s:%s/%s%s",
		modelInfo.tableInfo[tableName].modelName,
		modelInfo.tableInfo[tableName].modelName,
		tableName, keyCompStr))

		if (entryNode != nil) {
			return entryNode, CVL_SUCCESS
		}
	}

	return nil, CVL_ERROR
}

//Get the data from global cache
func dbCacheGet(tableName string) (*yparser.YParserNode, CVLRetCode) {

	TRACE_LOG(INFO_ALL, TRACE_CACHE, "Updating global cache for table %s", tableName)
	dbCacheTmp, existing := cvg.db[tableName]

	if (existing == false) {
		return  nil, CVL_FAILURE //not even empty cache present
	}

	if (dbCacheTmp.root != nil) {
		if (dbCacheTmp.expiry != 0) {
			//If cache is destroyable (i.e. expiry != 0), check if it has already expired.
			//If not expired update the time stamp
			if (time.Now().After(dbCacheTmp.startTime.Add(time.Second * time.Duration(dbCacheTmp.expiry)))) {
				//Cache expired, clear the cache
				dbCacheClear(tableName)

				return nil, CVL_ERROR
			}

			//Since the cache is used actively, update the timestamp
			dbCacheTmp.startTime = time.Now()
			cvg.db[tableName] = dbCacheTmp
		}

		return dbCacheTmp.root, CVL_SUCCESS
	} else {
		return  nil, CVL_SUCCESS // return success for no entry in Redis db and hencec empty cache
	}
}

//Get the table data from redis and cache it in yang node format
//expiry =0 never expire the cache
func dbCacheSet(update bool, tableName string, expiry uint16) CVLRetCode {

	cvg.mutex.Lock()

	//Get the data from redis and save it
	tableKeys, err:= redisClient.Keys(tableName +
	modelInfo.tableInfo[tableName].redisKeyDelim + "*").Result()

	if (err != nil) {
		cvg.mutex.Unlock()
		return CVL_FAILURE
	}

	TRACE_LOG(INFO_ALL, TRACE_CACHE, "Building global cache for table %s", tableName)

	tablePrefixLen := len(tableName + modelInfo.tableInfo[tableName].redisKeyDelim)
	for _, tableKey := range tableKeys {
		tableKey = tableKey[tablePrefixLen:] //remove table prefix
		if (cvg.cv.tmpDbCache[tableName] == nil) {
			cvg.cv.tmpDbCache[tableName] = map[string]interface{}{tableKey: nil}
		} else {
			tblMap := cvg.cv.tmpDbCache[tableName]
			tblMap.(map[string]interface{})[tableKey] =nil
			cvg.cv.tmpDbCache[tableName] = tblMap
		}
	}

	cvg.db[tableName] = dbCachedData{startTime:time.Now(), expiry: expiry,
	root: cvg.cv.fetchDataToTmpCache()}

	if (Tracing == true) {
		TRACE_LOG(INFO_ALL, TRACE_CACHE, "Cached Data = %v\n", cvg.cv.yp.NodeDump(cvg.db[tableName].root))
	}

	cvg.mutex.Unlock()

	//install keyspace notification for updating the cache
	if (update == false) {
		installDbChgNotif()
	}


	return CVL_SUCCESS
}

//Receive all updates for all tables on a single channel
func installDbChgNotif() {
	if (len(cvg.db) > 1) { //notif running for at least one table added previously
		cvg.stopChan <- 1 //stop active notification 
	}

	subList := make([]string, 0)
	for tableName, _ := range cvg.db {
		subList = append(subList,
		fmt.Sprintf("__keyspace@%d__:%s%s*", modelInfo.tableInfo[tableName].dbNum,
		tableName, modelInfo.tableInfo[tableName].redisKeyDelim))

	}

	//Listen on multiple channels
	cvg.pubsub = redisClient.PSubscribe(subList...)

	go func() {
		keySpacePrefixLen := len("__keyspace@4__:")

		notifCh := cvg.pubsub.Channel()
		for {
			select  {
			case <-cvg.stopChan:
				//stop this routine
				return
			case msg:= <-notifCh:
				//Handle update
				tbl, key := splitRedisKey(msg.Channel[keySpacePrefixLen:])
				if (tbl != "" && key != "") {
					dbCacheUpdate(tbl, key, msg.Payload)
				}
			}
		}
	}()
}

func dbCacheUpdate(tableName, key, op string) CVLRetCode {
	TRACE_LOG(INFO_ALL, TRACE_CACHE, "Updating global cache for table %s with key %s", tableName, key)

	//Find the node
	//Delete the entry in yang tree 

	cvg.mutex.Lock()

	node, _:= dbCacheEntryGet(tableName, key)
	if (node != nil) {
		//unlink and free the node
		cvg.cv.yp.FreeNode(node)
	}

	//Clear json map cache if any
	cvg.cv.clearTmpDbCache()

	tableKeys := []string {key}
	switch op {
	case "hset", "hmset", "hdel":
		//Get the entry from DB
		for _, tableKey := range tableKeys {
			cvg.cv.tmpDbCache[tableName] = map[string]interface{}{tableKey: nil}
		}

		//Get the translated Yang tree
		topNode := cvg.cv.fetchDataToTmpCache()

		//Merge the subtree with existing yang tree
		var errObj yparser.YParserError
		if (cvg.db[tableName].root != nil) {
			if topNode, errObj = cvg.cv.yp.MergeSubtree(cvg.db[tableName].root, topNode); errObj.ErrCode != yparser.YP_SUCCESS {
				cvg.mutex.Unlock()
				return CVL_ERROR
			}
		}

		//Update DB map
		db := cvg.db[tableName]
		db.root = topNode
		cvg.db[tableName] = db

	case "del":
		//NOP, alreday deleted the entry
	}

	cvg.mutex.Unlock()

	return CVL_SUCCESS
}

//Clear cache data for given table
func dbCacheClear(tableName string) CVLRetCode {
	cvg.cv.yp.FreeNode(cvg.db[tableName].root)
	delete(cvg.db, tableName)

	TRACE_LOG(INFO_ALL, TRACE_CACHE, "Clearing global cache for table %s", tableName)

	return CVL_SUCCESS
}


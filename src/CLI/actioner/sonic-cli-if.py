#!/usr/bin/python
import sys
import time
import json
import ast
import yaml
import openconfig_interfaces_client
from openconfig_interfaces_client.rest import ApiException
from scripts.render_cli import show_cli_output

import urllib3
urllib3.disable_warnings()


plugins = dict()

def register(func):
    """Register sdk client method as a plug-in"""
    plugins[func.__name__] = func
    return func


def call_method(name, args):
    method = plugins[name]
    return method(args)

def generate_body(func, args):
    body = None
    print("args[0] is ", args[0])
    if func.__name__ == 'patch_openconfig_interfaces_interfaces_interface':
       keypath = [ args[0] ]
       body = {}

    elif func.__name__ == 'delete_openconfig_interfaces_interfaces_interface':
       keypath = [ args[0] ]
    elif func.__name__ == 'patch_openconfig_interfaces_interfaces_interface_config_description':
       keypath = [ args[0] ]
       body = { "openconfig-interfaces:description": args[1] }
    elif func.__name__ == 'patch_openconfig_interfaces_interfaces_interface_config_enabled':
       keypath = [ args[0] ]
       if args[1] == "True":
           body = { "openconfig-interfaces:enabled": True }
       else:
           body = { "openconfig-interfaces:enabled": False }
    elif func.__name__ == 'patch_openconfig_interfaces_interfaces_interface_config_mtu':
       keypath = [ args[0] ]
       body = { "openconfig-interfaces:mtu":  int(args[1]) }
    elif func.__name__ == 'patch_openconfig_if_ip_interfaces_interface_subinterfaces_subinterface_ipv4_addresses_address_config':
       sp = args[1].split('/')
       keypath = [ args[0], 0, sp[0] ]
       body = { "openconfig-if-ip:config":  {"ip" : sp[0], "prefix-length" : int(sp[1])} }
    elif func.__name__ == 'patch_openconfig_if_ip_interfaces_interface_subinterfaces_subinterface_ipv6_addresses_address_config':
       sp = args[1].split('/')
       keypath = [ args[0], 0, sp[0] ]
       body = { "openconfig-if-ip:config":  {"ip" : sp[0], "prefix-length" : int(sp[1])} }
    elif func.__name__ == 'patch_openconfig_vlan_interfaces_interface_ethernet_switched_vlan_config':
       keypath = [args[0]]
       if args[1] == "ACCESS":
           body = {"openconfig-vlan:config": {"interface-mode": "ACCESS","access-vlan": int(args[2])}}
       else:
           body = {"openconfig-vlan:config": {"interface-mode": "TRUNK","trunk-vlans": [int(args[2])]}}
    elif func.__name__ == 'delete_openconfig_vlan_interfaces_interface_ethernet_switched_vlan_config_access_vlan':
        keypath = [args[0]]
    elif func.__name__ == 'delete_openconfig_if_ip_interfaces_interface_subinterfaces_subinterface_ipv4_addresses_address_config_prefix_length':
       keypath = [args[0], 0, args[1]]
    elif func.__name__ == 'delete_openconfig_if_ip_interfaces_interface_subinterfaces_subinterface_ipv6_addresses_address_config_prefix_length':
       keypath = [args[0], 0, args[1]]
    elif func.__name__ == 'get_openconfig_interfaces_interfaces_interface':
	keypath = [args[0]]
    elif func.__name__ == 'get_openconfig_interfaces_interfaces':
        keypath = []
    else:
       body = {} 

    return keypath,body

def getId(item):
    prfx = "Ethernet"
    state_dict = item['state']
    ifName = state_dict['name']

    if ifName.startswith(prfx):
        ifId = int(ifName[len(prfx):])
        return ifId
    return ifName

def run(func, args):
    print("args is---------", args)
    c = openconfig_interfaces_client.Configuration()
    c.verify_ssl = False
    aa = openconfig_interfaces_client.OpenconfigInterfacesApi(api_client=openconfig_interfaces_client.ApiClient(configuration=c))
        
    if "Portchannel" in args[0] and func.__name__ == 'patch_openconfig_interfaces_interfaces_interface':
        #print("check if not exist, then add body to json file") # read file check names
       #('args are', ['Portchannel11'])
        body = {
                    "name": args[0],
                    "min-links": 1,
                    "members": []
                }
        #append body to list 
        with open('dummy_data.json', 'r') as f:
            data= yaml.safe_load(f)    
        data['openconfig-interfaces:interface'].append(body)
        with open('dummy_data.json', 'w') as f:
            json.dump(data, f, sort_keys=True, indent=4)
        print ("Success")
        return    

    if "Portchannel" in args[0] and func.__name__ == 'get_openconfig_if_aggregate_interfaces_interface_aggregation_state':
        with open('dummy_data.json', 'r') as f:
            data= yaml.safe_load(f)
        show_cli_output("portchannel_show.j2", data)
        return    

    if func.__name__ == 'patch_openconfig_if_aggregate_interfaces_interface_ethernet_config_aggregate_id': #add members
        print ("Success --inside aggregate id fun")
        return   
 
    if func.__name__ == 'patch_openconfig_if_aggregate_interfaces_interface_aggregation_config_min_links': #change min-links value
        print ("Success")
        return  
 
    # create a body block
    keypath, body = generate_body(func, args)
        
    try:
        if body is not None:
           api_response = getattr(aa,func.__name__)(*keypath, body=body)
        else :
           api_response = getattr(aa,func.__name__)(*keypath)

        if api_response is None:
            print ("Success")
        else:
            # Get Command Output
            api_response = aa.api_client.sanitize_for_serialization(api_response)
            if 'openconfig-interfaces:interfaces' in api_response:
                value = api_response['openconfig-interfaces:interfaces']
                if 'interface' in value:
                    tup = value['interface']
                    value['interface'] = sorted(tup, key=getId)

            if api_response is None:
                print("Failed")
            else:
                if func.__name__ == 'get_openconfig_interfaces_interfaces_interface':
                     show_cli_output(args[1], api_response)
                elif func.__name__ == 'get_openconfig_interfaces_interfaces':
                     show_cli_output(args[0], api_response)
                else:
                     return
    except ApiException as e:
        #print("Exception when calling OpenconfigInterfacesApi->%s : %s\n" %(func.__name__, e))
        if e.body != "":
            body = json.loads(e.body)
            if "ietf-restconf:errors" in body:
                 err = body["ietf-restconf:errors"]
                 if "error" in err:
                     errList = err["error"]

                     errDict = {}
                     for dict in errList:
                         for k, v in dict.iteritems():
                              errDict[k] = v

                     if "error-message" in errDict:
                         print "%Error: " + errDict["error-message"]
                         return
                     print "%Error: Transaction Failure"
                     return
            print "%Error: Transaction Failure"
        else:
            print "Failed"

if __name__ == '__main__':

    func = eval(sys.argv[1], globals(), openconfig_interfaces_client.OpenconfigInterfacesApi.__dict__)
    print("argv is ----", sys.argv[0:])
    run(func, sys.argv[2:])

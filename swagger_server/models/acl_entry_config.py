# coding: utf-8

from __future__ import absolute_import
from datetime import date, datetime  # noqa: F401

from typing import List, Dict  # noqa: F401

from swagger_server.models.base_model_ import Model
from swagger_server.models.acl_openconfigaclacl_aclsets_aclentries_config import AclOpenconfigaclaclAclsetsAclentriesConfig  # noqa: F401,E501
from swagger_server import util


class AclEntryConfig(Model):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    def __init__(self, openconfig_aclconfig: AclOpenconfigaclaclAclsetsAclentriesConfig=None):  # noqa: E501
        """AclEntryConfig - a model defined in Swagger

        :param openconfig_aclconfig: The openconfig_aclconfig of this AclEntryConfig.  # noqa: E501
        :type openconfig_aclconfig: AclOpenconfigaclaclAclsetsAclentriesConfig
        """
        self.swagger_types = {
            'openconfig_aclconfig': AclOpenconfigaclaclAclsetsAclentriesConfig
        }

        self.attribute_map = {
            'openconfig_aclconfig': 'openconfig-acl:config'
        }

        self._openconfig_aclconfig = openconfig_aclconfig

    @classmethod
    def from_dict(cls, dikt) -> 'AclEntryConfig':
        """Returns the dict as a model

        :param dikt: A dict.
        :type: dict
        :return: The acl_entry_config of this AclEntryConfig.  # noqa: E501
        :rtype: AclEntryConfig
        """
        return util.deserialize_model(dikt, cls)

    @property
    def openconfig_aclconfig(self) -> AclOpenconfigaclaclAclsetsAclentriesConfig:
        """Gets the openconfig_aclconfig of this AclEntryConfig.


        :return: The openconfig_aclconfig of this AclEntryConfig.
        :rtype: AclOpenconfigaclaclAclsetsAclentriesConfig
        """
        return self._openconfig_aclconfig

    @openconfig_aclconfig.setter
    def openconfig_aclconfig(self, openconfig_aclconfig: AclOpenconfigaclaclAclsetsAclentriesConfig):
        """Sets the openconfig_aclconfig of this AclEntryConfig.


        :param openconfig_aclconfig: The openconfig_aclconfig of this AclEntryConfig.
        :type openconfig_aclconfig: AclOpenconfigaclaclAclsetsAclentriesConfig
        """

        self._openconfig_aclconfig = openconfig_aclconfig
# coding: utf-8

from __future__ import absolute_import
from datetime import date, datetime  # noqa: F401

from typing import List, Dict  # noqa: F401

from swagger_server.models.base_model_ import Model
from swagger_server.models.acl_openconfigaclacl_aclsets_aclentries_actions_config import AclOpenconfigaclaclAclsetsAclentriesActionsConfig  # noqa: F401,E501
from swagger_server.models.actions_config import ActionsConfig  # noqa: F401,E501
from swagger_server import util


class PostActionsConfig(Model):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    def __init__(self, openconfig_aclconfig: AclOpenconfigaclaclAclsetsAclentriesActionsConfig=None):  # noqa: E501
        """PostActionsConfig - a model defined in Swagger

        :param openconfig_aclconfig: The openconfig_aclconfig of this PostActionsConfig.  # noqa: E501
        :type openconfig_aclconfig: AclOpenconfigaclaclAclsetsAclentriesActionsConfig
        """
        self.swagger_types = {
            'openconfig_aclconfig': AclOpenconfigaclaclAclsetsAclentriesActionsConfig
        }

        self.attribute_map = {
            'openconfig_aclconfig': 'openconfig-acl:config'
        }

        self._openconfig_aclconfig = openconfig_aclconfig

    @classmethod
    def from_dict(cls, dikt) -> 'PostActionsConfig':
        """Returns the dict as a model

        :param dikt: A dict.
        :type: dict
        :return: The post_actions_config of this PostActionsConfig.  # noqa: E501
        :rtype: PostActionsConfig
        """
        return util.deserialize_model(dikt, cls)

    @property
    def openconfig_aclconfig(self) -> AclOpenconfigaclaclAclsetsAclentriesActionsConfig:
        """Gets the openconfig_aclconfig of this PostActionsConfig.


        :return: The openconfig_aclconfig of this PostActionsConfig.
        :rtype: AclOpenconfigaclaclAclsetsAclentriesActionsConfig
        """
        return self._openconfig_aclconfig

    @openconfig_aclconfig.setter
    def openconfig_aclconfig(self, openconfig_aclconfig: AclOpenconfigaclaclAclsetsAclentriesActionsConfig):
        """Sets the openconfig_aclconfig of this PostActionsConfig.


        :param openconfig_aclconfig: The openconfig_aclconfig of this PostActionsConfig.
        :type openconfig_aclconfig: AclOpenconfigaclaclAclsetsAclentriesActionsConfig
        """

        self._openconfig_aclconfig = openconfig_aclconfig
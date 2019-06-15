# coding: utf-8

from __future__ import absolute_import
from datetime import date, datetime  # noqa: F401

from typing import List, Dict  # noqa: F401

from swagger_server.models.base_model_ import Model
from swagger_server.models.acl_openconfigaclacl_aclsets_aclentries_actions_config import AclOpenconfigaclaclAclsetsAclentriesActionsConfig  # noqa: F401,E501
from swagger_server import util


class GetActionsState(Model):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    def __init__(self, openconfig_aclstate: AclOpenconfigaclaclAclsetsAclentriesActionsConfig=None):  # noqa: E501
        """GetActionsState - a model defined in Swagger

        :param openconfig_aclstate: The openconfig_aclstate of this GetActionsState.  # noqa: E501
        :type openconfig_aclstate: AclOpenconfigaclaclAclsetsAclentriesActionsConfig
        """
        self.swagger_types = {
            'openconfig_aclstate': AclOpenconfigaclaclAclsetsAclentriesActionsConfig
        }

        self.attribute_map = {
            'openconfig_aclstate': 'openconfig-acl:state'
        }

        self._openconfig_aclstate = openconfig_aclstate

    @classmethod
    def from_dict(cls, dikt) -> 'GetActionsState':
        """Returns the dict as a model

        :param dikt: A dict.
        :type: dict
        :return: The get_actions_state of this GetActionsState.  # noqa: E501
        :rtype: GetActionsState
        """
        return util.deserialize_model(dikt, cls)

    @property
    def openconfig_aclstate(self) -> AclOpenconfigaclaclAclsetsAclentriesActionsConfig:
        """Gets the openconfig_aclstate of this GetActionsState.


        :return: The openconfig_aclstate of this GetActionsState.
        :rtype: AclOpenconfigaclaclAclsetsAclentriesActionsConfig
        """
        return self._openconfig_aclstate

    @openconfig_aclstate.setter
    def openconfig_aclstate(self, openconfig_aclstate: AclOpenconfigaclaclAclsetsAclentriesActionsConfig):
        """Sets the openconfig_aclstate of this GetActionsState.


        :param openconfig_aclstate: The openconfig_aclstate of this GetActionsState.
        :type openconfig_aclstate: AclOpenconfigaclaclAclsetsAclentriesActionsConfig
        """

        self._openconfig_aclstate = openconfig_aclstate
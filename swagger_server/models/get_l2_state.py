# coding: utf-8

from __future__ import absolute_import
from datetime import date, datetime  # noqa: F401

from typing import List, Dict  # noqa: F401

from swagger_server.models.base_model_ import Model
from swagger_server.models.acl_openconfigaclacl_aclsets_aclentries_l2_config import AclOpenconfigaclaclAclsetsAclentriesL2Config  # noqa: F401,E501
from swagger_server import util


class GetL2State(Model):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    def __init__(self, openconfig_aclstate: AclOpenconfigaclaclAclsetsAclentriesL2Config=None):  # noqa: E501
        """GetL2State - a model defined in Swagger

        :param openconfig_aclstate: The openconfig_aclstate of this GetL2State.  # noqa: E501
        :type openconfig_aclstate: AclOpenconfigaclaclAclsetsAclentriesL2Config
        """
        self.swagger_types = {
            'openconfig_aclstate': AclOpenconfigaclaclAclsetsAclentriesL2Config
        }

        self.attribute_map = {
            'openconfig_aclstate': 'openconfig-acl:state'
        }

        self._openconfig_aclstate = openconfig_aclstate

    @classmethod
    def from_dict(cls, dikt) -> 'GetL2State':
        """Returns the dict as a model

        :param dikt: A dict.
        :type: dict
        :return: The get_l2_state of this GetL2State.  # noqa: E501
        :rtype: GetL2State
        """
        return util.deserialize_model(dikt, cls)

    @property
    def openconfig_aclstate(self) -> AclOpenconfigaclaclAclsetsAclentriesL2Config:
        """Gets the openconfig_aclstate of this GetL2State.


        :return: The openconfig_aclstate of this GetL2State.
        :rtype: AclOpenconfigaclaclAclsetsAclentriesL2Config
        """
        return self._openconfig_aclstate

    @openconfig_aclstate.setter
    def openconfig_aclstate(self, openconfig_aclstate: AclOpenconfigaclaclAclsetsAclentriesL2Config):
        """Sets the openconfig_aclstate of this GetL2State.


        :param openconfig_aclstate: The openconfig_aclstate of this GetL2State.
        :type openconfig_aclstate: AclOpenconfigaclaclAclsetsAclentriesL2Config
        """

        self._openconfig_aclstate = openconfig_aclstate
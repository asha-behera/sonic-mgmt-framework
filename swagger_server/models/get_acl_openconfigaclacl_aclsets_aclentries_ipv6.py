# coding: utf-8

from __future__ import absolute_import
from datetime import date, datetime  # noqa: F401

from typing import List, Dict  # noqa: F401

from swagger_server.models.base_model_ import Model
from swagger_server.models.acl_openconfigaclacl_aclsets_aclentries_ipv6_config import AclOpenconfigaclaclAclsetsAclentriesIpv6Config  # noqa: F401,E501
from swagger_server import util


class GetAclOpenconfigaclaclAclsetsAclentriesIpv6(Model):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    def __init__(self, config: AclOpenconfigaclaclAclsetsAclentriesIpv6Config=None, state: AclOpenconfigaclaclAclsetsAclentriesIpv6Config=None):  # noqa: E501
        """GetAclOpenconfigaclaclAclsetsAclentriesIpv6 - a model defined in Swagger

        :param config: The config of this GetAclOpenconfigaclaclAclsetsAclentriesIpv6.  # noqa: E501
        :type config: AclOpenconfigaclaclAclsetsAclentriesIpv6Config
        :param state: The state of this GetAclOpenconfigaclaclAclsetsAclentriesIpv6.  # noqa: E501
        :type state: AclOpenconfigaclaclAclsetsAclentriesIpv6Config
        """
        self.swagger_types = {
            'config': AclOpenconfigaclaclAclsetsAclentriesIpv6Config,
            'state': AclOpenconfigaclaclAclsetsAclentriesIpv6Config
        }

        self.attribute_map = {
            'config': 'config',
            'state': 'state'
        }

        self._config = config
        self._state = state

    @classmethod
    def from_dict(cls, dikt) -> 'GetAclOpenconfigaclaclAclsetsAclentriesIpv6':
        """Returns the dict as a model

        :param dikt: A dict.
        :type: dict
        :return: The get_acl_openconfigaclacl_aclsets_aclentries_ipv6 of this GetAclOpenconfigaclaclAclsetsAclentriesIpv6.  # noqa: E501
        :rtype: GetAclOpenconfigaclaclAclsetsAclentriesIpv6
        """
        return util.deserialize_model(dikt, cls)

    @property
    def config(self) -> AclOpenconfigaclaclAclsetsAclentriesIpv6Config:
        """Gets the config of this GetAclOpenconfigaclaclAclsetsAclentriesIpv6.


        :return: The config of this GetAclOpenconfigaclaclAclsetsAclentriesIpv6.
        :rtype: AclOpenconfigaclaclAclsetsAclentriesIpv6Config
        """
        return self._config

    @config.setter
    def config(self, config: AclOpenconfigaclaclAclsetsAclentriesIpv6Config):
        """Sets the config of this GetAclOpenconfigaclaclAclsetsAclentriesIpv6.


        :param config: The config of this GetAclOpenconfigaclaclAclsetsAclentriesIpv6.
        :type config: AclOpenconfigaclaclAclsetsAclentriesIpv6Config
        """

        self._config = config

    @property
    def state(self) -> AclOpenconfigaclaclAclsetsAclentriesIpv6Config:
        """Gets the state of this GetAclOpenconfigaclaclAclsetsAclentriesIpv6.


        :return: The state of this GetAclOpenconfigaclaclAclsetsAclentriesIpv6.
        :rtype: AclOpenconfigaclaclAclsetsAclentriesIpv6Config
        """
        return self._state

    @state.setter
    def state(self, state: AclOpenconfigaclaclAclsetsAclentriesIpv6Config):
        """Sets the state of this GetAclOpenconfigaclaclAclsetsAclentriesIpv6.


        :param state: The state of this GetAclOpenconfigaclaclAclsetsAclentriesIpv6.
        :type state: AclOpenconfigaclaclAclsetsAclentriesIpv6Config
        """

        self._state = state
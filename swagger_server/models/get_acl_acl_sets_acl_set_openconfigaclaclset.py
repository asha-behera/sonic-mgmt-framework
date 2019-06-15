# coding: utf-8

from __future__ import absolute_import
from datetime import date, datetime  # noqa: F401

from typing import List, Dict  # noqa: F401

from swagger_server.models.base_model_ import Model
from swagger_server.models.acl_openconfigaclacl_aclsets_config import AclOpenconfigaclaclAclsetsConfig  # noqa: F401,E501
from swagger_server.models.get_acl_openconfigaclacl_aclsets_aclentries import GetAclOpenconfigaclaclAclsetsAclentries  # noqa: F401,E501
from swagger_server import util


class GetAclAclSetsAclSetOpenconfigaclaclset(Model):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    def __init__(self, config: AclOpenconfigaclaclAclsetsConfig=None, state: AclOpenconfigaclaclAclsetsConfig=None, acl_entries: GetAclOpenconfigaclaclAclsetsAclentries=None):  # noqa: E501
        """GetAclAclSetsAclSetOpenconfigaclaclset - a model defined in Swagger

        :param config: The config of this GetAclAclSetsAclSetOpenconfigaclaclset.  # noqa: E501
        :type config: AclOpenconfigaclaclAclsetsConfig
        :param state: The state of this GetAclAclSetsAclSetOpenconfigaclaclset.  # noqa: E501
        :type state: AclOpenconfigaclaclAclsetsConfig
        :param acl_entries: The acl_entries of this GetAclAclSetsAclSetOpenconfigaclaclset.  # noqa: E501
        :type acl_entries: GetAclOpenconfigaclaclAclsetsAclentries
        """
        self.swagger_types = {
            'config': AclOpenconfigaclaclAclsetsConfig,
            'state': AclOpenconfigaclaclAclsetsConfig,
            'acl_entries': GetAclOpenconfigaclaclAclsetsAclentries
        }

        self.attribute_map = {
            'config': 'config',
            'state': 'state',
            'acl_entries': 'acl-entries'
        }

        self._config = config
        self._state = state
        self._acl_entries = acl_entries

    @classmethod
    def from_dict(cls, dikt) -> 'GetAclAclSetsAclSetOpenconfigaclaclset':
        """Returns the dict as a model

        :param dikt: A dict.
        :type: dict
        :return: The get_acl_acl_sets_acl_set_openconfigaclaclset of this GetAclAclSetsAclSetOpenconfigaclaclset.  # noqa: E501
        :rtype: GetAclAclSetsAclSetOpenconfigaclaclset
        """
        return util.deserialize_model(dikt, cls)

    @property
    def config(self) -> AclOpenconfigaclaclAclsetsConfig:
        """Gets the config of this GetAclAclSetsAclSetOpenconfigaclaclset.


        :return: The config of this GetAclAclSetsAclSetOpenconfigaclaclset.
        :rtype: AclOpenconfigaclaclAclsetsConfig
        """
        return self._config

    @config.setter
    def config(self, config: AclOpenconfigaclaclAclsetsConfig):
        """Sets the config of this GetAclAclSetsAclSetOpenconfigaclaclset.


        :param config: The config of this GetAclAclSetsAclSetOpenconfigaclaclset.
        :type config: AclOpenconfigaclaclAclsetsConfig
        """

        self._config = config

    @property
    def state(self) -> AclOpenconfigaclaclAclsetsConfig:
        """Gets the state of this GetAclAclSetsAclSetOpenconfigaclaclset.


        :return: The state of this GetAclAclSetsAclSetOpenconfigaclaclset.
        :rtype: AclOpenconfigaclaclAclsetsConfig
        """
        return self._state

    @state.setter
    def state(self, state: AclOpenconfigaclaclAclsetsConfig):
        """Sets the state of this GetAclAclSetsAclSetOpenconfigaclaclset.


        :param state: The state of this GetAclAclSetsAclSetOpenconfigaclaclset.
        :type state: AclOpenconfigaclaclAclsetsConfig
        """

        self._state = state

    @property
    def acl_entries(self) -> GetAclOpenconfigaclaclAclsetsAclentries:
        """Gets the acl_entries of this GetAclAclSetsAclSetOpenconfigaclaclset.


        :return: The acl_entries of this GetAclAclSetsAclSetOpenconfigaclaclset.
        :rtype: GetAclOpenconfigaclaclAclsetsAclentries
        """
        return self._acl_entries

    @acl_entries.setter
    def acl_entries(self, acl_entries: GetAclOpenconfigaclaclAclsetsAclentries):
        """Sets the acl_entries of this GetAclAclSetsAclSetOpenconfigaclaclset.


        :param acl_entries: The acl_entries of this GetAclAclSetsAclSetOpenconfigaclaclset.
        :type acl_entries: GetAclOpenconfigaclaclAclsetsAclentries
        """

        self._acl_entries = acl_entries
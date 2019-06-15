# coding: utf-8

from __future__ import absolute_import
from datetime import date, datetime  # noqa: F401

from typing import List, Dict  # noqa: F401

from swagger_server.models.base_model_ import Model
from swagger_server.models.acl_openconfigaclacl_aclsets_aclentries_transport_config import AclOpenconfigaclaclAclsetsAclentriesTransportConfig  # noqa: F401,E501
from swagger_server import util


class AclOpenconfigaclaclAclsetsAclentriesTransport(Model):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    def __init__(self, config: AclOpenconfigaclaclAclsetsAclentriesTransportConfig=None):  # noqa: E501
        """AclOpenconfigaclaclAclsetsAclentriesTransport - a model defined in Swagger

        :param config: The config of this AclOpenconfigaclaclAclsetsAclentriesTransport.  # noqa: E501
        :type config: AclOpenconfigaclaclAclsetsAclentriesTransportConfig
        """
        self.swagger_types = {
            'config': AclOpenconfigaclaclAclsetsAclentriesTransportConfig
        }

        self.attribute_map = {
            'config': 'config'
        }

        self._config = config

    @classmethod
    def from_dict(cls, dikt) -> 'AclOpenconfigaclaclAclsetsAclentriesTransport':
        """Returns the dict as a model

        :param dikt: A dict.
        :type: dict
        :return: The acl_openconfigaclacl_aclsets_aclentries_transport of this AclOpenconfigaclaclAclsetsAclentriesTransport.  # noqa: E501
        :rtype: AclOpenconfigaclaclAclsetsAclentriesTransport
        """
        return util.deserialize_model(dikt, cls)

    @property
    def config(self) -> AclOpenconfigaclaclAclsetsAclentriesTransportConfig:
        """Gets the config of this AclOpenconfigaclaclAclsetsAclentriesTransport.


        :return: The config of this AclOpenconfigaclaclAclsetsAclentriesTransport.
        :rtype: AclOpenconfigaclaclAclsetsAclentriesTransportConfig
        """
        return self._config

    @config.setter
    def config(self, config: AclOpenconfigaclaclAclsetsAclentriesTransportConfig):
        """Sets the config of this AclOpenconfigaclaclAclsetsAclentriesTransport.


        :param config: The config of this AclOpenconfigaclaclAclsetsAclentriesTransport.
        :type config: AclOpenconfigaclaclAclsetsAclentriesTransportConfig
        """

        self._config = config
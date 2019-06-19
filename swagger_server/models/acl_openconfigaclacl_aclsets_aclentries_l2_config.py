# coding: utf-8

from __future__ import absolute_import
from datetime import date, datetime  # noqa: F401

from typing import List, Dict  # noqa: F401

from swagger_server.models.base_model_ import Model
from swagger_server import util


class AclOpenconfigaclaclAclsetsAclentriesL2Config(Model):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    def __init__(self, source_mac: str=None, source_mac_mask: str=None, destination_mac: str=None, destination_mac_mask: str=None, ethertype: str=None):  # noqa: E501
        """AclOpenconfigaclaclAclsetsAclentriesL2Config - a model defined in Swagger

        :param source_mac: The source_mac of this AclOpenconfigaclaclAclsetsAclentriesL2Config.  # noqa: E501
        :type source_mac: str
        :param source_mac_mask: The source_mac_mask of this AclOpenconfigaclaclAclsetsAclentriesL2Config.  # noqa: E501
        :type source_mac_mask: str
        :param destination_mac: The destination_mac of this AclOpenconfigaclaclAclsetsAclentriesL2Config.  # noqa: E501
        :type destination_mac: str
        :param destination_mac_mask: The destination_mac_mask of this AclOpenconfigaclaclAclsetsAclentriesL2Config.  # noqa: E501
        :type destination_mac_mask: str
        :param ethertype: The ethertype of this AclOpenconfigaclaclAclsetsAclentriesL2Config.  # noqa: E501
        :type ethertype: str
        """
        self.swagger_types = {
            'source_mac': str,
            'source_mac_mask': str,
            'destination_mac': str,
            'destination_mac_mask': str,
            'ethertype': str
        }

        self.attribute_map = {
            'source_mac': 'source-mac',
            'source_mac_mask': 'source-mac-mask',
            'destination_mac': 'destination-mac',
            'destination_mac_mask': 'destination-mac-mask',
            'ethertype': 'ethertype'
        }

        self._source_mac = source_mac
        self._source_mac_mask = source_mac_mask
        self._destination_mac = destination_mac
        self._destination_mac_mask = destination_mac_mask
        self._ethertype = ethertype

    @classmethod
    def from_dict(cls, dikt) -> 'AclOpenconfigaclaclAclsetsAclentriesL2Config':
        """Returns the dict as a model

        :param dikt: A dict.
        :type: dict
        :return: The acl_openconfigaclacl_aclsets_aclentries_l2_config of this AclOpenconfigaclaclAclsetsAclentriesL2Config.  # noqa: E501
        :rtype: AclOpenconfigaclaclAclsetsAclentriesL2Config
        """
        return util.deserialize_model(dikt, cls)

    @property
    def source_mac(self) -> str:
        """Gets the source_mac of this AclOpenconfigaclaclAclsetsAclentriesL2Config.


        :return: The source_mac of this AclOpenconfigaclaclAclsetsAclentriesL2Config.
        :rtype: str
        """
        return self._source_mac

    @source_mac.setter
    def source_mac(self, source_mac: str):
        """Sets the source_mac of this AclOpenconfigaclaclAclsetsAclentriesL2Config.


        :param source_mac: The source_mac of this AclOpenconfigaclaclAclsetsAclentriesL2Config.
        :type source_mac: str
        """

        self._source_mac = source_mac

    @property
    def source_mac_mask(self) -> str:
        """Gets the source_mac_mask of this AclOpenconfigaclaclAclsetsAclentriesL2Config.


        :return: The source_mac_mask of this AclOpenconfigaclaclAclsetsAclentriesL2Config.
        :rtype: str
        """
        return self._source_mac_mask

    @source_mac_mask.setter
    def source_mac_mask(self, source_mac_mask: str):
        """Sets the source_mac_mask of this AclOpenconfigaclaclAclsetsAclentriesL2Config.


        :param source_mac_mask: The source_mac_mask of this AclOpenconfigaclaclAclsetsAclentriesL2Config.
        :type source_mac_mask: str
        """

        self._source_mac_mask = source_mac_mask

    @property
    def destination_mac(self) -> str:
        """Gets the destination_mac of this AclOpenconfigaclaclAclsetsAclentriesL2Config.


        :return: The destination_mac of this AclOpenconfigaclaclAclsetsAclentriesL2Config.
        :rtype: str
        """
        return self._destination_mac

    @destination_mac.setter
    def destination_mac(self, destination_mac: str):
        """Sets the destination_mac of this AclOpenconfigaclaclAclsetsAclentriesL2Config.


        :param destination_mac: The destination_mac of this AclOpenconfigaclaclAclsetsAclentriesL2Config.
        :type destination_mac: str
        """

        self._destination_mac = destination_mac

    @property
    def destination_mac_mask(self) -> str:
        """Gets the destination_mac_mask of this AclOpenconfigaclaclAclsetsAclentriesL2Config.


        :return: The destination_mac_mask of this AclOpenconfigaclaclAclsetsAclentriesL2Config.
        :rtype: str
        """
        return self._destination_mac_mask

    @destination_mac_mask.setter
    def destination_mac_mask(self, destination_mac_mask: str):
        """Sets the destination_mac_mask of this AclOpenconfigaclaclAclsetsAclentriesL2Config.


        :param destination_mac_mask: The destination_mac_mask of this AclOpenconfigaclaclAclsetsAclentriesL2Config.
        :type destination_mac_mask: str
        """

        self._destination_mac_mask = destination_mac_mask

    @property
    def ethertype(self) -> str:
        """Gets the ethertype of this AclOpenconfigaclaclAclsetsAclentriesL2Config.


        :return: The ethertype of this AclOpenconfigaclaclAclsetsAclentriesL2Config.
        :rtype: str
        """
        return self._ethertype

    @ethertype.setter
    def ethertype(self, ethertype: str):
        """Sets the ethertype of this AclOpenconfigaclaclAclsetsAclentriesL2Config.


        :param ethertype: The ethertype of this AclOpenconfigaclaclAclsetsAclentriesL2Config.
        :type ethertype: str
        """

        self._ethertype = ethertype
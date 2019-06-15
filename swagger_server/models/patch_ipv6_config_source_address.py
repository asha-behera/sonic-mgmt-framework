# coding: utf-8

from __future__ import absolute_import
from datetime import date, datetime  # noqa: F401

from typing import List, Dict  # noqa: F401

from swagger_server.models.base_model_ import Model
from swagger_server.models.ipv6_config_source_address import Ipv6ConfigSourceAddress  # noqa: F401,E501
from swagger_server import util


class PatchIpv6ConfigSourceAddress(Model):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    def __init__(self, source_address: str=None):  # noqa: E501
        """PatchIpv6ConfigSourceAddress - a model defined in Swagger

        :param source_address: The source_address of this PatchIpv6ConfigSourceAddress.  # noqa: E501
        :type source_address: str
        """
        self.swagger_types = {
            'source_address': str
        }

        self.attribute_map = {
            'source_address': 'source-address'
        }

        self._source_address = source_address

    @classmethod
    def from_dict(cls, dikt) -> 'PatchIpv6ConfigSourceAddress':
        """Returns the dict as a model

        :param dikt: A dict.
        :type: dict
        :return: The patch_ipv6_config_source_address of this PatchIpv6ConfigSourceAddress.  # noqa: E501
        :rtype: PatchIpv6ConfigSourceAddress
        """
        return util.deserialize_model(dikt, cls)

    @property
    def source_address(self) -> str:
        """Gets the source_address of this PatchIpv6ConfigSourceAddress.


        :return: The source_address of this PatchIpv6ConfigSourceAddress.
        :rtype: str
        """
        return self._source_address

    @source_address.setter
    def source_address(self, source_address: str):
        """Sets the source_address of this PatchIpv6ConfigSourceAddress.


        :param source_address: The source_address of this PatchIpv6ConfigSourceAddress.
        :type source_address: str
        """

        self._source_address = source_address
# coding: utf-8

from __future__ import absolute_import
from datetime import date, datetime  # noqa: F401

from typing import List, Dict  # noqa: F401

from swagger_server.models.base_model_ import Model
from swagger_server.models.config_destination_address import ConfigDestinationAddress  # noqa: F401,E501
from swagger_server import util


class PatchConfigDestinationAddress(Model):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    def __init__(self, destination_address: str=None):  # noqa: E501
        """PatchConfigDestinationAddress - a model defined in Swagger

        :param destination_address: The destination_address of this PatchConfigDestinationAddress.  # noqa: E501
        :type destination_address: str
        """
        self.swagger_types = {
            'destination_address': str
        }

        self.attribute_map = {
            'destination_address': 'destination-address'
        }

        self._destination_address = destination_address

    @classmethod
    def from_dict(cls, dikt) -> 'PatchConfigDestinationAddress':
        """Returns the dict as a model

        :param dikt: A dict.
        :type: dict
        :return: The patch_config_destination_address of this PatchConfigDestinationAddress.  # noqa: E501
        :rtype: PatchConfigDestinationAddress
        """
        return util.deserialize_model(dikt, cls)

    @property
    def destination_address(self) -> str:
        """Gets the destination_address of this PatchConfigDestinationAddress.


        :return: The destination_address of this PatchConfigDestinationAddress.
        :rtype: str
        """
        return self._destination_address

    @destination_address.setter
    def destination_address(self, destination_address: str):
        """Sets the destination_address of this PatchConfigDestinationAddress.


        :param destination_address: The destination_address of this PatchConfigDestinationAddress.
        :type destination_address: str
        """

        self._destination_address = destination_address
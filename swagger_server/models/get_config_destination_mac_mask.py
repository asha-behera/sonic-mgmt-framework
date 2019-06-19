# coding: utf-8

from __future__ import absolute_import
from datetime import date, datetime  # noqa: F401

from typing import List, Dict  # noqa: F401

from swagger_server.models.base_model_ import Model
from swagger_server import util


class GetConfigDestinationMacMask(Model):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    def __init__(self, destination_mac_mask: str=None):  # noqa: E501
        """GetConfigDestinationMacMask - a model defined in Swagger

        :param destination_mac_mask: The destination_mac_mask of this GetConfigDestinationMacMask.  # noqa: E501
        :type destination_mac_mask: str
        """
        self.swagger_types = {
            'destination_mac_mask': str
        }

        self.attribute_map = {
            'destination_mac_mask': 'destination-mac-mask'
        }

        self._destination_mac_mask = destination_mac_mask

    @classmethod
    def from_dict(cls, dikt) -> 'GetConfigDestinationMacMask':
        """Returns the dict as a model

        :param dikt: A dict.
        :type: dict
        :return: The get_config_destination_mac_mask of this GetConfigDestinationMacMask.  # noqa: E501
        :rtype: GetConfigDestinationMacMask
        """
        return util.deserialize_model(dikt, cls)

    @property
    def destination_mac_mask(self) -> str:
        """Gets the destination_mac_mask of this GetConfigDestinationMacMask.


        :return: The destination_mac_mask of this GetConfigDestinationMacMask.
        :rtype: str
        """
        return self._destination_mac_mask

    @destination_mac_mask.setter
    def destination_mac_mask(self, destination_mac_mask: str):
        """Sets the destination_mac_mask of this GetConfigDestinationMacMask.


        :param destination_mac_mask: The destination_mac_mask of this GetConfigDestinationMacMask.
        :type destination_mac_mask: str
        """

        self._destination_mac_mask = destination_mac_mask
# coding: utf-8

from __future__ import absolute_import
from datetime import date, datetime  # noqa: F401

from typing import List, Dict  # noqa: F401

from swagger_server.models.base_model_ import Model
from swagger_server.models.config_source_mac_mask import ConfigSourceMacMask  # noqa: F401,E501
from swagger_server import util


class PatchConfigSourceMacMask(Model):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    def __init__(self, source_mac_mask: str=None):  # noqa: E501
        """PatchConfigSourceMacMask - a model defined in Swagger

        :param source_mac_mask: The source_mac_mask of this PatchConfigSourceMacMask.  # noqa: E501
        :type source_mac_mask: str
        """
        self.swagger_types = {
            'source_mac_mask': str
        }

        self.attribute_map = {
            'source_mac_mask': 'source-mac-mask'
        }

        self._source_mac_mask = source_mac_mask

    @classmethod
    def from_dict(cls, dikt) -> 'PatchConfigSourceMacMask':
        """Returns the dict as a model

        :param dikt: A dict.
        :type: dict
        :return: The patch_config_source_mac_mask of this PatchConfigSourceMacMask.  # noqa: E501
        :rtype: PatchConfigSourceMacMask
        """
        return util.deserialize_model(dikt, cls)

    @property
    def source_mac_mask(self) -> str:
        """Gets the source_mac_mask of this PatchConfigSourceMacMask.


        :return: The source_mac_mask of this PatchConfigSourceMacMask.
        :rtype: str
        """
        return self._source_mac_mask

    @source_mac_mask.setter
    def source_mac_mask(self, source_mac_mask: str):
        """Sets the source_mac_mask of this PatchConfigSourceMacMask.


        :param source_mac_mask: The source_mac_mask of this PatchConfigSourceMacMask.
        :type source_mac_mask: str
        """

        self._source_mac_mask = source_mac_mask
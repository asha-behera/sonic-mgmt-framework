# coding: utf-8

from __future__ import absolute_import
from datetime import date, datetime  # noqa: F401

from typing import List, Dict  # noqa: F401

from swagger_server.models.base_model_ import Model
from swagger_server.models.config_hop_limit import ConfigHopLimit  # noqa: F401,E501
from swagger_server import util


class PatchConfigHopLimit(Model):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    def __init__(self, hop_limit: int=None):  # noqa: E501
        """PatchConfigHopLimit - a model defined in Swagger

        :param hop_limit: The hop_limit of this PatchConfigHopLimit.  # noqa: E501
        :type hop_limit: int
        """
        self.swagger_types = {
            'hop_limit': int
        }

        self.attribute_map = {
            'hop_limit': 'hop-limit'
        }

        self._hop_limit = hop_limit

    @classmethod
    def from_dict(cls, dikt) -> 'PatchConfigHopLimit':
        """Returns the dict as a model

        :param dikt: A dict.
        :type: dict
        :return: The patch_config_hop_limit of this PatchConfigHopLimit.  # noqa: E501
        :rtype: PatchConfigHopLimit
        """
        return util.deserialize_model(dikt, cls)

    @property
    def hop_limit(self) -> int:
        """Gets the hop_limit of this PatchConfigHopLimit.


        :return: The hop_limit of this PatchConfigHopLimit.
        :rtype: int
        """
        return self._hop_limit

    @hop_limit.setter
    def hop_limit(self, hop_limit: int):
        """Sets the hop_limit of this PatchConfigHopLimit.


        :param hop_limit: The hop_limit of this PatchConfigHopLimit.
        :type hop_limit: int
        """

        self._hop_limit = hop_limit
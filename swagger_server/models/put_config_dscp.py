# coding: utf-8

from __future__ import absolute_import
from datetime import date, datetime  # noqa: F401

from typing import List, Dict  # noqa: F401

from swagger_server.models.base_model_ import Model
from swagger_server.models.config_dscp import ConfigDscp  # noqa: F401,E501
from swagger_server import util


class PutConfigDscp(Model):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    def __init__(self, dscp: int=None):  # noqa: E501
        """PutConfigDscp - a model defined in Swagger

        :param dscp: The dscp of this PutConfigDscp.  # noqa: E501
        :type dscp: int
        """
        self.swagger_types = {
            'dscp': int
        }

        self.attribute_map = {
            'dscp': 'dscp'
        }

        self._dscp = dscp

    @classmethod
    def from_dict(cls, dikt) -> 'PutConfigDscp':
        """Returns the dict as a model

        :param dikt: A dict.
        :type: dict
        :return: The put_config_dscp of this PutConfigDscp.  # noqa: E501
        :rtype: PutConfigDscp
        """
        return util.deserialize_model(dikt, cls)

    @property
    def dscp(self) -> int:
        """Gets the dscp of this PutConfigDscp.


        :return: The dscp of this PutConfigDscp.
        :rtype: int
        """
        return self._dscp

    @dscp.setter
    def dscp(self, dscp: int):
        """Sets the dscp of this PutConfigDscp.


        :param dscp: The dscp of this PutConfigDscp.
        :type dscp: int
        """

        self._dscp = dscp
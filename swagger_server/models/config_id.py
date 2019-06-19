# coding: utf-8

from __future__ import absolute_import
from datetime import date, datetime  # noqa: F401

from typing import List, Dict  # noqa: F401

from swagger_server.models.base_model_ import Model
from swagger_server import util


class ConfigId(Model):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    def __init__(self, id: str=None):  # noqa: E501
        """ConfigId - a model defined in Swagger

        :param id: The id of this ConfigId.  # noqa: E501
        :type id: str
        """
        self.swagger_types = {
            'id': str
        }

        self.attribute_map = {
            'id': 'id'
        }

        self._id = id

    @classmethod
    def from_dict(cls, dikt) -> 'ConfigId':
        """Returns the dict as a model

        :param dikt: A dict.
        :type: dict
        :return: The config_id of this ConfigId.  # noqa: E501
        :rtype: ConfigId
        """
        return util.deserialize_model(dikt, cls)

    @property
    def id(self) -> str:
        """Gets the id of this ConfigId.


        :return: The id of this ConfigId.
        :rtype: str
        """
        return self._id

    @id.setter
    def id(self, id: str):
        """Sets the id of this ConfigId.


        :param id: The id of this ConfigId.
        :type id: str
        """

        self._id = id
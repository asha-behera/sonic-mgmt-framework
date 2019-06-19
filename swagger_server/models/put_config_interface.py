# coding: utf-8

from __future__ import absolute_import
from datetime import date, datetime  # noqa: F401

from typing import List, Dict  # noqa: F401

from swagger_server.models.base_model_ import Model
from swagger_server.models.config_interface import ConfigInterface  # noqa: F401,E501
from swagger_server import util


class PutConfigInterface(Model):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    def __init__(self, interface: str=None):  # noqa: E501
        """PutConfigInterface - a model defined in Swagger

        :param interface: The interface of this PutConfigInterface.  # noqa: E501
        :type interface: str
        """
        self.swagger_types = {
            'interface': str
        }

        self.attribute_map = {
            'interface': 'interface'
        }

        self._interface = interface

    @classmethod
    def from_dict(cls, dikt) -> 'PutConfigInterface':
        """Returns the dict as a model

        :param dikt: A dict.
        :type: dict
        :return: The put_config_interface of this PutConfigInterface.  # noqa: E501
        :rtype: PutConfigInterface
        """
        return util.deserialize_model(dikt, cls)

    @property
    def interface(self) -> str:
        """Gets the interface of this PutConfigInterface.


        :return: The interface of this PutConfigInterface.
        :rtype: str
        """
        return self._interface

    @interface.setter
    def interface(self, interface: str):
        """Sets the interface of this PutConfigInterface.


        :param interface: The interface of this PutConfigInterface.
        :type interface: str
        """

        self._interface = interface
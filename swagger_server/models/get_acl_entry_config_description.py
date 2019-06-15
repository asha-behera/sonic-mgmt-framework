# coding: utf-8

from __future__ import absolute_import
from datetime import date, datetime  # noqa: F401

from typing import List, Dict  # noqa: F401

from swagger_server.models.base_model_ import Model
from swagger_server import util


class GetAclEntryConfigDescription(Model):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    def __init__(self, description: str=None):  # noqa: E501
        """GetAclEntryConfigDescription - a model defined in Swagger

        :param description: The description of this GetAclEntryConfigDescription.  # noqa: E501
        :type description: str
        """
        self.swagger_types = {
            'description': str
        }

        self.attribute_map = {
            'description': 'description'
        }

        self._description = description

    @classmethod
    def from_dict(cls, dikt) -> 'GetAclEntryConfigDescription':
        """Returns the dict as a model

        :param dikt: A dict.
        :type: dict
        :return: The get_acl_entry_config_description of this GetAclEntryConfigDescription.  # noqa: E501
        :rtype: GetAclEntryConfigDescription
        """
        return util.deserialize_model(dikt, cls)

    @property
    def description(self) -> str:
        """Gets the description of this GetAclEntryConfigDescription.


        :return: The description of this GetAclEntryConfigDescription.
        :rtype: str
        """
        return self._description

    @description.setter
    def description(self, description: str):
        """Sets the description of this GetAclEntryConfigDescription.


        :param description: The description of this GetAclEntryConfigDescription.
        :type description: str
        """

        self._description = description
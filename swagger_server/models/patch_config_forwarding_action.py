# coding: utf-8

from __future__ import absolute_import
from datetime import date, datetime  # noqa: F401

from typing import List, Dict  # noqa: F401

from swagger_server.models.base_model_ import Model
from swagger_server.models.config_forwarding_action import ConfigForwardingAction  # noqa: F401,E501
from swagger_server import util


class PatchConfigForwardingAction(Model):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    def __init__(self, forwarding_action: str=None):  # noqa: E501
        """PatchConfigForwardingAction - a model defined in Swagger

        :param forwarding_action: The forwarding_action of this PatchConfigForwardingAction.  # noqa: E501
        :type forwarding_action: str
        """
        self.swagger_types = {
            'forwarding_action': str
        }

        self.attribute_map = {
            'forwarding_action': 'forwarding-action'
        }

        self._forwarding_action = forwarding_action

    @classmethod
    def from_dict(cls, dikt) -> 'PatchConfigForwardingAction':
        """Returns the dict as a model

        :param dikt: A dict.
        :type: dict
        :return: The patch_config_forwarding_action of this PatchConfigForwardingAction.  # noqa: E501
        :rtype: PatchConfigForwardingAction
        """
        return util.deserialize_model(dikt, cls)

    @property
    def forwarding_action(self) -> str:
        """Gets the forwarding_action of this PatchConfigForwardingAction.


        :return: The forwarding_action of this PatchConfigForwardingAction.
        :rtype: str
        """
        return self._forwarding_action

    @forwarding_action.setter
    def forwarding_action(self, forwarding_action: str):
        """Sets the forwarding_action of this PatchConfigForwardingAction.


        :param forwarding_action: The forwarding_action of this PatchConfigForwardingAction.
        :type forwarding_action: str
        """

        self._forwarding_action = forwarding_action
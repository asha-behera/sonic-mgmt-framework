# coding: utf-8

from __future__ import absolute_import
from datetime import date, datetime  # noqa: F401

from typing import List, Dict  # noqa: F401

from swagger_server.models.base_model_ import Model
from swagger_server.models.acl_entry_actions import AclEntryActions  # noqa: F401,E501
from swagger_server.models.acl_openconfigaclacl_aclsets_aclentries_actions import AclOpenconfigaclaclAclsetsAclentriesActions  # noqa: F401,E501
from swagger_server import util


class PatchAclEntryActions(Model):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    def __init__(self, openconfig_aclactions: AclOpenconfigaclaclAclsetsAclentriesActions=None):  # noqa: E501
        """PatchAclEntryActions - a model defined in Swagger

        :param openconfig_aclactions: The openconfig_aclactions of this PatchAclEntryActions.  # noqa: E501
        :type openconfig_aclactions: AclOpenconfigaclaclAclsetsAclentriesActions
        """
        self.swagger_types = {
            'openconfig_aclactions': AclOpenconfigaclaclAclsetsAclentriesActions
        }

        self.attribute_map = {
            'openconfig_aclactions': 'openconfig-acl:actions'
        }

        self._openconfig_aclactions = openconfig_aclactions

    @classmethod
    def from_dict(cls, dikt) -> 'PatchAclEntryActions':
        """Returns the dict as a model

        :param dikt: A dict.
        :type: dict
        :return: The patch_acl_entry_actions of this PatchAclEntryActions.  # noqa: E501
        :rtype: PatchAclEntryActions
        """
        return util.deserialize_model(dikt, cls)

    @property
    def openconfig_aclactions(self) -> AclOpenconfigaclaclAclsetsAclentriesActions:
        """Gets the openconfig_aclactions of this PatchAclEntryActions.


        :return: The openconfig_aclactions of this PatchAclEntryActions.
        :rtype: AclOpenconfigaclaclAclsetsAclentriesActions
        """
        return self._openconfig_aclactions

    @openconfig_aclactions.setter
    def openconfig_aclactions(self, openconfig_aclactions: AclOpenconfigaclaclAclsetsAclentriesActions):
        """Sets the openconfig_aclactions of this PatchAclEntryActions.


        :param openconfig_aclactions: The openconfig_aclactions of this PatchAclEntryActions.
        :type openconfig_aclactions: AclOpenconfigaclaclAclsetsAclentriesActions
        """

        self._openconfig_aclactions = openconfig_aclactions
# coding: utf-8

from __future__ import absolute_import
from datetime import date, datetime  # noqa: F401

from typing import List, Dict  # noqa: F401

from swagger_server.models.base_model_ import Model
from swagger_server.models.acl_openconfigaclacl_aclsets_aclentries_l2 import AclOpenconfigaclaclAclsetsAclentriesL2  # noqa: F401,E501
from swagger_server import util


class AclEntryL2(Model):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    def __init__(self, openconfig_acll2: AclOpenconfigaclaclAclsetsAclentriesL2=None):  # noqa: E501
        """AclEntryL2 - a model defined in Swagger

        :param openconfig_acll2: The openconfig_acll2 of this AclEntryL2.  # noqa: E501
        :type openconfig_acll2: AclOpenconfigaclaclAclsetsAclentriesL2
        """
        self.swagger_types = {
            'openconfig_acll2': AclOpenconfigaclaclAclsetsAclentriesL2
        }

        self.attribute_map = {
            'openconfig_acll2': 'openconfig-acl:l2'
        }

        self._openconfig_acll2 = openconfig_acll2

    @classmethod
    def from_dict(cls, dikt) -> 'AclEntryL2':
        """Returns the dict as a model

        :param dikt: A dict.
        :type: dict
        :return: The acl_entry_l2 of this AclEntryL2.  # noqa: E501
        :rtype: AclEntryL2
        """
        return util.deserialize_model(dikt, cls)

    @property
    def openconfig_acll2(self) -> AclOpenconfigaclaclAclsetsAclentriesL2:
        """Gets the openconfig_acll2 of this AclEntryL2.


        :return: The openconfig_acll2 of this AclEntryL2.
        :rtype: AclOpenconfigaclaclAclsetsAclentriesL2
        """
        return self._openconfig_acll2

    @openconfig_acll2.setter
    def openconfig_acll2(self, openconfig_acll2: AclOpenconfigaclaclAclsetsAclentriesL2):
        """Sets the openconfig_acll2 of this AclEntryL2.


        :param openconfig_acll2: The openconfig_acll2 of this AclEntryL2.
        :type openconfig_acll2: AclOpenconfigaclaclAclsetsAclentriesL2
        """

        self._openconfig_acll2 = openconfig_acll2
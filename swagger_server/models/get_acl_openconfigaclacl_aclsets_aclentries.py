# coding: utf-8

from __future__ import absolute_import
from datetime import date, datetime  # noqa: F401

from typing import List, Dict  # noqa: F401

from swagger_server.models.base_model_ import Model
from swagger_server.models.get_acl_openconfigaclacl_aclsets_aclentries_aclentry import GetAclOpenconfigaclaclAclsetsAclentriesAclentry  # noqa: F401,E501
from swagger_server import util


class GetAclOpenconfigaclaclAclsetsAclentries(Model):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    def __init__(self, acl_entry: List[GetAclOpenconfigaclaclAclsetsAclentriesAclentry]=None):  # noqa: E501
        """GetAclOpenconfigaclaclAclsetsAclentries - a model defined in Swagger

        :param acl_entry: The acl_entry of this GetAclOpenconfigaclaclAclsetsAclentries.  # noqa: E501
        :type acl_entry: List[GetAclOpenconfigaclaclAclsetsAclentriesAclentry]
        """
        self.swagger_types = {
            'acl_entry': List[GetAclOpenconfigaclaclAclsetsAclentriesAclentry]
        }

        self.attribute_map = {
            'acl_entry': 'acl-entry'
        }

        self._acl_entry = acl_entry

    @classmethod
    def from_dict(cls, dikt) -> 'GetAclOpenconfigaclaclAclsetsAclentries':
        """Returns the dict as a model

        :param dikt: A dict.
        :type: dict
        :return: The get_acl_openconfigaclacl_aclsets_aclentries of this GetAclOpenconfigaclaclAclsetsAclentries.  # noqa: E501
        :rtype: GetAclOpenconfigaclaclAclsetsAclentries
        """
        return util.deserialize_model(dikt, cls)

    @property
    def acl_entry(self) -> List[GetAclOpenconfigaclaclAclsetsAclentriesAclentry]:
        """Gets the acl_entry of this GetAclOpenconfigaclaclAclsetsAclentries.


        :return: The acl_entry of this GetAclOpenconfigaclaclAclsetsAclentries.
        :rtype: List[GetAclOpenconfigaclaclAclsetsAclentriesAclentry]
        """
        return self._acl_entry

    @acl_entry.setter
    def acl_entry(self, acl_entry: List[GetAclOpenconfigaclaclAclsetsAclentriesAclentry]):
        """Sets the acl_entry of this GetAclOpenconfigaclaclAclsetsAclentries.


        :param acl_entry: The acl_entry of this GetAclOpenconfigaclaclAclsetsAclentries.
        :type acl_entry: List[GetAclOpenconfigaclaclAclsetsAclentriesAclentry]
        """

        self._acl_entry = acl_entry
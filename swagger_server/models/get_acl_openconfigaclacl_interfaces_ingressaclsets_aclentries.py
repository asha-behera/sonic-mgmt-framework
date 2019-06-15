# coding: utf-8

from __future__ import absolute_import
from datetime import date, datetime  # noqa: F401

from typing import List, Dict  # noqa: F401

from swagger_server.models.base_model_ import Model
from swagger_server.models.get_acl_openconfigaclacl_interfaces_ingressaclsets_aclentries_aclentry import GetAclOpenconfigaclaclInterfacesIngressaclsetsAclentriesAclentry  # noqa: F401,E501
from swagger_server import util


class GetAclOpenconfigaclaclInterfacesIngressaclsetsAclentries(Model):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    def __init__(self, acl_entry: List[GetAclOpenconfigaclaclInterfacesIngressaclsetsAclentriesAclentry]=None):  # noqa: E501
        """GetAclOpenconfigaclaclInterfacesIngressaclsetsAclentries - a model defined in Swagger

        :param acl_entry: The acl_entry of this GetAclOpenconfigaclaclInterfacesIngressaclsetsAclentries.  # noqa: E501
        :type acl_entry: List[GetAclOpenconfigaclaclInterfacesIngressaclsetsAclentriesAclentry]
        """
        self.swagger_types = {
            'acl_entry': List[GetAclOpenconfigaclaclInterfacesIngressaclsetsAclentriesAclentry]
        }

        self.attribute_map = {
            'acl_entry': 'acl-entry'
        }

        self._acl_entry = acl_entry

    @classmethod
    def from_dict(cls, dikt) -> 'GetAclOpenconfigaclaclInterfacesIngressaclsetsAclentries':
        """Returns the dict as a model

        :param dikt: A dict.
        :type: dict
        :return: The get_acl_openconfigaclacl_interfaces_ingressaclsets_aclentries of this GetAclOpenconfigaclaclInterfacesIngressaclsetsAclentries.  # noqa: E501
        :rtype: GetAclOpenconfigaclaclInterfacesIngressaclsetsAclentries
        """
        return util.deserialize_model(dikt, cls)

    @property
    def acl_entry(self) -> List[GetAclOpenconfigaclaclInterfacesIngressaclsetsAclentriesAclentry]:
        """Gets the acl_entry of this GetAclOpenconfigaclaclInterfacesIngressaclsetsAclentries.


        :return: The acl_entry of this GetAclOpenconfigaclaclInterfacesIngressaclsetsAclentries.
        :rtype: List[GetAclOpenconfigaclaclInterfacesIngressaclsetsAclentriesAclentry]
        """
        return self._acl_entry

    @acl_entry.setter
    def acl_entry(self, acl_entry: List[GetAclOpenconfigaclaclInterfacesIngressaclsetsAclentriesAclentry]):
        """Sets the acl_entry of this GetAclOpenconfigaclaclInterfacesIngressaclsetsAclentries.


        :param acl_entry: The acl_entry of this GetAclOpenconfigaclaclInterfacesIngressaclsetsAclentries.
        :type acl_entry: List[GetAclOpenconfigaclaclInterfacesIngressaclsetsAclentriesAclentry]
        """

        self._acl_entry = acl_entry
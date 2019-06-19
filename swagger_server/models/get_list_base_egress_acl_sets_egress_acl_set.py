# coding: utf-8

from __future__ import absolute_import
from datetime import date, datetime  # noqa: F401

from typing import List, Dict  # noqa: F401

from swagger_server.models.base_model_ import Model
from swagger_server.models.get_acl_openconfigaclacl_interfaces_ingressaclsets_ingressaclset import GetAclOpenconfigaclaclInterfacesIngressaclsetsIngressaclset  # noqa: F401,E501
from swagger_server import util


class GetListBaseEgressAclSetsEgressAclSet(Model):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    def __init__(self, openconfig_aclegress_acl_set: List[GetAclOpenconfigaclaclInterfacesIngressaclsetsIngressaclset]=None):  # noqa: E501
        """GetListBaseEgressAclSetsEgressAclSet - a model defined in Swagger

        :param openconfig_aclegress_acl_set: The openconfig_aclegress_acl_set of this GetListBaseEgressAclSetsEgressAclSet.  # noqa: E501
        :type openconfig_aclegress_acl_set: List[GetAclOpenconfigaclaclInterfacesIngressaclsetsIngressaclset]
        """
        self.swagger_types = {
            'openconfig_aclegress_acl_set': List[GetAclOpenconfigaclaclInterfacesIngressaclsetsIngressaclset]
        }

        self.attribute_map = {
            'openconfig_aclegress_acl_set': 'openconfig-acl:egress-acl-set'
        }

        self._openconfig_aclegress_acl_set = openconfig_aclegress_acl_set

    @classmethod
    def from_dict(cls, dikt) -> 'GetListBaseEgressAclSetsEgressAclSet':
        """Returns the dict as a model

        :param dikt: A dict.
        :type: dict
        :return: The get_list_base_egress_acl_sets_egress_acl_set of this GetListBaseEgressAclSetsEgressAclSet.  # noqa: E501
        :rtype: GetListBaseEgressAclSetsEgressAclSet
        """
        return util.deserialize_model(dikt, cls)

    @property
    def openconfig_aclegress_acl_set(self) -> List[GetAclOpenconfigaclaclInterfacesIngressaclsetsIngressaclset]:
        """Gets the openconfig_aclegress_acl_set of this GetListBaseEgressAclSetsEgressAclSet.


        :return: The openconfig_aclegress_acl_set of this GetListBaseEgressAclSetsEgressAclSet.
        :rtype: List[GetAclOpenconfigaclaclInterfacesIngressaclsetsIngressaclset]
        """
        return self._openconfig_aclegress_acl_set

    @openconfig_aclegress_acl_set.setter
    def openconfig_aclegress_acl_set(self, openconfig_aclegress_acl_set: List[GetAclOpenconfigaclaclInterfacesIngressaclsetsIngressaclset]):
        """Sets the openconfig_aclegress_acl_set of this GetListBaseEgressAclSetsEgressAclSet.


        :param openconfig_aclegress_acl_set: The openconfig_aclegress_acl_set of this GetListBaseEgressAclSetsEgressAclSet.
        :type openconfig_aclegress_acl_set: List[GetAclOpenconfigaclaclInterfacesIngressaclsetsIngressaclset]
        """

        self._openconfig_aclegress_acl_set = openconfig_aclegress_acl_set
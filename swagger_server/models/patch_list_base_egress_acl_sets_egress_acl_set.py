# coding: utf-8

from __future__ import absolute_import
from datetime import date, datetime  # noqa: F401

from typing import List, Dict  # noqa: F401

from swagger_server.models.base_model_ import Model
from swagger_server.models.acl_openconfigaclacl_interfaces_ingressaclsets_ingressaclset import AclOpenconfigaclaclInterfacesIngressaclsetsIngressaclset  # noqa: F401,E501
from swagger_server.models.list_base_egress_acl_sets_egress_acl_set import ListBaseEgressAclSetsEgressAclSet  # noqa: F401,E501
from swagger_server import util


class PatchListBaseEgressAclSetsEgressAclSet(Model):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    def __init__(self, openconfig_aclegress_acl_set: List[AclOpenconfigaclaclInterfacesIngressaclsetsIngressaclset]=None):  # noqa: E501
        """PatchListBaseEgressAclSetsEgressAclSet - a model defined in Swagger

        :param openconfig_aclegress_acl_set: The openconfig_aclegress_acl_set of this PatchListBaseEgressAclSetsEgressAclSet.  # noqa: E501
        :type openconfig_aclegress_acl_set: List[AclOpenconfigaclaclInterfacesIngressaclsetsIngressaclset]
        """
        self.swagger_types = {
            'openconfig_aclegress_acl_set': List[AclOpenconfigaclaclInterfacesIngressaclsetsIngressaclset]
        }

        self.attribute_map = {
            'openconfig_aclegress_acl_set': 'openconfig-acl:egress-acl-set'
        }

        self._openconfig_aclegress_acl_set = openconfig_aclegress_acl_set

    @classmethod
    def from_dict(cls, dikt) -> 'PatchListBaseEgressAclSetsEgressAclSet':
        """Returns the dict as a model

        :param dikt: A dict.
        :type: dict
        :return: The patch_list_base_egress_acl_sets_egress_acl_set of this PatchListBaseEgressAclSetsEgressAclSet.  # noqa: E501
        :rtype: PatchListBaseEgressAclSetsEgressAclSet
        """
        return util.deserialize_model(dikt, cls)

    @property
    def openconfig_aclegress_acl_set(self) -> List[AclOpenconfigaclaclInterfacesIngressaclsetsIngressaclset]:
        """Gets the openconfig_aclegress_acl_set of this PatchListBaseEgressAclSetsEgressAclSet.


        :return: The openconfig_aclegress_acl_set of this PatchListBaseEgressAclSetsEgressAclSet.
        :rtype: List[AclOpenconfigaclaclInterfacesIngressaclsetsIngressaclset]
        """
        return self._openconfig_aclegress_acl_set

    @openconfig_aclegress_acl_set.setter
    def openconfig_aclegress_acl_set(self, openconfig_aclegress_acl_set: List[AclOpenconfigaclaclInterfacesIngressaclsetsIngressaclset]):
        """Sets the openconfig_aclegress_acl_set of this PatchListBaseEgressAclSetsEgressAclSet.


        :param openconfig_aclegress_acl_set: The openconfig_aclegress_acl_set of this PatchListBaseEgressAclSetsEgressAclSet.
        :type openconfig_aclegress_acl_set: List[AclOpenconfigaclaclInterfacesIngressaclsetsIngressaclset]
        """

        self._openconfig_aclegress_acl_set = openconfig_aclegress_acl_set
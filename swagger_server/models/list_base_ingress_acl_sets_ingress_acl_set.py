# coding: utf-8

from __future__ import absolute_import
from datetime import date, datetime  # noqa: F401

from typing import List, Dict  # noqa: F401

from swagger_server.models.base_model_ import Model
from swagger_server.models.acl_openconfigaclacl_interfaces_ingressaclsets_ingressaclset import AclOpenconfigaclaclInterfacesIngressaclsetsIngressaclset  # noqa: F401,E501
from swagger_server import util


class ListBaseIngressAclSetsIngressAclSet(Model):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    def __init__(self, openconfig_aclingress_acl_set: List[AclOpenconfigaclaclInterfacesIngressaclsetsIngressaclset]=None):  # noqa: E501
        """ListBaseIngressAclSetsIngressAclSet - a model defined in Swagger

        :param openconfig_aclingress_acl_set: The openconfig_aclingress_acl_set of this ListBaseIngressAclSetsIngressAclSet.  # noqa: E501
        :type openconfig_aclingress_acl_set: List[AclOpenconfigaclaclInterfacesIngressaclsetsIngressaclset]
        """
        self.swagger_types = {
            'openconfig_aclingress_acl_set': List[AclOpenconfigaclaclInterfacesIngressaclsetsIngressaclset]
        }

        self.attribute_map = {
            'openconfig_aclingress_acl_set': 'openconfig-acl:ingress-acl-set'
        }

        self._openconfig_aclingress_acl_set = openconfig_aclingress_acl_set

    @classmethod
    def from_dict(cls, dikt) -> 'ListBaseIngressAclSetsIngressAclSet':
        """Returns the dict as a model

        :param dikt: A dict.
        :type: dict
        :return: The list_base_ingress_acl_sets_ingress_acl_set of this ListBaseIngressAclSetsIngressAclSet.  # noqa: E501
        :rtype: ListBaseIngressAclSetsIngressAclSet
        """
        return util.deserialize_model(dikt, cls)

    @property
    def openconfig_aclingress_acl_set(self) -> List[AclOpenconfigaclaclInterfacesIngressaclsetsIngressaclset]:
        """Gets the openconfig_aclingress_acl_set of this ListBaseIngressAclSetsIngressAclSet.


        :return: The openconfig_aclingress_acl_set of this ListBaseIngressAclSetsIngressAclSet.
        :rtype: List[AclOpenconfigaclaclInterfacesIngressaclsetsIngressaclset]
        """
        return self._openconfig_aclingress_acl_set

    @openconfig_aclingress_acl_set.setter
    def openconfig_aclingress_acl_set(self, openconfig_aclingress_acl_set: List[AclOpenconfigaclaclInterfacesIngressaclsetsIngressaclset]):
        """Sets the openconfig_aclingress_acl_set of this ListBaseIngressAclSetsIngressAclSet.


        :param openconfig_aclingress_acl_set: The openconfig_aclingress_acl_set of this ListBaseIngressAclSetsIngressAclSet.
        :type openconfig_aclingress_acl_set: List[AclOpenconfigaclaclInterfacesIngressaclsetsIngressaclset]
        """

        self._openconfig_aclingress_acl_set = openconfig_aclingress_acl_set
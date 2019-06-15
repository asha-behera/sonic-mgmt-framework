# coding: utf-8

from __future__ import absolute_import
from datetime import date, datetime  # noqa: F401

from typing import List, Dict  # noqa: F401

from swagger_server.models.base_model_ import Model
from swagger_server.models.get_acl_openconfigaclacl_interfaces_ingressaclsets_ingressaclset import GetAclOpenconfigaclaclInterfacesIngressaclsetsIngressaclset  # noqa: F401,E501
from swagger_server import util


class GetAclOpenconfigaclaclInterfacesIngressaclsets(Model):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    def __init__(self, ingress_acl_set: List[GetAclOpenconfigaclaclInterfacesIngressaclsetsIngressaclset]=None):  # noqa: E501
        """GetAclOpenconfigaclaclInterfacesIngressaclsets - a model defined in Swagger

        :param ingress_acl_set: The ingress_acl_set of this GetAclOpenconfigaclaclInterfacesIngressaclsets.  # noqa: E501
        :type ingress_acl_set: List[GetAclOpenconfigaclaclInterfacesIngressaclsetsIngressaclset]
        """
        self.swagger_types = {
            'ingress_acl_set': List[GetAclOpenconfigaclaclInterfacesIngressaclsetsIngressaclset]
        }

        self.attribute_map = {
            'ingress_acl_set': 'ingress-acl-set'
        }

        self._ingress_acl_set = ingress_acl_set

    @classmethod
    def from_dict(cls, dikt) -> 'GetAclOpenconfigaclaclInterfacesIngressaclsets':
        """Returns the dict as a model

        :param dikt: A dict.
        :type: dict
        :return: The get_acl_openconfigaclacl_interfaces_ingressaclsets of this GetAclOpenconfigaclaclInterfacesIngressaclsets.  # noqa: E501
        :rtype: GetAclOpenconfigaclaclInterfacesIngressaclsets
        """
        return util.deserialize_model(dikt, cls)

    @property
    def ingress_acl_set(self) -> List[GetAclOpenconfigaclaclInterfacesIngressaclsetsIngressaclset]:
        """Gets the ingress_acl_set of this GetAclOpenconfigaclaclInterfacesIngressaclsets.


        :return: The ingress_acl_set of this GetAclOpenconfigaclaclInterfacesIngressaclsets.
        :rtype: List[GetAclOpenconfigaclaclInterfacesIngressaclsetsIngressaclset]
        """
        return self._ingress_acl_set

    @ingress_acl_set.setter
    def ingress_acl_set(self, ingress_acl_set: List[GetAclOpenconfigaclaclInterfacesIngressaclsetsIngressaclset]):
        """Sets the ingress_acl_set of this GetAclOpenconfigaclaclInterfacesIngressaclsets.


        :param ingress_acl_set: The ingress_acl_set of this GetAclOpenconfigaclaclInterfacesIngressaclsets.
        :type ingress_acl_set: List[GetAclOpenconfigaclaclInterfacesIngressaclsetsIngressaclset]
        """

        self._ingress_acl_set = ingress_acl_set
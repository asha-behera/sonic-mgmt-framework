# coding: utf-8

from __future__ import absolute_import
from datetime import date, datetime  # noqa: F401

from typing import List, Dict  # noqa: F401

from swagger_server.models.base_model_ import Model
from swagger_server.models.acl_openconfigaclacl_interfaces_ingressaclsets_ingressaclset import AclOpenconfigaclaclInterfacesIngressaclsetsIngressaclset  # noqa: F401,E501
from swagger_server import util


class AclOpenconfigaclaclInterfacesEgressaclsets(Model):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    def __init__(self, egress_acl_set: List[AclOpenconfigaclaclInterfacesIngressaclsetsIngressaclset]=None):  # noqa: E501
        """AclOpenconfigaclaclInterfacesEgressaclsets - a model defined in Swagger

        :param egress_acl_set: The egress_acl_set of this AclOpenconfigaclaclInterfacesEgressaclsets.  # noqa: E501
        :type egress_acl_set: List[AclOpenconfigaclaclInterfacesIngressaclsetsIngressaclset]
        """
        self.swagger_types = {
            'egress_acl_set': List[AclOpenconfigaclaclInterfacesIngressaclsetsIngressaclset]
        }

        self.attribute_map = {
            'egress_acl_set': 'egress-acl-set'
        }

        self._egress_acl_set = egress_acl_set

    @classmethod
    def from_dict(cls, dikt) -> 'AclOpenconfigaclaclInterfacesEgressaclsets':
        """Returns the dict as a model

        :param dikt: A dict.
        :type: dict
        :return: The acl_openconfigaclacl_interfaces_egressaclsets of this AclOpenconfigaclaclInterfacesEgressaclsets.  # noqa: E501
        :rtype: AclOpenconfigaclaclInterfacesEgressaclsets
        """
        return util.deserialize_model(dikt, cls)

    @property
    def egress_acl_set(self) -> List[AclOpenconfigaclaclInterfacesIngressaclsetsIngressaclset]:
        """Gets the egress_acl_set of this AclOpenconfigaclaclInterfacesEgressaclsets.


        :return: The egress_acl_set of this AclOpenconfigaclaclInterfacesEgressaclsets.
        :rtype: List[AclOpenconfigaclaclInterfacesIngressaclsetsIngressaclset]
        """
        return self._egress_acl_set

    @egress_acl_set.setter
    def egress_acl_set(self, egress_acl_set: List[AclOpenconfigaclaclInterfacesIngressaclsetsIngressaclset]):
        """Sets the egress_acl_set of this AclOpenconfigaclaclInterfacesEgressaclsets.


        :param egress_acl_set: The egress_acl_set of this AclOpenconfigaclaclInterfacesEgressaclsets.
        :type egress_acl_set: List[AclOpenconfigaclaclInterfacesIngressaclsetsIngressaclset]
        """

        self._egress_acl_set = egress_acl_set
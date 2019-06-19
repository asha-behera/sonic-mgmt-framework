# coding: utf-8

from __future__ import absolute_import
from datetime import date, datetime  # noqa: F401

from typing import List, Dict  # noqa: F401

from swagger_server.models.base_model_ import Model
from swagger_server.models.acl_openconfigaclacl_aclsets import AclOpenconfigaclaclAclsets  # noqa: F401,E501
from swagger_server import util


class AclAclSets(Model):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    def __init__(self, openconfig_aclacl_sets: AclOpenconfigaclaclAclsets=None):  # noqa: E501
        """AclAclSets - a model defined in Swagger

        :param openconfig_aclacl_sets: The openconfig_aclacl_sets of this AclAclSets.  # noqa: E501
        :type openconfig_aclacl_sets: AclOpenconfigaclaclAclsets
        """
        self.swagger_types = {
            'openconfig_aclacl_sets': AclOpenconfigaclaclAclsets
        }

        self.attribute_map = {
            'openconfig_aclacl_sets': 'openconfig-acl:acl-sets'
        }

        self._openconfig_aclacl_sets = openconfig_aclacl_sets

    @classmethod
    def from_dict(cls, dikt) -> 'AclAclSets':
        """Returns the dict as a model

        :param dikt: A dict.
        :type: dict
        :return: The acl_acl_sets of this AclAclSets.  # noqa: E501
        :rtype: AclAclSets
        """
        return util.deserialize_model(dikt, cls)

    @property
    def openconfig_aclacl_sets(self) -> AclOpenconfigaclaclAclsets:
        """Gets the openconfig_aclacl_sets of this AclAclSets.


        :return: The openconfig_aclacl_sets of this AclAclSets.
        :rtype: AclOpenconfigaclaclAclsets
        """
        return self._openconfig_aclacl_sets

    @openconfig_aclacl_sets.setter
    def openconfig_aclacl_sets(self, openconfig_aclacl_sets: AclOpenconfigaclaclAclsets):
        """Sets the openconfig_aclacl_sets of this AclAclSets.


        :param openconfig_aclacl_sets: The openconfig_aclacl_sets of this AclAclSets.
        :type openconfig_aclacl_sets: AclOpenconfigaclaclAclsets
        """

        self._openconfig_aclacl_sets = openconfig_aclacl_sets
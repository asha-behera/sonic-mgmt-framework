# coding: utf-8

from __future__ import absolute_import
from datetime import date, datetime  # noqa: F401

from typing import List, Dict  # noqa: F401

from swagger_server.models.base_model_ import Model
from swagger_server.models.acl_entry_config import AclEntryConfig  # noqa: F401,E501
from swagger_server.models.acl_entry_input_interface import AclEntryInputInterface  # noqa: F401,E501
from swagger_server.models.acl_entry_ipv4 import AclEntryIpv4  # noqa: F401,E501
from swagger_server.models.acl_entry_ipv6 import AclEntryIpv6  # noqa: F401,E501
from swagger_server.models.acl_entry_l2 import AclEntryL2  # noqa: F401,E501
from swagger_server.models.acl_entry_transport import AclEntryTransport  # noqa: F401,E501
from swagger_server.models.acl_openconfigaclacl_aclsets_aclentries_actions import AclOpenconfigaclaclAclsetsAclentriesActions  # noqa: F401,E501
from swagger_server.models.acl_openconfigaclacl_aclsets_aclentries_config import AclOpenconfigaclaclAclsetsAclentriesConfig  # noqa: F401,E501
from swagger_server.models.acl_openconfigaclacl_aclsets_aclentries_inputinterface import AclOpenconfigaclaclAclsetsAclentriesInputinterface  # noqa: F401,E501
from swagger_server.models.acl_openconfigaclacl_aclsets_aclentries_ipv4 import AclOpenconfigaclaclAclsetsAclentriesIpv4  # noqa: F401,E501
from swagger_server.models.acl_openconfigaclacl_aclsets_aclentries_ipv6 import AclOpenconfigaclaclAclsetsAclentriesIpv6  # noqa: F401,E501
from swagger_server.models.acl_openconfigaclacl_aclsets_aclentries_l2 import AclOpenconfigaclaclAclsetsAclentriesL2  # noqa: F401,E501
from swagger_server.models.acl_openconfigaclacl_aclsets_aclentries_transport import AclOpenconfigaclaclAclsetsAclentriesTransport  # noqa: F401,E501
from swagger_server import util


class PostAclEntryConfig(Model):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    def __init__(self, openconfig_acll2: AclOpenconfigaclaclAclsetsAclentriesL2=None, openconfig_aclipv4: AclOpenconfigaclaclAclsetsAclentriesIpv4=None, openconfig_aclipv6: AclOpenconfigaclaclAclsetsAclentriesIpv6=None, openconfig_acltransport: AclOpenconfigaclaclAclsetsAclentriesTransport=None, openconfig_aclinput_interface: AclOpenconfigaclaclAclsetsAclentriesInputinterface=None, openconfig_aclconfig: AclOpenconfigaclaclAclsetsAclentriesConfig=None, openconfig_aclactions: AclOpenconfigaclaclAclsetsAclentriesActions=None):  # noqa: E501
        """PostAclEntryConfig - a model defined in Swagger

        :param openconfig_acll2: The openconfig_acll2 of this PostAclEntryConfig.  # noqa: E501
        :type openconfig_acll2: AclOpenconfigaclaclAclsetsAclentriesL2
        :param openconfig_aclipv4: The openconfig_aclipv4 of this PostAclEntryConfig.  # noqa: E501
        :type openconfig_aclipv4: AclOpenconfigaclaclAclsetsAclentriesIpv4
        :param openconfig_aclipv6: The openconfig_aclipv6 of this PostAclEntryConfig.  # noqa: E501
        :type openconfig_aclipv6: AclOpenconfigaclaclAclsetsAclentriesIpv6
        :param openconfig_acltransport: The openconfig_acltransport of this PostAclEntryConfig.  # noqa: E501
        :type openconfig_acltransport: AclOpenconfigaclaclAclsetsAclentriesTransport
        :param openconfig_aclinput_interface: The openconfig_aclinput_interface of this PostAclEntryConfig.  # noqa: E501
        :type openconfig_aclinput_interface: AclOpenconfigaclaclAclsetsAclentriesInputinterface
        :param openconfig_aclconfig: The openconfig_aclconfig of this PostAclEntryConfig.  # noqa: E501
        :type openconfig_aclconfig: AclOpenconfigaclaclAclsetsAclentriesConfig
        :param openconfig_aclactions: The openconfig_aclactions of this PostAclEntryConfig.  # noqa: E501
        :type openconfig_aclactions: AclOpenconfigaclaclAclsetsAclentriesActions
        """
        self.swagger_types = {
            'openconfig_acll2': AclOpenconfigaclaclAclsetsAclentriesL2,
            'openconfig_aclipv4': AclOpenconfigaclaclAclsetsAclentriesIpv4,
            'openconfig_aclipv6': AclOpenconfigaclaclAclsetsAclentriesIpv6,
            'openconfig_acltransport': AclOpenconfigaclaclAclsetsAclentriesTransport,
            'openconfig_aclinput_interface': AclOpenconfigaclaclAclsetsAclentriesInputinterface,
            'openconfig_aclconfig': AclOpenconfigaclaclAclsetsAclentriesConfig,
            'openconfig_aclactions': AclOpenconfigaclaclAclsetsAclentriesActions
        }

        self.attribute_map = {
            'openconfig_acll2': 'openconfig-acl:l2',
            'openconfig_aclipv4': 'openconfig-acl:ipv4',
            'openconfig_aclipv6': 'openconfig-acl:ipv6',
            'openconfig_acltransport': 'openconfig-acl:transport',
            'openconfig_aclinput_interface': 'openconfig-acl:input-interface',
            'openconfig_aclconfig': 'openconfig-acl:config',
            'openconfig_aclactions': 'openconfig-acl:actions'
        }

        self._openconfig_acll2 = openconfig_acll2
        self._openconfig_aclipv4 = openconfig_aclipv4
        self._openconfig_aclipv6 = openconfig_aclipv6
        self._openconfig_acltransport = openconfig_acltransport
        self._openconfig_aclinput_interface = openconfig_aclinput_interface
        self._openconfig_aclconfig = openconfig_aclconfig
        self._openconfig_aclactions = openconfig_aclactions

    @classmethod
    def from_dict(cls, dikt) -> 'PostAclEntryConfig':
        """Returns the dict as a model

        :param dikt: A dict.
        :type: dict
        :return: The post_acl_entry_config of this PostAclEntryConfig.  # noqa: E501
        :rtype: PostAclEntryConfig
        """
        return util.deserialize_model(dikt, cls)

    @property
    def openconfig_acll2(self) -> AclOpenconfigaclaclAclsetsAclentriesL2:
        """Gets the openconfig_acll2 of this PostAclEntryConfig.


        :return: The openconfig_acll2 of this PostAclEntryConfig.
        :rtype: AclOpenconfigaclaclAclsetsAclentriesL2
        """
        return self._openconfig_acll2

    @openconfig_acll2.setter
    def openconfig_acll2(self, openconfig_acll2: AclOpenconfigaclaclAclsetsAclentriesL2):
        """Sets the openconfig_acll2 of this PostAclEntryConfig.


        :param openconfig_acll2: The openconfig_acll2 of this PostAclEntryConfig.
        :type openconfig_acll2: AclOpenconfigaclaclAclsetsAclentriesL2
        """

        self._openconfig_acll2 = openconfig_acll2

    @property
    def openconfig_aclipv4(self) -> AclOpenconfigaclaclAclsetsAclentriesIpv4:
        """Gets the openconfig_aclipv4 of this PostAclEntryConfig.


        :return: The openconfig_aclipv4 of this PostAclEntryConfig.
        :rtype: AclOpenconfigaclaclAclsetsAclentriesIpv4
        """
        return self._openconfig_aclipv4

    @openconfig_aclipv4.setter
    def openconfig_aclipv4(self, openconfig_aclipv4: AclOpenconfigaclaclAclsetsAclentriesIpv4):
        """Sets the openconfig_aclipv4 of this PostAclEntryConfig.


        :param openconfig_aclipv4: The openconfig_aclipv4 of this PostAclEntryConfig.
        :type openconfig_aclipv4: AclOpenconfigaclaclAclsetsAclentriesIpv4
        """

        self._openconfig_aclipv4 = openconfig_aclipv4

    @property
    def openconfig_aclipv6(self) -> AclOpenconfigaclaclAclsetsAclentriesIpv6:
        """Gets the openconfig_aclipv6 of this PostAclEntryConfig.


        :return: The openconfig_aclipv6 of this PostAclEntryConfig.
        :rtype: AclOpenconfigaclaclAclsetsAclentriesIpv6
        """
        return self._openconfig_aclipv6

    @openconfig_aclipv6.setter
    def openconfig_aclipv6(self, openconfig_aclipv6: AclOpenconfigaclaclAclsetsAclentriesIpv6):
        """Sets the openconfig_aclipv6 of this PostAclEntryConfig.


        :param openconfig_aclipv6: The openconfig_aclipv6 of this PostAclEntryConfig.
        :type openconfig_aclipv6: AclOpenconfigaclaclAclsetsAclentriesIpv6
        """

        self._openconfig_aclipv6 = openconfig_aclipv6

    @property
    def openconfig_acltransport(self) -> AclOpenconfigaclaclAclsetsAclentriesTransport:
        """Gets the openconfig_acltransport of this PostAclEntryConfig.


        :return: The openconfig_acltransport of this PostAclEntryConfig.
        :rtype: AclOpenconfigaclaclAclsetsAclentriesTransport
        """
        return self._openconfig_acltransport

    @openconfig_acltransport.setter
    def openconfig_acltransport(self, openconfig_acltransport: AclOpenconfigaclaclAclsetsAclentriesTransport):
        """Sets the openconfig_acltransport of this PostAclEntryConfig.


        :param openconfig_acltransport: The openconfig_acltransport of this PostAclEntryConfig.
        :type openconfig_acltransport: AclOpenconfigaclaclAclsetsAclentriesTransport
        """

        self._openconfig_acltransport = openconfig_acltransport

    @property
    def openconfig_aclinput_interface(self) -> AclOpenconfigaclaclAclsetsAclentriesInputinterface:
        """Gets the openconfig_aclinput_interface of this PostAclEntryConfig.


        :return: The openconfig_aclinput_interface of this PostAclEntryConfig.
        :rtype: AclOpenconfigaclaclAclsetsAclentriesInputinterface
        """
        return self._openconfig_aclinput_interface

    @openconfig_aclinput_interface.setter
    def openconfig_aclinput_interface(self, openconfig_aclinput_interface: AclOpenconfigaclaclAclsetsAclentriesInputinterface):
        """Sets the openconfig_aclinput_interface of this PostAclEntryConfig.


        :param openconfig_aclinput_interface: The openconfig_aclinput_interface of this PostAclEntryConfig.
        :type openconfig_aclinput_interface: AclOpenconfigaclaclAclsetsAclentriesInputinterface
        """

        self._openconfig_aclinput_interface = openconfig_aclinput_interface

    @property
    def openconfig_aclconfig(self) -> AclOpenconfigaclaclAclsetsAclentriesConfig:
        """Gets the openconfig_aclconfig of this PostAclEntryConfig.


        :return: The openconfig_aclconfig of this PostAclEntryConfig.
        :rtype: AclOpenconfigaclaclAclsetsAclentriesConfig
        """
        return self._openconfig_aclconfig

    @openconfig_aclconfig.setter
    def openconfig_aclconfig(self, openconfig_aclconfig: AclOpenconfigaclaclAclsetsAclentriesConfig):
        """Sets the openconfig_aclconfig of this PostAclEntryConfig.


        :param openconfig_aclconfig: The openconfig_aclconfig of this PostAclEntryConfig.
        :type openconfig_aclconfig: AclOpenconfigaclaclAclsetsAclentriesConfig
        """

        self._openconfig_aclconfig = openconfig_aclconfig

    @property
    def openconfig_aclactions(self) -> AclOpenconfigaclaclAclsetsAclentriesActions:
        """Gets the openconfig_aclactions of this PostAclEntryConfig.


        :return: The openconfig_aclactions of this PostAclEntryConfig.
        :rtype: AclOpenconfigaclaclAclsetsAclentriesActions
        """
        return self._openconfig_aclactions

    @openconfig_aclactions.setter
    def openconfig_aclactions(self, openconfig_aclactions: AclOpenconfigaclaclAclsetsAclentriesActions):
        """Sets the openconfig_aclactions of this PostAclEntryConfig.


        :param openconfig_aclactions: The openconfig_aclactions of this PostAclEntryConfig.
        :type openconfig_aclactions: AclOpenconfigaclaclAclsetsAclentriesActions
        """

        self._openconfig_aclactions = openconfig_aclactions
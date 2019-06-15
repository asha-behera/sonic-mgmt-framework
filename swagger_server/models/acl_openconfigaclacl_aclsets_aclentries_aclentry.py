# coding: utf-8

from __future__ import absolute_import
from datetime import date, datetime  # noqa: F401

from typing import List, Dict  # noqa: F401

from swagger_server.models.base_model_ import Model
from swagger_server.models.acl_openconfigaclacl_aclsets_aclentries_actions import AclOpenconfigaclaclAclsetsAclentriesActions  # noqa: F401,E501
from swagger_server.models.acl_openconfigaclacl_aclsets_aclentries_config import AclOpenconfigaclaclAclsetsAclentriesConfig  # noqa: F401,E501
from swagger_server.models.acl_openconfigaclacl_aclsets_aclentries_inputinterface import AclOpenconfigaclaclAclsetsAclentriesInputinterface  # noqa: F401,E501
from swagger_server.models.acl_openconfigaclacl_aclsets_aclentries_ipv4 import AclOpenconfigaclaclAclsetsAclentriesIpv4  # noqa: F401,E501
from swagger_server.models.acl_openconfigaclacl_aclsets_aclentries_ipv6 import AclOpenconfigaclaclAclsetsAclentriesIpv6  # noqa: F401,E501
from swagger_server.models.acl_openconfigaclacl_aclsets_aclentries_l2 import AclOpenconfigaclaclAclsetsAclentriesL2  # noqa: F401,E501
from swagger_server.models.acl_openconfigaclacl_aclsets_aclentries_transport import AclOpenconfigaclaclAclsetsAclentriesTransport  # noqa: F401,E501
from swagger_server import util


class AclOpenconfigaclaclAclsetsAclentriesAclentry(Model):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    def __init__(self, sequence_id: int=None, config: AclOpenconfigaclaclAclsetsAclentriesConfig=None, l2: AclOpenconfigaclaclAclsetsAclentriesL2=None, ipv4: AclOpenconfigaclaclAclsetsAclentriesIpv4=None, ipv6: AclOpenconfigaclaclAclsetsAclentriesIpv6=None, transport: AclOpenconfigaclaclAclsetsAclentriesTransport=None, input_interface: AclOpenconfigaclaclAclsetsAclentriesInputinterface=None, actions: AclOpenconfigaclaclAclsetsAclentriesActions=None):  # noqa: E501
        """AclOpenconfigaclaclAclsetsAclentriesAclentry - a model defined in Swagger

        :param sequence_id: The sequence_id of this AclOpenconfigaclaclAclsetsAclentriesAclentry.  # noqa: E501
        :type sequence_id: int
        :param config: The config of this AclOpenconfigaclaclAclsetsAclentriesAclentry.  # noqa: E501
        :type config: AclOpenconfigaclaclAclsetsAclentriesConfig
        :param l2: The l2 of this AclOpenconfigaclaclAclsetsAclentriesAclentry.  # noqa: E501
        :type l2: AclOpenconfigaclaclAclsetsAclentriesL2
        :param ipv4: The ipv4 of this AclOpenconfigaclaclAclsetsAclentriesAclentry.  # noqa: E501
        :type ipv4: AclOpenconfigaclaclAclsetsAclentriesIpv4
        :param ipv6: The ipv6 of this AclOpenconfigaclaclAclsetsAclentriesAclentry.  # noqa: E501
        :type ipv6: AclOpenconfigaclaclAclsetsAclentriesIpv6
        :param transport: The transport of this AclOpenconfigaclaclAclsetsAclentriesAclentry.  # noqa: E501
        :type transport: AclOpenconfigaclaclAclsetsAclentriesTransport
        :param input_interface: The input_interface of this AclOpenconfigaclaclAclsetsAclentriesAclentry.  # noqa: E501
        :type input_interface: AclOpenconfigaclaclAclsetsAclentriesInputinterface
        :param actions: The actions of this AclOpenconfigaclaclAclsetsAclentriesAclentry.  # noqa: E501
        :type actions: AclOpenconfigaclaclAclsetsAclentriesActions
        """
        self.swagger_types = {
            'sequence_id': int,
            'config': AclOpenconfigaclaclAclsetsAclentriesConfig,
            'l2': AclOpenconfigaclaclAclsetsAclentriesL2,
            'ipv4': AclOpenconfigaclaclAclsetsAclentriesIpv4,
            'ipv6': AclOpenconfigaclaclAclsetsAclentriesIpv6,
            'transport': AclOpenconfigaclaclAclsetsAclentriesTransport,
            'input_interface': AclOpenconfigaclaclAclsetsAclentriesInputinterface,
            'actions': AclOpenconfigaclaclAclsetsAclentriesActions
        }

        self.attribute_map = {
            'sequence_id': 'sequence-id',
            'config': 'config',
            'l2': 'l2',
            'ipv4': 'ipv4',
            'ipv6': 'ipv6',
            'transport': 'transport',
            'input_interface': 'input-interface',
            'actions': 'actions'
        }

        self._sequence_id = sequence_id
        self._config = config
        self._l2 = l2
        self._ipv4 = ipv4
        self._ipv6 = ipv6
        self._transport = transport
        self._input_interface = input_interface
        self._actions = actions

    @classmethod
    def from_dict(cls, dikt) -> 'AclOpenconfigaclaclAclsetsAclentriesAclentry':
        """Returns the dict as a model

        :param dikt: A dict.
        :type: dict
        :return: The acl_openconfigaclacl_aclsets_aclentries_aclentry of this AclOpenconfigaclaclAclsetsAclentriesAclentry.  # noqa: E501
        :rtype: AclOpenconfigaclaclAclsetsAclentriesAclentry
        """
        return util.deserialize_model(dikt, cls)

    @property
    def sequence_id(self) -> int:
        """Gets the sequence_id of this AclOpenconfigaclaclAclsetsAclentriesAclentry.


        :return: The sequence_id of this AclOpenconfigaclaclAclsetsAclentriesAclentry.
        :rtype: int
        """
        return self._sequence_id

    @sequence_id.setter
    def sequence_id(self, sequence_id: int):
        """Sets the sequence_id of this AclOpenconfigaclaclAclsetsAclentriesAclentry.


        :param sequence_id: The sequence_id of this AclOpenconfigaclaclAclsetsAclentriesAclentry.
        :type sequence_id: int
        """
        if sequence_id is None:
            raise ValueError("Invalid value for `sequence_id`, must not be `None`")  # noqa: E501

        self._sequence_id = sequence_id

    @property
    def config(self) -> AclOpenconfigaclaclAclsetsAclentriesConfig:
        """Gets the config of this AclOpenconfigaclaclAclsetsAclentriesAclentry.


        :return: The config of this AclOpenconfigaclaclAclsetsAclentriesAclentry.
        :rtype: AclOpenconfigaclaclAclsetsAclentriesConfig
        """
        return self._config

    @config.setter
    def config(self, config: AclOpenconfigaclaclAclsetsAclentriesConfig):
        """Sets the config of this AclOpenconfigaclaclAclsetsAclentriesAclentry.


        :param config: The config of this AclOpenconfigaclaclAclsetsAclentriesAclentry.
        :type config: AclOpenconfigaclaclAclsetsAclentriesConfig
        """

        self._config = config

    @property
    def l2(self) -> AclOpenconfigaclaclAclsetsAclentriesL2:
        """Gets the l2 of this AclOpenconfigaclaclAclsetsAclentriesAclentry.


        :return: The l2 of this AclOpenconfigaclaclAclsetsAclentriesAclentry.
        :rtype: AclOpenconfigaclaclAclsetsAclentriesL2
        """
        return self._l2

    @l2.setter
    def l2(self, l2: AclOpenconfigaclaclAclsetsAclentriesL2):
        """Sets the l2 of this AclOpenconfigaclaclAclsetsAclentriesAclentry.


        :param l2: The l2 of this AclOpenconfigaclaclAclsetsAclentriesAclentry.
        :type l2: AclOpenconfigaclaclAclsetsAclentriesL2
        """

        self._l2 = l2

    @property
    def ipv4(self) -> AclOpenconfigaclaclAclsetsAclentriesIpv4:
        """Gets the ipv4 of this AclOpenconfigaclaclAclsetsAclentriesAclentry.


        :return: The ipv4 of this AclOpenconfigaclaclAclsetsAclentriesAclentry.
        :rtype: AclOpenconfigaclaclAclsetsAclentriesIpv4
        """
        return self._ipv4

    @ipv4.setter
    def ipv4(self, ipv4: AclOpenconfigaclaclAclsetsAclentriesIpv4):
        """Sets the ipv4 of this AclOpenconfigaclaclAclsetsAclentriesAclentry.


        :param ipv4: The ipv4 of this AclOpenconfigaclaclAclsetsAclentriesAclentry.
        :type ipv4: AclOpenconfigaclaclAclsetsAclentriesIpv4
        """

        self._ipv4 = ipv4

    @property
    def ipv6(self) -> AclOpenconfigaclaclAclsetsAclentriesIpv6:
        """Gets the ipv6 of this AclOpenconfigaclaclAclsetsAclentriesAclentry.


        :return: The ipv6 of this AclOpenconfigaclaclAclsetsAclentriesAclentry.
        :rtype: AclOpenconfigaclaclAclsetsAclentriesIpv6
        """
        return self._ipv6

    @ipv6.setter
    def ipv6(self, ipv6: AclOpenconfigaclaclAclsetsAclentriesIpv6):
        """Sets the ipv6 of this AclOpenconfigaclaclAclsetsAclentriesAclentry.


        :param ipv6: The ipv6 of this AclOpenconfigaclaclAclsetsAclentriesAclentry.
        :type ipv6: AclOpenconfigaclaclAclsetsAclentriesIpv6
        """

        self._ipv6 = ipv6

    @property
    def transport(self) -> AclOpenconfigaclaclAclsetsAclentriesTransport:
        """Gets the transport of this AclOpenconfigaclaclAclsetsAclentriesAclentry.


        :return: The transport of this AclOpenconfigaclaclAclsetsAclentriesAclentry.
        :rtype: AclOpenconfigaclaclAclsetsAclentriesTransport
        """
        return self._transport

    @transport.setter
    def transport(self, transport: AclOpenconfigaclaclAclsetsAclentriesTransport):
        """Sets the transport of this AclOpenconfigaclaclAclsetsAclentriesAclentry.


        :param transport: The transport of this AclOpenconfigaclaclAclsetsAclentriesAclentry.
        :type transport: AclOpenconfigaclaclAclsetsAclentriesTransport
        """

        self._transport = transport

    @property
    def input_interface(self) -> AclOpenconfigaclaclAclsetsAclentriesInputinterface:
        """Gets the input_interface of this AclOpenconfigaclaclAclsetsAclentriesAclentry.


        :return: The input_interface of this AclOpenconfigaclaclAclsetsAclentriesAclentry.
        :rtype: AclOpenconfigaclaclAclsetsAclentriesInputinterface
        """
        return self._input_interface

    @input_interface.setter
    def input_interface(self, input_interface: AclOpenconfigaclaclAclsetsAclentriesInputinterface):
        """Sets the input_interface of this AclOpenconfigaclaclAclsetsAclentriesAclentry.


        :param input_interface: The input_interface of this AclOpenconfigaclaclAclsetsAclentriesAclentry.
        :type input_interface: AclOpenconfigaclaclAclsetsAclentriesInputinterface
        """

        self._input_interface = input_interface

    @property
    def actions(self) -> AclOpenconfigaclaclAclsetsAclentriesActions:
        """Gets the actions of this AclOpenconfigaclaclAclsetsAclentriesAclentry.


        :return: The actions of this AclOpenconfigaclaclAclsetsAclentriesAclentry.
        :rtype: AclOpenconfigaclaclAclsetsAclentriesActions
        """
        return self._actions

    @actions.setter
    def actions(self, actions: AclOpenconfigaclaclAclsetsAclentriesActions):
        """Sets the actions of this AclOpenconfigaclaclAclsetsAclentriesAclentry.


        :param actions: The actions of this AclOpenconfigaclaclAclsetsAclentriesAclentry.
        :type actions: AclOpenconfigaclaclAclsetsAclentriesActions
        """

        self._actions = actions
# coding: utf-8

"""
    Sonic NMS

    Network management Open APIs for Broadcom's Sonic.  # noqa: E501

    OpenAPI spec version: 1.0.0
    Contact: mohammed.faraaz@broadcom.com
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


import pprint
import re  # noqa: F401

import six

from swagger_client.models.config_destination_port import ConfigDestinationPort  # noqa: F401,E501
from swagger_client.models.config_source_port import ConfigSourcePort  # noqa: F401,E501


class PostConfigSourcePort(ConfigSourcePort):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    """
    Attributes:
      swagger_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    swagger_types = {
        'destination_port': 'str',
        'tcp_flags': 'str'
    }

    attribute_map = {
        'destination_port': 'destination-port',
        'tcp_flags': 'tcp-flags'
    }

    def __init__(self, destination_port=None, tcp_flags=None):  # noqa: E501
        """PostConfigSourcePort - a model defined in Swagger"""  # noqa: E501

        self._destination_port = None
        self._tcp_flags = None
        self.discriminator = None

        if destination_port is not None:
            self.destination_port = destination_port
        if tcp_flags is not None:
            self.tcp_flags = tcp_flags

    @property
    def destination_port(self):
        """Gets the destination_port of this PostConfigSourcePort.  # noqa: E501


        :return: The destination_port of this PostConfigSourcePort.  # noqa: E501
        :rtype: str
        """
        return self._destination_port

    @destination_port.setter
    def destination_port(self, destination_port):
        """Sets the destination_port of this PostConfigSourcePort.


        :param destination_port: The destination_port of this PostConfigSourcePort.  # noqa: E501
        :type: str
        """

        self._destination_port = destination_port

    @property
    def tcp_flags(self):
        """Gets the tcp_flags of this PostConfigSourcePort.  # noqa: E501


        :return: The tcp_flags of this PostConfigSourcePort.  # noqa: E501
        :rtype: str
        """
        return self._tcp_flags

    @tcp_flags.setter
    def tcp_flags(self, tcp_flags):
        """Sets the tcp_flags of this PostConfigSourcePort.


        :param tcp_flags: The tcp_flags of this PostConfigSourcePort.  # noqa: E501
        :type: str
        """

        self._tcp_flags = tcp_flags

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.swagger_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value
        if issubclass(PostConfigSourcePort, dict):
            for key, value in self.items():
                result[key] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, PostConfigSourcePort):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
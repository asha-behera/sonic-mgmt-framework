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


class Ipv6ConfigHopLimit(object):
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
        'hop_limit': 'int'
    }

    attribute_map = {
        'hop_limit': 'hop-limit'
    }

    def __init__(self, hop_limit=None):  # noqa: E501
        """Ipv6ConfigHopLimit - a model defined in Swagger"""  # noqa: E501

        self._hop_limit = None
        self.discriminator = None

        if hop_limit is not None:
            self.hop_limit = hop_limit

    @property
    def hop_limit(self):
        """Gets the hop_limit of this Ipv6ConfigHopLimit.  # noqa: E501


        :return: The hop_limit of this Ipv6ConfigHopLimit.  # noqa: E501
        :rtype: int
        """
        return self._hop_limit

    @hop_limit.setter
    def hop_limit(self, hop_limit):
        """Sets the hop_limit of this Ipv6ConfigHopLimit.


        :param hop_limit: The hop_limit of this Ipv6ConfigHopLimit.  # noqa: E501
        :type: int
        """

        self._hop_limit = hop_limit

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
        if issubclass(Ipv6ConfigHopLimit, dict):
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
        if not isinstance(other, Ipv6ConfigHopLimit):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
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


class GetAclEntryStateMatchedOctets(object):
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
        'matched_octets': 'int'
    }

    attribute_map = {
        'matched_octets': 'matched-octets'
    }

    def __init__(self, matched_octets=None):  # noqa: E501
        """GetAclEntryStateMatchedOctets - a model defined in Swagger"""  # noqa: E501

        self._matched_octets = None
        self.discriminator = None

        if matched_octets is not None:
            self.matched_octets = matched_octets

    @property
    def matched_octets(self):
        """Gets the matched_octets of this GetAclEntryStateMatchedOctets.  # noqa: E501


        :return: The matched_octets of this GetAclEntryStateMatchedOctets.  # noqa: E501
        :rtype: int
        """
        return self._matched_octets

    @matched_octets.setter
    def matched_octets(self, matched_octets):
        """Sets the matched_octets of this GetAclEntryStateMatchedOctets.


        :param matched_octets: The matched_octets of this GetAclEntryStateMatchedOctets.  # noqa: E501
        :type: int
        """

        self._matched_octets = matched_octets

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
        if issubclass(GetAclEntryStateMatchedOctets, dict):
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
        if not isinstance(other, GetAclEntryStateMatchedOctets):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
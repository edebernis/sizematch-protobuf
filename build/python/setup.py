#!/usr/bin/env python3
# -*- coding: utf-8 -*-

from setuptools import setup, find_packages


setup(
    name="sizematch-protobuf",
    version="0.0.7",
    packages=find_packages(),
    description="Sizematch protobuf libraries",
    author="Emeric de Bernis",
    author_email="emeric.debernis@gmail.com",
    install_requires=[
        "protobuf==3.11.3"
    ]
)

# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: items/items.proto

from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='items/items.proto',
  package='sizematch.protobuf.items',
  syntax='proto3',
  serialized_options=b'Z\005items',
  serialized_pb=b'\n\x11items/items.proto\x12\x18sizematch.protobuf.items\"\xa7\x02\n\x04Item\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0e\n\x06source\x18\x02 \x01(\t\x12\x0c\n\x04lang\x18\x03 \x01(\t\x12\x0c\n\x04urls\x18\x04 \x03(\t\x12\x0c\n\x04name\x18\x05 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x06 \x01(\t\x12\x12\n\ncategories\x18\x07 \x03(\t\x12\x12\n\nimage_urls\x18\x08 \x03(\t\x12\x42\n\ndimensions\x18\t \x03(\x0b\x32..sizematch.protobuf.items.Item.DimensionsEntry\x12\r\n\x05price\x18\n \x01(\x04\x12\x16\n\x0eprice_currency\x18\x0b \x01(\t\x1a\x31\n\x0f\x44imensionsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"\x9c\x02\n\x0eNormalizedItem\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0e\n\x06source\x18\x02 \x01(\t\x12,\n\x04lang\x18\x03 \x01(\x0e\x32\x1e.sizematch.protobuf.items.Lang\x12\x0c\n\x04urls\x18\x04 \x03(\t\x12\x0c\n\x04name\x18\x05 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x06 \x01(\t\x12\x12\n\ncategories\x18\x07 \x03(\t\x12\x12\n\nimage_urls\x18\x08 \x03(\t\x12\x37\n\ndimensions\x18\t \x03(\x0b\x32#.sizematch.protobuf.items.Dimension\x12.\n\x05price\x18\n \x01(\x0b\x32\x1f.sizematch.protobuf.items.Price\"\xee\x01\n\tDimension\x12\x36\n\x04name\x18\x01 \x01(\x0e\x32(.sizematch.protobuf.items.Dimension.Name\x12\r\n\x05value\x18\x02 \x01(\x04\x12\x36\n\x04unit\x18\x03 \x01(\x0e\x32(.sizematch.protobuf.items.Dimension.Unit\"4\n\x04Name\x12\n\n\x06HEIGHT\x10\x00\x12\t\n\x05WIDTH\x10\x01\x12\t\n\x05\x44\x45PTH\x10\x02\x12\n\n\x06WEIGHT\x10\x03\",\n\x04Unit\x12\x06\n\x02MM\x10\x00\x12\x06\n\x02\x43M\x10\x01\x12\x05\n\x01M\x10\x02\x12\x05\n\x01G\x10\x03\x12\x06\n\x02KG\x10\x04\"p\n\x05Price\x12\r\n\x05price\x18\x01 \x01(\x04\x12:\n\x08\x63urrency\x18\x02 \x01(\x0e\x32(.sizematch.protobuf.items.Price.Currency\"\x1c\n\x08\x43urrency\x12\x07\n\x03\x45UR\x10\x00\x12\x07\n\x03GBP\x10\x01*\x16\n\x04Lang\x12\x06\n\x02\x45N\x10\x00\x12\x06\n\x02\x46R\x10\x01\x42\x07Z\x05itemsb\x06proto3'
)

_LANG = _descriptor.EnumDescriptor(
  name='Lang',
  full_name='sizematch.protobuf.items.Lang',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='EN', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='FR', index=1, number=1,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=987,
  serialized_end=1009,
)
_sym_db.RegisterEnumDescriptor(_LANG)

Lang = enum_type_wrapper.EnumTypeWrapper(_LANG)
EN = 0
FR = 1


_DIMENSION_NAME = _descriptor.EnumDescriptor(
  name='Name',
  full_name='sizematch.protobuf.items.Dimension.Name',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='HEIGHT', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='WIDTH', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='DEPTH', index=2, number=2,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='WEIGHT', index=3, number=3,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=773,
  serialized_end=825,
)
_sym_db.RegisterEnumDescriptor(_DIMENSION_NAME)

_DIMENSION_UNIT = _descriptor.EnumDescriptor(
  name='Unit',
  full_name='sizematch.protobuf.items.Dimension.Unit',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='MM', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CM', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='M', index=2, number=2,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='G', index=3, number=3,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='KG', index=4, number=4,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=827,
  serialized_end=871,
)
_sym_db.RegisterEnumDescriptor(_DIMENSION_UNIT)

_PRICE_CURRENCY = _descriptor.EnumDescriptor(
  name='Currency',
  full_name='sizematch.protobuf.items.Price.Currency',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='EUR', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='GBP', index=1, number=1,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=957,
  serialized_end=985,
)
_sym_db.RegisterEnumDescriptor(_PRICE_CURRENCY)


_ITEM_DIMENSIONSENTRY = _descriptor.Descriptor(
  name='DimensionsEntry',
  full_name='sizematch.protobuf.items.Item.DimensionsEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='sizematch.protobuf.items.Item.DimensionsEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='sizematch.protobuf.items.Item.DimensionsEntry.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=b'8\001',
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=294,
  serialized_end=343,
)

_ITEM = _descriptor.Descriptor(
  name='Item',
  full_name='sizematch.protobuf.items.Item',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='sizematch.protobuf.items.Item.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='source', full_name='sizematch.protobuf.items.Item.source', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='lang', full_name='sizematch.protobuf.items.Item.lang', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='urls', full_name='sizematch.protobuf.items.Item.urls', index=3,
      number=4, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='name', full_name='sizematch.protobuf.items.Item.name', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='description', full_name='sizematch.protobuf.items.Item.description', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='categories', full_name='sizematch.protobuf.items.Item.categories', index=6,
      number=7, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='image_urls', full_name='sizematch.protobuf.items.Item.image_urls', index=7,
      number=8, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='dimensions', full_name='sizematch.protobuf.items.Item.dimensions', index=8,
      number=9, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='price', full_name='sizematch.protobuf.items.Item.price', index=9,
      number=10, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='price_currency', full_name='sizematch.protobuf.items.Item.price_currency', index=10,
      number=11, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_ITEM_DIMENSIONSENTRY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=48,
  serialized_end=343,
)


_NORMALIZEDITEM = _descriptor.Descriptor(
  name='NormalizedItem',
  full_name='sizematch.protobuf.items.NormalizedItem',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='sizematch.protobuf.items.NormalizedItem.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='source', full_name='sizematch.protobuf.items.NormalizedItem.source', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='lang', full_name='sizematch.protobuf.items.NormalizedItem.lang', index=2,
      number=3, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='urls', full_name='sizematch.protobuf.items.NormalizedItem.urls', index=3,
      number=4, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='name', full_name='sizematch.protobuf.items.NormalizedItem.name', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='description', full_name='sizematch.protobuf.items.NormalizedItem.description', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='categories', full_name='sizematch.protobuf.items.NormalizedItem.categories', index=6,
      number=7, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='image_urls', full_name='sizematch.protobuf.items.NormalizedItem.image_urls', index=7,
      number=8, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='dimensions', full_name='sizematch.protobuf.items.NormalizedItem.dimensions', index=8,
      number=9, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='price', full_name='sizematch.protobuf.items.NormalizedItem.price', index=9,
      number=10, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=346,
  serialized_end=630,
)


_DIMENSION = _descriptor.Descriptor(
  name='Dimension',
  full_name='sizematch.protobuf.items.Dimension',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='sizematch.protobuf.items.Dimension.name', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='sizematch.protobuf.items.Dimension.value', index=1,
      number=2, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='unit', full_name='sizematch.protobuf.items.Dimension.unit', index=2,
      number=3, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _DIMENSION_NAME,
    _DIMENSION_UNIT,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=633,
  serialized_end=871,
)


_PRICE = _descriptor.Descriptor(
  name='Price',
  full_name='sizematch.protobuf.items.Price',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='price', full_name='sizematch.protobuf.items.Price.price', index=0,
      number=1, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='currency', full_name='sizematch.protobuf.items.Price.currency', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _PRICE_CURRENCY,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=873,
  serialized_end=985,
)

_ITEM_DIMENSIONSENTRY.containing_type = _ITEM
_ITEM.fields_by_name['dimensions'].message_type = _ITEM_DIMENSIONSENTRY
_NORMALIZEDITEM.fields_by_name['lang'].enum_type = _LANG
_NORMALIZEDITEM.fields_by_name['dimensions'].message_type = _DIMENSION
_NORMALIZEDITEM.fields_by_name['price'].message_type = _PRICE
_DIMENSION.fields_by_name['name'].enum_type = _DIMENSION_NAME
_DIMENSION.fields_by_name['unit'].enum_type = _DIMENSION_UNIT
_DIMENSION_NAME.containing_type = _DIMENSION
_DIMENSION_UNIT.containing_type = _DIMENSION
_PRICE.fields_by_name['currency'].enum_type = _PRICE_CURRENCY
_PRICE_CURRENCY.containing_type = _PRICE
DESCRIPTOR.message_types_by_name['Item'] = _ITEM
DESCRIPTOR.message_types_by_name['NormalizedItem'] = _NORMALIZEDITEM
DESCRIPTOR.message_types_by_name['Dimension'] = _DIMENSION
DESCRIPTOR.message_types_by_name['Price'] = _PRICE
DESCRIPTOR.enum_types_by_name['Lang'] = _LANG
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Item = _reflection.GeneratedProtocolMessageType('Item', (_message.Message,), {

  'DimensionsEntry' : _reflection.GeneratedProtocolMessageType('DimensionsEntry', (_message.Message,), {
    'DESCRIPTOR' : _ITEM_DIMENSIONSENTRY,
    '__module__' : 'items.items_pb2'
    # @@protoc_insertion_point(class_scope:sizematch.protobuf.items.Item.DimensionsEntry)
    })
  ,
  'DESCRIPTOR' : _ITEM,
  '__module__' : 'items.items_pb2'
  # @@protoc_insertion_point(class_scope:sizematch.protobuf.items.Item)
  })
_sym_db.RegisterMessage(Item)
_sym_db.RegisterMessage(Item.DimensionsEntry)

NormalizedItem = _reflection.GeneratedProtocolMessageType('NormalizedItem', (_message.Message,), {
  'DESCRIPTOR' : _NORMALIZEDITEM,
  '__module__' : 'items.items_pb2'
  # @@protoc_insertion_point(class_scope:sizematch.protobuf.items.NormalizedItem)
  })
_sym_db.RegisterMessage(NormalizedItem)

Dimension = _reflection.GeneratedProtocolMessageType('Dimension', (_message.Message,), {
  'DESCRIPTOR' : _DIMENSION,
  '__module__' : 'items.items_pb2'
  # @@protoc_insertion_point(class_scope:sizematch.protobuf.items.Dimension)
  })
_sym_db.RegisterMessage(Dimension)

Price = _reflection.GeneratedProtocolMessageType('Price', (_message.Message,), {
  'DESCRIPTOR' : _PRICE,
  '__module__' : 'items.items_pb2'
  # @@protoc_insertion_point(class_scope:sizematch.protobuf.items.Price)
  })
_sym_db.RegisterMessage(Price)


DESCRIPTOR._options = None
_ITEM_DIMENSIONSENTRY._options = None
# @@protoc_insertion_point(module_scope)

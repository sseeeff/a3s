# Model
model:
  rest_name: mtlssource
  resource_name: mtlssources
  entity_name: MTLSSource
  package: a3s
  group: authn
  description: An MTLS Auth source can be used to issue tokens based on user certificates.
  get:
    description: Get a particular mtlssource object.
  update:
    description: Update a particular mtlssource object.
  delete:
    description: Delete a particular mtlssource object.
  extends:
  - '@sharded'
  - '@identifiable'

# Indexes
indexes:
- - namespace
  - name

# Attributes
attributes:
  v1:
  - name: certificateAuthority
    description: The Certificate authority to use to validate user certificates in PEM format.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: |-
      -----BEGIN CERTIFICATE-----
      MIIBZTCCAQugAwIBAgIRANYvXLTa16Ykvc9hQ4BBLJEwCgYIKoZIzj0EAwIwEjEQ
      MA4GA1UEAxMHQUNNRSBDQTAeFw0yMTExMDEyMzAwMTlaFw0zMTA5MTAyMzAwMTla
      MBIxEDAOBgNVBAMTB0FDTUUgQ0EwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASa
      7wknroxwB1znupZ67NzTG9Kuc+tNRlbI22eTDNMKYpIexzWDOyiQ95N3GQIdmAz5
      wVu9l2V3VuKUpD9mNgkRo0IwQDAOBgNVHQ8BAf8EBAMCAQYwDwYDVR0TAQH/BAUw
      AwEB/zAdBgNVHQ4EFgQURIT2kL76vMj9A3r9AUnaiHnHf4EwCgYIKoZIzj0EAwID
      SAAwRQIgS4SGaJ/B1Ul88Jal11Q5BwiY9bY2y9w+4xPNBxSyAIcCIQCSWVq+00xS
      bOmROq+EsxO4L/GzJx7MBbeJ6x142VKSBQ==
      -----END CERTIFICATE-----
    validations:
    - $pem

  - name: description
    description: The description of the object.
    type: string
    exposed: true
    stored: true

  - name: name
    description: The name of the source.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: mypki
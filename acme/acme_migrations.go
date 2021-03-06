package acme

import (
	"log"

	"github.com/hashicorp/terraform/terraform"
)

// resourceACMERegistrationMigrateState is the outer migration function for
// acme_registration, dispatching to specific incremental version upgraders as
// need be.
func resourceACMERegistrationMigrateState(version int, os *terraform.InstanceState, meta interface{}) (*terraform.InstanceState, error) {
	// Guard against a nil state.
	if os == nil {
		return nil, nil
	}

	// Guard against empty state, can't do anything with it
	if os.Empty() {
		return os, nil
	}

	var migrateFunc func(*terraform.InstanceState, interface{}) error
	switch version {
	case 0:
		log.Printf("[DEBUG] Migrating acme_registration state: old v%d state: %#v", version, os)
		migrateFunc = migrateACMERegistrationStateV1
	default:
		// Migration is complete
		log.Printf("[DEBUG] Migrating acme_registration state: completed v%d state: %#v", version, os)
		return os, nil
	}
	if err := migrateFunc(os, meta); err != nil {
		return nil, err
	}
	version++
	log.Printf("[DEBUG] Migrating acme_registration state: new v%d state: %#v", version, os)
	return resourceACMERegistrationMigrateState(version, os, meta)
}

// migrateACMERegistrationStateV1 handles migration of acme_registration from
// schema version 0 to version 1.
func migrateACMERegistrationStateV1(is *terraform.InstanceState, meta interface{}) error {
	delete(is.Attributes, "server_url")
	delete(is.Attributes, "registration_body")
	delete(is.Attributes, "registration_new_authz_url")
	delete(is.Attributes, "registration_tos_url")

	return nil
}

// resourceACMECertificateMigrateState is the outer migration function for
// acme_certificate, dispatching to specific incremental version upgraders as
// need be.
func resourceACMECertificateMigrateState(version int, os *terraform.InstanceState, meta interface{}) (*terraform.InstanceState, error) {
	// Guard against a nil state.
	if os == nil {
		return nil, nil
	}

	// Guard against empty state, can't do anything with it
	if os.Empty() {
		return os, nil
	}

	var migrateFunc func(*terraform.InstanceState, interface{}) error
	switch version {
	case 0:
		log.Printf("[DEBUG] Migrating acme_certificate state: old v%d state: %#v", version, os)
		migrateFunc = migrateACMECertificateStateV1
	default:
		// Migration is complete
		log.Printf("[DEBUG] Migrating acme_certificate state: completed v%d state: %#v", version, os)
		return os, nil
	}
	if err := migrateFunc(os, meta); err != nil {
		return nil, err
	}
	version++
	log.Printf("[DEBUG] Migrating acme_certificate state: new v%d state: %#v", version, os)
	return resourceACMECertificateMigrateState(version, os, meta)
}

// migrateACMECertificateStateV1 handles migration of acme_certificate from
// schema version 0 to version 1.
func migrateACMECertificateStateV1(is *terraform.InstanceState, meta interface{}) error {
	delete(is.Attributes, "server_url")
	delete(is.Attributes, "http_challenge_port")
	delete(is.Attributes, "tls_challenge_port")
	delete(is.Attributes, "registration_url")

	return nil
}

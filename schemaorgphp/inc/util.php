<?php
namespace SchemaOrg;

function so_json_serialize( $value ) {
	if ( empty( $value ) ) {
		return;
	}

	if ( is_array( $value ) ) {
		$out = array();

		foreach ( $value as $index => $element ) {
			if ( method_exists( $element, 'jsonSerialize' ) ) {
				array_push( $out, $element->jsonSerialize() );
			} else {
				array_push( $out, element );
			}
		}

		unset( $out['@context'] );
		return $out;
	}

	if ( is_object( $value ) ) {
		$out = array();

		if ( method_exists( $value, 'jsonSerialize' ) ) {
			$out = $value->jsonSerialize();
		}

		unset( $out['@context'] );
		return $out;
	}

	return $value;
}

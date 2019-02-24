<?php
namespace SchemaOrg;

function json_serialize( $value ) {
	if ( empty( $value ) ) {
		return;
	}

	if ( is_array( $value ) ) {
		$out = array();

		foreach ( $value as $index => $element ) {
			if ( method_exists( $element, 'jsonSerialize' ) ) {
				$out_part = $element->jsonSerialize();
				unset( $out_part['@context'] );

				array_push( $out, $out_part );
			} else {
				array_push( $out, $element );
			}
		}

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

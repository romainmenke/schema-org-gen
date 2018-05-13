<?php

class MusicComposition extends CreativeWork implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'MusicComposition';
	
	/**
	 * The person or organization who wrote a composition, or who is the composer of a work performed at some event.
	 * see : https://schema.org/composer
	 * @var \Organization|\Organization[]|\Person|\Person[]
	 */
	public var $composer;
	
	/**
	 * The date and place the work was first performed.
	 * see : https://schema.org/firstPerformance
	 * @var \Event|\Event[]
	 */
	public var $first_performance;
	
	/**
	 * Smaller compositions included in this work (e.g. a movement in a symphony).
	 * see : https://schema.org/includedComposition
	 * @var \MusicComposition|\MusicComposition[]
	 */
	public var $included_composition;
	
	/**
	 * The International Standard Musical Work Code for the composition.
	 * see : https://schema.org/iswcCode
	 * @var string|string[]
	 */
	public var $iswc_code;
	
	/**
	 * The person who wrote the words.
	 * see : https://schema.org/lyricist
	 * @var \Person|\Person[]
	 */
	public var $lyricist;
	
	/**
	 * The words in the song.
	 * see : https://schema.org/lyrics
	 * @var \CreativeWork|\CreativeWork[]
	 */
	public var $lyrics;
	
	/**
	 * An arrangement derived from the composition.
	 * see : https://schema.org/musicArrangement
	 * @var \MusicComposition|\MusicComposition[]
	 */
	public var $music_arrangement;
	
	/**
	 * The type of composition (e.g. overture, sonata, symphony, etc.).
	 * see : https://schema.org/musicCompositionForm
	 * @var string|string[]
	 */
	public var $music_composition_form;
	
	/**
	 * The key, mode, or scale this composition uses.
	 * see : https://schema.org/musicalKey
	 * @var string|string[]
	 */
	public var $musical_key;
	
	/**
	 * An audio recording of the work. Inverse property: recordingOf (see: https://schema.org/recordingOf).
	 * see : https://schema.org/recordedAs
	 * @var \MusicRecording|\MusicRecording[]
	 */
	public var $recorded_as;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'MusicComposition'
		);
		
		$serialized = so_json_serialize( $this->composer );
		if ( ! empty( $serialized ) ) {
			$out['composer'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->first_performance );
		if ( ! empty( $serialized ) ) {
			$out['firstPerformance'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->included_composition );
		if ( ! empty( $serialized ) ) {
			$out['includedComposition'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->iswc_code );
		if ( ! empty( $serialized ) ) {
			$out['iswcCode'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->lyricist );
		if ( ! empty( $serialized ) ) {
			$out['lyricist'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->lyrics );
		if ( ! empty( $serialized ) ) {
			$out['lyrics'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->music_arrangement );
		if ( ! empty( $serialized ) ) {
			$out['musicArrangement'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->music_composition_form );
		if ( ! empty( $serialized ) ) {
			$out['musicCompositionForm'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->musical_key );
		if ( ! empty( $serialized ) ) {
			$out['musicalKey'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->recorded_as );
		if ( ! empty( $serialized ) ) {
			$out['recordedAs'] = $serialized;
		}
		
		return $out;
	}
}

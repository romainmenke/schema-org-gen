<?php

class MusicVideoObject extends MediaObject implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'MusicVideoObject';
	
	/**
	 * A NewsArticle associated with the Media Object.
	 * see : https://schema.org/associatedArticle
	 * @var \NewsArticle|\NewsArticle[]
	 */
	public var $associated_article;
	
	/**
	 * The bitrate of the media object.
	 * see : https://schema.org/bitrate
	 * @var string|string[]
	 */
	public var $bitrate;
	
	/**
	 * File size in (mega/kilo) bytes.
	 * see : https://schema.org/contentSize
	 * @var string|string[]
	 */
	public var $content_size;
	
	/**
	 * Actual bytes of the media object, for example the image file or video file.
	 * see : https://schema.org/contentUrl
	 * @var string|string[]
	 */
	public var $content_url;
	
	/**
	 * The duration of the item (movie, audio recording, event, etc.) in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601).
	 * see : https://schema.org/duration
	 * @var \Duration|\Duration[]
	 */
	public var $duration;
	
	/**
	 * A URL pointing to a player for a specific video. In general, this is the information in the src element of an embed tag and should not be the same as the content of the loc tag.
	 * see : https://schema.org/embedUrl
	 * @var string|string[]
	 */
	public var $embed_url;
	
	/**
	 * The CreativeWork encoded by this media object.
	 * see : https://schema.org/encodesCreativeWork
	 * @var \CreativeWork|\CreativeWork[]
	 */
	public var $encodes_creative_work;
	
	/**
	 * mp3, mpeg4, etc.
	 * see : https://schema.org/encodingFormat
	 * @var string|string[]
	 */
	public var $encoding_format;
	
	/**
	 * The height of the item.
	 * see : https://schema.org/height
	 * @var \Distance|\Distance[]|\QuantitativeValue|\QuantitativeValue[]
	 */
	public var $height;
	
	/**
	 * Player type required—for example, Flash or Silverlight.
	 * see : https://schema.org/playerType
	 * @var string|string[]
	 */
	public var $player_type;
	
	/**
	 * The production company or studio responsible for the item e.g. series, video game, episode etc.
	 * see : https://schema.org/productionCompany
	 * @var \Organization|\Organization[]
	 */
	public var $production_company;
	
	/**
	 * The regions where the media is allowed. If not specified, then it&#39;s assumed to be allowed everywhere. Specify the countries in ISO 3166 format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_3166).
	 * see : https://schema.org/regionsAllowed
	 * @var \Place|\Place[]
	 */
	public var $regions_allowed;
	
	/**
	 * Indicates if use of the media require a subscription  (either paid or free). Allowed values are true or false (note that an earlier version had &#39;yes&#39;, &#39;no&#39;).
	 * see : https://schema.org/requiresSubscription
	 * @var boolean|boolean[]
	 */
	public var $requires_subscription;
	
	/**
	 * Date when this media object was uploaded to this site.
	 * see : https://schema.org/uploadDate
	 * @var string|string[]
	 */
	public var $upload_date;
	
	/**
	 * The width of the item.
	 * see : https://schema.org/width
	 * @var \Distance|\Distance[]|\QuantitativeValue|\QuantitativeValue[]
	 */
	public var $width;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'MusicVideoObject'
		);
		
		$serialized = so_json_serialize( $this->associated_article );
		if ( ! empty( $serialized ) ) {
			$out['associatedArticle'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->bitrate );
		if ( ! empty( $serialized ) ) {
			$out['bitrate'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->content_size );
		if ( ! empty( $serialized ) ) {
			$out['contentSize'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->content_url );
		if ( ! empty( $serialized ) ) {
			$out['contentUrl'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->duration );
		if ( ! empty( $serialized ) ) {
			$out['duration'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->embed_url );
		if ( ! empty( $serialized ) ) {
			$out['embedUrl'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->encodes_creative_work );
		if ( ! empty( $serialized ) ) {
			$out['encodesCreativeWork'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->encoding_format );
		if ( ! empty( $serialized ) ) {
			$out['encodingFormat'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->height );
		if ( ! empty( $serialized ) ) {
			$out['height'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->player_type );
		if ( ! empty( $serialized ) ) {
			$out['playerType'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->production_company );
		if ( ! empty( $serialized ) ) {
			$out['productionCompany'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->regions_allowed );
		if ( ! empty( $serialized ) ) {
			$out['regionsAllowed'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->requires_subscription );
		if ( ! empty( $serialized ) ) {
			$out['requiresSubscription'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->upload_date );
		if ( ! empty( $serialized ) ) {
			$out['uploadDate'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->width );
		if ( ! empty( $serialized ) ) {
			$out['width'] = $serialized;
		}
		
		return $out;
	}
}

package model

type YTContext struct {
	Context struct {
		Client struct {
			ClientName    string `json:"clientName"`
			ClientVersion string `json:"clientVersion"`
		} `json:"client"`
	} `json:"context"`
	Query string `json:"query"`
}

type YTSearch struct {
	ResponseContext struct {
		MaxAgeSeconds         int `json:"maxAgeSeconds"`
		ServiceTrackingParams []struct {
			Service string `json:"service"`
			Params  []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"params"`
		} `json:"serviceTrackingParams"`
	} `json:"responseContext"`
	TrackingParams string `json:"trackingParams"`
	Contents       struct {
		TabbedSearchResultsRenderer struct {
			Tabs []struct {
				TabRenderer struct {
					Title    string `json:"title"`
					Selected bool   `json:"selected"`
					Content  struct {
						SectionListRenderer struct {
							Contents []struct {
								MusicShelfRenderer struct {
									Title struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"title"`
									Contents []struct {
										MusicResponsiveListItemRenderer struct {
											TrackingParams string `json:"trackingParams"`
											Thumbnail      struct {
												MusicThumbnailRenderer struct {
													Thumbnail struct {
														Thumbnails []struct {
															Url    string `json:"url"`
															Width  int    `json:"width"`
															Height int    `json:"height"`
														} `json:"thumbnails"`
													} `json:"thumbnail"`
													ThumbnailCrop  string `json:"thumbnailCrop"`
													ThumbnailScale string `json:"thumbnailScale"`
													TrackingParams string `json:"trackingParams"`
												} `json:"musicThumbnailRenderer"`
											} `json:"thumbnail"`
											FlexColumns []struct {
												MusicResponsiveListItemFlexColumnRenderer struct {
													Text struct {
														Runs []struct {
															Text               string `json:"text"`
															NavigationEndpoint struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																BrowseEndpoint      struct {
																	BrowseId                              string `json:"browseId"`
																	BrowseEndpointContextSupportedConfigs struct {
																		BrowseEndpointContextMusicConfig struct {
																			PageType string `json:"pageType"`
																		} `json:"browseEndpointContextMusicConfig"`
																	} `json:"browseEndpointContextSupportedConfigs"`
																} `json:"browseEndpoint,omitempty"`
																WatchEndpoint struct {
																	VideoId        string `json:"videoId"`
																	PlaylistId     string `json:"playlistId"`
																	LoggingContext struct {
																		VssLoggingContext struct {
																			SerializedContextData string `json:"serializedContextData"`
																		} `json:"vssLoggingContext"`
																	} `json:"loggingContext"`
																	WatchEndpointMusicSupportedConfigs struct {
																		WatchEndpointMusicConfig struct {
																			MusicVideoType string `json:"musicVideoType"`
																		} `json:"watchEndpointMusicConfig"`
																	} `json:"watchEndpointMusicSupportedConfigs"`
																} `json:"watchEndpoint,omitempty"`
															} `json:"navigationEndpoint,omitempty"`
														} `json:"runs"`
													} `json:"text"`
													DisplayPriority string `json:"displayPriority"`
												} `json:"musicResponsiveListItemFlexColumnRenderer"`
											} `json:"flexColumns"`
											Menu struct {
												MenuRenderer struct {
													Items []struct {
														MenuNavigationItemRenderer struct {
															Text struct {
																Runs []struct {
																	Text string `json:"text"`
																} `json:"runs"`
															} `json:"text"`
															Icon struct {
																IconType string `json:"iconType"`
															} `json:"icon"`
															NavigationEndpoint struct {
																ClickTrackingParams   string `json:"clickTrackingParams"`
																WatchPlaylistEndpoint struct {
																	PlaylistId string `json:"playlistId"`
																	Params     string `json:"params"`
																} `json:"watchPlaylistEndpoint,omitempty"`
																ShareEntityEndpoint struct {
																	SerializedShareEntity string `json:"serializedShareEntity"`
																	SharePanelType        string `json:"sharePanelType"`
																} `json:"shareEntityEndpoint,omitempty"`
																ModalEndpoint struct {
																	Modal struct {
																		ModalWithTitleAndButtonRenderer struct {
																			Title struct {
																				Runs []struct {
																					Text string `json:"text"`
																				} `json:"runs"`
																			} `json:"title"`
																			Content struct {
																				Runs []struct {
																					Text string `json:"text"`
																				} `json:"runs"`
																			} `json:"content"`
																			Button struct {
																				ButtonRenderer struct {
																					Style      string `json:"style"`
																					IsDisabled bool   `json:"isDisabled"`
																					Text       struct {
																						Runs []struct {
																							Text string `json:"text"`
																						} `json:"runs"`
																					} `json:"text"`
																					NavigationEndpoint struct {
																						ClickTrackingParams string `json:"clickTrackingParams"`
																						SignInEndpoint      struct {
																							Hack bool `json:"hack"`
																						} `json:"signInEndpoint"`
																					} `json:"navigationEndpoint"`
																					TrackingParams string `json:"trackingParams"`
																				} `json:"buttonRenderer"`
																			} `json:"button"`
																		} `json:"modalWithTitleAndButtonRenderer"`
																	} `json:"modal"`
																} `json:"modalEndpoint,omitempty"`
																BrowseEndpoint struct {
																	BrowseId                              string `json:"browseId"`
																	BrowseEndpointContextSupportedConfigs struct {
																		BrowseEndpointContextMusicConfig struct {
																			PageType string `json:"pageType"`
																		} `json:"browseEndpointContextMusicConfig"`
																	} `json:"browseEndpointContextSupportedConfigs"`
																} `json:"browseEndpoint,omitempty"`
																WatchEndpoint struct {
																	VideoId        string `json:"videoId"`
																	PlaylistId     string `json:"playlistId"`
																	Params         string `json:"params"`
																	LoggingContext struct {
																		VssLoggingContext struct {
																			SerializedContextData string `json:"serializedContextData"`
																		} `json:"vssLoggingContext"`
																	} `json:"loggingContext"`
																	WatchEndpointMusicSupportedConfigs struct {
																		WatchEndpointMusicConfig struct {
																			MusicVideoType string `json:"musicVideoType"`
																		} `json:"watchEndpointMusicConfig"`
																	} `json:"watchEndpointMusicSupportedConfigs"`
																} `json:"watchEndpoint,omitempty"`
															} `json:"navigationEndpoint"`
															TrackingParams string `json:"trackingParams"`
														} `json:"menuNavigationItemRenderer,omitempty"`
														MenuServiceItemRenderer struct {
															Text struct {
																Runs []struct {
																	Text string `json:"text"`
																} `json:"runs"`
															} `json:"text"`
															Icon struct {
																IconType string `json:"iconType"`
															} `json:"icon"`
															ServiceEndpoint struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																QueueAddEndpoint    struct {
																	QueueTarget struct {
																		PlaylistId string `json:"playlistId,omitempty"`
																		VideoId    string `json:"videoId,omitempty"`
																	} `json:"queueTarget"`
																	QueueInsertPosition string `json:"queueInsertPosition"`
																	Commands            []struct {
																		ClickTrackingParams string `json:"clickTrackingParams"`
																		AddToToastAction    struct {
																			Item struct {
																				NotificationTextRenderer struct {
																					SuccessResponseText struct {
																						Runs []struct {
																							Text string `json:"text"`
																						} `json:"runs"`
																					} `json:"successResponseText"`
																					TrackingParams string `json:"trackingParams"`
																				} `json:"notificationTextRenderer"`
																			} `json:"item"`
																		} `json:"addToToastAction"`
																	} `json:"commands"`
																} `json:"queueAddEndpoint"`
															} `json:"serviceEndpoint"`
															TrackingParams string `json:"trackingParams"`
														} `json:"menuServiceItemRenderer,omitempty"`
														ToggleMenuServiceItemRenderer struct {
															DefaultText struct {
																Runs []struct {
																	Text string `json:"text"`
																} `json:"runs"`
															} `json:"defaultText"`
															DefaultIcon struct {
																IconType string `json:"iconType"`
															} `json:"defaultIcon"`
															DefaultServiceEndpoint struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																ModalEndpoint       struct {
																	Modal struct {
																		ModalWithTitleAndButtonRenderer struct {
																			Title struct {
																				Runs []struct {
																					Text string `json:"text"`
																				} `json:"runs"`
																			} `json:"title"`
																			Content struct {
																				Runs []struct {
																					Text string `json:"text"`
																				} `json:"runs"`
																			} `json:"content"`
																			Button struct {
																				ButtonRenderer struct {
																					Style      string `json:"style"`
																					IsDisabled bool   `json:"isDisabled"`
																					Text       struct {
																						Runs []struct {
																							Text string `json:"text"`
																						} `json:"runs"`
																					} `json:"text"`
																					NavigationEndpoint struct {
																						ClickTrackingParams string `json:"clickTrackingParams"`
																						SignInEndpoint      struct {
																							Hack bool `json:"hack"`
																						} `json:"signInEndpoint"`
																					} `json:"navigationEndpoint"`
																					TrackingParams string `json:"trackingParams"`
																				} `json:"buttonRenderer"`
																			} `json:"button"`
																		} `json:"modalWithTitleAndButtonRenderer"`
																	} `json:"modal"`
																} `json:"modalEndpoint"`
															} `json:"defaultServiceEndpoint"`
															ToggledText struct {
																Runs []struct {
																	Text string `json:"text"`
																} `json:"runs"`
															} `json:"toggledText"`
															ToggledIcon struct {
																IconType string `json:"iconType"`
															} `json:"toggledIcon"`
															ToggledServiceEndpoint struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																LikeEndpoint        struct {
																	Status string `json:"status"`
																	Target struct {
																		PlaylistId string `json:"playlistId"`
																	} `json:"target"`
																} `json:"likeEndpoint"`
															} `json:"toggledServiceEndpoint,omitempty"`
															TrackingParams string `json:"trackingParams"`
														} `json:"toggleMenuServiceItemRenderer,omitempty"`
													} `json:"items"`
													TrackingParams string `json:"trackingParams"`
													Accessibility  struct {
														AccessibilityData struct {
															Label string `json:"label"`
														} `json:"accessibilityData"`
													} `json:"accessibility"`
												} `json:"menuRenderer"`
											} `json:"menu"`
											FlexColumnDisplayStyle string `json:"flexColumnDisplayStyle"`
											NavigationEndpoint     struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												BrowseEndpoint      struct {
													BrowseId                              string `json:"browseId"`
													BrowseEndpointContextSupportedConfigs struct {
														BrowseEndpointContextMusicConfig struct {
															PageType string `json:"pageType"`
														} `json:"browseEndpointContextMusicConfig"`
													} `json:"browseEndpointContextSupportedConfigs"`
												} `json:"browseEndpoint"`
											} `json:"navigationEndpoint,omitempty"`
											ItemHeight string `json:"itemHeight"`
											Overlay    struct {
												MusicItemThumbnailOverlayRenderer struct {
													Background struct {
														VerticalGradient struct {
															GradientLayerColors []string `json:"gradientLayerColors"`
														} `json:"verticalGradient"`
													} `json:"background"`
													Content struct {
														MusicPlayButtonRenderer struct {
															PlayNavigationEndpoint struct {
																ClickTrackingParams   string `json:"clickTrackingParams"`
																WatchPlaylistEndpoint struct {
																	PlaylistId string `json:"playlistId"`
																	Params     string `json:"params,omitempty"`
																} `json:"watchPlaylistEndpoint,omitempty"`
																WatchEndpoint struct {
																	VideoId        string `json:"videoId"`
																	PlaylistId     string `json:"playlistId"`
																	Params         string `json:"params"`
																	LoggingContext struct {
																		VssLoggingContext struct {
																			SerializedContextData string `json:"serializedContextData"`
																		} `json:"vssLoggingContext"`
																	} `json:"loggingContext"`
																	WatchEndpointMusicSupportedConfigs struct {
																		WatchEndpointMusicConfig struct {
																			MusicVideoType string `json:"musicVideoType"`
																		} `json:"watchEndpointMusicConfig"`
																	} `json:"watchEndpointMusicSupportedConfigs"`
																} `json:"watchEndpoint,omitempty"`
															} `json:"playNavigationEndpoint"`
															TrackingParams string `json:"trackingParams"`
															PlayIcon       struct {
																IconType string `json:"iconType"`
															} `json:"playIcon"`
															PauseIcon struct {
																IconType string `json:"iconType"`
															} `json:"pauseIcon"`
															IconColor             int64 `json:"iconColor"`
															BackgroundColor       int   `json:"backgroundColor"`
															ActiveBackgroundColor int   `json:"activeBackgroundColor"`
															LoadingIndicatorColor int64 `json:"loadingIndicatorColor"`
															PlayingIcon           struct {
																IconType string `json:"iconType"`
															} `json:"playingIcon"`
															IconLoadingColor      int    `json:"iconLoadingColor"`
															ActiveScaleFactor     int    `json:"activeScaleFactor"`
															ButtonSize            string `json:"buttonSize"`
															RippleTarget          string `json:"rippleTarget"`
															AccessibilityPlayData struct {
																AccessibilityData struct {
																	Label string `json:"label"`
																} `json:"accessibilityData"`
															} `json:"accessibilityPlayData"`
															AccessibilityPauseData struct {
																AccessibilityData struct {
																	Label string `json:"label"`
																} `json:"accessibilityData"`
															} `json:"accessibilityPauseData"`
														} `json:"musicPlayButtonRenderer"`
													} `json:"content"`
													ContentPosition string `json:"contentPosition"`
													DisplayStyle    string `json:"displayStyle"`
												} `json:"musicItemThumbnailOverlayRenderer"`
											} `json:"overlay,omitempty"`
											PlaylistItemData struct {
												VideoId string `json:"videoId"`
											} `json:"playlistItemData,omitempty"`
										} `json:"musicResponsiveListItemRenderer"`
									} `json:"contents"`
									TrackingParams string `json:"trackingParams"`
									ShelfDivider   struct {
										MusicShelfDividerRenderer struct {
											Hidden bool `json:"hidden"`
										} `json:"musicShelfDividerRenderer"`
									} `json:"shelfDivider"`
									BottomText struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"bottomText,omitempty"`
									BottomEndpoint struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										SearchEndpoint      struct {
											Query  string `json:"query"`
											Params string `json:"params"`
										} `json:"searchEndpoint"`
									} `json:"bottomEndpoint,omitempty"`
								} `json:"musicShelfRenderer"`
							} `json:"contents"`
							TrackingParams string `json:"trackingParams"`
							Header         struct {
								ChipCloudRenderer struct {
									Chips []struct {
										ChipCloudChipRenderer struct {
											Style struct {
												StyleType string `json:"styleType"`
											} `json:"style"`
											Text struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"text"`
											NavigationEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												SearchEndpoint      struct {
													Query  string `json:"query"`
													Params string `json:"params"`
												} `json:"searchEndpoint"`
											} `json:"navigationEndpoint"`
											TrackingParams    string `json:"trackingParams"`
											AccessibilityData struct {
												AccessibilityData struct {
													Label string `json:"label"`
												} `json:"accessibilityData"`
											} `json:"accessibilityData"`
											IsSelected bool `json:"isSelected"`
										} `json:"chipCloudChipRenderer"`
									} `json:"chips"`
									CollapsedRowCount    int    `json:"collapsedRowCount"`
									TrackingParams       string `json:"trackingParams"`
									HorizontalScrollable bool   `json:"horizontalScrollable"`
								} `json:"chipCloudRenderer"`
							} `json:"header"`
						} `json:"sectionListRenderer"`
					} `json:"content"`
					TabIdentifier  string `json:"tabIdentifier"`
					TrackingParams string `json:"trackingParams"`
				} `json:"tabRenderer"`
			} `json:"tabs"`
		} `json:"tabbedSearchResultsRenderer"`
	} `json:"contents"`
}

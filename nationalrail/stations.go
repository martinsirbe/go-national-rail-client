package nationalrail

var StationCodeToNameMap = map[CRSCode]string{
	StationCodeAbbeyWood:                         StationNameAbbeyWood,
	StationCodeAber:                              StationNameAber,
	StationCodeAbercynon:                         StationNameAbercynon,
	StationCodeAberdare:                          StationNameAberdare,
	StationCodeAberdeen:                          StationNameAberdeen,
	StationCodeAberdour:                          StationNameAberdour,
	StationCodeAberdovey:                         StationNameAberdovey,
	StationCodeAbererch:                          StationNameAbererch,
	StationCodeAbergavenny:                       StationNameAbergavenny,
	StationCodeAbergeleAndPensarn:                StationNameAbergeleAndPensarn,
	StationCodeAberystwyth:                       StationNameAberystwyth,
	StationCodeAccrington:                        StationNameAccrington,
	StationCodeAchanalt:                          StationNameAchanalt,
	StationCodeAchnasheen:                        StationNameAchnasheen,
	StationCodeAchnashellach:                     StationNameAchnashellach,
	StationCodeAcklington:                        StationNameAcklington,
	StationCodeAcle:                              StationNameAcle,
	StationCodeAcocksGreen:                       StationNameAcocksGreen,
	StationCodeActonBridgeCheshire:               StationNameActonBridgeCheshire,
	StationCodeActonCentral:                      StationNameActonCentral,
	StationCodeActonMainLine:                     StationNameActonMainLine,
	StationCodeAdderleyPark:                      StationNameAdderleyPark,
	StationCodeAddiewell:                         StationNameAddiewell,
	StationCodeAddlestone:                        StationNameAddlestone,
	StationCodeAdisham:                           StationNameAdisham,
	StationCodeAdlingtonCheshire:                 StationNameAdlingtonCheshire,
	StationCodeAdlingtonLancs:                    StationNameAdlingtonLancs,
	StationCodeAdwick:                            StationNameAdwick,
	StationCodeAigburth:                          StationNameAigburth,
	StationCodeAinsdale:                          StationNameAinsdale,
	StationCodeAintree:                           StationNameAintree,
	StationCodeAirbles:                           StationNameAirbles,
	StationCodeAirdrie:                           StationNameAirdrie,
	StationCodeAlbanyPark:                        StationNameAlbanyPark,
	StationCodeAlbrighton:                        StationNameAlbrighton,
	StationCodeAlderleyEdge:                      StationNameAlderleyEdge,
	StationCodeAldermaston:                       StationNameAldermaston,
	StationCodeAldershot:                         StationNameAldershot,
	StationCodeAldrington:                        StationNameAldrington,
	StationCodeAlexandraPalace:                   StationNameAlexandraPalace,
	StationCodeAlexandraParade:                   StationNameAlexandraParade,
	StationCodeAlexandria:                        StationNameAlexandria,
	StationCodeAlfreton:                          StationNameAlfreton,
	StationCodeAllensWest:                        StationNameAllensWest,
	StationCodeAlloa:                             StationNameAlloa,
	StationCodeAlness:                            StationNameAlness,
	StationCodeAlnmouth:                          StationNameAlnmouth,
	StationCodeAlresfordEssex:                    StationNameAlresfordEssex,
	StationCodeAlsager:                           StationNameAlsager,
	StationCodeAlthorneEssex:                     StationNameAlthorneEssex,
	StationCodeAlthorpe:                          StationNameAlthorpe,
	StationCodeAltnabreac:                        StationNameAltnabreac,
	StationCodeAlton:                             StationNameAlton,
	StationCodeAltrincham:                        StationNameAltrincham,
	StationCodeAlvechurch:                        StationNameAlvechurch,
	StationCodeAmbergate:                         StationNameAmbergate,
	StationCodeAmberley:                          StationNameAmberley,
	StationCodeAmersham:                          StationNameAmersham,
	StationCodeAmmanford:                         StationNameAmmanford,
	StationCodeAncaster:                          StationNameAncaster,
	StationCodeAnderston:                         StationNameAnderston,
	StationCodeAndover:                           StationNameAndover,
	StationCodeAnerley:                           StationNameAnerley,
	StationCodeAngelRoad:                         StationNameAngelRoad,
	StationCodeAngmering:                         StationNameAngmering,
	StationCodeAnnan:                             StationNameAnnan,
	StationCodeAnniesland:                        StationNameAnniesland,
	StationCodeAnsdellAndFairhaven:               StationNameAnsdellAndFairhaven,
	StationCodeApperleyBridge:                    StationNameApperleyBridge,
	StationCodeAppleby:                           StationNameAppleby,
	StationCodeAppledoreKent:                     StationNameAppledoreKent,
	StationCodeAppleford:                         StationNameAppleford,
	StationCodeAppleyBridge:                      StationNameAppleyBridge,
	StationCodeApsley:                            StationNameApsley,
	StationCodeArbroath:                          StationNameArbroath,
	StationCodeArdgay:                            StationNameArdgay,
	StationCodeArdlui:                            StationNameArdlui,
	StationCodeArdrossanHarbour:                  StationNameArdrossanHarbour,
	StationCodeArdrossanSouthBeach:               StationNameArdrossanSouthBeach,
	StationCodeArdrossanTown:                     StationNameArdrossanTown,
	StationCodeArdwick:                           StationNameArdwick,
	StationCodeArgyleStreet:                      StationNameArgyleStreet,
	StationCodeArisaig:                           StationNameArisaig,
	StationCodeArlesey:                           StationNameArlesey,
	StationCodeArmadaleWestLothian:               StationNameArmadaleWestLothian,
	StationCodeArmathwaite:                       StationNameArmathwaite,
	StationCodeArnside:                           StationNameArnside,
	StationCodeArram:                             StationNameArram,
	StationCodeArrocharAndTarbet:                 StationNameArrocharAndTarbet,
	StationCodeArundel:                           StationNameArundel,
	StationCodeAscotBerks:                        StationNameAscotBerks,
	StationCodeAscottunderWychwood:               StationNameAscottunderWychwood,
	StationCodeAsh:                               StationNameAsh,
	StationCodeAshVale:                           StationNameAshVale,
	StationCodeAshburys:                          StationNameAshburys,
	StationCodeAshchurchforTewkesbury:            StationNameAshchurchforTewkesbury,
	StationCodeAshfield:                          StationNameAshfield,
	StationCodeAshfordInternational:              StationNameAshfordInternational,
	StationCodeAshfordInternationalEurostar:      StationNameAshfordInternationalEurostar,
	StationCodeAshfordSurrey:                     StationNameAshfordSurrey,
	StationCodeAshley:                            StationNameAshley,
	StationCodeAshtead:                           StationNameAshtead,
	StationCodeAshtonunderLyne:                   StationNameAshtonunderLyne,
	StationCodeAshurstKent:                       StationNameAshurstKent,
	StationCodeAshurstNewForest:                  StationNameAshurstNewForest,
	StationCodeAshwellAndMorden:                  StationNameAshwellAndMorden,
	StationCodeAskam:                             StationNameAskam,
	StationCodeAslockton:                         StationNameAslockton,
	StationCodeAspatria:                          StationNameAspatria,
	StationCodeAspleyGuise:                       StationNameAspleyGuise,
	StationCodeAston:                             StationNameAston,
	StationCodeAtherstone:                        StationNameAtherstone,
	StationCodeAtherton:                          StationNameAtherton,
	StationCodeAttadale:                          StationNameAttadale,
	StationCodeAttenborough:                      StationNameAttenborough,
	StationCodeAttleborough:                      StationNameAttleborough,
	StationCodeAuchinleck:                        StationNameAuchinleck,
	StationCodeAudleyEnd:                         StationNameAudleyEnd,
	StationCodeAughtonPark:                       StationNameAughtonPark,
	StationCodeAviemore:                          StationNameAviemore,
	StationCodeAvoncliff:                         StationNameAvoncliff,
	StationCodeAvonmouth:                         StationNameAvonmouth,
	StationCodeAxminster:                         StationNameAxminster,
	StationCodeAylesbury:                         StationNameAylesbury,
	StationCodeAylesburyValeParkway:              StationNameAylesburyValeParkway,
	StationCodeAylesford:                         StationNameAylesford,
	StationCodeAylesham:                          StationNameAylesham,
	StationCodeAyr:                               StationNameAyr,
	StationCodeBache:                             StationNameBache,
	StationCodeBaglan:                            StationNameBaglan,
	StationCodeBagshot:                           StationNameBagshot,
	StationCodeBaildon:                           StationNameBaildon,
	StationCodeBaillieston:                       StationNameBaillieston,
	StationCodeBalcombe:                          StationNameBalcombe,
	StationCodeBaldock:                           StationNameBaldock,
	StationCodeBalham:                            StationNameBalham,
	StationCodeBalloch:                           StationNameBalloch,
	StationCodeBalmossie:                         StationNameBalmossie,
	StationCodeBamberBridge:                      StationNameBamberBridge,
	StationCodeBamford:                           StationNameBamford,
	StationCodeBanavie:                           StationNameBanavie,
	StationCodeBanbury:                           StationNameBanbury,
	StationCodeBangorGwynedd:                     StationNameBangorGwynedd,
	StationCodeBankHall:                          StationNameBankHall,
	StationCodeBanstead:                          StationNameBanstead,
	StationCodeBarassie:                          StationNameBarassie,
	StationCodeBarbican:                          StationNameBarbican,
	StationCodeBardonMill:                        StationNameBardonMill,
	StationCodeBareLane:                          StationNameBareLane,
	StationCodeBargeddie:                         StationNameBargeddie,
	StationCodeBargoed:                           StationNameBargoed,
	StationCodeBarking:                           StationNameBarking,
	StationCodeBarlaston:                         StationNameBarlaston,
	StationCodeBarming:                           StationNameBarming,
	StationCodeBarmouth:                          StationNameBarmouth,
	StationCodeBarnehurst:                        StationNameBarnehurst,
	StationCodeBarnes:                            StationNameBarnes,
	StationCodeBarnesBridge:                      StationNameBarnesBridge,
	StationCodeBarnetby:                          StationNameBarnetby,
	StationCodeBarnham:                           StationNameBarnham,
	StationCodeBarnhill:                          StationNameBarnhill,
	StationCodeBarnsley:                          StationNameBarnsley,
	StationCodeBarnstaple:                        StationNameBarnstaple,
	StationCodeBarntGreen:                        StationNameBarntGreen,
	StationCodeBarrhead:                          StationNameBarrhead,
	StationCodeBarrhill:                          StationNameBarrhill,
	StationCodeBarrowHaven:                       StationNameBarrowHaven,
	StationCodeBarrowUponSoar:                    StationNameBarrowUponSoar,
	StationCodeBarrowinFurness:                   StationNameBarrowinFurness,
	StationCodeBarry:                             StationNameBarry,
	StationCodeBarryDocks:                        StationNameBarryDocks,
	StationCodeBarryIsland:                       StationNameBarryIsland,
	StationCodeBarryLinks:                        StationNameBarryLinks,
	StationCodeBartononHumber:                    StationNameBartononHumber,
	StationCodeBasildon:                          StationNameBasildon,
	StationCodeBasingstoke:                       StationNameBasingstoke,
	StationCodeBatAndBall:                        StationNameBatAndBall,
	StationCodeBathSpa:                           StationNameBathSpa,
	StationCodeBathgate:                          StationNameBathgate,
	StationCodeBatley:                            StationNameBatley,
	StationCodeBattersby:                         StationNameBattersby,
	StationCodeBatterseaPark:                     StationNameBatterseaPark,
	StationCodeBattle:                            StationNameBattle,
	StationCodeBattlesbridge:                     StationNameBattlesbridge,
	StationCodeBayford:                           StationNameBayford,
	StationCodeBeaconsfield:                      StationNameBeaconsfield,
	StationCodeBearley:                           StationNameBearley,
	StationCodeBearsden:                          StationNameBearsden,
	StationCodeBearsted:                          StationNameBearsted,
	StationCodeBeasdale:                          StationNameBeasdale,
	StationCodeBeaulieuRoad:                      StationNameBeaulieuRoad,
	StationCodeBeauly:                            StationNameBeauly,
	StationCodeBebington:                         StationNameBebington,
	StationCodeBeccles:                           StationNameBeccles,
	StationCodeBeckenhamHill:                     StationNameBeckenhamHill,
	StationCodeBeckenhamJunction:                 StationNameBeckenhamJunction,
	StationCodeBedford:                           StationNameBedford,
	StationCodeBedfordStJohns:                    StationNameBedfordStJohns,
	StationCodeBedhampton:                        StationNameBedhampton,
	StationCodeBedminster:                        StationNameBedminster,
	StationCodeBedworth:                          StationNameBedworth,
	StationCodeBedwyn:                            StationNameBedwyn,
	StationCodeBeeston:                           StationNameBeeston,
	StationCodeBekesbourne:                       StationNameBekesbourne,
	StationCodeBelleVue:                          StationNameBelleVue,
	StationCodeBellgrove:                         StationNameBellgrove,
	StationCodeBellingham:                        StationNameBellingham,
	StationCodeBellshill:                         StationNameBellshill,
	StationCodeBelmont:                           StationNameBelmont,
	StationCodeBelper:                            StationNameBelper,
	StationCodeBeltring:                          StationNameBeltring,
	StationCodeBelvedere:                         StationNameBelvedere,
	StationCodeBempton:                           StationNameBempton,
	StationCodeBenRhydding:                       StationNameBenRhydding,
	StationCodeBenfleet:                          StationNameBenfleet,
	StationCodeBentham:                           StationNameBentham,
	StationCodeBentleyHants:                      StationNameBentleyHants,
	StationCodeBentleySouthYorks:                 StationNameBentleySouthYorks,
	StationCodeBereAlston:                        StationNameBereAlston,
	StationCodeBereFerrers:                       StationNameBereFerrers,
	StationCodeBerkhamsted:                       StationNameBerkhamsted,
	StationCodeBerkswell:                         StationNameBerkswell,
	StationCodeBermudaPark:                       StationNameBermudaPark,
	StationCodeBerneyArms:                        StationNameBerneyArms,
	StationCodeBerryBrow:                         StationNameBerryBrow,
	StationCodeBerrylands:                        StationNameBerrylands,
	StationCodeBerwickSussex:                     StationNameBerwickSussex,
	StationCodeBerwickUponTweed:                  StationNameBerwickUponTweed,
	StationCodeBescarLane:                        StationNameBescarLane,
	StationCodeBescotStadium:                     StationNameBescotStadium,
	StationCodeBetchworth:                        StationNameBetchworth,
	StationCodeBethnalGreen:                      StationNameBethnalGreen,
	StationCodeBetwsyCoed:                        StationNameBetwsyCoed,
	StationCodeBeverley:                          StationNameBeverley,
	StationCodeBexhill:                           StationNameBexhill,
	StationCodeBexley:                            StationNameBexley,
	StationCodeBexleyheath:                       StationNameBexleyheath,
	StationCodeBicesterNorth:                     StationNameBicesterNorth,
	StationCodeBicesterVillage:                   StationNameBicesterVillage,
	StationCodeBickley:                           StationNameBickley,
	StationCodeBidston:                           StationNameBidston,
	StationCodeBiggleswade:                       StationNameBiggleswade,
	StationCodeBilbrook:                          StationNameBilbrook,
	StationCodeBillericay:                        StationNameBillericay,
	StationCodeBillinghamCleveland:               StationNameBillinghamCleveland,
	StationCodeBillingshurst:                     StationNameBillingshurst,
	StationCodeBingham:                           StationNameBingham,
	StationCodeBingley:                           StationNameBingley,
	StationCodeBirchgrove:                        StationNameBirchgrove,
	StationCodeBirchingtonOnSea:                  StationNameBirchingtonOnSea,
	StationCodeBirchwood:                         StationNameBirchwood,
	StationCodeBirkbeck:                          StationNameBirkbeck,
	StationCodeBirkdale:                          StationNameBirkdale,
	StationCodeBirkenheadCentral:                 StationNameBirkenheadCentral,
	StationCodeBirkenheadHamiltonSquare:          StationNameBirkenheadHamiltonSquare,
	StationCodeBirkenheadNorth:                   StationNameBirkenheadNorth,
	StationCodeBirkenheadPark:                    StationNameBirkenheadPark,
	StationCodeBirminghamInternational:           StationNameBirminghamInternational,
	StationCodeBirminghamMoorStreet:              StationNameBirminghamMoorStreet,
	StationCodeBirminghamNewStreet:               StationNameBirminghamNewStreet,
	StationCodeBirminghamSnowHill:                StationNameBirminghamSnowHill,
	StationCodeBishopAuckland:                    StationNameBishopAuckland,
	StationCodeBishopbriggs:                      StationNameBishopbriggs,
	StationCodeBishopsStortford:                  StationNameBishopsStortford,
	StationCodeBishopstoneSussex:                 StationNameBishopstoneSussex,
	StationCodeBishoptonStrathclyde:              StationNameBishoptonStrathclyde,
	StationCodeBitterne:                          StationNameBitterne,
	StationCodeBlackburn:                         StationNameBlackburn,
	StationCodeBlackheath:                        StationNameBlackheath,
	StationCodeBlackhorseRoad:                    StationNameBlackhorseRoad,
	StationCodeBlackpoolNorth:                    StationNameBlackpoolNorth,
	StationCodeBlackpoolPleasureBeach:            StationNameBlackpoolPleasureBeach,
	StationCodeBlackpoolSouth:                    StationNameBlackpoolSouth,
	StationCodeBlackridge:                        StationNameBlackridge,
	StationCodeBlackrod:                          StationNameBlackrod,
	StationCodeBlackwater:                        StationNameBlackwater,
	StationCodeBlaenauFfestiniog:                 StationNameBlaenauFfestiniog,
	StationCodeBlairAtholl:                       StationNameBlairAtholl,
	StationCodeBlairhill:                         StationNameBlairhill,
	StationCodeBlakeStreet:                       StationNameBlakeStreet,
	StationCodeBlakedown:                         StationNameBlakedown,
	StationCodeBlantyre:                          StationNameBlantyre,
	StationCodeBlaydon:                           StationNameBlaydon,
	StationCodeBleasby:                           StationNameBleasby,
	StationCodeBletchley:                         StationNameBletchley,
	StationCodeBloxwich:                          StationNameBloxwich,
	StationCodeBloxwichNorth:                     StationNameBloxwichNorth,
	StationCodeBlundellsandsAndCrosby:            StationNameBlundellsandsAndCrosby,
	StationCodeBlytheBridge:                      StationNameBlytheBridge,
	StationCodeBodminParkway:                     StationNameBodminParkway,
	StationCodeBodorgan:                          StationNameBodorgan,
	StationCodeBognorRegis:                       StationNameBognorRegis,
	StationCodeBogston:                           StationNameBogston,
	StationCodeBolton:                            StationNameBolton,
	StationCodeBoltonUponDearne:                  StationNameBoltonUponDearne,
	StationCodeBookham:                           StationNameBookham,
	StationCodeBootleCumbria:                     StationNameBootleCumbria,
	StationCodeBootleNewStrand:                   StationNameBootleNewStrand,
	StationCodeBootleOrielRoad:                   StationNameBootleOrielRoad,
	StationCodeBordesley:                         StationNameBordesley,
	StationCodeBoroughGreenAndWrotham:            StationNameBoroughGreenAndWrotham,
	StationCodeBorth:                             StationNameBorth,
	StationCodeBosham:                            StationNameBosham,
	StationCodeBoston:                            StationNameBoston,
	StationCodeBotley:                            StationNameBotley,
	StationCodeBottesford:                        StationNameBottesford,
	StationCodeBourneEnd:                         StationNameBourneEnd,
	StationCodeBournemouth:                       StationNameBournemouth,
	StationCodeBournville:                        StationNameBournville,
	StationCodeBowBrickhill:                      StationNameBowBrickhill,
	StationCodeBowesPark:                         StationNameBowesPark,
	StationCodeBowling:                           StationNameBowling,
	StationCodeBoxHillAndWesthumble:              StationNameBoxHillAndWesthumble,
	StationCodeBracknell:                         StationNameBracknell,
	StationCodeBradfordForsterSquare:             StationNameBradfordForsterSquare,
	StationCodeBradfordInterchange:               StationNameBradfordInterchange,
	StationCodeBradfordonAvon:                    StationNameBradfordonAvon,
	StationCodeBrading:                           StationNameBrading,
	StationCodeBraintree:                         StationNameBraintree,
	StationCodeBraintreeFreeport:                 StationNameBraintreeFreeport,
	StationCodeBramhall:                          StationNameBramhall,
	StationCodeBramleyHants:                      StationNameBramleyHants,
	StationCodeBramleyWYorks:                     StationNameBramleyWYorks,
	StationCodeBramptonCumbria:                   StationNameBramptonCumbria,
	StationCodeBramptonSuffolk:                   StationNameBramptonSuffolk,
	StationCodeBranchton:                         StationNameBranchton,
	StationCodeBrandon:                           StationNameBrandon,
	StationCodeBranksome:                         StationNameBranksome,
	StationCodeBraystonesCumbria:                 StationNameBraystonesCumbria,
	StationCodeBredbury:                          StationNameBredbury,
	StationCodeBreich:                            StationNameBreich,
	StationCodeBrentford:                         StationNameBrentford,
	StationCodeBrentwood:                         StationNameBrentwood,
	StationCodeBricketWood:                       StationNameBricketWood,
	StationCodeBridgend:                          StationNameBridgend,
	StationCodeBridgeofAllan:                     StationNameBridgeofAllan,
	StationCodeBridgeofOrchy:                     StationNameBridgeofOrchy,
	StationCodeBridgeton:                         StationNameBridgeton,
	StationCodeBridgwater:                        StationNameBridgwater,
	StationCodeBridlington:                       StationNameBridlington,
	StationCodeBrierfield:                        StationNameBrierfield,
	StationCodeBrigg:                             StationNameBrigg,
	StationCodeBrighouse:                         StationNameBrighouse,
	StationCodeBrightonEastSussex:                StationNameBrightonEastSussex,
	StationCodeBrimsdown:                         StationNameBrimsdown,
	StationCodeBrinnington:                       StationNameBrinnington,
	StationCodeBristolParkway:                    StationNameBristolParkway,
	StationCodeBristolTempleMeads:                StationNameBristolTempleMeads,
	StationCodeBrithdir:                          StationNameBrithdir,
	StationCodeBritonFerry:                       StationNameBritonFerry,
	StationCodeBrixton:                           StationNameBrixton,
	StationCodeBroadGreen:                        StationNameBroadGreen,
	StationCodeBroadbottom:                       StationNameBroadbottom,
	StationCodeBroadstairs:                       StationNameBroadstairs,
	StationCodeBrockenhurst:                      StationNameBrockenhurst,
	StationCodeBrockholes:                        StationNameBrockholes,
	StationCodeBrockley:                          StationNameBrockley,
	StationCodeBromborough:                       StationNameBromborough,
	StationCodeBromboroughRake:                   StationNameBromboroughRake,
	StationCodeBromleyCrossLancs:                 StationNameBromleyCrossLancs,
	StationCodeBromleyNorth:                      StationNameBromleyNorth,
	StationCodeBromleySouth:                      StationNameBromleySouth,
	StationCodeBromsgrove:                        StationNameBromsgrove,
	StationCodeBrondesbury:                       StationNameBrondesbury,
	StationCodeBrondesburyPark:                   StationNameBrondesburyPark,
	StationCodeBrookmansPark:                     StationNameBrookmansPark,
	StationCodeBrookwood:                         StationNameBrookwood,
	StationCodeBroome:                            StationNameBroome,
	StationCodeBroomfleet:                        StationNameBroomfleet,
	StationCodeBrora:                             StationNameBrora,
	StationCodeBrough:                            StationNameBrough,
	StationCodeBroughtyFerry:                     StationNameBroughtyFerry,
	StationCodeBroxbourne:                        StationNameBroxbourne,
	StationCodeBruceGrove:                        StationNameBruceGrove,
	StationCodeBrundall:                          StationNameBrundall,
	StationCodeBrundallGardens:                   StationNameBrundallGardens,
	StationCodeBrunstane:                         StationNameBrunstane,
	StationCodeBrunswick:                         StationNameBrunswick,
	StationCodeBruton:                            StationNameBruton,
	StationCodeBryn:                              StationNameBryn,
	StationCodeBuckenhamNorfolk:                  StationNameBuckenhamNorfolk,
	StationCodeBuckley:                           StationNameBuckley,
	StationCodeBucknell:                          StationNameBucknell,
	StationCodeBuckshawParkway:                   StationNameBuckshawParkway,
	StationCodeBugle:                             StationNameBugle,
	StationCodeBuilthRoad:                        StationNameBuilthRoad,
	StationCodeBulwell:                           StationNameBulwell,
	StationCodeBures:                             StationNameBures,
	StationCodeBurgessHill:                       StationNameBurgessHill,
	StationCodeBurleyPark:                        StationNameBurleyPark,
	StationCodeBurleyinWharfedale:                StationNameBurleyinWharfedale,
	StationCodeBurnage:                           StationNameBurnage,
	StationCodeBurnesideCumbria:                  StationNameBurnesideCumbria,
	StationCodeBurnhamBucks:                      StationNameBurnhamBucks,
	StationCodeBurnhamonCrouch:                   StationNameBurnhamonCrouch,
	StationCodeBurnleyBarracks:                   StationNameBurnleyBarracks,
	StationCodeBurnleyCentral:                    StationNameBurnleyCentral,
	StationCodeBurnleyManchesterRoad:             StationNameBurnleyManchesterRoad,
	StationCodeBurnsideStrathclyde:               StationNameBurnsideStrathclyde,
	StationCodeBurntisland:                       StationNameBurntisland,
	StationCodeBurscoughBridge:                   StationNameBurscoughBridge,
	StationCodeBurscoughJunction:                 StationNameBurscoughJunction,
	StationCodeBursledon:                         StationNameBursledon,
	StationCodeBurtonJoyce:                       StationNameBurtonJoyce,
	StationCodeBurtononTrent:                     StationNameBurtononTrent,
	StationCodeBuryStEdmunds:                     StationNameBuryStEdmunds,
	StationCodeBusby:                             StationNameBusby,
	StationCodeBushHillPark:                      StationNameBushHillPark,
	StationCodeBushey:                            StationNameBushey,
	StationCodeButlersLane:                       StationNameButlersLane,
	StationCodeBuxted:                            StationNameBuxted,
	StationCodeBuxton:                            StationNameBuxton,
	StationCodeByfleetAndNewHaw:                  StationNameByfleetAndNewHaw,
	StationCodeBynea:                             StationNameBynea,
	StationCodeCadoxton:                          StationNameCadoxton,
	StationCodeCaergwrle:                         StationNameCaergwrle,
	StationCodeCaerphilly:                        StationNameCaerphilly,
	StationCodeCaersws:                           StationNameCaersws,
	StationCodeCaldercruix:                       StationNameCaldercruix,
	StationCodeCaldicot:                          StationNameCaldicot,
	StationCodeCaledonianRdAndBarnsbury:          StationNameCaledonianRdAndBarnsbury,
	StationCodeCalstock:                          StationNameCalstock,
	StationCodeCamAndDursley:                     StationNameCamAndDursley,
	StationCodeCamberley:                         StationNameCamberley,
	StationCodeCamborne:                          StationNameCamborne,
	StationCodeCambridge:                         StationNameCambridge,
	StationCodeCambridgeHeath:                    StationNameCambridgeHeath,
	StationCodeCambuslang:                        StationNameCambuslang,
	StationCodeCamdenRoad:                        StationNameCamdenRoad,
	StationCodeCamelon:                           StationNameCamelon,
	StationCodeCanadaWater:                       StationNameCanadaWater,
	StationCodeCanley:                            StationNameCanley,
	StationCodeCannock:                           StationNameCannock,
	StationCodeCanonbury:                         StationNameCanonbury,
	StationCodeCanterburyEast:                    StationNameCanterburyEast,
	StationCodeCanterburyWest:                    StationNameCanterburyWest,
	StationCodeCantley:                           StationNameCantley,
	StationCodeCapenhurst:                        StationNameCapenhurst,
	StationCodeCarbisBay:                         StationNameCarbisBay,
	StationCodeCardenden:                         StationNameCardenden,
	StationCodeCardiffBay:                        StationNameCardiffBay,
	StationCodeCardiffCentral:                    StationNameCardiffCentral,
	StationCodeCardiffQueenStreet:                StationNameCardiffQueenStreet,
	StationCodeCardonald:                         StationNameCardonald,
	StationCodeCardross:                          StationNameCardross,
	StationCodeCarfin:                            StationNameCarfin,
	StationCodeCarkAndCartmel:                    StationNameCarkAndCartmel,
	StationCodeCarlisle:                          StationNameCarlisle,
	StationCodeCarlton:                           StationNameCarlton,
	StationCodeCarluke:                           StationNameCarluke,
	StationCodeCarmarthen:                        StationNameCarmarthen,
	StationCodeCarmyle:                           StationNameCarmyle,
	StationCodeCarnforth:                         StationNameCarnforth,
	StationCodeCarnoustie:                        StationNameCarnoustie,
	StationCodeCarntyne:                          StationNameCarntyne,
	StationCodeCarpendersPark:                    StationNameCarpendersPark,
	StationCodeCarrbridge:                        StationNameCarrbridge,
	StationCodeCarshalton:                        StationNameCarshalton,
	StationCodeCarshaltonBeeches:                 StationNameCarshaltonBeeches,
	StationCodeCarstairs:                         StationNameCarstairs,
	StationCodeCartsdyke:                         StationNameCartsdyke,
	StationCodeCastleBarPark:                     StationNameCastleBarPark,
	StationCodeCastleCary:                        StationNameCastleCary,
	StationCodeCastleford:                        StationNameCastleford,
	StationCodeCastletonManchester:               StationNameCastletonManchester,
	StationCodeCastletonMoor:                     StationNameCastletonMoor,
	StationCodeCaterham:                          StationNameCaterham,
	StationCodeCatford:                           StationNameCatford,
	StationCodeCatfordBridge:                     StationNameCatfordBridge,
	StationCodeCathays:                           StationNameCathays,
	StationCodeCathcart:                          StationNameCathcart,
	StationCodeCattal:                            StationNameCattal,
	StationCodeCauseland:                         StationNameCauseland,
	StationCodeCefnyBedd:                         StationNameCefnyBedd,
	StationCodeChadwellHeath:                     StationNameChadwellHeath,
	StationCodeChaffordHundredLakeside:           StationNameChaffordHundredLakeside,
	StationCodeChalfontAndLatimer:                StationNameChalfontAndLatimer,
	StationCodeChalkwell:                         StationNameChalkwell,
	StationCodeChandlersFord:                     StationNameChandlersFord,
	StationCodeChapelenleFrith:                   StationNameChapelenleFrith,
	StationCodeChapeltonDevon:                    StationNameChapeltonDevon,
	StationCodeChapeltownSouthYorks:              StationNameChapeltownSouthYorks,
	StationCodeChappelAndWakesColne:              StationNameChappelAndWakesColne,
	StationCodeCharingCrossGlasgow:               StationNameCharingCrossGlasgow,
	StationCodeCharingKent:                       StationNameCharingKent,
	StationCodeCharlbury:                         StationNameCharlbury,
	StationCodeCharlton:                          StationNameCharlton,
	StationCodeChartham:                          StationNameChartham,
	StationCodeChassenRoad:                       StationNameChassenRoad,
	StationCodeChatelherault:                     StationNameChatelherault,
	StationCodeChatham:                           StationNameChatham,
	StationCodeChathill:                          StationNameChathill,
	StationCodeCheadleHulme:                      StationNameCheadleHulme,
	StationCodeCheam:                             StationNameCheam,
	StationCodeCheddington:                       StationNameCheddington,
	StationCodeChelfordCheshire:                  StationNameChelfordCheshire,
	StationCodeChelmsford:                        StationNameChelmsford,
	StationCodeChelsfield:                        StationNameChelsfield,
	StationCodeCheltenhamSpa:                     StationNameCheltenhamSpa,
	StationCodeChepstow:                          StationNameChepstow,
	StationCodeCherryTree:                        StationNameCherryTree,
	StationCodeChertsey:                          StationNameChertsey,
	StationCodeCheshunt:                          StationNameCheshunt,
	StationCodeChessingtonNorth:                  StationNameChessingtonNorth,
	StationCodeChessingtonSouth:                  StationNameChessingtonSouth,
	StationCodeChester:                           StationNameChester,
	StationCodeChesterRoad:                       StationNameChesterRoad,
	StationCodeChesterfield:                      StationNameChesterfield,
	StationCodeChesterleStreet:                   StationNameChesterleStreet,
	StationCodeChestfieldAndSwalecliffe:          StationNameChestfieldAndSwalecliffe,
	StationCodeChetnole:                          StationNameChetnole,
	StationCodeChichester:                        StationNameChichester,
	StationCodeChilham:                           StationNameChilham,
	StationCodeChilworth:                         StationNameChilworth,
	StationCodeChingford:                         StationNameChingford,
	StationCodeChinley:                           StationNameChinley,
	StationCodeChippenham:                        StationNameChippenham,
	StationCodeChipstead:                         StationNameChipstead,
	StationCodeChirk:                             StationNameChirk,
	StationCodeChislehurst:                       StationNameChislehurst,
	StationCodeChiswick:                          StationNameChiswick,
	StationCodeCholsey:                           StationNameCholsey,
	StationCodeChorley:                           StationNameChorley,
	StationCodeChorleywood:                       StationNameChorleywood,
	StationCodeChristchurch:                      StationNameChristchurch,
	StationCodeChristsHospital:                   StationNameChristsHospital,
	StationCodeChurchAndOswaldtwistle:            StationNameChurchAndOswaldtwistle,
	StationCodeChurchFenton:                      StationNameChurchFenton,
	StationCodeChurchStretton:                    StationNameChurchStretton,
	StationCodeCilmeri:                           StationNameCilmeri,
	StationCodeCityThameslink:                    StationNameCityThameslink,
	StationCodeClactononSea:                      StationNameClactononSea,
	StationCodeClandon:                           StationNameClandon,
	StationCodeClaphamHighStreet:                 StationNameClaphamHighStreet,
	StationCodeClaphamJunction:                   StationNameClaphamJunction,
	StationCodeClaphamNorthYorkshire:             StationNameClaphamNorthYorkshire,
	StationCodeClapton:                           StationNameClapton,
	StationCodeClarbestonRoad:                    StationNameClarbestonRoad,
	StationCodeClarkston:                         StationNameClarkston,
	StationCodeClaverdon:                         StationNameClaverdon,
	StationCodeClaygate:                          StationNameClaygate,
	StationCodeCleethorpes:                       StationNameCleethorpes,
	StationCodeCleland:                           StationNameCleland,
	StationCodeCliftonDown:                       StationNameCliftonDown,
	StationCodeCliftonManchester:                 StationNameCliftonManchester,
	StationCodeClitheroe:                         StationNameClitheroe,
	StationCodeClockHouse:                        StationNameClockHouse,
	StationCodeClunderwen:                        StationNameClunderwen,
	StationCodeClydebank:                         StationNameClydebank,
	StationCodeCoatbridgeCentral:                 StationNameCoatbridgeCentral,
	StationCodeCoatbridgeSunnyside:               StationNameCoatbridgeSunnyside,
	StationCodeCoatdyke:                          StationNameCoatdyke,
	StationCodeCobhamAndStokedAbernon:            StationNameCobhamAndStokedAbernon,
	StationCodeCodsall:                           StationNameCodsall,
	StationCodeCogan:                             StationNameCogan,
	StationCodeColchester:                        StationNameColchester,
	StationCodeColchesterTown:                    StationNameColchesterTown,
	StationCodeColeshillParkway:                  StationNameColeshillParkway,
	StationCodeCollingham:                        StationNameCollingham,
	StationCodeCollington:                        StationNameCollington,
	StationCodeColne:                             StationNameColne,
	StationCodeColwall:                           StationNameColwall,
	StationCodeColwynBay:                         StationNameColwynBay,
	StationCodeCombeOxon:                         StationNameCombeOxon,
	StationCodeCommondale:                        StationNameCommondale,
	StationCodeCongleton:                         StationNameCongleton,
	StationCodeConisbrough:                       StationNameConisbrough,
	StationCodeConnelFerry:                       StationNameConnelFerry,
	StationCodeCononBridge:                       StationNameCononBridge,
	StationCodeCononley:                          StationNameCononley,
	StationCodeConwayPark:                        StationNameConwayPark,
	StationCodeConwy:                             StationNameConwy,
	StationCodeCoodenBeach:                       StationNameCoodenBeach,
	StationCodeCookham:                           StationNameCookham,
	StationCodeCooksbridge:                       StationNameCooksbridge,
	StationCodeCoombeJunctionHalt:                StationNameCoombeJunctionHalt,
	StationCodeCopplestone:                       StationNameCopplestone,
	StationCodeCorbridge:                         StationNameCorbridge,
	StationCodeCorby:                             StationNameCorby,
	StationCodeCorkerhill:                        StationNameCorkerhill,
	StationCodeCorkickle:                         StationNameCorkickle,
	StationCodeCorpach:                           StationNameCorpach,
	StationCodeCorrour:                           StationNameCorrour,
	StationCodeCoryton:                           StationNameCoryton,
	StationCodeCoseley:                           StationNameCoseley,
	StationCodeCosford:                           StationNameCosford,
	StationCodeCosham:                            StationNameCosham,
	StationCodeCottingham:                        StationNameCottingham,
	StationCodeCottingley:                        StationNameCottingley,
	StationCodeCoulsdonSouth:                     StationNameCoulsdonSouth,
	StationCodeCoulsdonTown:                      StationNameCoulsdonTown,
	StationCodeCoventry:                          StationNameCoventry,
	StationCodeCoventryArena:                     StationNameCoventryArena,
	StationCodeCowdenKent:                        StationNameCowdenKent,
	StationCodeCowdenbeath:                       StationNameCowdenbeath,
	StationCodeCradleyHeath:                      StationNameCradleyHeath,
	StationCodeCraigendoran:                      StationNameCraigendoran,
	StationCodeCramlington:                       StationNameCramlington,
	StationCodeCranbrookDevon:                    StationNameCranbrookDevon,
	StationCodeCravenArms:                        StationNameCravenArms,
	StationCodeCrawley:                           StationNameCrawley,
	StationCodeCrayford:                          StationNameCrayford,
	StationCodeCrediton:                          StationNameCrediton,
	StationCodeCressingEssex:                     StationNameCressingEssex,
	StationCodeCressington:                       StationNameCressington,
	StationCodeCreswell:                          StationNameCreswell,
	StationCodeCrewe:                             StationNameCrewe,
	StationCodeCrewkerne:                         StationNameCrewkerne,
	StationCodeCrewsHill:                         StationNameCrewsHill,
	StationCodeCrianlarich:                       StationNameCrianlarich,
	StationCodeCriccieth:                         StationNameCriccieth,
	StationCodeCricklewood:                       StationNameCricklewood,
	StationCodeCroftfoot:                         StationNameCroftfoot,
	StationCodeCroftonPark:                       StationNameCroftonPark,
	StationCodeCromer:                            StationNameCromer,
	StationCodeCromford:                          StationNameCromford,
	StationCodeCrookston:                         StationNameCrookston,
	StationCodeCrossGates:                        StationNameCrossGates,
	StationCodeCrossflatts:                       StationNameCrossflatts,
	StationCodeCrosshill:                         StationNameCrosshill,
	StationCodeCrosskeys:                         StationNameCrosskeys,
	StationCodeCrossmyloof:                       StationNameCrossmyloof,
	StationCodeCroston:                           StationNameCroston,
	StationCodeCrouchHill:                        StationNameCrouchHill,
	StationCodeCrowborough:                       StationNameCrowborough,
	StationCodeCrowhurst:                         StationNameCrowhurst,
	StationCodeCrowle:                            StationNameCrowle,
	StationCodeCrowthorne:                        StationNameCrowthorne,
	StationCodeCroy:                              StationNameCroy,
	StationCodeCrystalPalace:                     StationNameCrystalPalace,
	StationCodeCuddington:                        StationNameCuddington,
	StationCodeCuffley:                           StationNameCuffley,
	StationCodeCulham:                            StationNameCulham,
	StationCodeCulrain:                           StationNameCulrain,
	StationCodeCumbernauld:                       StationNameCumbernauld,
	StationCodeCupar:                             StationNameCupar,
	StationCodeCurriehill:                        StationNameCurriehill,
	StationCodeCuxton:                            StationNameCuxton,
	StationCodeCwmbach:                           StationNameCwmbach,
	StationCodeCwmbran:                           StationNameCwmbran,
	StationCodeCynghordy:                         StationNameCynghordy,
	StationCodeDagenhamDock:                      StationNameDagenhamDock,
	StationCodeDaisyHill:                         StationNameDaisyHill,
	StationCodeDalgetyBay:                        StationNameDalgetyBay,
	StationCodeDalmally:                          StationNameDalmally,
	StationCodeDalmarnock:                        StationNameDalmarnock,
	StationCodeDalmeny:                           StationNameDalmeny,
	StationCodeDalmuir:                           StationNameDalmuir,
	StationCodeDalreoch:                          StationNameDalreoch,
	StationCodeDalry:                             StationNameDalry,
	StationCodeDalstonCumbria:                    StationNameDalstonCumbria,
	StationCodeDalstonJunction:                   StationNameDalstonJunction,
	StationCodeDalstonKingsland:                  StationNameDalstonKingsland,
	StationCodeDaltonCumbria:                     StationNameDaltonCumbria,
	StationCodeDalwhinnie:                        StationNameDalwhinnie,
	StationCodeDanby:                             StationNameDanby,
	StationCodeDanescourt:                        StationNameDanescourt,
	StationCodeDanzey:                            StationNameDanzey,
	StationCodeDarlington:                        StationNameDarlington,
	StationCodeDarnall:                           StationNameDarnall,
	StationCodeDarsham:                           StationNameDarsham,
	StationCodeDartford:                          StationNameDartford,
	StationCodeDarton:                            StationNameDarton,
	StationCodeDarwen:                            StationNameDarwen,
	StationCodeDatchet:                           StationNameDatchet,
	StationCodeDavenport:                         StationNameDavenport,
	StationCodeDawlish:                           StationNameDawlish,
	StationCodeDawlishWarren:                     StationNameDawlishWarren,
	StationCodeDeal:                              StationNameDeal,
	StationCodeDeanWilts:                         StationNameDeanWilts,
	StationCodeDeansgate:                         StationNameDeansgate,
	StationCodeDeganwy:                           StationNameDeganwy,
	StationCodeDeighton:                          StationNameDeighton,
	StationCodeDelamere:                          StationNameDelamere,
	StationCodeDenbyDale:                         StationNameDenbyDale,
	StationCodeDenham:                            StationNameDenham,
	StationCodeDenhamGolfClub:                    StationNameDenhamGolfClub,
	StationCodeDenmarkHill:                       StationNameDenmarkHill,
	StationCodeDent:                              StationNameDent,
	StationCodeDenton:                            StationNameDenton,
	StationCodeDeptford:                          StationNameDeptford,
	StationCodeDerby:                             StationNameDerby,
	StationCodeDerbyRoadIpswich:                  StationNameDerbyRoadIpswich,
	StationCodeDevonportDevon:                    StationNameDevonportDevon,
	StationCodeDevonportDockyard:                 StationNameDevonportDockyard,
	StationCodeDewsbury:                          StationNameDewsbury,
	StationCodeDidcotParkway:                     StationNameDidcotParkway,
	StationCodeDigbyAndSowton:                    StationNameDigbyAndSowton,
	StationCodeDiltonMarsh:                       StationNameDiltonMarsh,
	StationCodeDinasPowys:                        StationNameDinasPowys,
	StationCodeDinasRhondda:                      StationNameDinasRhondda,
	StationCodeDingleRoad:                        StationNameDingleRoad,
	StationCodeDingwall:                          StationNameDingwall,
	StationCodeDinsdale:                          StationNameDinsdale,
	StationCodeDinting:                           StationNameDinting,
	StationCodeDisley:                            StationNameDisley,
	StationCodeDiss:                              StationNameDiss,
	StationCodeDodworth:                          StationNameDodworth,
	StationCodeDolau:                             StationNameDolau,
	StationCodeDoleham:                           StationNameDoleham,
	StationCodeDolgarrog:                         StationNameDolgarrog,
	StationCodeDolwyddelan:                       StationNameDolwyddelan,
	StationCodeDoncaster:                         StationNameDoncaster,
	StationCodeDorchesterSouth:                   StationNameDorchesterSouth,
	StationCodeDorchesterWest:                    StationNameDorchesterWest,
	StationCodeDoreAndTotley:                     StationNameDoreAndTotley,
	StationCodeDorkingDeepdene:                   StationNameDorkingDeepdene,
	StationCodeDorkingMain:                       StationNameDorkingMain,
	StationCodeDorkingWest:                       StationNameDorkingWest,
	StationCodeDormans:                           StationNameDormans,
	StationCodeDorridge:                          StationNameDorridge,
	StationCodeDoveHoles:                         StationNameDoveHoles,
	StationCodeDoverPriory:                       StationNameDoverPriory,
	StationCodeDovercourt:                        StationNameDovercourt,
	StationCodeDoveyJunction:                     StationNameDoveyJunction,
	StationCodeDownhamMarket:                     StationNameDownhamMarket,
	StationCodeDraytonGreen:                      StationNameDraytonGreen,
	StationCodeDraytonPark:                       StationNameDraytonPark,
	StationCodeDrem:                              StationNameDrem,
	StationCodeDriffield:                         StationNameDriffield,
	StationCodeDrigg:                             StationNameDrigg,
	StationCodeDroitwichSpa:                      StationNameDroitwichSpa,
	StationCodeDronfield:                         StationNameDronfield,
	StationCodeDrumchapel:                        StationNameDrumchapel,
	StationCodeDrumfrochar:                       StationNameDrumfrochar,
	StationCodeDrumgelloch:                       StationNameDrumgelloch,
	StationCodeDrumry:                            StationNameDrumry,
	StationCodeDublinFerryport:                   StationNameDublinFerryport,
	StationCodeDublinPortStena:                   StationNameDublinPortStena,
	StationCodeDuddeston:                         StationNameDuddeston,
	StationCodeDudleyPort:                        StationNameDudleyPort,
	StationCodeDuffield:                          StationNameDuffield,
	StationCodeDuirinish:                         StationNameDuirinish,
	StationCodeDukeStreet:                        StationNameDukeStreet,
	StationCodeDullingham:                        StationNameDullingham,
	StationCodeDumbartonCentral:                  StationNameDumbartonCentral,
	StationCodeDumbartonEast:                     StationNameDumbartonEast,
	StationCodeDumbreck:                          StationNameDumbreck,
	StationCodeDumfries:                          StationNameDumfries,
	StationCodeDumptonPark:                       StationNameDumptonPark,
	StationCodeDunbar:                            StationNameDunbar,
	StationCodeDunblane:                          StationNameDunblane,
	StationCodeDuncraig:                          StationNameDuncraig,
	StationCodeDundee:                            StationNameDundee,
	StationCodeDunfermlineQueenMargaret:          StationNameDunfermlineQueenMargaret,
	StationCodeDunfermlineTown:                   StationNameDunfermlineTown,
	StationCodeDunkeldAndBirnam:                  StationNameDunkeldAndBirnam,
	StationCodeDunlop:                            StationNameDunlop,
	StationCodeDunrobinCastle:                    StationNameDunrobinCastle,
	StationCodeDunston:                           StationNameDunston,
	StationCodeDuntonGreen:                       StationNameDuntonGreen,
	StationCodeDurham:                            StationNameDurham,
	StationCodeDurringtononSea:                   StationNameDurringtononSea,
	StationCodeDyce:                              StationNameDyce,
	StationCodeDyffrynArdudwy:                    StationNameDyffrynArdudwy,
	StationCodeEaglescliffe:                      StationNameEaglescliffe,
	StationCodeEalingBroadway:                    StationNameEalingBroadway,
	StationCodeEarlestown:                        StationNameEarlestown,
	StationCodeEarley:                            StationNameEarley,
	StationCodeEarlsfield:                        StationNameEarlsfield,
	StationCodeEarlswoodSurrey:                   StationNameEarlswoodSurrey,
	StationCodeEarlswoodWestMidlands:             StationNameEarlswoodWestMidlands,
	StationCodeEastCroydon:                       StationNameEastCroydon,
	StationCodeEastDidsbury:                      StationNameEastDidsbury,
	StationCodeEastDulwich:                       StationNameEastDulwich,
	StationCodeEastFarleigh:                      StationNameEastFarleigh,
	StationCodeEastGarforth:                      StationNameEastGarforth,
	StationCodeEastGrinstead:                     StationNameEastGrinstead,
	StationCodeEastKilbride:                      StationNameEastKilbride,
	StationCodeEastMalling:                       StationNameEastMalling,
	StationCodeEastMidlandsParkway:               StationNameEastMidlandsParkway,
	StationCodeEastTilbury:                       StationNameEastTilbury,
	StationCodeEastWorthing:                      StationNameEastWorthing,
	StationCodeEastbourne:                        StationNameEastbourne,
	StationCodeEastbrook:                         StationNameEastbrook,
	StationCodeEasterhouse:                       StationNameEasterhouse,
	StationCodeEasthamRake:                       StationNameEasthamRake,
	StationCodeEastleigh:                         StationNameEastleigh,
	StationCodeEastrington:                       StationNameEastrington,
	StationCodeEbbsfleetInternational:            StationNameEbbsfleetInternational,
	StationCodeEbbwValeParkway:                   StationNameEbbwValeParkway,
	StationCodeEbbwValeTown:                      StationNameEbbwValeTown,
	StationCodeEcclesManchester:                  StationNameEcclesManchester,
	StationCodeEcclesRoad:                        StationNameEcclesRoad,
	StationCodeEcclestonPark:                     StationNameEcclestonPark,
	StationCodeEdale:                             StationNameEdale,
	StationCodeEdenPark:                          StationNameEdenPark,
	StationCodeEdenbridge:                        StationNameEdenbridge,
	StationCodeEdenbridgeTown:                    StationNameEdenbridgeTown,
	StationCodeEdgeHill:                          StationNameEdgeHill,
	StationCodeEdinburgh:                         StationNameEdinburgh,
	StationCodeEdinburghGateway:                  StationNameEdinburghGateway,
	StationCodeEdinburghPark:                     StationNameEdinburghPark,
	StationCodeEdmontonGreen:                     StationNameEdmontonGreen,
	StationCodeEffinghamJunction:                 StationNameEffinghamJunction,
	StationCodeEggesford:                         StationNameEggesford,
	StationCodeEgham:                             StationNameEgham,
	StationCodeEgton:                             StationNameEgton,
	StationCodeElephantAndCastle:                 StationNameElephantAndCastle,
	StationCodeElephantAndCastleUnderground:      StationNameElephantAndCastleUnderground,
	StationCodeElgin:                             StationNameElgin,
	StationCodeEllesmerePort:                     StationNameEllesmerePort,
	StationCodeElmersEnd:                         StationNameElmersEnd,
	StationCodeElmsteadWoods:                     StationNameElmsteadWoods,
	StationCodeElmswell:                          StationNameElmswell,
	StationCodeElsecar:                           StationNameElsecar,
	StationCodeElsenhamEssex:                     StationNameElsenhamEssex,
	StationCodeElstreeAndBorehamwood:             StationNameElstreeAndBorehamwood,
	StationCodeEltham:                            StationNameEltham,
	StationCodeEltonAndOrston:                    StationNameEltonAndOrston,
	StationCodeEly:                               StationNameEly,
	StationCodeEmersonPark:                       StationNameEmersonPark,
	StationCodeEmsworth:                          StationNameEmsworth,
	StationCodeEnerglynAndChurchillPark:          StationNameEnerglynAndChurchillPark,
	StationCodeEnfieldChase:                      StationNameEnfieldChase,
	StationCodeEnfieldLock:                       StationNameEnfieldLock,
	StationCodeEnfieldTown:                       StationNameEnfieldTown,
	StationCodeEntwistle:                         StationNameEntwistle,
	StationCodeEpsomDowns:                        StationNameEpsomDowns,
	StationCodeEpsomSurrey:                       StationNameEpsomSurrey,
	StationCodeErdington:                         StationNameErdington,
	StationCodeEridge:                            StationNameEridge,
	StationCodeErith:                             StationNameErith,
	StationCodeEsher:                             StationNameEsher,
	StationCodeEskbank:                           StationNameEskbank,
	StationCodeEssexRoad:                         StationNameEssexRoad,
	StationCodeEtchingham:                        StationNameEtchingham,
	StationCodeEuxtonBalshawLane:                 StationNameEuxtonBalshawLane,
	StationCodeEvesham:                           StationNameEvesham,
	StationCodeEwellEast:                         StationNameEwellEast,
	StationCodeEwellWest:                         StationNameEwellWest,
	StationCodeExeterCentral:                     StationNameExeterCentral,
	StationCodeExeterStDavids:                    StationNameExeterStDavids,
	StationCodeExeterStThomas:                    StationNameExeterStThomas,
	StationCodeExhibitionCentreGlasgow:           StationNameExhibitionCentreGlasgow,
	StationCodeExmouth:                           StationNameExmouth,
	StationCodeExton:                             StationNameExton,
	StationCodeEynsford:                          StationNameEynsford,
	StationCodeFairbourne:                        StationNameFairbourne,
	StationCodeFairfield:                         StationNameFairfield,
	StationCodeFairlie:                           StationNameFairlie,
	StationCodeFairwater:                         StationNameFairwater,
	StationCodeFalconwood:                        StationNameFalconwood,
	StationCodeFalkirkGrahamston:                 StationNameFalkirkGrahamston,
	StationCodeFalkirkHigh:                       StationNameFalkirkHigh,
	StationCodeFallsofCruachan:                   StationNameFallsofCruachan,
	StationCodeFalmer:                            StationNameFalmer,
	StationCodeFalmouthDocks:                     StationNameFalmouthDocks,
	StationCodeFalmouthTown:                      StationNameFalmouthTown,
	StationCodeFareham:                           StationNameFareham,
	StationCodeFarnboroughMain:                   StationNameFarnboroughMain,
	StationCodeFarnboroughNorth:                  StationNameFarnboroughNorth,
	StationCodeFarncombe:                         StationNameFarncombe,
	StationCodeFarnham:                           StationNameFarnham,
	StationCodeFarninghamRoad:                    StationNameFarninghamRoad,
	StationCodeFarnworth:                         StationNameFarnworth,
	StationCodeFarringdon:                        StationNameFarringdon,
	StationCodeFauldhouse:                        StationNameFauldhouse,
	StationCodeFaversham:                         StationNameFaversham,
	StationCodeFaygate:                           StationNameFaygate,
	StationCodeFazakerley:                        StationNameFazakerley,
	StationCodeFearn:                             StationNameFearn,
	StationCodeFeatherstone:                      StationNameFeatherstone,
	StationCodeFelixstowe:                        StationNameFelixstowe,
	StationCodeFeltham:                           StationNameFeltham,
	StationCodeFeniton:                           StationNameFeniton,
	StationCodeFennyStratford:                    StationNameFennyStratford,
	StationCodeFernhill:                          StationNameFernhill,
	StationCodeFerriby:                           StationNameFerriby,
	StationCodeFerryside:                         StationNameFerryside,
	StationCodeFfairfach:                         StationNameFfairfach,
	StationCodeFiley:                             StationNameFiley,
	StationCodeFiltonAbbeyWood:                   StationNameFiltonAbbeyWood,
	StationCodeFinchleyRoadAndFrognal:            StationNameFinchleyRoadAndFrognal,
	StationCodeFinsburyPark:                      StationNameFinsburyPark,
	StationCodeFinstock:                          StationNameFinstock,
	StationCodeFishbourneSussex:                  StationNameFishbourneSussex,
	StationCodeFishersgate:                       StationNameFishersgate,
	StationCodeFishguardAndGoodwick:              StationNameFishguardAndGoodwick,
	StationCodeFishguardHarbour:                  StationNameFishguardHarbour,
	StationCodeFiskerton:                         StationNameFiskerton,
	StationCodeFitzwilliam:                       StationNameFitzwilliam,
	StationCodeFiveWays:                          StationNameFiveWays,
	StationCodeFleet:                             StationNameFleet,
	StationCodeFlimby:                            StationNameFlimby,
	StationCodeFlint:                             StationNameFlint,
	StationCodeFlitwick:                          StationNameFlitwick,
	StationCodeFlixton:                           StationNameFlixton,
	StationCodeFloweryField:                      StationNameFloweryField,
	StationCodeFolkestoneCentral:                 StationNameFolkestoneCentral,
	StationCodeFolkestoneWest:                    StationNameFolkestoneWest,
	StationCodeFord:                              StationNameFord,
	StationCodeForestGate:                        StationNameForestGate,
	StationCodeForestHill:                        StationNameForestHill,
	StationCodeFormby:                            StationNameFormby,
	StationCodeForres:                            StationNameForres,
	StationCodeForsinard:                         StationNameForsinard,
	StationCodeFortMatilda:                       StationNameFortMatilda,
	StationCodeFortWilliam:                       StationNameFortWilliam,
	StationCodeFourOaks:                          StationNameFourOaks,
	StationCodeFoxfield:                          StationNameFoxfield,
	StationCodeFoxton:                            StationNameFoxton,
	StationCodeFrant:                             StationNameFrant,
	StationCodeFratton:                           StationNameFratton,
	StationCodeFreshfield:                        StationNameFreshfield,
	StationCodeFreshford:                         StationNameFreshford,
	StationCodeFrimley:                           StationNameFrimley,
	StationCodeFrintononSea:                      StationNameFrintononSea,
	StationCodeFrizinghall:                       StationNameFrizinghall,
	StationCodeFrodsham:                          StationNameFrodsham,
	StationCodeFrome:                             StationNameFrome,
	StationCodeFulwell:                           StationNameFulwell,
	StationCodeFurnessVale:                       StationNameFurnessVale,
	StationCodeFurzePlatt:                        StationNameFurzePlatt,
	StationCodeGainsboroughCentral:               StationNameGainsboroughCentral,
	StationCodeGainsboroughLeaRoad:               StationNameGainsboroughLeaRoad,
	StationCodeGalashiels:                        StationNameGalashiels,
	StationCodeGarelochhead:                      StationNameGarelochhead,
	StationCodeGarforth:                          StationNameGarforth,
	StationCodeGargrave:                          StationNameGargrave,
	StationCodeGarrowhill:                        StationNameGarrowhill,
	StationCodeGarscadden:                        StationNameGarscadden,
	StationCodeGarsdale:                          StationNameGarsdale,
	StationCodeGarstonHertfordshire:              StationNameGarstonHertfordshire,
	StationCodeGarswood:                          StationNameGarswood,
	StationCodeGartcosh:                          StationNameGartcosh,
	StationCodeGarthMidGlamorgan:                 StationNameGarthMidGlamorgan,
	StationCodeGarthPowys:                        StationNameGarthPowys,
	StationCodeGarve:                             StationNameGarve,
	StationCodeGathurst:                          StationNameGathurst,
	StationCodeGatley:                            StationNameGatley,
	StationCodeGatwickAirport:                    StationNameGatwickAirport,
	StationCodeGeorgemasJunction:                 StationNameGeorgemasJunction,
	StationCodeGerrardsCross:                     StationNameGerrardsCross,
	StationCodeGideaPark:                         StationNameGideaPark,
	StationCodeGiffnock:                          StationNameGiffnock,
	StationCodeGiggleswick:                       StationNameGiggleswick,
	StationCodeGilberdyke:                        StationNameGilberdyke,
	StationCodeGilfachFargoed:                    StationNameGilfachFargoed,
	StationCodeGillinghamDorset:                  StationNameGillinghamDorset,
	StationCodeGillinghamKent:                    StationNameGillinghamKent,
	StationCodeGilshochill:                       StationNameGilshochill,
	StationCodeGipsyHill:                         StationNameGipsyHill,
	StationCodeGirvan:                            StationNameGirvan,
	StationCodeGlaisdale:                         StationNameGlaisdale,
	StationCodeGlanConwy:                         StationNameGlanConwy,
	StationCodeGlasgowCentral:                    StationNameGlasgowCentral,
	StationCodeGlasgowQueenStreet:                StationNameGlasgowQueenStreet,
	StationCodeGlasshoughton:                     StationNameGlasshoughton,
	StationCodeGlazebrook:                        StationNameGlazebrook,
	StationCodeGleneagles:                        StationNameGleneagles,
	StationCodeGlenfinnan:                        StationNameGlenfinnan,
	StationCodeGlengarnock:                       StationNameGlengarnock,
	StationCodeGlenrotheswithThornton:            StationNameGlenrotheswithThornton,
	StationCodeGlossop:                           StationNameGlossop,
	StationCodeGloucester:                        StationNameGloucester,
	StationCodeGlynde:                            StationNameGlynde,
	StationCodeGobowen:                           StationNameGobowen,
	StationCodeGodalming:                         StationNameGodalming,
	StationCodeGodley:                            StationNameGodley,
	StationCodeGodstone:                          StationNameGodstone,
	StationCodeGoldthorpe:                        StationNameGoldthorpe,
	StationCodeGolfStreet:                        StationNameGolfStreet,
	StationCodeGolspie:                           StationNameGolspie,
	StationCodeGomshall:                          StationNameGomshall,
	StationCodeGoodmayes:                         StationNameGoodmayes,
	StationCodeGoole:                             StationNameGoole,
	StationCodeGoostrey:                          StationNameGoostrey,
	StationCodeGordonHill:                        StationNameGordonHill,
	StationCodeGorebridge:                        StationNameGorebridge,
	StationCodeGoringAndStreatley:                StationNameGoringAndStreatley,
	StationCodeGoringbySea:                       StationNameGoringbySea,
	StationCodeGorton:                            StationNameGorton,
	StationCodeGospelOak:                         StationNameGospelOak,
	StationCodeGourock:                           StationNameGourock,
	StationCodeGowerton:                          StationNameGowerton,
	StationCodeGoxhill:                           StationNameGoxhill,
	StationCodeGrangeOverSands:                   StationNameGrangeOverSands,
	StationCodeGrangePark:                        StationNameGrangePark,
	StationCodeGrangetownCardiff:                 StationNameGrangetownCardiff,
	StationCodeGrantham:                          StationNameGrantham,
	StationCodeGrateley:                          StationNameGrateley,
	StationCodeGravellyHill:                      StationNameGravellyHill,
	StationCodeGravesend:                         StationNameGravesend,
	StationCodeGrays:                             StationNameGrays,
	StationCodeGreatAyton:                        StationNameGreatAyton,
	StationCodeGreatBentley:                      StationNameGreatBentley,
	StationCodeGreatChesterford:                  StationNameGreatChesterford,
	StationCodeGreatCoates:                       StationNameGreatCoates,
	StationCodeGreatMalvern:                      StationNameGreatMalvern,
	StationCodeGreatMissenden:                    StationNameGreatMissenden,
	StationCodeGreatYarmouth:                     StationNameGreatYarmouth,
	StationCodeGreenLane:                         StationNameGreenLane,
	StationCodeGreenRoad:                         StationNameGreenRoad,
	StationCodeGreenbank:                         StationNameGreenbank,
	StationCodeGreenfaulds:                       StationNameGreenfaulds,
	StationCodeGreenfield:                        StationNameGreenfield,
	StationCodeGreenford:                         StationNameGreenford,
	StationCodeGreenhitheForBluewater:            StationNameGreenhitheForBluewater,
	StationCodeGreenockCentral:                   StationNameGreenockCentral,
	StationCodeGreenockWest:                      StationNameGreenockWest,
	StationCodeGreenwich:                         StationNameGreenwich,
	StationCodeGretnaGreen:                       StationNameGretnaGreen,
	StationCodeGrimsbyDocks:                      StationNameGrimsbyDocks,
	StationCodeGrimsbyTown:                       StationNameGrimsbyTown,
	StationCodeGrindleford:                       StationNameGrindleford,
	StationCodeGrosmont:                          StationNameGrosmont,
	StationCodeGrovePark:                         StationNameGrovePark,
	StationCodeGuideBridge:                       StationNameGuideBridge,
	StationCodeGuildford:                         StationNameGuildford,
	StationCodeGuiseley:                          StationNameGuiseley,
	StationCodeGunnersbury:                       StationNameGunnersbury,
	StationCodeGunnislake:                        StationNameGunnislake,
	StationCodeGunton:                            StationNameGunton,
	StationCodeGwersyllt:                         StationNameGwersyllt,
	StationCodeGypsyLane:                         StationNameGypsyLane,
	StationCodeHabrough:                          StationNameHabrough,
	StationCodeHackbridge:                        StationNameHackbridge,
	StationCodeHackneyCentral:                    StationNameHackneyCentral,
	StationCodeHackneyDowns:                      StationNameHackneyDowns,
	StationCodeHackneyWick:                       StationNameHackneyWick,
	StationCodeHaddenhamAndThameParkway:          StationNameHaddenhamAndThameParkway,
	StationCodeHaddiscoe:                         StationNameHaddiscoe,
	StationCodeHadfield:                          StationNameHadfield,
	StationCodeHadleyWood:                        StationNameHadleyWood,
	StationCodeHagFold:                           StationNameHagFold,
	StationCodeHaggerston:                        StationNameHaggerston,
	StationCodeHagley:                            StationNameHagley,
	StationCodeHairmyres:                         StationNameHairmyres,
	StationCodeHaleManchester:                    StationNameHaleManchester,
	StationCodeHalesworth:                        StationNameHalesworth,
	StationCodeHalewood:                          StationNameHalewood,
	StationCodeHalifax:                           StationNameHalifax,
	StationCodeHallGreen:                         StationNameHallGreen,
	StationCodeHallRoad:                          StationNameHallRoad,
	StationCodeHalling:                           StationNameHalling,
	StationCodeHallithWood:                       StationNameHallithWood,
	StationCodeHaltwhistle:                       StationNameHaltwhistle,
	StationCodeHamStreet:                         StationNameHamStreet,
	StationCodeHamble:                            StationNameHamble,
	StationCodeHamiltonCentral:                   StationNameHamiltonCentral,
	StationCodeHamiltonWest:                      StationNameHamiltonWest,
	StationCodeHammerton:                         StationNameHammerton,
	StationCodeHampdenParkSussex:                 StationNameHampdenParkSussex,
	StationCodeHampsteadHeath:                    StationNameHampsteadHeath,
	StationCodeHamptonCourt:                      StationNameHamptonCourt,
	StationCodeHamptonLondon:                     StationNameHamptonLondon,
	StationCodeHamptonWick:                       StationNameHamptonWick,
	StationCodeHamptoninArden:                    StationNameHamptoninArden,
	StationCodeHamsteadBirmingham:                StationNameHamsteadBirmingham,
	StationCodeHamworthy:                         StationNameHamworthy,
	StationCodeHanborough:                        StationNameHanborough,
	StationCodeHandforth:                         StationNameHandforth,
	StationCodeHanwell:                           StationNameHanwell,
	StationCodeHapton:                            StationNameHapton,
	StationCodeHarlech:                           StationNameHarlech,
	StationCodeHarlesden:                         StationNameHarlesden,
	StationCodeHarlingRoad:                       StationNameHarlingRoad,
	StationCodeHarlingtonBeds:                    StationNameHarlingtonBeds,
	StationCodeHarlowMill:                        StationNameHarlowMill,
	StationCodeHarlowTown:                        StationNameHarlowTown,
	StationCodeHaroldWood:                        StationNameHaroldWood,
	StationCodeHarpenden:                         StationNameHarpenden,
	StationCodeHarrietsham:                       StationNameHarrietsham,
	StationCodeHarringay:                         StationNameHarringay,
	StationCodeHarringayGreenLanes:               StationNameHarringayGreenLanes,
	StationCodeHarrington:                        StationNameHarrington,
	StationCodeHarrogate:                         StationNameHarrogate,
	StationCodeHarrowAndWealdstone:               StationNameHarrowAndWealdstone,
	StationCodeHarrowontheHill:                   StationNameHarrowontheHill,
	StationCodeHartfordCheshire:                  StationNameHartfordCheshire,
	StationCodeHartlebury:                        StationNameHartlebury,
	StationCodeHartlepool:                        StationNameHartlepool,
	StationCodeHartwood:                          StationNameHartwood,
	StationCodeHarwichInternational:              StationNameHarwichInternational,
	StationCodeHarwichTown:                       StationNameHarwichTown,
	StationCodeHaslemere:                         StationNameHaslemere,
	StationCodeHassocks:                          StationNameHassocks,
	StationCodeHastings:                          StationNameHastings,
	StationCodeHatchEnd:                          StationNameHatchEnd,
	StationCodeHatfieldAndStainforth:             StationNameHatfieldAndStainforth,
	StationCodeHatfieldHerts:                     StationNameHatfieldHerts,
	StationCodeHatfieldPeverel:                   StationNameHatfieldPeverel,
	StationCodeHathersage:                        StationNameHathersage,
	StationCodeHattersley:                        StationNameHattersley,
	StationCodeHatton:                            StationNameHatton,
	StationCodeHavant:                            StationNameHavant,
	StationCodeHavenhouse:                        StationNameHavenhouse,
	StationCodeHaverfordwest:                     StationNameHaverfordwest,
	StationCodeHawarden:                          StationNameHawarden,
	StationCodeHawardenBridge:                    StationNameHawardenBridge,
	StationCodeHawkhead:                          StationNameHawkhead,
	StationCodeHaydonBridge:                      StationNameHaydonBridge,
	StationCodeHaydonsRoad:                       StationNameHaydonsRoad,
	StationCodeHayesAndHarlington:                StationNameHayesAndHarlington,
	StationCodeHayesKent:                         StationNameHayesKent,
	StationCodeHayle:                             StationNameHayle,
	StationCodeHaymarket:                         StationNameHaymarket,
	StationCodeHaywardsHeath:                     StationNameHaywardsHeath,
	StationCodeHazelGrove:                        StationNameHazelGrove,
	StationCodeHeadcorn:                          StationNameHeadcorn,
	StationCodeHeadingley:                        StationNameHeadingley,
	StationCodeHeadstoneLane:                     StationNameHeadstoneLane,
	StationCodeHealdGreen:                        StationNameHealdGreen,
	StationCodeHealing:                           StationNameHealing,
	StationCodeHeathHighLevel:                    StationNameHeathHighLevel,
	StationCodeHeathLowLevel:                     StationNameHeathLowLevel,
	StationCodeHeathrowAirportTerminal4:          StationNameHeathrowAirportTerminal4,
	StationCodeHeathrowAirportTerminal5:          StationNameHeathrowAirportTerminal5,
	StationCodeHeathrowAirportTerminals12and3:    StationNameHeathrowAirportTerminals12and3,
	StationCodeHeatonChapel:                      StationNameHeatonChapel,
	StationCodeHebdenBridge:                      StationNameHebdenBridge,
	StationCodeHeckington:                        StationNameHeckington,
	StationCodeHedgeEnd:                          StationNameHedgeEnd,
	StationCodeHednesford:                        StationNameHednesford,
	StationCodeHeighington:                       StationNameHeighington,
	StationCodeHelensburghCentral:                StationNameHelensburghCentral,
	StationCodeHelensburghUpper:                  StationNameHelensburghUpper,
	StationCodeHellifield:                        StationNameHellifield,
	StationCodeHelmsdale:                         StationNameHelmsdale,
	StationCodeHelsby:                            StationNameHelsby,
	StationCodeHemelHempstead:                    StationNameHemelHempstead,
	StationCodeHendon:                            StationNameHendon,
	StationCodeHengoed:                           StationNameHengoed,
	StationCodeHenleyinArden:                     StationNameHenleyinArden,
	StationCodeHenleyonThames:                    StationNameHenleyonThames,
	StationCodeHensall:                           StationNameHensall,
	StationCodeHereford:                          StationNameHereford,
	StationCodeHerneBay:                          StationNameHerneBay,
	StationCodeHerneHill:                         StationNameHerneHill,
	StationCodeHersham:                           StationNameHersham,
	StationCodeHertfordEast:                      StationNameHertfordEast,
	StationCodeHertfordNorth:                     StationNameHertfordNorth,
	StationCodeHessle:                            StationNameHessle,
	StationCodeHeswall:                           StationNameHeswall,
	StationCodeHever:                             StationNameHever,
	StationCodeHeworth:                           StationNameHeworth,
	StationCodeHexham:                            StationNameHexham,
	StationCodeHeyford:                           StationNameHeyford,
	StationCodeHeyshamPort:                       StationNameHeyshamPort,
	StationCodeHighBrooms:                        StationNameHighBrooms,
	StationCodeHighStreetGlasgow:                 StationNameHighStreetGlasgow,
	StationCodeHighStreetKensingtonUnderground:   StationNameHighStreetKensingtonUnderground,
	StationCodeHighWycombe:                       StationNameHighWycombe,
	StationCodeHigham:                            StationNameHigham,
	StationCodeHighamsPark:                       StationNameHighamsPark,
	StationCodeHighbridgeAndBurnham:              StationNameHighbridgeAndBurnham,
	StationCodeHighburyAndIslington:              StationNameHighburyAndIslington,
	StationCodeHightown:                          StationNameHightown,
	StationCodeHildenborough:                     StationNameHildenborough,
	StationCodeHillfoot:                          StationNameHillfoot,
	StationCodeHillingtonEast:                    StationNameHillingtonEast,
	StationCodeHillingtonWest:                    StationNameHillingtonWest,
	StationCodeHillside:                          StationNameHillside,
	StationCodeHilsea:                            StationNameHilsea,
	StationCodeHinchleyWood:                      StationNameHinchleyWood,
	StationCodeHinckleyLeics:                     StationNameHinckleyLeics,
	StationCodeHindley:                           StationNameHindley,
	StationCodeHintonAdmiral:                     StationNameHintonAdmiral,
	StationCodeHitchin:                           StationNameHitchin,
	StationCodeHitherGreen:                       StationNameHitherGreen,
	StationCodeHockley:                           StationNameHockley,
	StationCodeHollingbourne:                     StationNameHollingbourne,
	StationCodeHolmesChapel:                      StationNameHolmesChapel,
	StationCodeHolmwood:                          StationNameHolmwood,
	StationCodeHoltonHeath:                       StationNameHoltonHeath,
	StationCodeHolyhead:                          StationNameHolyhead,
	StationCodeHolytown:                          StationNameHolytown,
	StationCodeHomerton:                          StationNameHomerton,
	StationCodeHoneybourne:                       StationNameHoneybourne,
	StationCodeHoniton:                           StationNameHoniton,
	StationCodeHonley:                            StationNameHonley,
	StationCodeHonorOakPark:                      StationNameHonorOakPark,
	StationCodeHook:                              StationNameHook,
	StationCodeHooton:                            StationNameHooton,
	StationCodeHopeDerbyshire:                    StationNameHopeDerbyshire,
	StationCodeHopeFlintshire:                    StationNameHopeFlintshire,
	StationCodeHoptonHeath:                       StationNameHoptonHeath,
	StationCodeHorley:                            StationNameHorley,
	StationCodeHornbeamPark:                      StationNameHornbeamPark,
	StationCodeHornsey:                           StationNameHornsey,
	StationCodeHorsforth:                         StationNameHorsforth,
	StationCodeHorsham:                           StationNameHorsham,
	StationCodeHorsley:                           StationNameHorsley,
	StationCodeHortoninRibblesdale:               StationNameHortoninRibblesdale,
	StationCodeHorwichParkway:                    StationNameHorwichParkway,
	StationCodeHoscar:                            StationNameHoscar,
	StationCodeHoughGreen:                        StationNameHoughGreen,
	StationCodeHounslow:                          StationNameHounslow,
	StationCodeHove:                              StationNameHove,
	StationCodeHovetonAndWroxham:                 StationNameHovetonAndWroxham,
	StationCodeHowWoodHerts:                      StationNameHowWoodHerts,
	StationCodeHowden:                            StationNameHowden,
	StationCodeHowwoodRenfrewshire:               StationNameHowwoodRenfrewshire,
	StationCodeHoxton:                            StationNameHoxton,
	StationCodeHoylake:                           StationNameHoylake,
	StationCodeHubbertsBridge:                    StationNameHubbertsBridge,
	StationCodeHucknall:                          StationNameHucknall,
	StationCodeHuddersfield:                      StationNameHuddersfield,
	StationCodeHull:                              StationNameHull,
	StationCodeHumphreyPark:                      StationNameHumphreyPark,
	StationCodeHuncoat:                           StationNameHuncoat,
	StationCodeHungerford:                        StationNameHungerford,
	StationCodeHunmanby:                          StationNameHunmanby,
	StationCodeHuntingdon:                        StationNameHuntingdon,
	StationCodeHuntly:                            StationNameHuntly,
	StationCodeHuntsCross:                        StationNameHuntsCross,
	StationCodeHurstGreen:                        StationNameHurstGreen,
	StationCodeHuttonCranswick:                   StationNameHuttonCranswick,
	StationCodeHuyton:                            StationNameHuyton,
	StationCodeHydeCentral:                       StationNameHydeCentral,
	StationCodeHydeNorth:                         StationNameHydeNorth,
	StationCodeHykeham:                           StationNameHykeham,
	StationCodeHyndland:                          StationNameHyndland,
	StationCodeHytheEssex:                        StationNameHytheEssex,
	StationCodeIBMHalt:                           StationNameIBMHalt,
	StationCodeIfield:                            StationNameIfield,
	StationCodeIlford:                            StationNameIlford,
	StationCodeIlkley:                            StationNameIlkley,
	StationCodeImperialWharf:                     StationNameImperialWharf,
	StationCodeInceAndElton:                      StationNameInceAndElton,
	StationCodeInceManchester:                    StationNameInceManchester,
	StationCodeIngatestone:                       StationNameIngatestone,
	StationCodeInsch:                             StationNameInsch,
	StationCodeInvergordon:                       StationNameInvergordon,
	StationCodeInvergowrie:                       StationNameInvergowrie,
	StationCodeInverkeithing:                     StationNameInverkeithing,
	StationCodeInverkip:                          StationNameInverkip,
	StationCodeInverness:                         StationNameInverness,
	StationCodeInvershin:                         StationNameInvershin,
	StationCodeInverurie:                         StationNameInverurie,
	StationCodeIpswich:                           StationNameIpswich,
	StationCodeIrlam:                             StationNameIrlam,
	StationCodeIrvine:                            StationNameIrvine,
	StationCodeIsleworth:                         StationNameIsleworth,
	StationCodeIslip:                             StationNameIslip,
	StationCodeIver:                              StationNameIver,
	StationCodeIvybridge:                         StationNameIvybridge,
	StationCodeJamesCook:                         StationNameJamesCook,
	StationCodeJewelleryQuarter:                  StationNameJewelleryQuarter,
	StationCodeJohnstonPembs:                     StationNameJohnstonPembs,
	StationCodeJohnstoneStrathclyde:              StationNameJohnstoneStrathclyde,
	StationCodeJordanhill:                        StationNameJordanhill,
	StationCodeKearsleyManchester:                StationNameKearsleyManchester,
	StationCodeKearsney:                          StationNameKearsney,
	StationCodeKeighley:                          StationNameKeighley,
	StationCodeKeith:                             StationNameKeith,
	StationCodeKelvedon:                          StationNameKelvedon,
	StationCodeKelvindale:                        StationNameKelvindale,
	StationCodeKemble:                            StationNameKemble,
	StationCodeKempstonHardwick:                  StationNameKempstonHardwick,
	StationCodeKemptonParkRacecourse:             StationNameKemptonParkRacecourse,
	StationCodeKemsing:                           StationNameKemsing,
	StationCodeKemsley:                           StationNameKemsley,
	StationCodeKendal:                            StationNameKendal,
	StationCodeKenley:                            StationNameKenley,
	StationCodeKennett:                           StationNameKennett,
	StationCodeKennishead:                        StationNameKennishead,
	StationCodeKensalGreen:                       StationNameKensalGreen,
	StationCodeKensalRise:                        StationNameKensalRise,
	StationCodeKensingtonOlympia:                 StationNameKensingtonOlympia,
	StationCodeKentHouse:                         StationNameKentHouse,
	StationCodeKentishTown:                       StationNameKentishTown,
	StationCodeKentishTownWest:                   StationNameKentishTownWest,
	StationCodeKenton:                            StationNameKenton,
	StationCodeKentsBank:                         StationNameKentsBank,
	StationCodeKettering:                         StationNameKettering,
	StationCodeKewBridge:                         StationNameKewBridge,
	StationCodeKewGardens:                        StationNameKewGardens,
	StationCodeKeyham:                            StationNameKeyham,
	StationCodeKeynsham:                          StationNameKeynsham,
	StationCodeKidbrooke:                         StationNameKidbrooke,
	StationCodeKidderminster:                     StationNameKidderminster,
	StationCodeKidsgrove:                         StationNameKidsgrove,
	StationCodeKidwelly:                          StationNameKidwelly,
	StationCodeKilburnHighRoad:                   StationNameKilburnHighRoad,
	StationCodeKildale:                           StationNameKildale,
	StationCodeKildonan:                          StationNameKildonan,
	StationCodeKilgetty:                          StationNameKilgetty,
	StationCodeKilmarnock:                        StationNameKilmarnock,
	StationCodeKilmaurs:                          StationNameKilmaurs,
	StationCodeKilpatrick:                        StationNameKilpatrick,
	StationCodeKilwinning:                        StationNameKilwinning,
	StationCodeKinbrace:                          StationNameKinbrace,
	StationCodeKingham:                           StationNameKingham,
	StationCodeKinghorn:                          StationNameKinghorn,
	StationCodeKingsLangley:                      StationNameKingsLangley,
	StationCodeKingsLynn:                         StationNameKingsLynn,
	StationCodeKingsNorton:                       StationNameKingsNorton,
	StationCodeKingsNympton:                      StationNameKingsNympton,
	StationCodeKingsPark:                         StationNameKingsPark,
	StationCodeKingsSutton:                       StationNameKingsSutton,
	StationCodeKingsknowe:                        StationNameKingsknowe,
	StationCodeKingston:                          StationNameKingston,
	StationCodeKingswood:                         StationNameKingswood,
	StationCodeKingussie:                         StationNameKingussie,
	StationCodeKintbury:                          StationNameKintbury,
	StationCodeKirbyCross:                        StationNameKirbyCross,
	StationCodeKirkSandall:                       StationNameKirkSandall,
	StationCodeKirkbyMerseyside:                  StationNameKirkbyMerseyside,
	StationCodeKirkbyStephen:                     StationNameKirkbyStephen,
	StationCodeKirkbyinAshfield:                  StationNameKirkbyinAshfield,
	StationCodeKirkbyinFurness:                   StationNameKirkbyinFurness,
	StationCodeKirkcaldy:                         StationNameKirkcaldy,
	StationCodeKirkconnel:                        StationNameKirkconnel,
	StationCodeKirkdale:                          StationNameKirkdale,
	StationCodeKirkhamAndWesham:                  StationNameKirkhamAndWesham,
	StationCodeKirkhill:                          StationNameKirkhill,
	StationCodeKirknewton:                        StationNameKirknewton,
	StationCodeKirkstallForge:                    StationNameKirkstallForge,
	StationCodeKirkwood:                          StationNameKirkwood,
	StationCodeKirtonLindsey:                     StationNameKirtonLindsey,
	StationCodeKivetonBridge:                     StationNameKivetonBridge,
	StationCodeKivetonPark:                       StationNameKivetonPark,
	StationCodeKnaresborough:                     StationNameKnaresborough,
	StationCodeKnebworth:                         StationNameKnebworth,
	StationCodeKnighton:                          StationNameKnighton,
	StationCodeKnockholt:                         StationNameKnockholt,
	StationCodeKnottingley:                       StationNameKnottingley,
	StationCodeKnucklas:                          StationNameKnucklas,
	StationCodeKnutsford:                         StationNameKnutsford,
	StationCodeKyleofLochalsh:                    StationNameKyleofLochalsh,
	StationCodeLadybank:                          StationNameLadybank,
	StationCodeLadywell:                          StationNameLadywell,
	StationCodeLaindon:                           StationNameLaindon,
	StationCodeLairg:                             StationNameLairg,
	StationCodeLake:                              StationNameLake,
	StationCodeLakenheath:                        StationNameLakenheath,
	StationCodeLamphey:                           StationNameLamphey,
	StationCodeLanark:                            StationNameLanark,
	StationCodeLancaster:                         StationNameLancaster,
	StationCodeLancing:                           StationNameLancing,
	StationCodeLandywood:                         StationNameLandywood,
	StationCodeLangbank:                          StationNameLangbank,
	StationCodeLangho:                            StationNameLangho,
	StationCodeLangleyBerks:                      StationNameLangleyBerks,
	StationCodeLangleyGreen:                      StationNameLangleyGreen,
	StationCodeLangleyMill:                       StationNameLangleyMill,
	StationCodeLangside:                          StationNameLangside,
	StationCodeLangwathby:                        StationNameLangwathby,
	StationCodeLangwithWhaleyThorns:              StationNameLangwithWhaleyThorns,
	StationCodeLapford:                           StationNameLapford,
	StationCodeLapworth:                          StationNameLapworth,
	StationCodeLarbert:                           StationNameLarbert,
	StationCodeLargs:                             StationNameLargs,
	StationCodeLarkhall:                          StationNameLarkhall,
	StationCodeLaurencekirk:                      StationNameLaurencekirk,
	StationCodeLawrenceHill:                      StationNameLawrenceHill,
	StationCodeLaytonLancs:                       StationNameLaytonLancs,
	StationCodeLazonbyAndKirkoswald:              StationNameLazonbyAndKirkoswald,
	StationCodeLeaBridge:                         StationNameLeaBridge,
	StationCodeLeaGreen:                          StationNameLeaGreen,
	StationCodeLeaHall:                           StationNameLeaHall,
	StationCodeLeagrave:                          StationNameLeagrave,
	StationCodeLealholm:                          StationNameLealholm,
	StationCodeLeamingtonSpa:                     StationNameLeamingtonSpa,
	StationCodeLeasowe:                           StationNameLeasowe,
	StationCodeLeatherhead:                       StationNameLeatherhead,
	StationCodeLedbury:                           StationNameLedbury,
	StationCodeLeeLondon:                         StationNameLeeLondon,
	StationCodeLeeds:                             StationNameLeeds,
	StationCodeLeicester:                         StationNameLeicester,
	StationCodeLeighKent:                         StationNameLeighKent,
	StationCodeLeighonSea:                        StationNameLeighonSea,
	StationCodeLeightonBuzzard:                   StationNameLeightonBuzzard,
	StationCodeLelant:                            StationNameLelant,
	StationCodeLelantSaltings:                    StationNameLelantSaltings,
	StationCodeLenham:                            StationNameLenham,
	StationCodeLenzie:                            StationNameLenzie,
	StationCodeLeominster:                        StationNameLeominster,
	StationCodeLetchworthGardenCity:              StationNameLetchworthGardenCity,
	StationCodeLeucharsforStAndrews:              StationNameLeucharsforStAndrews,
	StationCodeLevenshulme:                       StationNameLevenshulme,
	StationCodeLewes:                             StationNameLewes,
	StationCodeLewisham:                          StationNameLewisham,
	StationCodeLeyland:                           StationNameLeyland,
	StationCodeLeytonMidlandRoad:                 StationNameLeytonMidlandRoad,
	StationCodeLeytonstoneHighRoad:               StationNameLeytonstoneHighRoad,
	StationCodeLichfieldCity:                     StationNameLichfieldCity,
	StationCodeLichfieldTrentValley:              StationNameLichfieldTrentValley,
	StationCodeLidlington:                        StationNameLidlington,
	StationCodeLimehouse:                         StationNameLimehouse,
	StationCodeLincolnCentral:                    StationNameLincolnCentral,
	StationCodeLingfield:                         StationNameLingfield,
	StationCodeLingwood:                          StationNameLingwood,
	StationCodeLinlithgow:                        StationNameLinlithgow,
	StationCodeLiphook:                           StationNameLiphook,
	StationCodeLiskeard:                          StationNameLiskeard,
	StationCodeLiss:                              StationNameLiss,
	StationCodeLisvaneAndThornhill:               StationNameLisvaneAndThornhill,
	StationCodeLittleKimble:                      StationNameLittleKimble,
	StationCodeLittleSutton:                      StationNameLittleSutton,
	StationCodeLittleborough:                     StationNameLittleborough,
	StationCodeLittlehampton:                     StationNameLittlehampton,
	StationCodeLittlehaven:                       StationNameLittlehaven,
	StationCodeLittleport:                        StationNameLittleport,
	StationCodeLiverpoolCentral:                  StationNameLiverpoolCentral,
	StationCodeLiverpoolJamesStreet:              StationNameLiverpoolJamesStreet,
	StationCodeLiverpoolLimeStreet:               StationNameLiverpoolLimeStreet,
	StationCodeLiverpoolSouthParkway:             StationNameLiverpoolSouthParkway,
	StationCodeLivingstonNorth:                   StationNameLivingstonNorth,
	StationCodeLivingstonSouth:                   StationNameLivingstonSouth,
	StationCodeLlanaber:                          StationNameLlanaber,
	StationCodeLlanbedr:                          StationNameLlanbedr,
	StationCodeLlanbisterRoad:                    StationNameLlanbisterRoad,
	StationCodeLlanbradach:                       StationNameLlanbradach,
	StationCodeLlandaf:                           StationNameLlandaf,
	StationCodeLlandanwg:                         StationNameLlandanwg,
	StationCodeLlandecwyn:                        StationNameLlandecwyn,
	StationCodeLlandeilo:                         StationNameLlandeilo,
	StationCodeLlandovery:                        StationNameLlandovery,
	StationCodeLlandrindod:                       StationNameLlandrindod,
	StationCodeLlandudno:                         StationNameLlandudno,
	StationCodeLlandudnoJunction:                 StationNameLlandudnoJunction,
	StationCodeLlandybie:                         StationNameLlandybie,
	StationCodeLlanelli:                          StationNameLlanelli,
	StationCodeLlanfairfechan:                    StationNameLlanfairfechan,
	StationCodeLlanfairpwll:                      StationNameLlanfairpwll,
	StationCodeLlangadog:                         StationNameLlangadog,
	StationCodeLlangammarch:                      StationNameLlangammarch,
	StationCodeLlangennech:                       StationNameLlangennech,
	StationCodeLlangynllo:                        StationNameLlangynllo,
	StationCodeLlanharan:                         StationNameLlanharan,
	StationCodeLlanhilleth:                       StationNameLlanhilleth,
	StationCodeLlanishen:                         StationNameLlanishen,
	StationCodeLlanrwst:                          StationNameLlanrwst,
	StationCodeLlansamlet:                        StationNameLlansamlet,
	StationCodeLlantwitMajor:                     StationNameLlantwitMajor,
	StationCodeLlanwrda:                          StationNameLlanwrda,
	StationCodeLlanwrtyd:                         StationNameLlanwrtyd,
	StationCodeLlwyngwril:                        StationNameLlwyngwril,
	StationCodeLlwynypia:                         StationNameLlwynypia,
	StationCodeLochAwe:                           StationNameLochAwe,
	StationCodeLochEilOutwardBound:               StationNameLochEilOutwardBound,
	StationCodeLochailort:                        StationNameLochailort,
	StationCodeLocheilside:                       StationNameLocheilside,
	StationCodeLochgelly:                         StationNameLochgelly,
	StationCodeLochluichart:                      StationNameLochluichart,
	StationCodeLochwinnoch:                       StationNameLochwinnoch,
	StationCodeLockerbie:                         StationNameLockerbie,
	StationCodeLockwood:                          StationNameLockwood,
	StationCodeLondonBlackfriars:                 StationNameLondonBlackfriars,
	StationCodeLondonBridge:                      StationNameLondonBridge,
	StationCodeLondonCannonStreet:                StationNameLondonCannonStreet,
	StationCodeLondonCharingCross:                StationNameLondonCharingCross,
	StationCodeLondonEuston:                      StationNameLondonEuston,
	StationCodeLondonFenchurchStreet:             StationNameLondonFenchurchStreet,
	StationCodeLondonFields:                      StationNameLondonFields,
	StationCodeLondonKingsCross:                  StationNameLondonKingsCross,
	StationCodeLondonLiverpoolStreet:             StationNameLondonLiverpoolStreet,
	StationCodeLondonMarylebone:                  StationNameLondonMarylebone,
	StationCodeLondonPaddington:                  StationNameLondonPaddington,
	StationCodeLondonRoadBrighton:                StationNameLondonRoadBrighton,
	StationCodeLondonRoadGuildford:               StationNameLondonRoadGuildford,
	StationCodeLondonStPancrasIntl:               StationNameLondonStPancrasIntl,
	StationCodeLondonVictoria:                    StationNameLondonVictoria,
	StationCodeLondonWaterloo:                    StationNameLondonWaterloo,
	StationCodeLondonWaterlooEast:                StationNameLondonWaterlooEast,
	StationCodeLongBuckby:                        StationNameLongBuckby,
	StationCodeLongEaton:                         StationNameLongEaton,
	StationCodeLongPreston:                       StationNameLongPreston,
	StationCodeLongbeck:                          StationNameLongbeck,
	StationCodeLongbridge:                        StationNameLongbridge,
	StationCodeLongcross:                         StationNameLongcross,
	StationCodeLongfield:                         StationNameLongfield,
	StationCodeLongniddry:                        StationNameLongniddry,
	StationCodeLongport:                          StationNameLongport,
	StationCodeLongton:                           StationNameLongton,
	StationCodeLooe:                              StationNameLooe,
	StationCodeLostock:                           StationNameLostock,
	StationCodeLostockGralam:                     StationNameLostockGralam,
	StationCodeLostockHall:                       StationNameLostockHall,
	StationCodeLostwithiel:                       StationNameLostwithiel,
	StationCodeLoughborough:                      StationNameLoughborough,
	StationCodeLoughboroughJunction:              StationNameLoughboroughJunction,
	StationCodeLowdham:                           StationNameLowdham,
	StationCodeLowerSydenham:                     StationNameLowerSydenham,
	StationCodeLowestoft:                         StationNameLowestoft,
	StationCodeLudlow:                            StationNameLudlow,
	StationCodeLuton:                             StationNameLuton,
	StationCodeLutonAirportParkway:               StationNameLutonAirportParkway,
	StationCodeLuxulyan:                          StationNameLuxulyan,
	StationCodeLydney:                            StationNameLydney,
	StationCodeLyeWestMidlands:                   StationNameLyeWestMidlands,
	StationCodeLymingtonPier:                     StationNameLymingtonPier,
	StationCodeLymingtonTown:                     StationNameLymingtonTown,
	StationCodeLympstoneCommando:                 StationNameLympstoneCommando,
	StationCodeLympstoneVillage:                  StationNameLympstoneVillage,
	StationCodeLytham:                            StationNameLytham,
	StationCodeMacclesfield:                      StationNameMacclesfield,
	StationCodeMachynlleth:                       StationNameMachynlleth,
	StationCodeMaesteg:                           StationNameMaesteg,
	StationCodeMaestegEwennyRoad:                 StationNameMaestegEwennyRoad,
	StationCodeMaghull:                           StationNameMaghull,
	StationCodeMaidenNewton:                      StationNameMaidenNewton,
	StationCodeMaidenhead:                        StationNameMaidenhead,
	StationCodeMaidstoneBarracks:                 StationNameMaidstoneBarracks,
	StationCodeMaidstoneEast:                     StationNameMaidstoneEast,
	StationCodeMaidstoneWest:                     StationNameMaidstoneWest,
	StationCodeMaldenManor:                       StationNameMaldenManor,
	StationCodeMallaig:                           StationNameMallaig,
	StationCodeMalton:                            StationNameMalton,
	StationCodeMalvernLink:                       StationNameMalvernLink,
	StationCodeManchesterAirport:                 StationNameManchesterAirport,
	StationCodeManchesterOxfordRoad:              StationNameManchesterOxfordRoad,
	StationCodeManchesterPiccadilly:              StationNameManchesterPiccadilly,
	StationCodeManchesterUnitedFootballGround:    StationNameManchesterUnitedFootballGround,
	StationCodeManchesterVictoria:                StationNameManchesterVictoria,
	StationCodeManea:                             StationNameManea,
	StationCodeManningtree:                       StationNameManningtree,
	StationCodeManorPark:                         StationNameManorPark,
	StationCodeManorRoad:                         StationNameManorRoad,
	StationCodeManorbier:                         StationNameManorbier,
	StationCodeManors:                            StationNameManors,
	StationCodeMansfield:                         StationNameMansfield,
	StationCodeMansfieldWoodhouse:                StationNameMansfieldWoodhouse,
	StationCodeMarch:                             StationNameMarch,
	StationCodeMardenKent:                        StationNameMardenKent,
	StationCodeMargate:                           StationNameMargate,
	StationCodeMarketHarborough:                  StationNameMarketHarborough,
	StationCodeMarketRasen:                       StationNameMarketRasen,
	StationCodeMarkinch:                          StationNameMarkinch,
	StationCodeMarksTey:                          StationNameMarksTey,
	StationCodeMarlow:                            StationNameMarlow,
	StationCodeMarple:                            StationNameMarple,
	StationCodeMarsdenYorks:                      StationNameMarsdenYorks,
	StationCodeMarske:                            StationNameMarske,
	StationCodeMarstonGreen:                      StationNameMarstonGreen,
	StationCodeMartinMill:                        StationNameMartinMill,
	StationCodeMartinsHeron:                      StationNameMartinsHeron,
	StationCodeMarton:                            StationNameMarton,
	StationCodeMaryhill:                          StationNameMaryhill,
	StationCodeMaryland:                          StationNameMaryland,
	StationCodeMaryport:                          StationNameMaryport,
	StationCodeMatlock:                           StationNameMatlock,
	StationCodeMatlockBath:                       StationNameMatlockBath,
	StationCodeMauldethRoad:                      StationNameMauldethRoad,
	StationCodeMaxwellPark:                       StationNameMaxwellPark,
	StationCodeMaybole:                           StationNameMaybole,
	StationCodeMazeHill:                          StationNameMazeHill,
	StationCodeMeadowhall:                        StationNameMeadowhall,
	StationCodeMeldreth:                          StationNameMeldreth,
	StationCodeMelksham:                          StationNameMelksham,
	StationCodeMeltonMowbray:                     StationNameMeltonMowbray,
	StationCodeMeltonSuffolk:                     StationNameMeltonSuffolk,
	StationCodeMenheniot:                         StationNameMenheniot,
	StationCodeMenston:                           StationNameMenston,
	StationCodeMeols:                             StationNameMeols,
	StationCodeMeolsCop:                          StationNameMeolsCop,
	StationCodeMeopham:                           StationNameMeopham,
	StationCodeMerryton:                          StationNameMerryton,
	StationCodeMerstham:                          StationNameMerstham,
	StationCodeMerthyrTydfil:                     StationNameMerthyrTydfil,
	StationCodeMerthyrVale:                       StationNameMerthyrVale,
	StationCodeMetheringham:                      StationNameMetheringham,
	StationCodeMetroCentre:                       StationNameMetroCentre,
	StationCodeMexborough:                        StationNameMexborough,
	StationCodeMicheldever:                       StationNameMicheldever,
	StationCodeMicklefield:                       StationNameMicklefield,
	StationCodeMiddlesbrough:                     StationNameMiddlesbrough,
	StationCodeMiddlewood:                        StationNameMiddlewood,
	StationCodeMidgham:                           StationNameMidgham,
	StationCodeMilfordHaven:                      StationNameMilfordHaven,
	StationCodeMilfordSurrey:                     StationNameMilfordSurrey,
	StationCodeMillHillBroadway:                  StationNameMillHillBroadway,
	StationCodeMillHillLancs:                     StationNameMillHillLancs,
	StationCodeMillbrookBeds:                     StationNameMillbrookBeds,
	StationCodeMillbrookHants:                    StationNameMillbrookHants,
	StationCodeMillikenPark:                      StationNameMillikenPark,
	StationCodeMillom:                            StationNameMillom,
	StationCodeMillsHillManchester:               StationNameMillsHillManchester,
	StationCodeMilngavie:                         StationNameMilngavie,
	StationCodeMiltonKeynesCentral:               StationNameMiltonKeynesCentral,
	StationCodeMinffordd:                         StationNameMinffordd,
	StationCodeMinster:                           StationNameMinster,
	StationCodeMirfield:                          StationNameMirfield,
	StationCodeMistley:                           StationNameMistley,
	StationCodeMitchamEastfields:                 StationNameMitchamEastfields,
	StationCodeMitchamJunction:                   StationNameMitchamJunction,
	StationCodeMobberley:                         StationNameMobberley,
	StationCodeMonifieth:                         StationNameMonifieth,
	StationCodeMonksRisborough:                   StationNameMonksRisborough,
	StationCodeMontpelier:                        StationNameMontpelier,
	StationCodeMontrose:                          StationNameMontrose,
	StationCodeMoorfields:                        StationNameMoorfields,
	StationCodeMoorgate:                          StationNameMoorgate,
	StationCodeMoorside:                          StationNameMoorside,
	StationCodeMoorthorpe:                        StationNameMoorthorpe,
	StationCodeMorar:                             StationNameMorar,
	StationCodeMorchardRoad:                      StationNameMorchardRoad,
	StationCodeMordenSouth:                       StationNameMordenSouth,
	StationCodeMorecambe:                         StationNameMorecambe,
	StationCodeMoretonDorset:                     StationNameMoretonDorset,
	StationCodeMoretonMerseyside:                 StationNameMoretonMerseyside,
	StationCodeMoretoninMarsh:                    StationNameMoretoninMarsh,
	StationCodeMorfaMawddach:                     StationNameMorfaMawddach,
	StationCodeMorley:                            StationNameMorley,
	StationCodeMorpeth:                           StationNameMorpeth,
	StationCodeMortimer:                          StationNameMortimer,
	StationCodeMortlake:                          StationNameMortlake,
	StationCodeMosesGate:                         StationNameMosesGate,
	StationCodeMossSide:                          StationNameMossSide,
	StationCodeMossleyHill:                       StationNameMossleyHill,
	StationCodeMossleyManchester:                 StationNameMossleyManchester,
	StationCodeMosspark:                          StationNameMosspark,
	StationCodeMoston:                            StationNameMoston,
	StationCodeMotherwell:                        StationNameMotherwell,
	StationCodeMotspurPark:                       StationNameMotspurPark,
	StationCodeMottingham:                        StationNameMottingham,
	StationCodeMottisfontAndDunbridge:            StationNameMottisfontAndDunbridge,
	StationCodeMouldsworth:                       StationNameMouldsworth,
	StationCodeMoulsecoomb:                       StationNameMoulsecoomb,
	StationCodeMountFlorida:                      StationNameMountFlorida,
	StationCodeMountVernon:                       StationNameMountVernon,
	StationCodeMountainAsh:                       StationNameMountainAsh,
	StationCodeMuirend:                           StationNameMuirend,
	StationCodeMuirofOrd:                         StationNameMuirofOrd,
	StationCodeMusselburgh:                       StationNameMusselburgh,
	StationCodeMytholmroyd:                       StationNameMytholmroyd,
	StationCodeNafferton:                         StationNameNafferton,
	StationCodeNailseaAndBackwell:                StationNameNailseaAndBackwell,
	StationCodeNairn:                             StationNameNairn,
	StationCodeNantwich:                          StationNameNantwich,
	StationCodeNarberth:                          StationNameNarberth,
	StationCodeNarborough:                        StationNameNarborough,
	StationCodeNavigationRoad:                    StationNameNavigationRoad,
	StationCodeNeath:                             StationNameNeath,
	StationCodeNeedhamMarket:                     StationNameNeedhamMarket,
	StationCodeNeilston:                          StationNameNeilston,
	StationCodeNelson:                            StationNameNelson,
	StationCodeNeston:                            StationNameNeston,
	StationCodeNetherfield:                       StationNameNetherfield,
	StationCodeNethertown:                        StationNameNethertown,
	StationCodeNetley:                            StationNameNetley,
	StationCodeNewBarnet:                         StationNameNewBarnet,
	StationCodeNewBeckenham:                      StationNameNewBeckenham,
	StationCodeNewBrighton:                       StationNameNewBrighton,
	StationCodeNewClee:                           StationNameNewClee,
	StationCodeNewCross:                          StationNameNewCross,
	StationCodeNewCrossGate:                      StationNameNewCrossGate,
	StationCodeNewCumnock:                        StationNameNewCumnock,
	StationCodeNewEltham:                         StationNameNewEltham,
	StationCodeNewHolland:                        StationNameNewHolland,
	StationCodeNewHythe:                          StationNameNewHythe,
	StationCodeNewLane:                           StationNameNewLane,
	StationCodeNewMalden:                         StationNameNewMalden,
	StationCodeNewMillsCentral:                   StationNameNewMillsCentral,
	StationCodeNewMillsNewtown:                   StationNameNewMillsNewtown,
	StationCodeNewMilton:                         StationNameNewMilton,
	StationCodeNewPudsey:                         StationNameNewPudsey,
	StationCodeNewSouthgate:                      StationNameNewSouthgate,
	StationCodeNewarkCastle:                      StationNameNewarkCastle,
	StationCodeNewarkNorthGate:                   StationNameNewarkNorthGate,
	StationCodeNewbridge:                         StationNameNewbridge,
	StationCodeNewbury:                           StationNameNewbury,
	StationCodeNewburyRacecourse:                 StationNameNewburyRacecourse,
	StationCodeNewcastle:                         StationNameNewcastle,
	StationCodeNewcourt:                          StationNameNewcourt,
	StationCodeNewcraighall:                      StationNameNewcraighall,
	StationCodeNewhavenHarbour:                   StationNameNewhavenHarbour,
	StationCodeNewhavenTown:                      StationNameNewhavenTown,
	StationCodeNewington:                         StationNameNewington,
	StationCodeNewmarket:                         StationNameNewmarket,
	StationCodeNewportEssex:                      StationNameNewportEssex,
	StationCodeNewportSouthWales:                 StationNameNewportSouthWales,
	StationCodeNewquay:                           StationNameNewquay,
	StationCodeNewstead:                          StationNameNewstead,
	StationCodeNewtonAbbot:                       StationNameNewtonAbbot,
	StationCodeNewtonAycliffe:                    StationNameNewtonAycliffe,
	StationCodeNewtonLanark:                      StationNameNewtonLanark,
	StationCodeNewtonStCyres:                     StationNameNewtonStCyres,
	StationCodeNewtonforHyde:                     StationNameNewtonforHyde,
	StationCodeNewtongrange:                      StationNameNewtongrange,
	StationCodeNewtonleWillows:                   StationNameNewtonleWillows,
	StationCodeNewtonmore:                        StationNameNewtonmore,
	StationCodeNewtononAyr:                       StationNameNewtononAyr,
	StationCodeNewtownPowys:                      StationNameNewtownPowys,
	StationCodeNinianPark:                        StationNameNinianPark,
	StationCodeNitshill:                          StationNameNitshill,
	StationCodeNorbiton:                          StationNameNorbiton,
	StationCodeNorbury:                           StationNameNorbury,
	StationCodeNormansBay:                        StationNameNormansBay,
	StationCodeNormanton:                         StationNameNormanton,
	StationCodeNorthBerwick:                      StationNameNorthBerwick,
	StationCodeNorthCamp:                         StationNameNorthCamp,
	StationCodeNorthDulwich:                      StationNameNorthDulwich,
	StationCodeNorthFambridge:                    StationNameNorthFambridge,
	StationCodeNorthLlanrwst:                     StationNameNorthLlanrwst,
	StationCodeNorthQueensferry:                  StationNameNorthQueensferry,
	StationCodeNorthRoadDarlington:               StationNameNorthRoadDarlington,
	StationCodeNorthSheen:                        StationNameNorthSheen,
	StationCodeNorthWalsham:                      StationNameNorthWalsham,
	StationCodeNorthWembley:                      StationNameNorthWembley,
	StationCodeNorthallerton:                     StationNameNorthallerton,
	StationCodeNorthampton:                       StationNameNorthampton,
	StationCodeNorthfield:                        StationNameNorthfield,
	StationCodeNorthfleet:                        StationNameNorthfleet,
	StationCodeNortholtPark:                      StationNameNortholtPark,
	StationCodeNorthumberlandPark:                StationNameNorthumberlandPark,
	StationCodeNorthwich:                         StationNameNorthwich,
	StationCodeNortonBridge:                      StationNameNortonBridge,
	StationCodeNorwich:                           StationNameNorwich,
	StationCodeNorwoodJunction:                   StationNameNorwoodJunction,
	StationCodeNottingham:                        StationNameNottingham,
	StationCodeNuneaton:                          StationNameNuneaton,
	StationCodeNunhead:                           StationNameNunhead,
	StationCodeNunthorpe:                         StationNameNunthorpe,
	StationCodeNutbourne:                         StationNameNutbourne,
	StationCodeNutfield:                          StationNameNutfield,
	StationCodeOakengates:                        StationNameOakengates,
	StationCodeOakham:                            StationNameOakham,
	StationCodeOakleighPark:                      StationNameOakleighPark,
	StationCodeOban:                              StationNameOban,
	StationCodeOckendon:                          StationNameOckendon,
	StationCodeOckley:                            StationNameOckley,
	StationCodeOkehampton:                        StationNameOkehampton,
	StationCodeOldHill:                           StationNameOldHill,
	StationCodeOldRoan:                           StationNameOldRoan,
	StationCodeOldStreet:                         StationNameOldStreet,
	StationCodeOldfieldPark:                      StationNameOldfieldPark,
	StationCodeOlton:                             StationNameOlton,
	StationCodeOre:                               StationNameOre,
	StationCodeOrmskirk:                          StationNameOrmskirk,
	StationCodeOrpington:                         StationNameOrpington,
	StationCodeOrrell:                            StationNameOrrell,
	StationCodeOrrellPark:                        StationNameOrrellPark,
	StationCodeOtford:                            StationNameOtford,
	StationCodeOultonBroadNorth:                  StationNameOultonBroadNorth,
	StationCodeOultonBroadSouth:                  StationNameOultonBroadSouth,
	StationCodeOutwood:                           StationNameOutwood,
	StationCodeOverpool:                          StationNameOverpool,
	StationCodeOverton:                           StationNameOverton,
	StationCodeOxenholmeLakeDistrict:             StationNameOxenholmeLakeDistrict,
	StationCodeOxford:                            StationNameOxford,
	StationCodeOxfordParkway:                     StationNameOxfordParkway,
	StationCodeOxshott:                           StationNameOxshott,
	StationCodeOxted:                             StationNameOxted,
	StationCodePaddockWood:                       StationNamePaddockWood,
	StationCodePadgate:                           StationNamePadgate,
	StationCodePaignton:                          StationNamePaignton,
	StationCodePaisleyCanal:                      StationNamePaisleyCanal,
	StationCodePaisleyGilmourStreet:              StationNamePaisleyGilmourStreet,
	StationCodePaisleyStJames:                    StationNamePaisleyStJames,
	StationCodePalmersGreen:                      StationNamePalmersGreen,
	StationCodePangbourne:                        StationNamePangbourne,
	StationCodePannal:                            StationNamePannal,
	StationCodePantyffynnon:                      StationNamePantyffynnon,
	StationCodePar:                               StationNamePar,
	StationCodeParbold:                           StationNameParbold,
	StationCodeParkStreet:                        StationNameParkStreet,
	StationCodeParkstoneDorset:                   StationNameParkstoneDorset,
	StationCodeParsonStreet:                      StationNameParsonStreet,
	StationCodePartick:                           StationNamePartick,
	StationCodeParton:                            StationNameParton,
	StationCodePatchway:                          StationNamePatchway,
	StationCodePatricroft:                        StationNamePatricroft,
	StationCodePatterton:                         StationNamePatterton,
	StationCodePeartree:                          StationNamePeartree,
	StationCodePeckhamRye:                        StationNamePeckhamRye,
	StationCodePegswood:                          StationNamePegswood,
	StationCodePemberton:                         StationNamePemberton,
	StationCodePembreyAndBurryPort:               StationNamePembreyAndBurryPort,
	StationCodePembroke:                          StationNamePembroke,
	StationCodePembrokeDock:                      StationNamePembrokeDock,
	StationCodePenally:                           StationNamePenally,
	StationCodePenarth:                           StationNamePenarth,
	StationCodePencoed:                           StationNamePencoed,
	StationCodePengam:                            StationNamePengam,
	StationCodePengeEast:                         StationNamePengeEast,
	StationCodePengeWest:                         StationNamePengeWest,
	StationCodePenhelig:                          StationNamePenhelig,
	StationCodePenistone:                         StationNamePenistone,
	StationCodePenkridge:                         StationNamePenkridge,
	StationCodePenmaenmawr:                       StationNamePenmaenmawr,
	StationCodePenmere:                           StationNamePenmere,
	StationCodePenrhiwceiber:                     StationNamePenrhiwceiber,
	StationCodePenrhyndeudraeth:                  StationNamePenrhyndeudraeth,
	StationCodePenrithNorthLakes:                 StationNamePenrithNorthLakes,
	StationCodePenrynCornwall:                    StationNamePenrynCornwall,
	StationCodePensarnGwynedd:                    StationNamePensarnGwynedd,
	StationCodePenshurst:                         StationNamePenshurst,
	StationCodePentreBach:                        StationNamePentreBach,
	StationCodePenyBont:                          StationNamePenyBont,
	StationCodePenychain:                         StationNamePenychain,
	StationCodePenyffordd:                        StationNamePenyffordd,
	StationCodePenzance:                          StationNamePenzance,
	StationCodePerranwell:                        StationNamePerranwell,
	StationCodePerryBarr:                         StationNamePerryBarr,
	StationCodePershore:                          StationNamePershore,
	StationCodePerth:                             StationNamePerth,
	StationCodePeterborough:                      StationNamePeterborough,
	StationCodePetersfield:                       StationNamePetersfield,
	StationCodePettsWood:                         StationNamePettsWood,
	StationCodePevenseyAndWestham:                StationNamePevenseyAndWestham,
	StationCodePevenseyBay:                       StationNamePevenseyBay,
	StationCodePewsey:                            StationNamePewsey,
	StationCodePilning:                           StationNamePilning,
	StationCodePinhoe:                            StationNamePinhoe,
	StationCodePitlochry:                         StationNamePitlochry,
	StationCodePitsea:                            StationNamePitsea,
	StationCodePleasington:                       StationNamePleasington,
	StationCodePlockton:                          StationNamePlockton,
	StationCodePluckley:                          StationNamePluckley,
	StationCodePlumley:                           StationNamePlumley,
	StationCodePlumpton:                          StationNamePlumpton,
	StationCodePlumstead:                         StationNamePlumstead,
	StationCodePlymouth:                          StationNamePlymouth,
	StationCodePokesdown:                         StationNamePokesdown,
	StationCodePolegate:                          StationNamePolegate,
	StationCodePolesworth:                        StationNamePolesworth,
	StationCodePollokshawsEast:                   StationNamePollokshawsEast,
	StationCodePollokshawsWest:                   StationNamePollokshawsWest,
	StationCodePollokshieldsEast:                 StationNamePollokshieldsEast,
	StationCodePollokshieldsWest:                 StationNamePollokshieldsWest,
	StationCodePolmont:                           StationNamePolmont,
	StationCodePolsloeBridge:                     StationNamePolsloeBridge,
	StationCodePondersEnd:                        StationNamePondersEnd,
	StationCodePontarddulais:                     StationNamePontarddulais,
	StationCodePontefractBaghill:                 StationNamePontefractBaghill,
	StationCodePontefractMonkhill:                StationNamePontefractMonkhill,
	StationCodePontefractTanshelf:                StationNamePontefractTanshelf,
	StationCodePontlottyn:                        StationNamePontlottyn,
	StationCodePontyPant:                         StationNamePontyPant,
	StationCodePontyclun:                         StationNamePontyclun,
	StationCodePontypoolAndNewInn:                StationNamePontypoolAndNewInn,
	StationCodePontypridd:                        StationNamePontypridd,
	StationCodePoole:                             StationNamePoole,
	StationCodePoppleton:                         StationNamePoppleton,
	StationCodePortGlasgow:                       StationNamePortGlasgow,
	StationCodePortSunlight:                      StationNamePortSunlight,
	StationCodePortTalbotParkway:                 StationNamePortTalbotParkway,
	StationCodePortchester:                       StationNamePortchester,
	StationCodePorth:                             StationNamePorth,
	StationCodePorthmadog:                        StationNamePorthmadog,
	StationCodePortlethen:                        StationNamePortlethen,
	StationCodePortslade:                         StationNamePortslade,
	StationCodePortsmouthAndSouthsea:             StationNamePortsmouthAndSouthsea,
	StationCodePortsmouthArms:                    StationNamePortsmouthArms,
	StationCodePortsmouthHarbour:                 StationNamePortsmouthHarbour,
	StationCodePossilparkAndParkhouse:            StationNamePossilparkAndParkhouse,
	StationCodePottersBar:                        StationNamePottersBar,
	StationCodePoultonleFylde:                    StationNamePoultonleFylde,
	StationCodePoynton:                           StationNamePoynton,
	StationCodePrees:                             StationNamePrees,
	StationCodePrescot:                           StationNamePrescot,
	StationCodePrestatyn:                         StationNamePrestatyn,
	StationCodePrestbury:                         StationNamePrestbury,
	StationCodePrestonLancs:                      StationNamePrestonLancs,
	StationCodePrestonPark:                       StationNamePrestonPark,
	StationCodePrestonpans:                       StationNamePrestonpans,
	StationCodePrestwickInternationalAirport:     StationNamePrestwickInternationalAirport,
	StationCodePrestwickTown:                     StationNamePrestwickTown,
	StationCodePriesthillAndDarnley:              StationNamePriesthillAndDarnley,
	StationCodePrincesRisborough:                 StationNamePrincesRisborough,
	StationCodePrittlewell:                       StationNamePrittlewell,
	StationCodePrudhoe:                           StationNamePrudhoe,
	StationCodePulborough:                        StationNamePulborough,
	StationCodePurfleet:                          StationNamePurfleet,
	StationCodePurley:                            StationNamePurley,
	StationCodePurleyOaks:                        StationNamePurleyOaks,
	StationCodePutney:                            StationNamePutney,
	StationCodePwllheli:                          StationNamePwllheli,
	StationCodePyeCorner:                         StationNamePyeCorner,
	StationCodePyle:                              StationNamePyle,
	StationCodeQuakersYard:                       StationNameQuakersYard,
	StationCodeQueenborough:                      StationNameQueenborough,
	StationCodeQueensParkGlasgow:                 StationNameQueensParkGlasgow,
	StationCodeQueensParkLondon:                  StationNameQueensParkLondon,
	StationCodeQueensRoadPeckham:                 StationNameQueensRoadPeckham,
	StationCodeQueenstownRoadBattersea:           StationNameQueenstownRoadBattersea,
	StationCodeQuintrellDowns:                    StationNameQuintrellDowns,
	StationCodeRadcliffeonTrent:                  StationNameRadcliffeonTrent,
	StationCodeRadlett:                           StationNameRadlett,
	StationCodeRadley:                            StationNameRadley,
	StationCodeRadyr:                             StationNameRadyr,
	StationCodeRainford:                          StationNameRainford,
	StationCodeRainhamEssex:                      StationNameRainhamEssex,
	StationCodeRainhamKent:                       StationNameRainhamKent,
	StationCodeRainhill:                          StationNameRainhill,
	StationCodeRamsgate:                          StationNameRamsgate,
	StationCodeRamsgreaveAndWilpshire:            StationNameRamsgreaveAndWilpshire,
	StationCodeRannoch:                           StationNameRannoch,
	StationCodeRauceby:                           StationNameRauceby,
	StationCodeRavenglassforEskdale:              StationNameRavenglassforEskdale,
	StationCodeRavensbourne:                      StationNameRavensbourne,
	StationCodeRavensthorpe:                      StationNameRavensthorpe,
	StationCodeRawcliffe:                         StationNameRawcliffe,
	StationCodeRayleigh:                          StationNameRayleigh,
	StationCodeRaynesPark:                        StationNameRaynesPark,
	StationCodeReading:                           StationNameReading,
	StationCodeReadingWest:                       StationNameReadingWest,
	StationCodeRectoryRoad:                       StationNameRectoryRoad,
	StationCodeRedbridge:                         StationNameRedbridge,
	StationCodeRedcarBritishSteel:                StationNameRedcarBritishSteel,
	StationCodeRedcarCentral:                     StationNameRedcarCentral,
	StationCodeRedcarEast:                        StationNameRedcarEast,
	StationCodeReddishNorth:                      StationNameReddishNorth,
	StationCodeReddishSouth:                      StationNameReddishSouth,
	StationCodeRedditch:                          StationNameRedditch,
	StationCodeRedhill:                           StationNameRedhill,
	StationCodeRedland:                           StationNameRedland,
	StationCodeRedruth:                           StationNameRedruth,
	StationCodeReedhamNorfolk:                    StationNameReedhamNorfolk,
	StationCodeReedhamSurrey:                     StationNameReedhamSurrey,
	StationCodeReigate:                           StationNameReigate,
	StationCodeRenton:                            StationNameRenton,
	StationCodeRetford:                           StationNameRetford,
	StationCodeRhiwbina:                          StationNameRhiwbina,
	StationCodeRhooseCardiffInternationalAirport: StationNameRhooseCardiffInternationalAirport,
	StationCodeRhosneigr:                         StationNameRhosneigr,
	StationCodeRhyl:                              StationNameRhyl,
	StationCodeRhymney:                           StationNameRhymney,
	StationCodeRibblehead:                        StationNameRibblehead,
	StationCodeRiceLane:                          StationNameRiceLane,
	StationCodeRichmondLondon:                    StationNameRichmondLondon,
	StationCodeRickmansworth:                     StationNameRickmansworth,
	StationCodeRiddlesdown:                       StationNameRiddlesdown,
	StationCodeRidgmont:                          StationNameRidgmont,
	StationCodeRidingMill:                        StationNameRidingMill,
	StationCodeRiscaAndPontymister:               StationNameRiscaAndPontymister,
	StationCodeRishton:                           StationNameRishton,
	StationCodeRobertsbridge:                     StationNameRobertsbridge,
	StationCodeRoby:                              StationNameRoby,
	StationCodeRochdale:                          StationNameRochdale,
	StationCodeRoche:                             StationNameRoche,
	StationCodeRochester:                         StationNameRochester,
	StationCodeRochford:                          StationNameRochford,
	StationCodeRockFerry:                         StationNameRockFerry,
	StationCodeRogart:                            StationNameRogart,
	StationCodeRogerstone:                        StationNameRogerstone,
	StationCodeRolleston:                         StationNameRolleston,
	StationCodeRomanBridge:                       StationNameRomanBridge,
	StationCodeRomford:                           StationNameRomford,
	StationCodeRomiley:                           StationNameRomiley,
	StationCodeRomsey:                            StationNameRomsey,
	StationCodeRoose:                             StationNameRoose,
	StationCodeRoseGrove:                         StationNameRoseGrove,
	StationCodeRoseHillMarple:                    StationNameRoseHillMarple,
	StationCodeRosyth:                            StationNameRosyth,
	StationCodeRotherhamCentral:                  StationNameRotherhamCentral,
	StationCodeRotherhithe:                       StationNameRotherhithe,
	StationCodeRoughtonRoad:                      StationNameRoughtonRoad,
	StationCodeRowlandsCastle:                    StationNameRowlandsCastle,
	StationCodeRowleyRegis:                       StationNameRowleyRegis,
	StationCodeRoyBridge:                         StationNameRoyBridge,
	StationCodeRoydon:                            StationNameRoydon,
	StationCodeRoyston:                           StationNameRoyston,
	StationCodeRuabon:                            StationNameRuabon,
	StationCodeRufford:                           StationNameRufford,
	StationCodeRugby:                             StationNameRugby,
	StationCodeRugeleyTown:                       StationNameRugeleyTown,
	StationCodeRugeleyTrentValley:                StationNameRugeleyTrentValley,
	StationCodeRuncorn:                           StationNameRuncorn,
	StationCodeRuncornEast:                       StationNameRuncornEast,
	StationCodeRuskington:                        StationNameRuskington,
	StationCodeRuswarp:                           StationNameRuswarp,
	StationCodeRutherglen:                        StationNameRutherglen,
	StationCodeRydeEsplanade:                     StationNameRydeEsplanade,
	StationCodeRydePierHead:                      StationNameRydePierHead,
	StationCodeRydeStJohnsRoad:                   StationNameRydeStJohnsRoad,
	StationCodeRyderBrow:                         StationNameRyderBrow,
	StationCodeRyeHouse:                          StationNameRyeHouse,
	StationCodeRyeSussex:                         StationNameRyeSussex,
	StationCodeSalfordCentral:                    StationNameSalfordCentral,
	StationCodeSalfordCrescent:                   StationNameSalfordCrescent,
	StationCodeSalfordsSurrey:                    StationNameSalfordsSurrey,
	StationCodeSalhouse:                          StationNameSalhouse,
	StationCodeSalisbury:                         StationNameSalisbury,
	StationCodeSaltaire:                          StationNameSaltaire,
	StationCodeSaltash:                           StationNameSaltash,
	StationCodeSaltburn:                          StationNameSaltburn,
	StationCodeSaltcoats:                         StationNameSaltcoats,
	StationCodeSaltmarshe:                        StationNameSaltmarshe,
	StationCodeSalwick:                           StationNameSalwick,
	StationCodeSampfordCourtenay:                 StationNameSampfordCourtenay,
	StationCodeSandalAndAgbrigg:                  StationNameSandalAndAgbrigg,
	StationCodeSandbach:                          StationNameSandbach,
	StationCodeSanderstead:                       StationNameSanderstead,
	StationCodeSandhills:                         StationNameSandhills,
	StationCodeSandhurstBerks:                    StationNameSandhurstBerks,
	StationCodeSandling:                          StationNameSandling,
	StationCodeSandown:                           StationNameSandown,
	StationCodeSandplace:                         StationNameSandplace,
	StationCodeSandwellAndDudley:                 StationNameSandwellAndDudley,
	StationCodeSandwich:                          StationNameSandwich,
	StationCodeSandy:                             StationNameSandy,
	StationCodeSankeyforPenketh:                  StationNameSankeyforPenketh,
	StationCodeSanquhar:                          StationNameSanquhar,
	StationCodeSarn:                              StationNameSarn,
	StationCodeSaundersfoot:                      StationNameSaundersfoot,
	StationCodeSaunderton:                        StationNameSaunderton,
	StationCodeSawbridgeworth:                    StationNameSawbridgeworth,
	StationCodeSaxilby:                           StationNameSaxilby,
	StationCodeSaxmundham:                        StationNameSaxmundham,
	StationCodeScarborough:                       StationNameScarborough,
	StationCodeScotscalder:                       StationNameScotscalder,
	StationCodeScotstounhill:                     StationNameScotstounhill,
	StationCodeScunthorpe:                        StationNameScunthorpe,
	StationCodeSeaMills:                          StationNameSeaMills,
	StationCodeSeafordSussex:                     StationNameSeafordSussex,
	StationCodeSeaforthAndLitherland:             StationNameSeaforthAndLitherland,
	StationCodeSeaham:                            StationNameSeaham,
	StationCodeSeamer:                            StationNameSeamer,
	StationCodeSeascale:                          StationNameSeascale,
	StationCodeSeatonCarew:                       StationNameSeatonCarew,
	StationCodeSeerGreenAndJordans:               StationNameSeerGreenAndJordans,
	StationCodeSelby:                             StationNameSelby,
	StationCodeSelhurst:                          StationNameSelhurst,
	StationCodeSellafield:                        StationNameSellafield,
	StationCodeSelling:                           StationNameSelling,
	StationCodeSellyOak:                          StationNameSellyOak,
	StationCodeSettle:                            StationNameSettle,
	StationCodeSevenKings:                        StationNameSevenKings,
	StationCodeSevenSisters:                      StationNameSevenSisters,
	StationCodeSevenoaks:                         StationNameSevenoaks,
	StationCodeSevernBeach:                       StationNameSevernBeach,
	StationCodeSevernTunnelJunction:              StationNameSevernTunnelJunction,
	StationCodeShadwell:                          StationNameShadwell,
	StationCodeShalfordSurrey:                    StationNameShalfordSurrey,
	StationCodeShanklin:                          StationNameShanklin,
	StationCodeShawfair:                          StationNameShawfair,
	StationCodeShawford:                          StationNameShawford,
	StationCodeShawlands:                         StationNameShawlands,
	StationCodeSheernessonSea:                    StationNameSheernessonSea,
	StationCodeSheffield:                         StationNameSheffield,
	StationCodeShelfordCambs:                     StationNameShelfordCambs,
	StationCodeShenfield:                         StationNameShenfield,
	StationCodeShenstone:                         StationNameShenstone,
	StationCodeShepherdsBush:                     StationNameShepherdsBush,
	StationCodeShepherdsWell:                     StationNameShepherdsWell,
	StationCodeShepley:                           StationNameShepley,
	StationCodeShepperton:                        StationNameShepperton,
	StationCodeShepreth:                          StationNameShepreth,
	StationCodeSherborne:                         StationNameSherborne,
	StationCodeSherburninElmet:                   StationNameSherburninElmet,
	StationCodeSheringham:                        StationNameSheringham,
	StationCodeShettleston:                       StationNameShettleston,
	StationCodeShieldmuir:                        StationNameShieldmuir,
	StationCodeShifnal:                           StationNameShifnal,
	StationCodeShildon:                           StationNameShildon,
	StationCodeShiplake:                          StationNameShiplake,
	StationCodeShipleyYorks:                      StationNameShipleyYorks,
	StationCodeShippeaHill:                       StationNameShippeaHill,
	StationCodeShipton:                           StationNameShipton,
	StationCodeShirebrook:                        StationNameShirebrook,
	StationCodeShirehampton:                      StationNameShirehampton,
	StationCodeShireoaks:                         StationNameShireoaks,
	StationCodeShirley:                           StationNameShirley,
	StationCodeShoeburyness:                      StationNameShoeburyness,
	StationCodeSholing:                           StationNameSholing,
	StationCodeShoreditchHighStreet:              StationNameShoreditchHighStreet,
	StationCodeShorehamKent:                      StationNameShorehamKent,
	StationCodeShorehambySea:                     StationNameShorehambySea,
	StationCodeShortlands:                        StationNameShortlands,
	StationCodeShotton:                           StationNameShotton,
	StationCodeShotts:                            StationNameShotts,
	StationCodeShrewsbury:                        StationNameShrewsbury,
	StationCodeSidcup:                            StationNameSidcup,
	StationCodeSileby:                            StationNameSileby,
	StationCodeSilecroft:                         StationNameSilecroft,
	StationCodeSilkstoneCommon:                   StationNameSilkstoneCommon,
	StationCodeSilverStreet:                      StationNameSilverStreet,
	StationCodeSilverdale:                        StationNameSilverdale,
	StationCodeSinger:                            StationNameSinger,
	StationCodeSittingbourne:                     StationNameSittingbourne,
	StationCodeSkegness:                          StationNameSkegness,
	StationCodeSkewen:                            StationNameSkewen,
	StationCodeSkipton:                           StationNameSkipton,
	StationCodeSladeGreen:                        StationNameSladeGreen,
	StationCodeSlaithwaite:                       StationNameSlaithwaite,
	StationCodeSlateford:                         StationNameSlateford,
	StationCodeSleaford:                          StationNameSleaford,
	StationCodeSleights:                          StationNameSleights,
	StationCodeSlough:                            StationNameSlough,
	StationCodeSmallHeath:                        StationNameSmallHeath,
	StationCodeSmallbrookJunction:                StationNameSmallbrookJunction,
	StationCodeSmethwickGaltonBridge:             StationNameSmethwickGaltonBridge,
	StationCodeSmethwickRolfeStreet:              StationNameSmethwickRolfeStreet,
	StationCodeSmithyBridge:                      StationNameSmithyBridge,
	StationCodeSnaith:                            StationNameSnaith,
	StationCodeSnodland:                          StationNameSnodland,
	StationCodeSnowdown:                          StationNameSnowdown,
	StationCodeSoleStreet:                        StationNameSoleStreet,
	StationCodeSolihull:                          StationNameSolihull,
	StationCodeSomerleyton:                       StationNameSomerleyton,
	StationCodeSouthActon:                        StationNameSouthActon,
	StationCodeSouthBank:                         StationNameSouthBank,
	StationCodeSouthBermondsey:                   StationNameSouthBermondsey,
	StationCodeSouthCroydon:                      StationNameSouthCroydon,
	StationCodeSouthElmsall:                      StationNameSouthElmsall,
	StationCodeSouthGreenford:                    StationNameSouthGreenford,
	StationCodeSouthGyle:                         StationNameSouthGyle,
	StationCodeSouthHampstead:                    StationNameSouthHampstead,
	StationCodeSouthKenton:                       StationNameSouthKenton,
	StationCodeSouthMerton:                       StationNameSouthMerton,
	StationCodeSouthMilford:                      StationNameSouthMilford,
	StationCodeSouthRuislip:                      StationNameSouthRuislip,
	StationCodeSouthTottenham:                    StationNameSouthTottenham,
	StationCodeSouthWigston:                      StationNameSouthWigston,
	StationCodeSouthWoodhamFerrers:               StationNameSouthWoodhamFerrers,
	StationCodeSouthall:                          StationNameSouthall,
	StationCodeSouthamptonAirportParkway:         StationNameSouthamptonAirportParkway,
	StationCodeSouthamptonCentral:                StationNameSouthamptonCentral,
	StationCodeSouthbourne:                       StationNameSouthbourne,
	StationCodeSouthbury:                         StationNameSouthbury,
	StationCodeSouthease:                         StationNameSouthease,
	StationCodeSouthendAirport:                   StationNameSouthendAirport,
	StationCodeSouthendCentral:                   StationNameSouthendCentral,
	StationCodeSouthendEast:                      StationNameSouthendEast,
	StationCodeSouthendVictoria:                  StationNameSouthendVictoria,
	StationCodeSouthminster:                      StationNameSouthminster,
	StationCodeSouthport:                         StationNameSouthport,
	StationCodeSouthwick:                         StationNameSouthwick,
	StationCodeSowerbyBridge:                     StationNameSowerbyBridge,
	StationCodeSpalding:                          StationNameSpalding,
	StationCodeSpeanBridge:                       StationNameSpeanBridge,
	StationCodeSpital:                            StationNameSpital,
	StationCodeSpondon:                           StationNameSpondon,
	StationCodeSpoonerRow:                        StationNameSpoonerRow,
	StationCodeSpringRoad:                        StationNameSpringRoad,
	StationCodeSpringburn:                        StationNameSpringburn,
	StationCodeSpringfield:                       StationNameSpringfield,
	StationCodeSquiresGate:                       StationNameSquiresGate,
	StationCodeStAlbansAbbey:                     StationNameStAlbansAbbey,
	StationCodeStAlbansCity:                      StationNameStAlbansCity,
	StationCodeStAndrewsRoad:                     StationNameStAndrewsRoad,
	StationCodeStAnnesonSea:                      StationNameStAnnesonSea,
	StationCodeStAustell:                         StationNameStAustell,
	StationCodeStBees:                            StationNameStBees,
	StationCodeStBudeauxFerryRoad:                StationNameStBudeauxFerryRoad,
	StationCodeStBudeauxVictoriaRoad:             StationNameStBudeauxVictoriaRoad,
	StationCodeStColumbRoad:                      StationNameStColumbRoad,
	StationCodeStDenys:                           StationNameStDenys,
	StationCodeStErth:                            StationNameStErth,
	StationCodeStGermans:                         StationNameStGermans,
	StationCodeStHelensCentral:                   StationNameStHelensCentral,
	StationCodeStHelensJunction:                  StationNameStHelensJunction,
	StationCodeStHelierSurrey:                    StationNameStHelierSurrey,
	StationCodeStIvesCornwall:                    StationNameStIvesCornwall,
	StationCodeStJamesParkExeter:                 StationNameStJamesParkExeter,
	StationCodeStJamesStreetWalthamstow:          StationNameStJamesStreetWalthamstow,
	StationCodeStJohnsLondon:                     StationNameStJohnsLondon,
	StationCodeStKeyneWishingWellHalt:            StationNameStKeyneWishingWellHalt,
	StationCodeStLeonardsWarriorSquare:           StationNameStLeonardsWarriorSquare,
	StationCodeStMargaretsHerts:                  StationNameStMargaretsHerts,
	StationCodeStMargaretsLondon:                 StationNameStMargaretsLondon,
	StationCodeStMaryCray:                        StationNameStMaryCray,
	StationCodeStMichaels:                        StationNameStMichaels,
	StationCodeStNeots:                           StationNameStNeots,
	StationCodeStafford:                          StationNameStafford,
	StationCodeStaines:                           StationNameStaines,
	StationCodeStallingborough:                   StationNameStallingborough,
	StationCodeStalybridge:                       StationNameStalybridge,
	StationCodeStamfordHill:                      StationNameStamfordHill,
	StationCodeStamfordLincs:                     StationNameStamfordLincs,
	StationCodeStanfordleHope:                    StationNameStanfordleHope,
	StationCodeStanlowAndThornton:                StationNameStanlowAndThornton,
	StationCodeStanstedAirport:                   StationNameStanstedAirport,
	StationCodeStanstedMountfitchet:              StationNameStanstedMountfitchet,
	StationCodeStaplehurst:                       StationNameStaplehurst,
	StationCodeStapletonRoad:                     StationNameStapletonRoad,
	StationCodeStarbeck:                          StationNameStarbeck,
	StationCodeStarcross:                         StationNameStarcross,
	StationCodeStaveleyCumbria:                   StationNameStaveleyCumbria,
	StationCodeStechford:                         StationNameStechford,
	StationCodeSteetonAndSilsden:                 StationNameSteetonAndSilsden,
	StationCodeStepps:                            StationNameStepps,
	StationCodeStevenage:                         StationNameStevenage,
	StationCodeStevenston:                        StationNameStevenston,
	StationCodeStewartby:                         StationNameStewartby,
	StationCodeStewarton:                         StationNameStewarton,
	StationCodeStirling:                          StationNameStirling,
	StationCodeStockport:                         StationNameStockport,
	StationCodeStocksfield:                       StationNameStocksfield,
	StationCodeStocksmoor:                        StationNameStocksmoor,
	StationCodeStockton:                          StationNameStockton,
	StationCodeStokeMandeville:                   StationNameStokeMandeville,
	StationCodeStokeNewington:                    StationNameStokeNewington,
	StationCodeStokeonTrent:                      StationNameStokeonTrent,
	StationCodeStoneCrossing:                     StationNameStoneCrossing,
	StationCodeStoneStaffs:                       StationNameStoneStaffs,
	StationCodeStonebridgePark:                   StationNameStonebridgePark,
	StationCodeStonegate:                         StationNameStonegate,
	StationCodeStonehaven:                        StationNameStonehaven,
	StationCodeStonehouse:                        StationNameStonehouse,
	StationCodeStoneleigh:                        StationNameStoneleigh,
	StationCodeStourbridgeJunction:               StationNameStourbridgeJunction,
	StationCodeStourbridgeTown:                   StationNameStourbridgeTown,
	StationCodeStow:                              StationNameStow,
	StationCodeStowmarket:                        StationNameStowmarket,
	StationCodeStranraer:                         StationNameStranraer,
	StationCodeStratfordInternational:            StationNameStratfordInternational,
	StationCodeStratfordLondon:                   StationNameStratfordLondon,
	StationCodeStratfordUponAvon:                 StationNameStratfordUponAvon,
	StationCodeStratfordUponAvonParkway:          StationNameStratfordUponAvonParkway,
	StationCodeStrathcarron:                      StationNameStrathcarron,
	StationCodeStrawberryHill:                    StationNameStrawberryHill,
	StationCodeStreathamCommon:                   StationNameStreathamCommon,
	StationCodeStreathamGreaterLondon:            StationNameStreathamGreaterLondon,
	StationCodeStreathamHill:                     StationNameStreathamHill,
	StationCodeStreethouse:                       StationNameStreethouse,
	StationCodeStrines:                           StationNameStrines,
	StationCodeStromeferry:                       StationNameStromeferry,
	StationCodeStrood:                            StationNameStrood,
	StationCodeStroudGloucs:                      StationNameStroudGloucs,
	StationCodeSturry:                            StationNameSturry,
	StationCodeStyal:                             StationNameStyal,
	StationCodeSudburyAndHarrowRoad:              StationNameSudburyAndHarrowRoad,
	StationCodeSudburyHillHarrow:                 StationNameSudburyHillHarrow,
	StationCodeSudburySuffolk:                    StationNameSudburySuffolk,
	StationCodeSugarLoaf:                         StationNameSugarLoaf,
	StationCodeSummerston:                        StationNameSummerston,
	StationCodeSunbury:                           StationNameSunbury,
	StationCodeSunderland:                        StationNameSunderland,
	StationCodeSundridgePark:                     StationNameSundridgePark,
	StationCodeSunningdale:                       StationNameSunningdale,
	StationCodeSunnymeads:                        StationNameSunnymeads,
	StationCodeSurbiton:                          StationNameSurbiton,
	StationCodeSurreyQuays:                       StationNameSurreyQuays,
	StationCodeSuttonColdfield:                   StationNameSuttonColdfield,
	StationCodeSuttonCommon:                      StationNameSuttonCommon,
	StationCodeSuttonParkway:                     StationNameSuttonParkway,
	StationCodeSuttonSurrey:                      StationNameSuttonSurrey,
	StationCodeSwale:                             StationNameSwale,
	StationCodeSwanley:                           StationNameSwanley,
	StationCodeSwanscombe:                        StationNameSwanscombe,
	StationCodeSwansea:                           StationNameSwansea,
	StationCodeSwanwick:                          StationNameSwanwick,
	StationCodeSway:                              StationNameSway,
	StationCodeSwaythling:                        StationNameSwaythling,
	StationCodeSwinderby:                         StationNameSwinderby,
	StationCodeSwindonWilts:                      StationNameSwindonWilts,
	StationCodeSwineshead:                        StationNameSwineshead,
	StationCodeSwintonManchester:                 StationNameSwintonManchester,
	StationCodeSwintonSouthYorks:                 StationNameSwintonSouthYorks,
	StationCodeSydenhamHill:                      StationNameSydenhamHill,
	StationCodeSydenhamLondon:                    StationNameSydenhamLondon,
	StationCodeSyonLane:                          StationNameSyonLane,
	StationCodeSyston:                            StationNameSyston,
	StationCodeTackley:                           StationNameTackley,
	StationCodeTadworth:                          StationNameTadworth,
	StationCodeTaffsWell:                         StationNameTaffsWell,
	StationCodeTain:                              StationNameTain,
	StationCodeTalsarnau:                         StationNameTalsarnau,
	StationCodeTalyCafn:                          StationNameTalyCafn,
	StationCodeTalybont:                          StationNameTalybont,
	StationCodeTameBridgeParkway:                 StationNameTameBridgeParkway,
	StationCodeTamworth:                          StationNameTamworth,
	StationCodeTaplow:                            StationNameTaplow,
	StationCodeTattenhamCorner:                   StationNameTattenhamCorner,
	StationCodeTaunton:                           StationNameTaunton,
	StationCodeTaynuilt:                          StationNameTaynuilt,
	StationCodeTeddington:                        StationNameTeddington,
	StationCodeTeessideAirport:                   StationNameTeessideAirport,
	StationCodeTeignmouth:                        StationNameTeignmouth,
	StationCodeTelfordCentral:                    StationNameTelfordCentral,
	StationCodeTemplecombe:                       StationNameTemplecombe,
	StationCodeTenby:                             StationNameTenby,
	StationCodeTeynham:                           StationNameTeynham,
	StationCodeThamesDitton:                      StationNameThamesDitton,
	StationCodeThatcham:                          StationNameThatcham,
	StationCodeThattoHeath:                       StationNameThattoHeath,
	StationCodeTheHawthorns:                      StationNameTheHawthorns,
	StationCodeTheLakesWarks:                     StationNameTheLakesWarks,
	StationCodeTheale:                            StationNameTheale,
	StationCodeTheobaldsGrove:                    StationNameTheobaldsGrove,
	StationCodeThetford:                          StationNameThetford,
	StationCodeThirsk:                            StationNameThirsk,
	StationCodeThornaby:                          StationNameThornaby,
	StationCodeThorneNorth:                       StationNameThorneNorth,
	StationCodeThorneSouth:                       StationNameThorneSouth,
	StationCodeThornford:                         StationNameThornford,
	StationCodeThornliebank:                      StationNameThornliebank,
	StationCodeThorntonAbbey:                     StationNameThorntonAbbey,
	StationCodeThorntonHeath:                     StationNameThorntonHeath,
	StationCodeThorntonhall:                      StationNameThorntonhall,
	StationCodeThorpeBay:                         StationNameThorpeBay,
	StationCodeThorpeCulvert:                     StationNameThorpeCulvert,
	StationCodeThorpeleSoken:                     StationNameThorpeleSoken,
	StationCodeThreeBridges:                      StationNameThreeBridges,
	StationCodeThreeOaks:                         StationNameThreeOaks,
	StationCodeThurgarton:                        StationNameThurgarton,
	StationCodeThurnscoe:                         StationNameThurnscoe,
	StationCodeThurso:                            StationNameThurso,
	StationCodeThurston:                          StationNameThurston,
	StationCodeTilburyTown:                       StationNameTilburyTown,
	StationCodeTileHill:                          StationNameTileHill,
	StationCodeTilehurst:                         StationNameTilehurst,
	StationCodeTipton:                            StationNameTipton,
	StationCodeTirPhil:                           StationNameTirPhil,
	StationCodeTisbury:                           StationNameTisbury,
	StationCodeTivertonParkway:                   StationNameTivertonParkway,
	StationCodeTodmorden:                         StationNameTodmorden,
	StationCodeTolworth:                          StationNameTolworth,
	StationCodeTonPentre:                         StationNameTonPentre,
	StationCodeTonbridge:                         StationNameTonbridge,
	StationCodeTondu:                             StationNameTondu,
	StationCodeTonfanau:                          StationNameTonfanau,
	StationCodeTonypandy:                         StationNameTonypandy,
	StationCodeTooting:                           StationNameTooting,
	StationCodeTopsham:                           StationNameTopsham,
	StationCodeTorquay:                           StationNameTorquay,
	StationCodeTorre:                             StationNameTorre,
	StationCodeTotnes:                            StationNameTotnes,
	StationCodeTottenhamHale:                     StationNameTottenhamHale,
	StationCodeTotton:                            StationNameTotton,
	StationCodeTownGreen:                         StationNameTownGreen,
	StationCodeTraffordPark:                      StationNameTraffordPark,
	StationCodeTrefforest:                        StationNameTrefforest,
	StationCodeTrefforestEstate:                  StationNameTrefforestEstate,
	StationCodeTrehafod:                          StationNameTrehafod,
	StationCodeTreherbert:                        StationNameTreherbert,
	StationCodeTreorchy:                          StationNameTreorchy,
	StationCodeTrimley:                           StationNameTrimley,
	StationCodeTring:                             StationNameTring,
	StationCodeTroedyrhiw:                        StationNameTroedyrhiw,
	StationCodeTroon:                             StationNameTroon,
	StationCodeTrowbridge:                        StationNameTrowbridge,
	StationCodeTruro:                             StationNameTruro,
	StationCodeTulloch:                           StationNameTulloch,
	StationCodeTulseHill:                         StationNameTulseHill,
	StationCodeTunbridgeWells:                    StationNameTunbridgeWells,
	StationCodeTurkeyStreet:                      StationNameTurkeyStreet,
	StationCodeTutburyAndHatton:                  StationNameTutburyAndHatton,
	StationCodeTweedbank:                         StationNameTweedbank,
	StationCodeTwickenham:                        StationNameTwickenham,
	StationCodeTwyford:                           StationNameTwyford,
	StationCodeTyCroes:                           StationNameTyCroes,
	StationCodeTyGlas:                            StationNameTyGlas,
	StationCodeTygwyn:                            StationNameTygwyn,
	StationCodeTyndrumLower:                      StationNameTyndrumLower,
	StationCodeTyseley:                           StationNameTyseley,
	StationCodeTywyn:                             StationNameTywyn,
	StationCodeUckfield:                          StationNameUckfield,
	StationCodeUddingston:                        StationNameUddingston,
	StationCodeUlceby:                            StationNameUlceby,
	StationCodeUlleskelf:                         StationNameUlleskelf,
	StationCodeUlverston:                         StationNameUlverston,
	StationCodeUmberleigh:                        StationNameUmberleigh,
	StationCodeUniversityBirmingham:              StationNameUniversityBirmingham,
	StationCodeUphall:                            StationNameUphall,
	StationCodeUpholland:                         StationNameUpholland,
	StationCodeUpminster:                         StationNameUpminster,
	StationCodeUpperHalliford:                    StationNameUpperHalliford,
	StationCodeUpperHolloway:                     StationNameUpperHolloway,
	StationCodeUpperTyndrum:                      StationNameUpperTyndrum,
	StationCodeUpperWarlingham:                   StationNameUpperWarlingham,
	StationCodeUptonMerseyside:                   StationNameUptonMerseyside,
	StationCodeUpwey:                             StationNameUpwey,
	StationCodeUrmston:                           StationNameUrmston,
	StationCodeUttoxeter:                         StationNameUttoxeter,
	StationCodeValley:                            StationNameValley,
	StationCodeVauxhall:                          StationNameVauxhall,
	StationCodeVirginiaWater:                     StationNameVirginiaWater,
	StationCodeWaddon:                            StationNameWaddon,
	StationCodeWadhurst:                          StationNameWadhurst,
	StationCodeWainfleet:                         StationNameWainfleet,
	StationCodeWakefieldKirkgate:                 StationNameWakefieldKirkgate,
	StationCodeWakefieldWestgate:                 StationNameWakefieldWestgate,
	StationCodeWalkden:                           StationNameWalkden,
	StationCodeWallaseyGroveRoad:                 StationNameWallaseyGroveRoad,
	StationCodeWallaseyVillage:                   StationNameWallaseyVillage,
	StationCodeWallington:                        StationNameWallington,
	StationCodeWallyford:                         StationNameWallyford,
	StationCodeWalmer:                            StationNameWalmer,
	StationCodeWalsall:                           StationNameWalsall,
	StationCodeWalsden:                           StationNameWalsden,
	StationCodeWalthamCross:                      StationNameWalthamCross,
	StationCodeWalthamstowCentral:                StationNameWalthamstowCentral,
	StationCodeWalthamstowQueensRoad:             StationNameWalthamstowQueensRoad,
	StationCodeWaltonMerseyside:                  StationNameWaltonMerseyside,
	StationCodeWaltononThames:                    StationNameWaltononThames,
	StationCodeWaltonontheNaze:                   StationNameWaltonontheNaze,
	StationCodeWanborough:                        StationNameWanborough,
	StationCodeWandsworthCommon:                  StationNameWandsworthCommon,
	StationCodeWandsworthRoad:                    StationNameWandsworthRoad,
	StationCodeWandsworthTown:                    StationNameWandsworthTown,
	StationCodeWansteadPark:                      StationNameWansteadPark,
	StationCodeWapping:                           StationNameWapping,
	StationCodeWarblington:                       StationNameWarblington,
	StationCodeWareHerts:                         StationNameWareHerts,
	StationCodeWarehamDorset:                     StationNameWarehamDorset,
	StationCodeWargrave:                          StationNameWargrave,
	StationCodeWarminster:                        StationNameWarminster,
	StationCodeWarnham:                           StationNameWarnham,
	StationCodeWarringtonBankQuay:                StationNameWarringtonBankQuay,
	StationCodeWarringtonCentral:                 StationNameWarringtonCentral,
	StationCodeWarwick:                           StationNameWarwick,
	StationCodeWarwickParkway:                    StationNameWarwickParkway,
	StationCodeWaterOrton:                        StationNameWaterOrton,
	StationCodeWaterbeach:                        StationNameWaterbeach,
	StationCodeWateringbury:                      StationNameWateringbury,
	StationCodeWaterlooMerseyside:                StationNameWaterlooMerseyside,
	StationCodeWatfordHighStreet:                 StationNameWatfordHighStreet,
	StationCodeWatfordJunction:                   StationNameWatfordJunction,
	StationCodeWatfordNorth:                      StationNameWatfordNorth,
	StationCodeWatlington:                        StationNameWatlington,
	StationCodeWattonatStone:                     StationNameWattonatStone,
	StationCodeWaunGronPark:                      StationNameWaunGronPark,
	StationCodeWavertreeTechnologyPark:           StationNameWavertreeTechnologyPark,
	StationCodeWedgwood:                          StationNameWedgwood,
	StationCodeWeeley:                            StationNameWeeley,
	StationCodeWeeton:                            StationNameWeeton,
	StationCodeWelhamGreen:                       StationNameWelhamGreen,
	StationCodeWelling:                           StationNameWelling,
	StationCodeWellingborough:                    StationNameWellingborough,
	StationCodeWellingtonShropshire:              StationNameWellingtonShropshire,
	StationCodeWelshpool:                         StationNameWelshpool,
	StationCodeWelwynGardenCity:                  StationNameWelwynGardenCity,
	StationCodeWelwynNorth:                       StationNameWelwynNorth,
	StationCodeWem:                               StationNameWem,
	StationCodeWembleyCentral:                    StationNameWembleyCentral,
	StationCodeWembleyStadium:                    StationNameWembleyStadium,
	StationCodeWemyssBay:                         StationNameWemyssBay,
	StationCodeWendover:                          StationNameWendover,
	StationCodeWennington:                        StationNameWennington,
	StationCodeWestAllerton:                      StationNameWestAllerton,
	StationCodeWestBrompton:                      StationNameWestBrompton,
	StationCodeWestByfleet:                       StationNameWestByfleet,
	StationCodeWestCalder:                        StationNameWestCalder,
	StationCodeWestCroydon:                       StationNameWestCroydon,
	StationCodeWestDrayton:                       StationNameWestDrayton,
	StationCodeWestDulwich:                       StationNameWestDulwich,
	StationCodeWestEaling:                        StationNameWestEaling,
	StationCodeWestHam:                           StationNameWestHam,
	StationCodeWestHampstead:                     StationNameWestHampstead,
	StationCodeWestHampsteadThameslink:           StationNameWestHampsteadThameslink,
	StationCodeWestHorndon:                       StationNameWestHorndon,
	StationCodeWestKilbride:                      StationNameWestKilbride,
	StationCodeWestKirby:                         StationNameWestKirby,
	StationCodeWestMalling:                       StationNameWestMalling,
	StationCodeWestNorwood:                       StationNameWestNorwood,
	StationCodeWestRuislip:                       StationNameWestRuislip,
	StationCodeWestRunton:                        StationNameWestRunton,
	StationCodeWestStLeonards:                    StationNameWestStLeonards,
	StationCodeWestSutton:                        StationNameWestSutton,
	StationCodeWestWickham:                       StationNameWestWickham,
	StationCodeWestWorthing:                      StationNameWestWorthing,
	StationCodeWestburyWilts:                     StationNameWestburyWilts,
	StationCodeWestcliff:                         StationNameWestcliff,
	StationCodeWestcombePark:                     StationNameWestcombePark,
	StationCodeWestenhanger:                      StationNameWestenhanger,
	StationCodeWesterHailes:                      StationNameWesterHailes,
	StationCodeWesterfield:                       StationNameWesterfield,
	StationCodeWesterton:                         StationNameWesterton,
	StationCodeWestgateOnSea:                     StationNameWestgateOnSea,
	StationCodeWesthoughton:                      StationNameWesthoughton,
	StationCodeWestonMilton:                      StationNameWestonMilton,
	StationCodeWestonsuperMare:                   StationNameWestonsuperMare,
	StationCodeWetheral:                          StationNameWetheral,
	StationCodeWeybridge:                         StationNameWeybridge,
	StationCodeWeymouth:                          StationNameWeymouth,
	StationCodeWhaleyBridge:                      StationNameWhaleyBridge,
	StationCodeWhalleyLancs:                      StationNameWhalleyLancs,
	StationCodeWhatstandwell:                     StationNameWhatstandwell,
	StationCodeWhifflet:                          StationNameWhifflet,
	StationCodeWhimple:                           StationNameWhimple,
	StationCodeWhinhill:                          StationNameWhinhill,
	StationCodeWhiston:                           StationNameWhiston,
	StationCodeWhitby:                            StationNameWhitby,
	StationCodeWhitchurchCardiff:                 StationNameWhitchurchCardiff,
	StationCodeWhitchurchHants:                   StationNameWhitchurchHants,
	StationCodeWhitchurchShropshire:              StationNameWhitchurchShropshire,
	StationCodeWhiteHartLane:                     StationNameWhiteHartLane,
	StationCodeWhiteNotley:                       StationNameWhiteNotley,
	StationCodeWhitechapel:                       StationNameWhitechapel,
	StationCodeWhitecraigs:                       StationNameWhitecraigs,
	StationCodeWhitehaven:                        StationNameWhitehaven,
	StationCodeWhitland:                          StationNameWhitland,
	StationCodeWhitleyBridge:                     StationNameWhitleyBridge,
	StationCodeWhitlocksEnd:                      StationNameWhitlocksEnd,
	StationCodeWhitstable:                        StationNameWhitstable,
	StationCodeWhittlesea:                        StationNameWhittlesea,
	StationCodeWhittlesfordParkway:               StationNameWhittlesfordParkway,
	StationCodeWhittonLondon:                     StationNameWhittonLondon,
	StationCodeWhitwellDerbyshire:                StationNameWhitwellDerbyshire,
	StationCodeWhyteleafe:                        StationNameWhyteleafe,
	StationCodeWhyteleafeSouth:                   StationNameWhyteleafeSouth,
	StationCodeWick:                              StationNameWick,
	StationCodeWickford:                          StationNameWickford,
	StationCodeWickhamMarket:                     StationNameWickhamMarket,
	StationCodeWiddrington:                       StationNameWiddrington,
	StationCodeWidnes:                            StationNameWidnes,
	StationCodeWidneyManor:                       StationNameWidneyManor,
	StationCodeWiganNorthWestern:                 StationNameWiganNorthWestern,
	StationCodeWiganWallgate:                     StationNameWiganWallgate,
	StationCodeWigton:                            StationNameWigton,
	StationCodeWildmill:                          StationNameWildmill,
	StationCodeWillesdenJunction:                 StationNameWillesdenJunction,
	StationCodeWilliamwood:                       StationNameWilliamwood,
	StationCodeWillington:                        StationNameWillington,
	StationCodeWilmcote:                          StationNameWilmcote,
	StationCodeWilmslow:                          StationNameWilmslow,
	StationCodeWilnecoteStaffs:                   StationNameWilnecoteStaffs,
	StationCodeWimbledon:                         StationNameWimbledon,
	StationCodeWimbledonChase:                    StationNameWimbledonChase,
	StationCodeWinchelsea:                        StationNameWinchelsea,
	StationCodeWinchester:                        StationNameWinchester,
	StationCodeWinchfield:                        StationNameWinchfield,
	StationCodeWinchmoreHill:                     StationNameWinchmoreHill,
	StationCodeWindermere:                        StationNameWindermere,
	StationCodeWindsorAndEtonCentral:             StationNameWindsorAndEtonCentral,
	StationCodeWindsorAndEtonRiverside:           StationNameWindsorAndEtonRiverside,
	StationCodeWinnersh:                          StationNameWinnersh,
	StationCodeWinnershTriangle:                  StationNameWinnershTriangle,
	StationCodeWinsford:                          StationNameWinsford,
	StationCodeWishaw:                            StationNameWishaw,
	StationCodeWitham:                            StationNameWitham,
	StationCodeWitley:                            StationNameWitley,
	StationCodeWittonWestMidlands:                StationNameWittonWestMidlands,
	StationCodeWivelsfield:                       StationNameWivelsfield,
	StationCodeWivenhoe:                          StationNameWivenhoe,
	StationCodeWoburnSands:                       StationNameWoburnSands,
	StationCodeWoking:                            StationNameWoking,
	StationCodeWokingham:                         StationNameWokingham,
	StationCodeWoldingham:                        StationNameWoldingham,
	StationCodeWolverhampton:                     StationNameWolverhampton,
	StationCodeWolverton:                         StationNameWolverton,
	StationCodeWombwell:                          StationNameWombwell,
	StationCodeWoodEnd:                           StationNameWoodEnd,
	StationCodeWoodStreet:                        StationNameWoodStreet,
	StationCodeWoodbridge:                        StationNameWoodbridge,
	StationCodeWoodgrangePark:                    StationNameWoodgrangePark,
	StationCodeWoodhall:                          StationNameWoodhall,
	StationCodeWoodhouse:                         StationNameWoodhouse,
	StationCodeWoodlesford:                       StationNameWoodlesford,
	StationCodeWoodley:                           StationNameWoodley,
	StationCodeWoodmansterne:                     StationNameWoodmansterne,
	StationCodeWoodsmoor:                         StationNameWoodsmoor,
	StationCodeWool:                              StationNameWool,
	StationCodeWoolston:                          StationNameWoolston,
	StationCodeWoolwichArsenal:                   StationNameWoolwichArsenal,
	StationCodeWoolwichDockyard:                  StationNameWoolwichDockyard,
	StationCodeWoottonWawen:                      StationNameWoottonWawen,
	StationCodeWorcesterForegateStreet:           StationNameWorcesterForegateStreet,
	StationCodeWorcesterPark:                     StationNameWorcesterPark,
	StationCodeWorcesterShrubHill:                StationNameWorcesterShrubHill,
	StationCodeWorkington:                        StationNameWorkington,
	StationCodeWorksop:                           StationNameWorksop,
	StationCodeWorle:                             StationNameWorle,
	StationCodeWorplesdon:                        StationNameWorplesdon,
	StationCodeWorstead:                          StationNameWorstead,
	StationCodeWorthing:                          StationNameWorthing,
	StationCodeWrabness:                          StationNameWrabness,
	StationCodeWraysbury:                         StationNameWraysbury,
	StationCodeWrenbury:                          StationNameWrenbury,
	StationCodeWressle:                           StationNameWressle,
	StationCodeWrexhamCentral:                    StationNameWrexhamCentral,
	StationCodeWrexhamGeneral:                    StationNameWrexhamGeneral,
	StationCodeWye:                               StationNameWye,
	StationCodeWylam:                             StationNameWylam,
	StationCodeWyldeGreen:                        StationNameWyldeGreen,
	StationCodeWymondham:                         StationNameWymondham,
	StationCodeWythall:                           StationNameWythall,
	StationCodeYalding:                           StationNameYalding,
	StationCodeYardleyWood:                       StationNameYardleyWood,
	StationCodeYarm:                              StationNameYarm,
	StationCodeYate:                              StationNameYate,
	StationCodeYatton:                            StationNameYatton,
	StationCodeYeoford:                           StationNameYeoford,
	StationCodeYeovilJunction:                    StationNameYeovilJunction,
	StationCodeYeovilPenMill:                     StationNameYeovilPenMill,
	StationCodeYetminster:                        StationNameYetminster,
	StationCodeYnyswen:                           StationNameYnyswen,
	StationCodeYoker:                             StationNameYoker,
	StationCodeYork:                              StationNameYork,
	StationCodeYorton:                            StationNameYorton,
	StationCodeYstradMynach:                      StationNameYstradMynach,
	StationCodeYstradRhondda:                     StationNameYstradRhondda,
}

var StationNameToCodeMap = map[string]CRSCode{
	StationNameAbbeyWood:                         StationCodeAbbeyWood,
	StationNameAber:                              StationCodeAber,
	StationNameAbercynon:                         StationCodeAbercynon,
	StationNameAberdare:                          StationCodeAberdare,
	StationNameAberdeen:                          StationCodeAberdeen,
	StationNameAberdour:                          StationCodeAberdour,
	StationNameAberdovey:                         StationCodeAberdovey,
	StationNameAbererch:                          StationCodeAbererch,
	StationNameAbergavenny:                       StationCodeAbergavenny,
	StationNameAbergeleAndPensarn:                StationCodeAbergeleAndPensarn,
	StationNameAberystwyth:                       StationCodeAberystwyth,
	StationNameAccrington:                        StationCodeAccrington,
	StationNameAchanalt:                          StationCodeAchanalt,
	StationNameAchnasheen:                        StationCodeAchnasheen,
	StationNameAchnashellach:                     StationCodeAchnashellach,
	StationNameAcklington:                        StationCodeAcklington,
	StationNameAcle:                              StationCodeAcle,
	StationNameAcocksGreen:                       StationCodeAcocksGreen,
	StationNameActonBridgeCheshire:               StationCodeActonBridgeCheshire,
	StationNameActonCentral:                      StationCodeActonCentral,
	StationNameActonMainLine:                     StationCodeActonMainLine,
	StationNameAdderleyPark:                      StationCodeAdderleyPark,
	StationNameAddiewell:                         StationCodeAddiewell,
	StationNameAddlestone:                        StationCodeAddlestone,
	StationNameAdisham:                           StationCodeAdisham,
	StationNameAdlingtonCheshire:                 StationCodeAdlingtonCheshire,
	StationNameAdlingtonLancs:                    StationCodeAdlingtonLancs,
	StationNameAdwick:                            StationCodeAdwick,
	StationNameAigburth:                          StationCodeAigburth,
	StationNameAinsdale:                          StationCodeAinsdale,
	StationNameAintree:                           StationCodeAintree,
	StationNameAirbles:                           StationCodeAirbles,
	StationNameAirdrie:                           StationCodeAirdrie,
	StationNameAlbanyPark:                        StationCodeAlbanyPark,
	StationNameAlbrighton:                        StationCodeAlbrighton,
	StationNameAlderleyEdge:                      StationCodeAlderleyEdge,
	StationNameAldermaston:                       StationCodeAldermaston,
	StationNameAldershot:                         StationCodeAldershot,
	StationNameAldrington:                        StationCodeAldrington,
	StationNameAlexandraPalace:                   StationCodeAlexandraPalace,
	StationNameAlexandraParade:                   StationCodeAlexandraParade,
	StationNameAlexandria:                        StationCodeAlexandria,
	StationNameAlfreton:                          StationCodeAlfreton,
	StationNameAllensWest:                        StationCodeAllensWest,
	StationNameAlloa:                             StationCodeAlloa,
	StationNameAlness:                            StationCodeAlness,
	StationNameAlnmouth:                          StationCodeAlnmouth,
	StationNameAlresfordEssex:                    StationCodeAlresfordEssex,
	StationNameAlsager:                           StationCodeAlsager,
	StationNameAlthorneEssex:                     StationCodeAlthorneEssex,
	StationNameAlthorpe:                          StationCodeAlthorpe,
	StationNameAltnabreac:                        StationCodeAltnabreac,
	StationNameAlton:                             StationCodeAlton,
	StationNameAltrincham:                        StationCodeAltrincham,
	StationNameAlvechurch:                        StationCodeAlvechurch,
	StationNameAmbergate:                         StationCodeAmbergate,
	StationNameAmberley:                          StationCodeAmberley,
	StationNameAmersham:                          StationCodeAmersham,
	StationNameAmmanford:                         StationCodeAmmanford,
	StationNameAncaster:                          StationCodeAncaster,
	StationNameAnderston:                         StationCodeAnderston,
	StationNameAndover:                           StationCodeAndover,
	StationNameAnerley:                           StationCodeAnerley,
	StationNameAngelRoad:                         StationCodeAngelRoad,
	StationNameAngmering:                         StationCodeAngmering,
	StationNameAnnan:                             StationCodeAnnan,
	StationNameAnniesland:                        StationCodeAnniesland,
	StationNameAnsdellAndFairhaven:               StationCodeAnsdellAndFairhaven,
	StationNameApperleyBridge:                    StationCodeApperleyBridge,
	StationNameAppleby:                           StationCodeAppleby,
	StationNameAppledoreKent:                     StationCodeAppledoreKent,
	StationNameAppleford:                         StationCodeAppleford,
	StationNameAppleyBridge:                      StationCodeAppleyBridge,
	StationNameApsley:                            StationCodeApsley,
	StationNameArbroath:                          StationCodeArbroath,
	StationNameArdgay:                            StationCodeArdgay,
	StationNameArdlui:                            StationCodeArdlui,
	StationNameArdrossanHarbour:                  StationCodeArdrossanHarbour,
	StationNameArdrossanSouthBeach:               StationCodeArdrossanSouthBeach,
	StationNameArdrossanTown:                     StationCodeArdrossanTown,
	StationNameArdwick:                           StationCodeArdwick,
	StationNameArgyleStreet:                      StationCodeArgyleStreet,
	StationNameArisaig:                           StationCodeArisaig,
	StationNameArlesey:                           StationCodeArlesey,
	StationNameArmadaleWestLothian:               StationCodeArmadaleWestLothian,
	StationNameArmathwaite:                       StationCodeArmathwaite,
	StationNameArnside:                           StationCodeArnside,
	StationNameArram:                             StationCodeArram,
	StationNameArrocharAndTarbet:                 StationCodeArrocharAndTarbet,
	StationNameArundel:                           StationCodeArundel,
	StationNameAscotBerks:                        StationCodeAscotBerks,
	StationNameAscottunderWychwood:               StationCodeAscottunderWychwood,
	StationNameAsh:                               StationCodeAsh,
	StationNameAshVale:                           StationCodeAshVale,
	StationNameAshburys:                          StationCodeAshburys,
	StationNameAshchurchforTewkesbury:            StationCodeAshchurchforTewkesbury,
	StationNameAshfield:                          StationCodeAshfield,
	StationNameAshfordInternational:              StationCodeAshfordInternational,
	StationNameAshfordInternationalEurostar:      StationCodeAshfordInternationalEurostar,
	StationNameAshfordSurrey:                     StationCodeAshfordSurrey,
	StationNameAshley:                            StationCodeAshley,
	StationNameAshtead:                           StationCodeAshtead,
	StationNameAshtonunderLyne:                   StationCodeAshtonunderLyne,
	StationNameAshurstKent:                       StationCodeAshurstKent,
	StationNameAshurstNewForest:                  StationCodeAshurstNewForest,
	StationNameAshwellAndMorden:                  StationCodeAshwellAndMorden,
	StationNameAskam:                             StationCodeAskam,
	StationNameAslockton:                         StationCodeAslockton,
	StationNameAspatria:                          StationCodeAspatria,
	StationNameAspleyGuise:                       StationCodeAspleyGuise,
	StationNameAston:                             StationCodeAston,
	StationNameAtherstone:                        StationCodeAtherstone,
	StationNameAtherton:                          StationCodeAtherton,
	StationNameAttadale:                          StationCodeAttadale,
	StationNameAttenborough:                      StationCodeAttenborough,
	StationNameAttleborough:                      StationCodeAttleborough,
	StationNameAuchinleck:                        StationCodeAuchinleck,
	StationNameAudleyEnd:                         StationCodeAudleyEnd,
	StationNameAughtonPark:                       StationCodeAughtonPark,
	StationNameAviemore:                          StationCodeAviemore,
	StationNameAvoncliff:                         StationCodeAvoncliff,
	StationNameAvonmouth:                         StationCodeAvonmouth,
	StationNameAxminster:                         StationCodeAxminster,
	StationNameAylesbury:                         StationCodeAylesbury,
	StationNameAylesburyValeParkway:              StationCodeAylesburyValeParkway,
	StationNameAylesford:                         StationCodeAylesford,
	StationNameAylesham:                          StationCodeAylesham,
	StationNameAyr:                               StationCodeAyr,
	StationNameBache:                             StationCodeBache,
	StationNameBaglan:                            StationCodeBaglan,
	StationNameBagshot:                           StationCodeBagshot,
	StationNameBaildon:                           StationCodeBaildon,
	StationNameBaillieston:                       StationCodeBaillieston,
	StationNameBalcombe:                          StationCodeBalcombe,
	StationNameBaldock:                           StationCodeBaldock,
	StationNameBalham:                            StationCodeBalham,
	StationNameBalloch:                           StationCodeBalloch,
	StationNameBalmossie:                         StationCodeBalmossie,
	StationNameBamberBridge:                      StationCodeBamberBridge,
	StationNameBamford:                           StationCodeBamford,
	StationNameBanavie:                           StationCodeBanavie,
	StationNameBanbury:                           StationCodeBanbury,
	StationNameBangorGwynedd:                     StationCodeBangorGwynedd,
	StationNameBankHall:                          StationCodeBankHall,
	StationNameBanstead:                          StationCodeBanstead,
	StationNameBarassie:                          StationCodeBarassie,
	StationNameBarbican:                          StationCodeBarbican,
	StationNameBardonMill:                        StationCodeBardonMill,
	StationNameBareLane:                          StationCodeBareLane,
	StationNameBargeddie:                         StationCodeBargeddie,
	StationNameBargoed:                           StationCodeBargoed,
	StationNameBarking:                           StationCodeBarking,
	StationNameBarlaston:                         StationCodeBarlaston,
	StationNameBarming:                           StationCodeBarming,
	StationNameBarmouth:                          StationCodeBarmouth,
	StationNameBarnehurst:                        StationCodeBarnehurst,
	StationNameBarnes:                            StationCodeBarnes,
	StationNameBarnesBridge:                      StationCodeBarnesBridge,
	StationNameBarnetby:                          StationCodeBarnetby,
	StationNameBarnham:                           StationCodeBarnham,
	StationNameBarnhill:                          StationCodeBarnhill,
	StationNameBarnsley:                          StationCodeBarnsley,
	StationNameBarnstaple:                        StationCodeBarnstaple,
	StationNameBarntGreen:                        StationCodeBarntGreen,
	StationNameBarrhead:                          StationCodeBarrhead,
	StationNameBarrhill:                          StationCodeBarrhill,
	StationNameBarrowHaven:                       StationCodeBarrowHaven,
	StationNameBarrowUponSoar:                    StationCodeBarrowUponSoar,
	StationNameBarrowinFurness:                   StationCodeBarrowinFurness,
	StationNameBarry:                             StationCodeBarry,
	StationNameBarryDocks:                        StationCodeBarryDocks,
	StationNameBarryIsland:                       StationCodeBarryIsland,
	StationNameBarryLinks:                        StationCodeBarryLinks,
	StationNameBartononHumber:                    StationCodeBartononHumber,
	StationNameBasildon:                          StationCodeBasildon,
	StationNameBasingstoke:                       StationCodeBasingstoke,
	StationNameBatAndBall:                        StationCodeBatAndBall,
	StationNameBathSpa:                           StationCodeBathSpa,
	StationNameBathgate:                          StationCodeBathgate,
	StationNameBatley:                            StationCodeBatley,
	StationNameBattersby:                         StationCodeBattersby,
	StationNameBatterseaPark:                     StationCodeBatterseaPark,
	StationNameBattle:                            StationCodeBattle,
	StationNameBattlesbridge:                     StationCodeBattlesbridge,
	StationNameBayford:                           StationCodeBayford,
	StationNameBeaconsfield:                      StationCodeBeaconsfield,
	StationNameBearley:                           StationCodeBearley,
	StationNameBearsden:                          StationCodeBearsden,
	StationNameBearsted:                          StationCodeBearsted,
	StationNameBeasdale:                          StationCodeBeasdale,
	StationNameBeaulieuRoad:                      StationCodeBeaulieuRoad,
	StationNameBeauly:                            StationCodeBeauly,
	StationNameBebington:                         StationCodeBebington,
	StationNameBeccles:                           StationCodeBeccles,
	StationNameBeckenhamHill:                     StationCodeBeckenhamHill,
	StationNameBeckenhamJunction:                 StationCodeBeckenhamJunction,
	StationNameBedford:                           StationCodeBedford,
	StationNameBedfordStJohns:                    StationCodeBedfordStJohns,
	StationNameBedhampton:                        StationCodeBedhampton,
	StationNameBedminster:                        StationCodeBedminster,
	StationNameBedworth:                          StationCodeBedworth,
	StationNameBedwyn:                            StationCodeBedwyn,
	StationNameBeeston:                           StationCodeBeeston,
	StationNameBekesbourne:                       StationCodeBekesbourne,
	StationNameBelleVue:                          StationCodeBelleVue,
	StationNameBellgrove:                         StationCodeBellgrove,
	StationNameBellingham:                        StationCodeBellingham,
	StationNameBellshill:                         StationCodeBellshill,
	StationNameBelmont:                           StationCodeBelmont,
	StationNameBelper:                            StationCodeBelper,
	StationNameBeltring:                          StationCodeBeltring,
	StationNameBelvedere:                         StationCodeBelvedere,
	StationNameBempton:                           StationCodeBempton,
	StationNameBenRhydding:                       StationCodeBenRhydding,
	StationNameBenfleet:                          StationCodeBenfleet,
	StationNameBentham:                           StationCodeBentham,
	StationNameBentleyHants:                      StationCodeBentleyHants,
	StationNameBentleySouthYorks:                 StationCodeBentleySouthYorks,
	StationNameBereAlston:                        StationCodeBereAlston,
	StationNameBereFerrers:                       StationCodeBereFerrers,
	StationNameBerkhamsted:                       StationCodeBerkhamsted,
	StationNameBerkswell:                         StationCodeBerkswell,
	StationNameBermudaPark:                       StationCodeBermudaPark,
	StationNameBerneyArms:                        StationCodeBerneyArms,
	StationNameBerryBrow:                         StationCodeBerryBrow,
	StationNameBerrylands:                        StationCodeBerrylands,
	StationNameBerwickSussex:                     StationCodeBerwickSussex,
	StationNameBerwickUponTweed:                  StationCodeBerwickUponTweed,
	StationNameBescarLane:                        StationCodeBescarLane,
	StationNameBescotStadium:                     StationCodeBescotStadium,
	StationNameBetchworth:                        StationCodeBetchworth,
	StationNameBethnalGreen:                      StationCodeBethnalGreen,
	StationNameBetwsyCoed:                        StationCodeBetwsyCoed,
	StationNameBeverley:                          StationCodeBeverley,
	StationNameBexhill:                           StationCodeBexhill,
	StationNameBexley:                            StationCodeBexley,
	StationNameBexleyheath:                       StationCodeBexleyheath,
	StationNameBicesterNorth:                     StationCodeBicesterNorth,
	StationNameBicesterVillage:                   StationCodeBicesterVillage,
	StationNameBickley:                           StationCodeBickley,
	StationNameBidston:                           StationCodeBidston,
	StationNameBiggleswade:                       StationCodeBiggleswade,
	StationNameBilbrook:                          StationCodeBilbrook,
	StationNameBillericay:                        StationCodeBillericay,
	StationNameBillinghamCleveland:               StationCodeBillinghamCleveland,
	StationNameBillingshurst:                     StationCodeBillingshurst,
	StationNameBingham:                           StationCodeBingham,
	StationNameBingley:                           StationCodeBingley,
	StationNameBirchgrove:                        StationCodeBirchgrove,
	StationNameBirchingtonOnSea:                  StationCodeBirchingtonOnSea,
	StationNameBirchwood:                         StationCodeBirchwood,
	StationNameBirkbeck:                          StationCodeBirkbeck,
	StationNameBirkdale:                          StationCodeBirkdale,
	StationNameBirkenheadCentral:                 StationCodeBirkenheadCentral,
	StationNameBirkenheadHamiltonSquare:          StationCodeBirkenheadHamiltonSquare,
	StationNameBirkenheadNorth:                   StationCodeBirkenheadNorth,
	StationNameBirkenheadPark:                    StationCodeBirkenheadPark,
	StationNameBirminghamInternational:           StationCodeBirminghamInternational,
	StationNameBirminghamMoorStreet:              StationCodeBirminghamMoorStreet,
	StationNameBirminghamNewStreet:               StationCodeBirminghamNewStreet,
	StationNameBirminghamSnowHill:                StationCodeBirminghamSnowHill,
	StationNameBishopAuckland:                    StationCodeBishopAuckland,
	StationNameBishopbriggs:                      StationCodeBishopbriggs,
	StationNameBishopsStortford:                  StationCodeBishopsStortford,
	StationNameBishopstoneSussex:                 StationCodeBishopstoneSussex,
	StationNameBishoptonStrathclyde:              StationCodeBishoptonStrathclyde,
	StationNameBitterne:                          StationCodeBitterne,
	StationNameBlackburn:                         StationCodeBlackburn,
	StationNameBlackheath:                        StationCodeBlackheath,
	StationNameBlackhorseRoad:                    StationCodeBlackhorseRoad,
	StationNameBlackpoolNorth:                    StationCodeBlackpoolNorth,
	StationNameBlackpoolPleasureBeach:            StationCodeBlackpoolPleasureBeach,
	StationNameBlackpoolSouth:                    StationCodeBlackpoolSouth,
	StationNameBlackridge:                        StationCodeBlackridge,
	StationNameBlackrod:                          StationCodeBlackrod,
	StationNameBlackwater:                        StationCodeBlackwater,
	StationNameBlaenauFfestiniog:                 StationCodeBlaenauFfestiniog,
	StationNameBlairAtholl:                       StationCodeBlairAtholl,
	StationNameBlairhill:                         StationCodeBlairhill,
	StationNameBlakeStreet:                       StationCodeBlakeStreet,
	StationNameBlakedown:                         StationCodeBlakedown,
	StationNameBlantyre:                          StationCodeBlantyre,
	StationNameBlaydon:                           StationCodeBlaydon,
	StationNameBleasby:                           StationCodeBleasby,
	StationNameBletchley:                         StationCodeBletchley,
	StationNameBloxwich:                          StationCodeBloxwich,
	StationNameBloxwichNorth:                     StationCodeBloxwichNorth,
	StationNameBlundellsandsAndCrosby:            StationCodeBlundellsandsAndCrosby,
	StationNameBlytheBridge:                      StationCodeBlytheBridge,
	StationNameBodminParkway:                     StationCodeBodminParkway,
	StationNameBodorgan:                          StationCodeBodorgan,
	StationNameBognorRegis:                       StationCodeBognorRegis,
	StationNameBogston:                           StationCodeBogston,
	StationNameBolton:                            StationCodeBolton,
	StationNameBoltonUponDearne:                  StationCodeBoltonUponDearne,
	StationNameBookham:                           StationCodeBookham,
	StationNameBootleCumbria:                     StationCodeBootleCumbria,
	StationNameBootleNewStrand:                   StationCodeBootleNewStrand,
	StationNameBootleOrielRoad:                   StationCodeBootleOrielRoad,
	StationNameBordesley:                         StationCodeBordesley,
	StationNameBoroughGreenAndWrotham:            StationCodeBoroughGreenAndWrotham,
	StationNameBorth:                             StationCodeBorth,
	StationNameBosham:                            StationCodeBosham,
	StationNameBoston:                            StationCodeBoston,
	StationNameBotley:                            StationCodeBotley,
	StationNameBottesford:                        StationCodeBottesford,
	StationNameBourneEnd:                         StationCodeBourneEnd,
	StationNameBournemouth:                       StationCodeBournemouth,
	StationNameBournville:                        StationCodeBournville,
	StationNameBowBrickhill:                      StationCodeBowBrickhill,
	StationNameBowesPark:                         StationCodeBowesPark,
	StationNameBowling:                           StationCodeBowling,
	StationNameBoxHillAndWesthumble:              StationCodeBoxHillAndWesthumble,
	StationNameBracknell:                         StationCodeBracknell,
	StationNameBradfordForsterSquare:             StationCodeBradfordForsterSquare,
	StationNameBradfordInterchange:               StationCodeBradfordInterchange,
	StationNameBradfordonAvon:                    StationCodeBradfordonAvon,
	StationNameBrading:                           StationCodeBrading,
	StationNameBraintree:                         StationCodeBraintree,
	StationNameBraintreeFreeport:                 StationCodeBraintreeFreeport,
	StationNameBramhall:                          StationCodeBramhall,
	StationNameBramleyHants:                      StationCodeBramleyHants,
	StationNameBramleyWYorks:                     StationCodeBramleyWYorks,
	StationNameBramptonCumbria:                   StationCodeBramptonCumbria,
	StationNameBramptonSuffolk:                   StationCodeBramptonSuffolk,
	StationNameBranchton:                         StationCodeBranchton,
	StationNameBrandon:                           StationCodeBrandon,
	StationNameBranksome:                         StationCodeBranksome,
	StationNameBraystonesCumbria:                 StationCodeBraystonesCumbria,
	StationNameBredbury:                          StationCodeBredbury,
	StationNameBreich:                            StationCodeBreich,
	StationNameBrentford:                         StationCodeBrentford,
	StationNameBrentwood:                         StationCodeBrentwood,
	StationNameBricketWood:                       StationCodeBricketWood,
	StationNameBridgend:                          StationCodeBridgend,
	StationNameBridgeofAllan:                     StationCodeBridgeofAllan,
	StationNameBridgeofOrchy:                     StationCodeBridgeofOrchy,
	StationNameBridgeton:                         StationCodeBridgeton,
	StationNameBridgwater:                        StationCodeBridgwater,
	StationNameBridlington:                       StationCodeBridlington,
	StationNameBrierfield:                        StationCodeBrierfield,
	StationNameBrigg:                             StationCodeBrigg,
	StationNameBrighouse:                         StationCodeBrighouse,
	StationNameBrightonEastSussex:                StationCodeBrightonEastSussex,
	StationNameBrimsdown:                         StationCodeBrimsdown,
	StationNameBrinnington:                       StationCodeBrinnington,
	StationNameBristolParkway:                    StationCodeBristolParkway,
	StationNameBristolTempleMeads:                StationCodeBristolTempleMeads,
	StationNameBrithdir:                          StationCodeBrithdir,
	StationNameBritonFerry:                       StationCodeBritonFerry,
	StationNameBrixton:                           StationCodeBrixton,
	StationNameBroadGreen:                        StationCodeBroadGreen,
	StationNameBroadbottom:                       StationCodeBroadbottom,
	StationNameBroadstairs:                       StationCodeBroadstairs,
	StationNameBrockenhurst:                      StationCodeBrockenhurst,
	StationNameBrockholes:                        StationCodeBrockholes,
	StationNameBrockley:                          StationCodeBrockley,
	StationNameBromborough:                       StationCodeBromborough,
	StationNameBromboroughRake:                   StationCodeBromboroughRake,
	StationNameBromleyCrossLancs:                 StationCodeBromleyCrossLancs,
	StationNameBromleyNorth:                      StationCodeBromleyNorth,
	StationNameBromleySouth:                      StationCodeBromleySouth,
	StationNameBromsgrove:                        StationCodeBromsgrove,
	StationNameBrondesbury:                       StationCodeBrondesbury,
	StationNameBrondesburyPark:                   StationCodeBrondesburyPark,
	StationNameBrookmansPark:                     StationCodeBrookmansPark,
	StationNameBrookwood:                         StationCodeBrookwood,
	StationNameBroome:                            StationCodeBroome,
	StationNameBroomfleet:                        StationCodeBroomfleet,
	StationNameBrora:                             StationCodeBrora,
	StationNameBrough:                            StationCodeBrough,
	StationNameBroughtyFerry:                     StationCodeBroughtyFerry,
	StationNameBroxbourne:                        StationCodeBroxbourne,
	StationNameBruceGrove:                        StationCodeBruceGrove,
	StationNameBrundall:                          StationCodeBrundall,
	StationNameBrundallGardens:                   StationCodeBrundallGardens,
	StationNameBrunstane:                         StationCodeBrunstane,
	StationNameBrunswick:                         StationCodeBrunswick,
	StationNameBruton:                            StationCodeBruton,
	StationNameBryn:                              StationCodeBryn,
	StationNameBuckenhamNorfolk:                  StationCodeBuckenhamNorfolk,
	StationNameBuckley:                           StationCodeBuckley,
	StationNameBucknell:                          StationCodeBucknell,
	StationNameBuckshawParkway:                   StationCodeBuckshawParkway,
	StationNameBugle:                             StationCodeBugle,
	StationNameBuilthRoad:                        StationCodeBuilthRoad,
	StationNameBulwell:                           StationCodeBulwell,
	StationNameBures:                             StationCodeBures,
	StationNameBurgessHill:                       StationCodeBurgessHill,
	StationNameBurleyPark:                        StationCodeBurleyPark,
	StationNameBurleyinWharfedale:                StationCodeBurleyinWharfedale,
	StationNameBurnage:                           StationCodeBurnage,
	StationNameBurnesideCumbria:                  StationCodeBurnesideCumbria,
	StationNameBurnhamBucks:                      StationCodeBurnhamBucks,
	StationNameBurnhamonCrouch:                   StationCodeBurnhamonCrouch,
	StationNameBurnleyBarracks:                   StationCodeBurnleyBarracks,
	StationNameBurnleyCentral:                    StationCodeBurnleyCentral,
	StationNameBurnleyManchesterRoad:             StationCodeBurnleyManchesterRoad,
	StationNameBurnsideStrathclyde:               StationCodeBurnsideStrathclyde,
	StationNameBurntisland:                       StationCodeBurntisland,
	StationNameBurscoughBridge:                   StationCodeBurscoughBridge,
	StationNameBurscoughJunction:                 StationCodeBurscoughJunction,
	StationNameBursledon:                         StationCodeBursledon,
	StationNameBurtonJoyce:                       StationCodeBurtonJoyce,
	StationNameBurtononTrent:                     StationCodeBurtononTrent,
	StationNameBuryStEdmunds:                     StationCodeBuryStEdmunds,
	StationNameBusby:                             StationCodeBusby,
	StationNameBushHillPark:                      StationCodeBushHillPark,
	StationNameBushey:                            StationCodeBushey,
	StationNameButlersLane:                       StationCodeButlersLane,
	StationNameBuxted:                            StationCodeBuxted,
	StationNameBuxton:                            StationCodeBuxton,
	StationNameByfleetAndNewHaw:                  StationCodeByfleetAndNewHaw,
	StationNameBynea:                             StationCodeBynea,
	StationNameCadoxton:                          StationCodeCadoxton,
	StationNameCaergwrle:                         StationCodeCaergwrle,
	StationNameCaerphilly:                        StationCodeCaerphilly,
	StationNameCaersws:                           StationCodeCaersws,
	StationNameCaldercruix:                       StationCodeCaldercruix,
	StationNameCaldicot:                          StationCodeCaldicot,
	StationNameCaledonianRdAndBarnsbury:          StationCodeCaledonianRdAndBarnsbury,
	StationNameCalstock:                          StationCodeCalstock,
	StationNameCamAndDursley:                     StationCodeCamAndDursley,
	StationNameCamberley:                         StationCodeCamberley,
	StationNameCamborne:                          StationCodeCamborne,
	StationNameCambridge:                         StationCodeCambridge,
	StationNameCambridgeHeath:                    StationCodeCambridgeHeath,
	StationNameCambuslang:                        StationCodeCambuslang,
	StationNameCamdenRoad:                        StationCodeCamdenRoad,
	StationNameCamelon:                           StationCodeCamelon,
	StationNameCanadaWater:                       StationCodeCanadaWater,
	StationNameCanley:                            StationCodeCanley,
	StationNameCannock:                           StationCodeCannock,
	StationNameCanonbury:                         StationCodeCanonbury,
	StationNameCanterburyEast:                    StationCodeCanterburyEast,
	StationNameCanterburyWest:                    StationCodeCanterburyWest,
	StationNameCantley:                           StationCodeCantley,
	StationNameCapenhurst:                        StationCodeCapenhurst,
	StationNameCarbisBay:                         StationCodeCarbisBay,
	StationNameCardenden:                         StationCodeCardenden,
	StationNameCardiffBay:                        StationCodeCardiffBay,
	StationNameCardiffCentral:                    StationCodeCardiffCentral,
	StationNameCardiffQueenStreet:                StationCodeCardiffQueenStreet,
	StationNameCardonald:                         StationCodeCardonald,
	StationNameCardross:                          StationCodeCardross,
	StationNameCarfin:                            StationCodeCarfin,
	StationNameCarkAndCartmel:                    StationCodeCarkAndCartmel,
	StationNameCarlisle:                          StationCodeCarlisle,
	StationNameCarlton:                           StationCodeCarlton,
	StationNameCarluke:                           StationCodeCarluke,
	StationNameCarmarthen:                        StationCodeCarmarthen,
	StationNameCarmyle:                           StationCodeCarmyle,
	StationNameCarnforth:                         StationCodeCarnforth,
	StationNameCarnoustie:                        StationCodeCarnoustie,
	StationNameCarntyne:                          StationCodeCarntyne,
	StationNameCarpendersPark:                    StationCodeCarpendersPark,
	StationNameCarrbridge:                        StationCodeCarrbridge,
	StationNameCarshalton:                        StationCodeCarshalton,
	StationNameCarshaltonBeeches:                 StationCodeCarshaltonBeeches,
	StationNameCarstairs:                         StationCodeCarstairs,
	StationNameCartsdyke:                         StationCodeCartsdyke,
	StationNameCastleBarPark:                     StationCodeCastleBarPark,
	StationNameCastleCary:                        StationCodeCastleCary,
	StationNameCastleford:                        StationCodeCastleford,
	StationNameCastletonManchester:               StationCodeCastletonManchester,
	StationNameCastletonMoor:                     StationCodeCastletonMoor,
	StationNameCaterham:                          StationCodeCaterham,
	StationNameCatford:                           StationCodeCatford,
	StationNameCatfordBridge:                     StationCodeCatfordBridge,
	StationNameCathays:                           StationCodeCathays,
	StationNameCathcart:                          StationCodeCathcart,
	StationNameCattal:                            StationCodeCattal,
	StationNameCauseland:                         StationCodeCauseland,
	StationNameCefnyBedd:                         StationCodeCefnyBedd,
	StationNameChadwellHeath:                     StationCodeChadwellHeath,
	StationNameChaffordHundredLakeside:           StationCodeChaffordHundredLakeside,
	StationNameChalfontAndLatimer:                StationCodeChalfontAndLatimer,
	StationNameChalkwell:                         StationCodeChalkwell,
	StationNameChandlersFord:                     StationCodeChandlersFord,
	StationNameChapelenleFrith:                   StationCodeChapelenleFrith,
	StationNameChapeltonDevon:                    StationCodeChapeltonDevon,
	StationNameChapeltownSouthYorks:              StationCodeChapeltownSouthYorks,
	StationNameChappelAndWakesColne:              StationCodeChappelAndWakesColne,
	StationNameCharingCrossGlasgow:               StationCodeCharingCrossGlasgow,
	StationNameCharingKent:                       StationCodeCharingKent,
	StationNameCharlbury:                         StationCodeCharlbury,
	StationNameCharlton:                          StationCodeCharlton,
	StationNameChartham:                          StationCodeChartham,
	StationNameChassenRoad:                       StationCodeChassenRoad,
	StationNameChatelherault:                     StationCodeChatelherault,
	StationNameChatham:                           StationCodeChatham,
	StationNameChathill:                          StationCodeChathill,
	StationNameCheadleHulme:                      StationCodeCheadleHulme,
	StationNameCheam:                             StationCodeCheam,
	StationNameCheddington:                       StationCodeCheddington,
	StationNameChelfordCheshire:                  StationCodeChelfordCheshire,
	StationNameChelmsford:                        StationCodeChelmsford,
	StationNameChelsfield:                        StationCodeChelsfield,
	StationNameCheltenhamSpa:                     StationCodeCheltenhamSpa,
	StationNameChepstow:                          StationCodeChepstow,
	StationNameCherryTree:                        StationCodeCherryTree,
	StationNameChertsey:                          StationCodeChertsey,
	StationNameCheshunt:                          StationCodeCheshunt,
	StationNameChessingtonNorth:                  StationCodeChessingtonNorth,
	StationNameChessingtonSouth:                  StationCodeChessingtonSouth,
	StationNameChester:                           StationCodeChester,
	StationNameChesterRoad:                       StationCodeChesterRoad,
	StationNameChesterfield:                      StationCodeChesterfield,
	StationNameChesterleStreet:                   StationCodeChesterleStreet,
	StationNameChestfieldAndSwalecliffe:          StationCodeChestfieldAndSwalecliffe,
	StationNameChetnole:                          StationCodeChetnole,
	StationNameChichester:                        StationCodeChichester,
	StationNameChilham:                           StationCodeChilham,
	StationNameChilworth:                         StationCodeChilworth,
	StationNameChingford:                         StationCodeChingford,
	StationNameChinley:                           StationCodeChinley,
	StationNameChippenham:                        StationCodeChippenham,
	StationNameChipstead:                         StationCodeChipstead,
	StationNameChirk:                             StationCodeChirk,
	StationNameChislehurst:                       StationCodeChislehurst,
	StationNameChiswick:                          StationCodeChiswick,
	StationNameCholsey:                           StationCodeCholsey,
	StationNameChorley:                           StationCodeChorley,
	StationNameChorleywood:                       StationCodeChorleywood,
	StationNameChristchurch:                      StationCodeChristchurch,
	StationNameChristsHospital:                   StationCodeChristsHospital,
	StationNameChurchAndOswaldtwistle:            StationCodeChurchAndOswaldtwistle,
	StationNameChurchFenton:                      StationCodeChurchFenton,
	StationNameChurchStretton:                    StationCodeChurchStretton,
	StationNameCilmeri:                           StationCodeCilmeri,
	StationNameCityThameslink:                    StationCodeCityThameslink,
	StationNameClactononSea:                      StationCodeClactononSea,
	StationNameClandon:                           StationCodeClandon,
	StationNameClaphamHighStreet:                 StationCodeClaphamHighStreet,
	StationNameClaphamJunction:                   StationCodeClaphamJunction,
	StationNameClaphamNorthYorkshire:             StationCodeClaphamNorthYorkshire,
	StationNameClapton:                           StationCodeClapton,
	StationNameClarbestonRoad:                    StationCodeClarbestonRoad,
	StationNameClarkston:                         StationCodeClarkston,
	StationNameClaverdon:                         StationCodeClaverdon,
	StationNameClaygate:                          StationCodeClaygate,
	StationNameCleethorpes:                       StationCodeCleethorpes,
	StationNameCleland:                           StationCodeCleland,
	StationNameCliftonDown:                       StationCodeCliftonDown,
	StationNameCliftonManchester:                 StationCodeCliftonManchester,
	StationNameClitheroe:                         StationCodeClitheroe,
	StationNameClockHouse:                        StationCodeClockHouse,
	StationNameClunderwen:                        StationCodeClunderwen,
	StationNameClydebank:                         StationCodeClydebank,
	StationNameCoatbridgeCentral:                 StationCodeCoatbridgeCentral,
	StationNameCoatbridgeSunnyside:               StationCodeCoatbridgeSunnyside,
	StationNameCoatdyke:                          StationCodeCoatdyke,
	StationNameCobhamAndStokedAbernon:            StationCodeCobhamAndStokedAbernon,
	StationNameCodsall:                           StationCodeCodsall,
	StationNameCogan:                             StationCodeCogan,
	StationNameColchester:                        StationCodeColchester,
	StationNameColchesterTown:                    StationCodeColchesterTown,
	StationNameColeshillParkway:                  StationCodeColeshillParkway,
	StationNameCollingham:                        StationCodeCollingham,
	StationNameCollington:                        StationCodeCollington,
	StationNameColne:                             StationCodeColne,
	StationNameColwall:                           StationCodeColwall,
	StationNameColwynBay:                         StationCodeColwynBay,
	StationNameCombeOxon:                         StationCodeCombeOxon,
	StationNameCommondale:                        StationCodeCommondale,
	StationNameCongleton:                         StationCodeCongleton,
	StationNameConisbrough:                       StationCodeConisbrough,
	StationNameConnelFerry:                       StationCodeConnelFerry,
	StationNameCononBridge:                       StationCodeCononBridge,
	StationNameCononley:                          StationCodeCononley,
	StationNameConwayPark:                        StationCodeConwayPark,
	StationNameConwy:                             StationCodeConwy,
	StationNameCoodenBeach:                       StationCodeCoodenBeach,
	StationNameCookham:                           StationCodeCookham,
	StationNameCooksbridge:                       StationCodeCooksbridge,
	StationNameCoombeJunctionHalt:                StationCodeCoombeJunctionHalt,
	StationNameCopplestone:                       StationCodeCopplestone,
	StationNameCorbridge:                         StationCodeCorbridge,
	StationNameCorby:                             StationCodeCorby,
	StationNameCorkerhill:                        StationCodeCorkerhill,
	StationNameCorkickle:                         StationCodeCorkickle,
	StationNameCorpach:                           StationCodeCorpach,
	StationNameCorrour:                           StationCodeCorrour,
	StationNameCoryton:                           StationCodeCoryton,
	StationNameCoseley:                           StationCodeCoseley,
	StationNameCosford:                           StationCodeCosford,
	StationNameCosham:                            StationCodeCosham,
	StationNameCottingham:                        StationCodeCottingham,
	StationNameCottingley:                        StationCodeCottingley,
	StationNameCoulsdonSouth:                     StationCodeCoulsdonSouth,
	StationNameCoulsdonTown:                      StationCodeCoulsdonTown,
	StationNameCoventry:                          StationCodeCoventry,
	StationNameCoventryArena:                     StationCodeCoventryArena,
	StationNameCowdenKent:                        StationCodeCowdenKent,
	StationNameCowdenbeath:                       StationCodeCowdenbeath,
	StationNameCradleyHeath:                      StationCodeCradleyHeath,
	StationNameCraigendoran:                      StationCodeCraigendoran,
	StationNameCramlington:                       StationCodeCramlington,
	StationNameCranbrookDevon:                    StationCodeCranbrookDevon,
	StationNameCravenArms:                        StationCodeCravenArms,
	StationNameCrawley:                           StationCodeCrawley,
	StationNameCrayford:                          StationCodeCrayford,
	StationNameCrediton:                          StationCodeCrediton,
	StationNameCressingEssex:                     StationCodeCressingEssex,
	StationNameCressington:                       StationCodeCressington,
	StationNameCreswell:                          StationCodeCreswell,
	StationNameCrewe:                             StationCodeCrewe,
	StationNameCrewkerne:                         StationCodeCrewkerne,
	StationNameCrewsHill:                         StationCodeCrewsHill,
	StationNameCrianlarich:                       StationCodeCrianlarich,
	StationNameCriccieth:                         StationCodeCriccieth,
	StationNameCricklewood:                       StationCodeCricklewood,
	StationNameCroftfoot:                         StationCodeCroftfoot,
	StationNameCroftonPark:                       StationCodeCroftonPark,
	StationNameCromer:                            StationCodeCromer,
	StationNameCromford:                          StationCodeCromford,
	StationNameCrookston:                         StationCodeCrookston,
	StationNameCrossGates:                        StationCodeCrossGates,
	StationNameCrossflatts:                       StationCodeCrossflatts,
	StationNameCrosshill:                         StationCodeCrosshill,
	StationNameCrosskeys:                         StationCodeCrosskeys,
	StationNameCrossmyloof:                       StationCodeCrossmyloof,
	StationNameCroston:                           StationCodeCroston,
	StationNameCrouchHill:                        StationCodeCrouchHill,
	StationNameCrowborough:                       StationCodeCrowborough,
	StationNameCrowhurst:                         StationCodeCrowhurst,
	StationNameCrowle:                            StationCodeCrowle,
	StationNameCrowthorne:                        StationCodeCrowthorne,
	StationNameCroy:                              StationCodeCroy,
	StationNameCrystalPalace:                     StationCodeCrystalPalace,
	StationNameCuddington:                        StationCodeCuddington,
	StationNameCuffley:                           StationCodeCuffley,
	StationNameCulham:                            StationCodeCulham,
	StationNameCulrain:                           StationCodeCulrain,
	StationNameCumbernauld:                       StationCodeCumbernauld,
	StationNameCupar:                             StationCodeCupar,
	StationNameCurriehill:                        StationCodeCurriehill,
	StationNameCuxton:                            StationCodeCuxton,
	StationNameCwmbach:                           StationCodeCwmbach,
	StationNameCwmbran:                           StationCodeCwmbran,
	StationNameCynghordy:                         StationCodeCynghordy,
	StationNameDagenhamDock:                      StationCodeDagenhamDock,
	StationNameDaisyHill:                         StationCodeDaisyHill,
	StationNameDalgetyBay:                        StationCodeDalgetyBay,
	StationNameDalmally:                          StationCodeDalmally,
	StationNameDalmarnock:                        StationCodeDalmarnock,
	StationNameDalmeny:                           StationCodeDalmeny,
	StationNameDalmuir:                           StationCodeDalmuir,
	StationNameDalreoch:                          StationCodeDalreoch,
	StationNameDalry:                             StationCodeDalry,
	StationNameDalstonCumbria:                    StationCodeDalstonCumbria,
	StationNameDalstonJunction:                   StationCodeDalstonJunction,
	StationNameDalstonKingsland:                  StationCodeDalstonKingsland,
	StationNameDaltonCumbria:                     StationCodeDaltonCumbria,
	StationNameDalwhinnie:                        StationCodeDalwhinnie,
	StationNameDanby:                             StationCodeDanby,
	StationNameDanescourt:                        StationCodeDanescourt,
	StationNameDanzey:                            StationCodeDanzey,
	StationNameDarlington:                        StationCodeDarlington,
	StationNameDarnall:                           StationCodeDarnall,
	StationNameDarsham:                           StationCodeDarsham,
	StationNameDartford:                          StationCodeDartford,
	StationNameDarton:                            StationCodeDarton,
	StationNameDarwen:                            StationCodeDarwen,
	StationNameDatchet:                           StationCodeDatchet,
	StationNameDavenport:                         StationCodeDavenport,
	StationNameDawlish:                           StationCodeDawlish,
	StationNameDawlishWarren:                     StationCodeDawlishWarren,
	StationNameDeal:                              StationCodeDeal,
	StationNameDeanWilts:                         StationCodeDeanWilts,
	StationNameDeansgate:                         StationCodeDeansgate,
	StationNameDeganwy:                           StationCodeDeganwy,
	StationNameDeighton:                          StationCodeDeighton,
	StationNameDelamere:                          StationCodeDelamere,
	StationNameDenbyDale:                         StationCodeDenbyDale,
	StationNameDenham:                            StationCodeDenham,
	StationNameDenhamGolfClub:                    StationCodeDenhamGolfClub,
	StationNameDenmarkHill:                       StationCodeDenmarkHill,
	StationNameDent:                              StationCodeDent,
	StationNameDenton:                            StationCodeDenton,
	StationNameDeptford:                          StationCodeDeptford,
	StationNameDerby:                             StationCodeDerby,
	StationNameDerbyRoadIpswich:                  StationCodeDerbyRoadIpswich,
	StationNameDevonportDevon:                    StationCodeDevonportDevon,
	StationNameDevonportDockyard:                 StationCodeDevonportDockyard,
	StationNameDewsbury:                          StationCodeDewsbury,
	StationNameDidcotParkway:                     StationCodeDidcotParkway,
	StationNameDigbyAndSowton:                    StationCodeDigbyAndSowton,
	StationNameDiltonMarsh:                       StationCodeDiltonMarsh,
	StationNameDinasPowys:                        StationCodeDinasPowys,
	StationNameDinasRhondda:                      StationCodeDinasRhondda,
	StationNameDingleRoad:                        StationCodeDingleRoad,
	StationNameDingwall:                          StationCodeDingwall,
	StationNameDinsdale:                          StationCodeDinsdale,
	StationNameDinting:                           StationCodeDinting,
	StationNameDisley:                            StationCodeDisley,
	StationNameDiss:                              StationCodeDiss,
	StationNameDodworth:                          StationCodeDodworth,
	StationNameDolau:                             StationCodeDolau,
	StationNameDoleham:                           StationCodeDoleham,
	StationNameDolgarrog:                         StationCodeDolgarrog,
	StationNameDolwyddelan:                       StationCodeDolwyddelan,
	StationNameDoncaster:                         StationCodeDoncaster,
	StationNameDorchesterSouth:                   StationCodeDorchesterSouth,
	StationNameDorchesterWest:                    StationCodeDorchesterWest,
	StationNameDoreAndTotley:                     StationCodeDoreAndTotley,
	StationNameDorkingDeepdene:                   StationCodeDorkingDeepdene,
	StationNameDorkingMain:                       StationCodeDorkingMain,
	StationNameDorkingWest:                       StationCodeDorkingWest,
	StationNameDormans:                           StationCodeDormans,
	StationNameDorridge:                          StationCodeDorridge,
	StationNameDoveHoles:                         StationCodeDoveHoles,
	StationNameDoverPriory:                       StationCodeDoverPriory,
	StationNameDovercourt:                        StationCodeDovercourt,
	StationNameDoveyJunction:                     StationCodeDoveyJunction,
	StationNameDownhamMarket:                     StationCodeDownhamMarket,
	StationNameDraytonGreen:                      StationCodeDraytonGreen,
	StationNameDraytonPark:                       StationCodeDraytonPark,
	StationNameDrem:                              StationCodeDrem,
	StationNameDriffield:                         StationCodeDriffield,
	StationNameDrigg:                             StationCodeDrigg,
	StationNameDroitwichSpa:                      StationCodeDroitwichSpa,
	StationNameDronfield:                         StationCodeDronfield,
	StationNameDrumchapel:                        StationCodeDrumchapel,
	StationNameDrumfrochar:                       StationCodeDrumfrochar,
	StationNameDrumgelloch:                       StationCodeDrumgelloch,
	StationNameDrumry:                            StationCodeDrumry,
	StationNameDublinFerryport:                   StationCodeDublinFerryport,
	StationNameDublinPortStena:                   StationCodeDublinPortStena,
	StationNameDuddeston:                         StationCodeDuddeston,
	StationNameDudleyPort:                        StationCodeDudleyPort,
	StationNameDuffield:                          StationCodeDuffield,
	StationNameDuirinish:                         StationCodeDuirinish,
	StationNameDukeStreet:                        StationCodeDukeStreet,
	StationNameDullingham:                        StationCodeDullingham,
	StationNameDumbartonCentral:                  StationCodeDumbartonCentral,
	StationNameDumbartonEast:                     StationCodeDumbartonEast,
	StationNameDumbreck:                          StationCodeDumbreck,
	StationNameDumfries:                          StationCodeDumfries,
	StationNameDumptonPark:                       StationCodeDumptonPark,
	StationNameDunbar:                            StationCodeDunbar,
	StationNameDunblane:                          StationCodeDunblane,
	StationNameDuncraig:                          StationCodeDuncraig,
	StationNameDundee:                            StationCodeDundee,
	StationNameDunfermlineQueenMargaret:          StationCodeDunfermlineQueenMargaret,
	StationNameDunfermlineTown:                   StationCodeDunfermlineTown,
	StationNameDunkeldAndBirnam:                  StationCodeDunkeldAndBirnam,
	StationNameDunlop:                            StationCodeDunlop,
	StationNameDunrobinCastle:                    StationCodeDunrobinCastle,
	StationNameDunston:                           StationCodeDunston,
	StationNameDuntonGreen:                       StationCodeDuntonGreen,
	StationNameDurham:                            StationCodeDurham,
	StationNameDurringtononSea:                   StationCodeDurringtononSea,
	StationNameDyce:                              StationCodeDyce,
	StationNameDyffrynArdudwy:                    StationCodeDyffrynArdudwy,
	StationNameEaglescliffe:                      StationCodeEaglescliffe,
	StationNameEalingBroadway:                    StationCodeEalingBroadway,
	StationNameEarlestown:                        StationCodeEarlestown,
	StationNameEarley:                            StationCodeEarley,
	StationNameEarlsfield:                        StationCodeEarlsfield,
	StationNameEarlswoodSurrey:                   StationCodeEarlswoodSurrey,
	StationNameEarlswoodWestMidlands:             StationCodeEarlswoodWestMidlands,
	StationNameEastCroydon:                       StationCodeEastCroydon,
	StationNameEastDidsbury:                      StationCodeEastDidsbury,
	StationNameEastDulwich:                       StationCodeEastDulwich,
	StationNameEastFarleigh:                      StationCodeEastFarleigh,
	StationNameEastGarforth:                      StationCodeEastGarforth,
	StationNameEastGrinstead:                     StationCodeEastGrinstead,
	StationNameEastKilbride:                      StationCodeEastKilbride,
	StationNameEastMalling:                       StationCodeEastMalling,
	StationNameEastMidlandsParkway:               StationCodeEastMidlandsParkway,
	StationNameEastTilbury:                       StationCodeEastTilbury,
	StationNameEastWorthing:                      StationCodeEastWorthing,
	StationNameEastbourne:                        StationCodeEastbourne,
	StationNameEastbrook:                         StationCodeEastbrook,
	StationNameEasterhouse:                       StationCodeEasterhouse,
	StationNameEasthamRake:                       StationCodeEasthamRake,
	StationNameEastleigh:                         StationCodeEastleigh,
	StationNameEastrington:                       StationCodeEastrington,
	StationNameEbbsfleetInternational:            StationCodeEbbsfleetInternational,
	StationNameEbbwValeParkway:                   StationCodeEbbwValeParkway,
	StationNameEbbwValeTown:                      StationCodeEbbwValeTown,
	StationNameEcclesManchester:                  StationCodeEcclesManchester,
	StationNameEcclesRoad:                        StationCodeEcclesRoad,
	StationNameEcclestonPark:                     StationCodeEcclestonPark,
	StationNameEdale:                             StationCodeEdale,
	StationNameEdenPark:                          StationCodeEdenPark,
	StationNameEdenbridge:                        StationCodeEdenbridge,
	StationNameEdenbridgeTown:                    StationCodeEdenbridgeTown,
	StationNameEdgeHill:                          StationCodeEdgeHill,
	StationNameEdinburgh:                         StationCodeEdinburgh,
	StationNameEdinburghGateway:                  StationCodeEdinburghGateway,
	StationNameEdinburghPark:                     StationCodeEdinburghPark,
	StationNameEdmontonGreen:                     StationCodeEdmontonGreen,
	StationNameEffinghamJunction:                 StationCodeEffinghamJunction,
	StationNameEggesford:                         StationCodeEggesford,
	StationNameEgham:                             StationCodeEgham,
	StationNameEgton:                             StationCodeEgton,
	StationNameElephantAndCastle:                 StationCodeElephantAndCastle,
	StationNameElephantAndCastleUnderground:      StationCodeElephantAndCastleUnderground,
	StationNameElgin:                             StationCodeElgin,
	StationNameEllesmerePort:                     StationCodeEllesmerePort,
	StationNameElmersEnd:                         StationCodeElmersEnd,
	StationNameElmsteadWoods:                     StationCodeElmsteadWoods,
	StationNameElmswell:                          StationCodeElmswell,
	StationNameElsecar:                           StationCodeElsecar,
	StationNameElsenhamEssex:                     StationCodeElsenhamEssex,
	StationNameElstreeAndBorehamwood:             StationCodeElstreeAndBorehamwood,
	StationNameEltham:                            StationCodeEltham,
	StationNameEltonAndOrston:                    StationCodeEltonAndOrston,
	StationNameEly:                               StationCodeEly,
	StationNameEmersonPark:                       StationCodeEmersonPark,
	StationNameEmsworth:                          StationCodeEmsworth,
	StationNameEnerglynAndChurchillPark:          StationCodeEnerglynAndChurchillPark,
	StationNameEnfieldChase:                      StationCodeEnfieldChase,
	StationNameEnfieldLock:                       StationCodeEnfieldLock,
	StationNameEnfieldTown:                       StationCodeEnfieldTown,
	StationNameEntwistle:                         StationCodeEntwistle,
	StationNameEpsomDowns:                        StationCodeEpsomDowns,
	StationNameEpsomSurrey:                       StationCodeEpsomSurrey,
	StationNameErdington:                         StationCodeErdington,
	StationNameEridge:                            StationCodeEridge,
	StationNameErith:                             StationCodeErith,
	StationNameEsher:                             StationCodeEsher,
	StationNameEskbank:                           StationCodeEskbank,
	StationNameEssexRoad:                         StationCodeEssexRoad,
	StationNameEtchingham:                        StationCodeEtchingham,
	StationNameEuxtonBalshawLane:                 StationCodeEuxtonBalshawLane,
	StationNameEvesham:                           StationCodeEvesham,
	StationNameEwellEast:                         StationCodeEwellEast,
	StationNameEwellWest:                         StationCodeEwellWest,
	StationNameExeterCentral:                     StationCodeExeterCentral,
	StationNameExeterStDavids:                    StationCodeExeterStDavids,
	StationNameExeterStThomas:                    StationCodeExeterStThomas,
	StationNameExhibitionCentreGlasgow:           StationCodeExhibitionCentreGlasgow,
	StationNameExmouth:                           StationCodeExmouth,
	StationNameExton:                             StationCodeExton,
	StationNameEynsford:                          StationCodeEynsford,
	StationNameFairbourne:                        StationCodeFairbourne,
	StationNameFairfield:                         StationCodeFairfield,
	StationNameFairlie:                           StationCodeFairlie,
	StationNameFairwater:                         StationCodeFairwater,
	StationNameFalconwood:                        StationCodeFalconwood,
	StationNameFalkirkGrahamston:                 StationCodeFalkirkGrahamston,
	StationNameFalkirkHigh:                       StationCodeFalkirkHigh,
	StationNameFallsofCruachan:                   StationCodeFallsofCruachan,
	StationNameFalmer:                            StationCodeFalmer,
	StationNameFalmouthDocks:                     StationCodeFalmouthDocks,
	StationNameFalmouthTown:                      StationCodeFalmouthTown,
	StationNameFareham:                           StationCodeFareham,
	StationNameFarnboroughMain:                   StationCodeFarnboroughMain,
	StationNameFarnboroughNorth:                  StationCodeFarnboroughNorth,
	StationNameFarncombe:                         StationCodeFarncombe,
	StationNameFarnham:                           StationCodeFarnham,
	StationNameFarninghamRoad:                    StationCodeFarninghamRoad,
	StationNameFarnworth:                         StationCodeFarnworth,
	StationNameFarringdon:                        StationCodeFarringdon,
	StationNameFauldhouse:                        StationCodeFauldhouse,
	StationNameFaversham:                         StationCodeFaversham,
	StationNameFaygate:                           StationCodeFaygate,
	StationNameFazakerley:                        StationCodeFazakerley,
	StationNameFearn:                             StationCodeFearn,
	StationNameFeatherstone:                      StationCodeFeatherstone,
	StationNameFelixstowe:                        StationCodeFelixstowe,
	StationNameFeltham:                           StationCodeFeltham,
	StationNameFeniton:                           StationCodeFeniton,
	StationNameFennyStratford:                    StationCodeFennyStratford,
	StationNameFernhill:                          StationCodeFernhill,
	StationNameFerriby:                           StationCodeFerriby,
	StationNameFerryside:                         StationCodeFerryside,
	StationNameFfairfach:                         StationCodeFfairfach,
	StationNameFiley:                             StationCodeFiley,
	StationNameFiltonAbbeyWood:                   StationCodeFiltonAbbeyWood,
	StationNameFinchleyRoadAndFrognal:            StationCodeFinchleyRoadAndFrognal,
	StationNameFinsburyPark:                      StationCodeFinsburyPark,
	StationNameFinstock:                          StationCodeFinstock,
	StationNameFishbourneSussex:                  StationCodeFishbourneSussex,
	StationNameFishersgate:                       StationCodeFishersgate,
	StationNameFishguardAndGoodwick:              StationCodeFishguardAndGoodwick,
	StationNameFishguardHarbour:                  StationCodeFishguardHarbour,
	StationNameFiskerton:                         StationCodeFiskerton,
	StationNameFitzwilliam:                       StationCodeFitzwilliam,
	StationNameFiveWays:                          StationCodeFiveWays,
	StationNameFleet:                             StationCodeFleet,
	StationNameFlimby:                            StationCodeFlimby,
	StationNameFlint:                             StationCodeFlint,
	StationNameFlitwick:                          StationCodeFlitwick,
	StationNameFlixton:                           StationCodeFlixton,
	StationNameFloweryField:                      StationCodeFloweryField,
	StationNameFolkestoneCentral:                 StationCodeFolkestoneCentral,
	StationNameFolkestoneWest:                    StationCodeFolkestoneWest,
	StationNameFord:                              StationCodeFord,
	StationNameForestGate:                        StationCodeForestGate,
	StationNameForestHill:                        StationCodeForestHill,
	StationNameFormby:                            StationCodeFormby,
	StationNameForres:                            StationCodeForres,
	StationNameForsinard:                         StationCodeForsinard,
	StationNameFortMatilda:                       StationCodeFortMatilda,
	StationNameFortWilliam:                       StationCodeFortWilliam,
	StationNameFourOaks:                          StationCodeFourOaks,
	StationNameFoxfield:                          StationCodeFoxfield,
	StationNameFoxton:                            StationCodeFoxton,
	StationNameFrant:                             StationCodeFrant,
	StationNameFratton:                           StationCodeFratton,
	StationNameFreshfield:                        StationCodeFreshfield,
	StationNameFreshford:                         StationCodeFreshford,
	StationNameFrimley:                           StationCodeFrimley,
	StationNameFrintononSea:                      StationCodeFrintononSea,
	StationNameFrizinghall:                       StationCodeFrizinghall,
	StationNameFrodsham:                          StationCodeFrodsham,
	StationNameFrome:                             StationCodeFrome,
	StationNameFulwell:                           StationCodeFulwell,
	StationNameFurnessVale:                       StationCodeFurnessVale,
	StationNameFurzePlatt:                        StationCodeFurzePlatt,
	StationNameGainsboroughCentral:               StationCodeGainsboroughCentral,
	StationNameGainsboroughLeaRoad:               StationCodeGainsboroughLeaRoad,
	StationNameGalashiels:                        StationCodeGalashiels,
	StationNameGarelochhead:                      StationCodeGarelochhead,
	StationNameGarforth:                          StationCodeGarforth,
	StationNameGargrave:                          StationCodeGargrave,
	StationNameGarrowhill:                        StationCodeGarrowhill,
	StationNameGarscadden:                        StationCodeGarscadden,
	StationNameGarsdale:                          StationCodeGarsdale,
	StationNameGarstonHertfordshire:              StationCodeGarstonHertfordshire,
	StationNameGarswood:                          StationCodeGarswood,
	StationNameGartcosh:                          StationCodeGartcosh,
	StationNameGarthMidGlamorgan:                 StationCodeGarthMidGlamorgan,
	StationNameGarthPowys:                        StationCodeGarthPowys,
	StationNameGarve:                             StationCodeGarve,
	StationNameGathurst:                          StationCodeGathurst,
	StationNameGatley:                            StationCodeGatley,
	StationNameGatwickAirport:                    StationCodeGatwickAirport,
	StationNameGeorgemasJunction:                 StationCodeGeorgemasJunction,
	StationNameGerrardsCross:                     StationCodeGerrardsCross,
	StationNameGideaPark:                         StationCodeGideaPark,
	StationNameGiffnock:                          StationCodeGiffnock,
	StationNameGiggleswick:                       StationCodeGiggleswick,
	StationNameGilberdyke:                        StationCodeGilberdyke,
	StationNameGilfachFargoed:                    StationCodeGilfachFargoed,
	StationNameGillinghamDorset:                  StationCodeGillinghamDorset,
	StationNameGillinghamKent:                    StationCodeGillinghamKent,
	StationNameGilshochill:                       StationCodeGilshochill,
	StationNameGipsyHill:                         StationCodeGipsyHill,
	StationNameGirvan:                            StationCodeGirvan,
	StationNameGlaisdale:                         StationCodeGlaisdale,
	StationNameGlanConwy:                         StationCodeGlanConwy,
	StationNameGlasgowCentral:                    StationCodeGlasgowCentral,
	StationNameGlasgowQueenStreet:                StationCodeGlasgowQueenStreet,
	StationNameGlasshoughton:                     StationCodeGlasshoughton,
	StationNameGlazebrook:                        StationCodeGlazebrook,
	StationNameGleneagles:                        StationCodeGleneagles,
	StationNameGlenfinnan:                        StationCodeGlenfinnan,
	StationNameGlengarnock:                       StationCodeGlengarnock,
	StationNameGlenrotheswithThornton:            StationCodeGlenrotheswithThornton,
	StationNameGlossop:                           StationCodeGlossop,
	StationNameGloucester:                        StationCodeGloucester,
	StationNameGlynde:                            StationCodeGlynde,
	StationNameGobowen:                           StationCodeGobowen,
	StationNameGodalming:                         StationCodeGodalming,
	StationNameGodley:                            StationCodeGodley,
	StationNameGodstone:                          StationCodeGodstone,
	StationNameGoldthorpe:                        StationCodeGoldthorpe,
	StationNameGolfStreet:                        StationCodeGolfStreet,
	StationNameGolspie:                           StationCodeGolspie,
	StationNameGomshall:                          StationCodeGomshall,
	StationNameGoodmayes:                         StationCodeGoodmayes,
	StationNameGoole:                             StationCodeGoole,
	StationNameGoostrey:                          StationCodeGoostrey,
	StationNameGordonHill:                        StationCodeGordonHill,
	StationNameGorebridge:                        StationCodeGorebridge,
	StationNameGoringAndStreatley:                StationCodeGoringAndStreatley,
	StationNameGoringbySea:                       StationCodeGoringbySea,
	StationNameGorton:                            StationCodeGorton,
	StationNameGospelOak:                         StationCodeGospelOak,
	StationNameGourock:                           StationCodeGourock,
	StationNameGowerton:                          StationCodeGowerton,
	StationNameGoxhill:                           StationCodeGoxhill,
	StationNameGrangeOverSands:                   StationCodeGrangeOverSands,
	StationNameGrangePark:                        StationCodeGrangePark,
	StationNameGrangetownCardiff:                 StationCodeGrangetownCardiff,
	StationNameGrantham:                          StationCodeGrantham,
	StationNameGrateley:                          StationCodeGrateley,
	StationNameGravellyHill:                      StationCodeGravellyHill,
	StationNameGravesend:                         StationCodeGravesend,
	StationNameGrays:                             StationCodeGrays,
	StationNameGreatAyton:                        StationCodeGreatAyton,
	StationNameGreatBentley:                      StationCodeGreatBentley,
	StationNameGreatChesterford:                  StationCodeGreatChesterford,
	StationNameGreatCoates:                       StationCodeGreatCoates,
	StationNameGreatMalvern:                      StationCodeGreatMalvern,
	StationNameGreatMissenden:                    StationCodeGreatMissenden,
	StationNameGreatYarmouth:                     StationCodeGreatYarmouth,
	StationNameGreenLane:                         StationCodeGreenLane,
	StationNameGreenRoad:                         StationCodeGreenRoad,
	StationNameGreenbank:                         StationCodeGreenbank,
	StationNameGreenfaulds:                       StationCodeGreenfaulds,
	StationNameGreenfield:                        StationCodeGreenfield,
	StationNameGreenford:                         StationCodeGreenford,
	StationNameGreenhitheForBluewater:            StationCodeGreenhitheForBluewater,
	StationNameGreenockCentral:                   StationCodeGreenockCentral,
	StationNameGreenockWest:                      StationCodeGreenockWest,
	StationNameGreenwich:                         StationCodeGreenwich,
	StationNameGretnaGreen:                       StationCodeGretnaGreen,
	StationNameGrimsbyDocks:                      StationCodeGrimsbyDocks,
	StationNameGrimsbyTown:                       StationCodeGrimsbyTown,
	StationNameGrindleford:                       StationCodeGrindleford,
	StationNameGrosmont:                          StationCodeGrosmont,
	StationNameGrovePark:                         StationCodeGrovePark,
	StationNameGuideBridge:                       StationCodeGuideBridge,
	StationNameGuildford:                         StationCodeGuildford,
	StationNameGuiseley:                          StationCodeGuiseley,
	StationNameGunnersbury:                       StationCodeGunnersbury,
	StationNameGunnislake:                        StationCodeGunnislake,
	StationNameGunton:                            StationCodeGunton,
	StationNameGwersyllt:                         StationCodeGwersyllt,
	StationNameGypsyLane:                         StationCodeGypsyLane,
	StationNameHabrough:                          StationCodeHabrough,
	StationNameHackbridge:                        StationCodeHackbridge,
	StationNameHackneyCentral:                    StationCodeHackneyCentral,
	StationNameHackneyDowns:                      StationCodeHackneyDowns,
	StationNameHackneyWick:                       StationCodeHackneyWick,
	StationNameHaddenhamAndThameParkway:          StationCodeHaddenhamAndThameParkway,
	StationNameHaddiscoe:                         StationCodeHaddiscoe,
	StationNameHadfield:                          StationCodeHadfield,
	StationNameHadleyWood:                        StationCodeHadleyWood,
	StationNameHagFold:                           StationCodeHagFold,
	StationNameHaggerston:                        StationCodeHaggerston,
	StationNameHagley:                            StationCodeHagley,
	StationNameHairmyres:                         StationCodeHairmyres,
	StationNameHaleManchester:                    StationCodeHaleManchester,
	StationNameHalesworth:                        StationCodeHalesworth,
	StationNameHalewood:                          StationCodeHalewood,
	StationNameHalifax:                           StationCodeHalifax,
	StationNameHallGreen:                         StationCodeHallGreen,
	StationNameHallRoad:                          StationCodeHallRoad,
	StationNameHalling:                           StationCodeHalling,
	StationNameHallithWood:                       StationCodeHallithWood,
	StationNameHaltwhistle:                       StationCodeHaltwhistle,
	StationNameHamStreet:                         StationCodeHamStreet,
	StationNameHamble:                            StationCodeHamble,
	StationNameHamiltonCentral:                   StationCodeHamiltonCentral,
	StationNameHamiltonWest:                      StationCodeHamiltonWest,
	StationNameHammerton:                         StationCodeHammerton,
	StationNameHampdenParkSussex:                 StationCodeHampdenParkSussex,
	StationNameHampsteadHeath:                    StationCodeHampsteadHeath,
	StationNameHamptonCourt:                      StationCodeHamptonCourt,
	StationNameHamptonLondon:                     StationCodeHamptonLondon,
	StationNameHamptonWick:                       StationCodeHamptonWick,
	StationNameHamptoninArden:                    StationCodeHamptoninArden,
	StationNameHamsteadBirmingham:                StationCodeHamsteadBirmingham,
	StationNameHamworthy:                         StationCodeHamworthy,
	StationNameHanborough:                        StationCodeHanborough,
	StationNameHandforth:                         StationCodeHandforth,
	StationNameHanwell:                           StationCodeHanwell,
	StationNameHapton:                            StationCodeHapton,
	StationNameHarlech:                           StationCodeHarlech,
	StationNameHarlesden:                         StationCodeHarlesden,
	StationNameHarlingRoad:                       StationCodeHarlingRoad,
	StationNameHarlingtonBeds:                    StationCodeHarlingtonBeds,
	StationNameHarlowMill:                        StationCodeHarlowMill,
	StationNameHarlowTown:                        StationCodeHarlowTown,
	StationNameHaroldWood:                        StationCodeHaroldWood,
	StationNameHarpenden:                         StationCodeHarpenden,
	StationNameHarrietsham:                       StationCodeHarrietsham,
	StationNameHarringay:                         StationCodeHarringay,
	StationNameHarringayGreenLanes:               StationCodeHarringayGreenLanes,
	StationNameHarrington:                        StationCodeHarrington,
	StationNameHarrogate:                         StationCodeHarrogate,
	StationNameHarrowAndWealdstone:               StationCodeHarrowAndWealdstone,
	StationNameHarrowontheHill:                   StationCodeHarrowontheHill,
	StationNameHartfordCheshire:                  StationCodeHartfordCheshire,
	StationNameHartlebury:                        StationCodeHartlebury,
	StationNameHartlepool:                        StationCodeHartlepool,
	StationNameHartwood:                          StationCodeHartwood,
	StationNameHarwichInternational:              StationCodeHarwichInternational,
	StationNameHarwichTown:                       StationCodeHarwichTown,
	StationNameHaslemere:                         StationCodeHaslemere,
	StationNameHassocks:                          StationCodeHassocks,
	StationNameHastings:                          StationCodeHastings,
	StationNameHatchEnd:                          StationCodeHatchEnd,
	StationNameHatfieldAndStainforth:             StationCodeHatfieldAndStainforth,
	StationNameHatfieldHerts:                     StationCodeHatfieldHerts,
	StationNameHatfieldPeverel:                   StationCodeHatfieldPeverel,
	StationNameHathersage:                        StationCodeHathersage,
	StationNameHattersley:                        StationCodeHattersley,
	StationNameHatton:                            StationCodeHatton,
	StationNameHavant:                            StationCodeHavant,
	StationNameHavenhouse:                        StationCodeHavenhouse,
	StationNameHaverfordwest:                     StationCodeHaverfordwest,
	StationNameHawarden:                          StationCodeHawarden,
	StationNameHawardenBridge:                    StationCodeHawardenBridge,
	StationNameHawkhead:                          StationCodeHawkhead,
	StationNameHaydonBridge:                      StationCodeHaydonBridge,
	StationNameHaydonsRoad:                       StationCodeHaydonsRoad,
	StationNameHayesAndHarlington:                StationCodeHayesAndHarlington,
	StationNameHayesKent:                         StationCodeHayesKent,
	StationNameHayle:                             StationCodeHayle,
	StationNameHaymarket:                         StationCodeHaymarket,
	StationNameHaywardsHeath:                     StationCodeHaywardsHeath,
	StationNameHazelGrove:                        StationCodeHazelGrove,
	StationNameHeadcorn:                          StationCodeHeadcorn,
	StationNameHeadingley:                        StationCodeHeadingley,
	StationNameHeadstoneLane:                     StationCodeHeadstoneLane,
	StationNameHealdGreen:                        StationCodeHealdGreen,
	StationNameHealing:                           StationCodeHealing,
	StationNameHeathHighLevel:                    StationCodeHeathHighLevel,
	StationNameHeathLowLevel:                     StationCodeHeathLowLevel,
	StationNameHeathrowAirportTerminal4:          StationCodeHeathrowAirportTerminal4,
	StationNameHeathrowAirportTerminal5:          StationCodeHeathrowAirportTerminal5,
	StationNameHeathrowAirportTerminals12and3:    StationCodeHeathrowAirportTerminals12and3,
	StationNameHeatonChapel:                      StationCodeHeatonChapel,
	StationNameHebdenBridge:                      StationCodeHebdenBridge,
	StationNameHeckington:                        StationCodeHeckington,
	StationNameHedgeEnd:                          StationCodeHedgeEnd,
	StationNameHednesford:                        StationCodeHednesford,
	StationNameHeighington:                       StationCodeHeighington,
	StationNameHelensburghCentral:                StationCodeHelensburghCentral,
	StationNameHelensburghUpper:                  StationCodeHelensburghUpper,
	StationNameHellifield:                        StationCodeHellifield,
	StationNameHelmsdale:                         StationCodeHelmsdale,
	StationNameHelsby:                            StationCodeHelsby,
	StationNameHemelHempstead:                    StationCodeHemelHempstead,
	StationNameHendon:                            StationCodeHendon,
	StationNameHengoed:                           StationCodeHengoed,
	StationNameHenleyinArden:                     StationCodeHenleyinArden,
	StationNameHenleyonThames:                    StationCodeHenleyonThames,
	StationNameHensall:                           StationCodeHensall,
	StationNameHereford:                          StationCodeHereford,
	StationNameHerneBay:                          StationCodeHerneBay,
	StationNameHerneHill:                         StationCodeHerneHill,
	StationNameHersham:                           StationCodeHersham,
	StationNameHertfordEast:                      StationCodeHertfordEast,
	StationNameHertfordNorth:                     StationCodeHertfordNorth,
	StationNameHessle:                            StationCodeHessle,
	StationNameHeswall:                           StationCodeHeswall,
	StationNameHever:                             StationCodeHever,
	StationNameHeworth:                           StationCodeHeworth,
	StationNameHexham:                            StationCodeHexham,
	StationNameHeyford:                           StationCodeHeyford,
	StationNameHeyshamPort:                       StationCodeHeyshamPort,
	StationNameHighBrooms:                        StationCodeHighBrooms,
	StationNameHighStreetGlasgow:                 StationCodeHighStreetGlasgow,
	StationNameHighStreetKensingtonUnderground:   StationCodeHighStreetKensingtonUnderground,
	StationNameHighWycombe:                       StationCodeHighWycombe,
	StationNameHigham:                            StationCodeHigham,
	StationNameHighamsPark:                       StationCodeHighamsPark,
	StationNameHighbridgeAndBurnham:              StationCodeHighbridgeAndBurnham,
	StationNameHighburyAndIslington:              StationCodeHighburyAndIslington,
	StationNameHightown:                          StationCodeHightown,
	StationNameHildenborough:                     StationCodeHildenborough,
	StationNameHillfoot:                          StationCodeHillfoot,
	StationNameHillingtonEast:                    StationCodeHillingtonEast,
	StationNameHillingtonWest:                    StationCodeHillingtonWest,
	StationNameHillside:                          StationCodeHillside,
	StationNameHilsea:                            StationCodeHilsea,
	StationNameHinchleyWood:                      StationCodeHinchleyWood,
	StationNameHinckleyLeics:                     StationCodeHinckleyLeics,
	StationNameHindley:                           StationCodeHindley,
	StationNameHintonAdmiral:                     StationCodeHintonAdmiral,
	StationNameHitchin:                           StationCodeHitchin,
	StationNameHitherGreen:                       StationCodeHitherGreen,
	StationNameHockley:                           StationCodeHockley,
	StationNameHollingbourne:                     StationCodeHollingbourne,
	StationNameHolmesChapel:                      StationCodeHolmesChapel,
	StationNameHolmwood:                          StationCodeHolmwood,
	StationNameHoltonHeath:                       StationCodeHoltonHeath,
	StationNameHolyhead:                          StationCodeHolyhead,
	StationNameHolytown:                          StationCodeHolytown,
	StationNameHomerton:                          StationCodeHomerton,
	StationNameHoneybourne:                       StationCodeHoneybourne,
	StationNameHoniton:                           StationCodeHoniton,
	StationNameHonley:                            StationCodeHonley,
	StationNameHonorOakPark:                      StationCodeHonorOakPark,
	StationNameHook:                              StationCodeHook,
	StationNameHooton:                            StationCodeHooton,
	StationNameHopeDerbyshire:                    StationCodeHopeDerbyshire,
	StationNameHopeFlintshire:                    StationCodeHopeFlintshire,
	StationNameHoptonHeath:                       StationCodeHoptonHeath,
	StationNameHorley:                            StationCodeHorley,
	StationNameHornbeamPark:                      StationCodeHornbeamPark,
	StationNameHornsey:                           StationCodeHornsey,
	StationNameHorsforth:                         StationCodeHorsforth,
	StationNameHorsham:                           StationCodeHorsham,
	StationNameHorsley:                           StationCodeHorsley,
	StationNameHortoninRibblesdale:               StationCodeHortoninRibblesdale,
	StationNameHorwichParkway:                    StationCodeHorwichParkway,
	StationNameHoscar:                            StationCodeHoscar,
	StationNameHoughGreen:                        StationCodeHoughGreen,
	StationNameHounslow:                          StationCodeHounslow,
	StationNameHove:                              StationCodeHove,
	StationNameHovetonAndWroxham:                 StationCodeHovetonAndWroxham,
	StationNameHowWoodHerts:                      StationCodeHowWoodHerts,
	StationNameHowden:                            StationCodeHowden,
	StationNameHowwoodRenfrewshire:               StationCodeHowwoodRenfrewshire,
	StationNameHoxton:                            StationCodeHoxton,
	StationNameHoylake:                           StationCodeHoylake,
	StationNameHubbertsBridge:                    StationCodeHubbertsBridge,
	StationNameHucknall:                          StationCodeHucknall,
	StationNameHuddersfield:                      StationCodeHuddersfield,
	StationNameHull:                              StationCodeHull,
	StationNameHumphreyPark:                      StationCodeHumphreyPark,
	StationNameHuncoat:                           StationCodeHuncoat,
	StationNameHungerford:                        StationCodeHungerford,
	StationNameHunmanby:                          StationCodeHunmanby,
	StationNameHuntingdon:                        StationCodeHuntingdon,
	StationNameHuntly:                            StationCodeHuntly,
	StationNameHuntsCross:                        StationCodeHuntsCross,
	StationNameHurstGreen:                        StationCodeHurstGreen,
	StationNameHuttonCranswick:                   StationCodeHuttonCranswick,
	StationNameHuyton:                            StationCodeHuyton,
	StationNameHydeCentral:                       StationCodeHydeCentral,
	StationNameHydeNorth:                         StationCodeHydeNorth,
	StationNameHykeham:                           StationCodeHykeham,
	StationNameHyndland:                          StationCodeHyndland,
	StationNameHytheEssex:                        StationCodeHytheEssex,
	StationNameIBMHalt:                           StationCodeIBMHalt,
	StationNameIfield:                            StationCodeIfield,
	StationNameIlford:                            StationCodeIlford,
	StationNameIlkley:                            StationCodeIlkley,
	StationNameImperialWharf:                     StationCodeImperialWharf,
	StationNameInceAndElton:                      StationCodeInceAndElton,
	StationNameInceManchester:                    StationCodeInceManchester,
	StationNameIngatestone:                       StationCodeIngatestone,
	StationNameInsch:                             StationCodeInsch,
	StationNameInvergordon:                       StationCodeInvergordon,
	StationNameInvergowrie:                       StationCodeInvergowrie,
	StationNameInverkeithing:                     StationCodeInverkeithing,
	StationNameInverkip:                          StationCodeInverkip,
	StationNameInverness:                         StationCodeInverness,
	StationNameInvershin:                         StationCodeInvershin,
	StationNameInverurie:                         StationCodeInverurie,
	StationNameIpswich:                           StationCodeIpswich,
	StationNameIrlam:                             StationCodeIrlam,
	StationNameIrvine:                            StationCodeIrvine,
	StationNameIsleworth:                         StationCodeIsleworth,
	StationNameIslip:                             StationCodeIslip,
	StationNameIver:                              StationCodeIver,
	StationNameIvybridge:                         StationCodeIvybridge,
	StationNameJamesCook:                         StationCodeJamesCook,
	StationNameJewelleryQuarter:                  StationCodeJewelleryQuarter,
	StationNameJohnstonPembs:                     StationCodeJohnstonPembs,
	StationNameJohnstoneStrathclyde:              StationCodeJohnstoneStrathclyde,
	StationNameJordanhill:                        StationCodeJordanhill,
	StationNameKearsleyManchester:                StationCodeKearsleyManchester,
	StationNameKearsney:                          StationCodeKearsney,
	StationNameKeighley:                          StationCodeKeighley,
	StationNameKeith:                             StationCodeKeith,
	StationNameKelvedon:                          StationCodeKelvedon,
	StationNameKelvindale:                        StationCodeKelvindale,
	StationNameKemble:                            StationCodeKemble,
	StationNameKempstonHardwick:                  StationCodeKempstonHardwick,
	StationNameKemptonParkRacecourse:             StationCodeKemptonParkRacecourse,
	StationNameKemsing:                           StationCodeKemsing,
	StationNameKemsley:                           StationCodeKemsley,
	StationNameKendal:                            StationCodeKendal,
	StationNameKenley:                            StationCodeKenley,
	StationNameKennett:                           StationCodeKennett,
	StationNameKennishead:                        StationCodeKennishead,
	StationNameKensalGreen:                       StationCodeKensalGreen,
	StationNameKensalRise:                        StationCodeKensalRise,
	StationNameKensingtonOlympia:                 StationCodeKensingtonOlympia,
	StationNameKentHouse:                         StationCodeKentHouse,
	StationNameKentishTown:                       StationCodeKentishTown,
	StationNameKentishTownWest:                   StationCodeKentishTownWest,
	StationNameKenton:                            StationCodeKenton,
	StationNameKentsBank:                         StationCodeKentsBank,
	StationNameKettering:                         StationCodeKettering,
	StationNameKewBridge:                         StationCodeKewBridge,
	StationNameKewGardens:                        StationCodeKewGardens,
	StationNameKeyham:                            StationCodeKeyham,
	StationNameKeynsham:                          StationCodeKeynsham,
	StationNameKidbrooke:                         StationCodeKidbrooke,
	StationNameKidderminster:                     StationCodeKidderminster,
	StationNameKidsgrove:                         StationCodeKidsgrove,
	StationNameKidwelly:                          StationCodeKidwelly,
	StationNameKilburnHighRoad:                   StationCodeKilburnHighRoad,
	StationNameKildale:                           StationCodeKildale,
	StationNameKildonan:                          StationCodeKildonan,
	StationNameKilgetty:                          StationCodeKilgetty,
	StationNameKilmarnock:                        StationCodeKilmarnock,
	StationNameKilmaurs:                          StationCodeKilmaurs,
	StationNameKilpatrick:                        StationCodeKilpatrick,
	StationNameKilwinning:                        StationCodeKilwinning,
	StationNameKinbrace:                          StationCodeKinbrace,
	StationNameKingham:                           StationCodeKingham,
	StationNameKinghorn:                          StationCodeKinghorn,
	StationNameKingsLangley:                      StationCodeKingsLangley,
	StationNameKingsLynn:                         StationCodeKingsLynn,
	StationNameKingsNorton:                       StationCodeKingsNorton,
	StationNameKingsNympton:                      StationCodeKingsNympton,
	StationNameKingsPark:                         StationCodeKingsPark,
	StationNameKingsSutton:                       StationCodeKingsSutton,
	StationNameKingsknowe:                        StationCodeKingsknowe,
	StationNameKingston:                          StationCodeKingston,
	StationNameKingswood:                         StationCodeKingswood,
	StationNameKingussie:                         StationCodeKingussie,
	StationNameKintbury:                          StationCodeKintbury,
	StationNameKirbyCross:                        StationCodeKirbyCross,
	StationNameKirkSandall:                       StationCodeKirkSandall,
	StationNameKirkbyMerseyside:                  StationCodeKirkbyMerseyside,
	StationNameKirkbyStephen:                     StationCodeKirkbyStephen,
	StationNameKirkbyinAshfield:                  StationCodeKirkbyinAshfield,
	StationNameKirkbyinFurness:                   StationCodeKirkbyinFurness,
	StationNameKirkcaldy:                         StationCodeKirkcaldy,
	StationNameKirkconnel:                        StationCodeKirkconnel,
	StationNameKirkdale:                          StationCodeKirkdale,
	StationNameKirkhamAndWesham:                  StationCodeKirkhamAndWesham,
	StationNameKirkhill:                          StationCodeKirkhill,
	StationNameKirknewton:                        StationCodeKirknewton,
	StationNameKirkstallForge:                    StationCodeKirkstallForge,
	StationNameKirkwood:                          StationCodeKirkwood,
	StationNameKirtonLindsey:                     StationCodeKirtonLindsey,
	StationNameKivetonBridge:                     StationCodeKivetonBridge,
	StationNameKivetonPark:                       StationCodeKivetonPark,
	StationNameKnaresborough:                     StationCodeKnaresborough,
	StationNameKnebworth:                         StationCodeKnebworth,
	StationNameKnighton:                          StationCodeKnighton,
	StationNameKnockholt:                         StationCodeKnockholt,
	StationNameKnottingley:                       StationCodeKnottingley,
	StationNameKnucklas:                          StationCodeKnucklas,
	StationNameKnutsford:                         StationCodeKnutsford,
	StationNameKyleofLochalsh:                    StationCodeKyleofLochalsh,
	StationNameLadybank:                          StationCodeLadybank,
	StationNameLadywell:                          StationCodeLadywell,
	StationNameLaindon:                           StationCodeLaindon,
	StationNameLairg:                             StationCodeLairg,
	StationNameLake:                              StationCodeLake,
	StationNameLakenheath:                        StationCodeLakenheath,
	StationNameLamphey:                           StationCodeLamphey,
	StationNameLanark:                            StationCodeLanark,
	StationNameLancaster:                         StationCodeLancaster,
	StationNameLancing:                           StationCodeLancing,
	StationNameLandywood:                         StationCodeLandywood,
	StationNameLangbank:                          StationCodeLangbank,
	StationNameLangho:                            StationCodeLangho,
	StationNameLangleyBerks:                      StationCodeLangleyBerks,
	StationNameLangleyGreen:                      StationCodeLangleyGreen,
	StationNameLangleyMill:                       StationCodeLangleyMill,
	StationNameLangside:                          StationCodeLangside,
	StationNameLangwathby:                        StationCodeLangwathby,
	StationNameLangwithWhaleyThorns:              StationCodeLangwithWhaleyThorns,
	StationNameLapford:                           StationCodeLapford,
	StationNameLapworth:                          StationCodeLapworth,
	StationNameLarbert:                           StationCodeLarbert,
	StationNameLargs:                             StationCodeLargs,
	StationNameLarkhall:                          StationCodeLarkhall,
	StationNameLaurencekirk:                      StationCodeLaurencekirk,
	StationNameLawrenceHill:                      StationCodeLawrenceHill,
	StationNameLaytonLancs:                       StationCodeLaytonLancs,
	StationNameLazonbyAndKirkoswald:              StationCodeLazonbyAndKirkoswald,
	StationNameLeaBridge:                         StationCodeLeaBridge,
	StationNameLeaGreen:                          StationCodeLeaGreen,
	StationNameLeaHall:                           StationCodeLeaHall,
	StationNameLeagrave:                          StationCodeLeagrave,
	StationNameLealholm:                          StationCodeLealholm,
	StationNameLeamingtonSpa:                     StationCodeLeamingtonSpa,
	StationNameLeasowe:                           StationCodeLeasowe,
	StationNameLeatherhead:                       StationCodeLeatherhead,
	StationNameLedbury:                           StationCodeLedbury,
	StationNameLeeLondon:                         StationCodeLeeLondon,
	StationNameLeeds:                             StationCodeLeeds,
	StationNameLeicester:                         StationCodeLeicester,
	StationNameLeighKent:                         StationCodeLeighKent,
	StationNameLeighonSea:                        StationCodeLeighonSea,
	StationNameLeightonBuzzard:                   StationCodeLeightonBuzzard,
	StationNameLelant:                            StationCodeLelant,
	StationNameLelantSaltings:                    StationCodeLelantSaltings,
	StationNameLenham:                            StationCodeLenham,
	StationNameLenzie:                            StationCodeLenzie,
	StationNameLeominster:                        StationCodeLeominster,
	StationNameLetchworthGardenCity:              StationCodeLetchworthGardenCity,
	StationNameLeucharsforStAndrews:              StationCodeLeucharsforStAndrews,
	StationNameLevenshulme:                       StationCodeLevenshulme,
	StationNameLewes:                             StationCodeLewes,
	StationNameLewisham:                          StationCodeLewisham,
	StationNameLeyland:                           StationCodeLeyland,
	StationNameLeytonMidlandRoad:                 StationCodeLeytonMidlandRoad,
	StationNameLeytonstoneHighRoad:               StationCodeLeytonstoneHighRoad,
	StationNameLichfieldCity:                     StationCodeLichfieldCity,
	StationNameLichfieldTrentValley:              StationCodeLichfieldTrentValley,
	StationNameLidlington:                        StationCodeLidlington,
	StationNameLimehouse:                         StationCodeLimehouse,
	StationNameLincolnCentral:                    StationCodeLincolnCentral,
	StationNameLingfield:                         StationCodeLingfield,
	StationNameLingwood:                          StationCodeLingwood,
	StationNameLinlithgow:                        StationCodeLinlithgow,
	StationNameLiphook:                           StationCodeLiphook,
	StationNameLiskeard:                          StationCodeLiskeard,
	StationNameLiss:                              StationCodeLiss,
	StationNameLisvaneAndThornhill:               StationCodeLisvaneAndThornhill,
	StationNameLittleKimble:                      StationCodeLittleKimble,
	StationNameLittleSutton:                      StationCodeLittleSutton,
	StationNameLittleborough:                     StationCodeLittleborough,
	StationNameLittlehampton:                     StationCodeLittlehampton,
	StationNameLittlehaven:                       StationCodeLittlehaven,
	StationNameLittleport:                        StationCodeLittleport,
	StationNameLiverpoolCentral:                  StationCodeLiverpoolCentral,
	StationNameLiverpoolJamesStreet:              StationCodeLiverpoolJamesStreet,
	StationNameLiverpoolLimeStreet:               StationCodeLiverpoolLimeStreet,
	StationNameLiverpoolSouthParkway:             StationCodeLiverpoolSouthParkway,
	StationNameLivingstonNorth:                   StationCodeLivingstonNorth,
	StationNameLivingstonSouth:                   StationCodeLivingstonSouth,
	StationNameLlanaber:                          StationCodeLlanaber,
	StationNameLlanbedr:                          StationCodeLlanbedr,
	StationNameLlanbisterRoad:                    StationCodeLlanbisterRoad,
	StationNameLlanbradach:                       StationCodeLlanbradach,
	StationNameLlandaf:                           StationCodeLlandaf,
	StationNameLlandanwg:                         StationCodeLlandanwg,
	StationNameLlandecwyn:                        StationCodeLlandecwyn,
	StationNameLlandeilo:                         StationCodeLlandeilo,
	StationNameLlandovery:                        StationCodeLlandovery,
	StationNameLlandrindod:                       StationCodeLlandrindod,
	StationNameLlandudno:                         StationCodeLlandudno,
	StationNameLlandudnoJunction:                 StationCodeLlandudnoJunction,
	StationNameLlandybie:                         StationCodeLlandybie,
	StationNameLlanelli:                          StationCodeLlanelli,
	StationNameLlanfairfechan:                    StationCodeLlanfairfechan,
	StationNameLlanfairpwll:                      StationCodeLlanfairpwll,
	StationNameLlangadog:                         StationCodeLlangadog,
	StationNameLlangammarch:                      StationCodeLlangammarch,
	StationNameLlangennech:                       StationCodeLlangennech,
	StationNameLlangynllo:                        StationCodeLlangynllo,
	StationNameLlanharan:                         StationCodeLlanharan,
	StationNameLlanhilleth:                       StationCodeLlanhilleth,
	StationNameLlanishen:                         StationCodeLlanishen,
	StationNameLlanrwst:                          StationCodeLlanrwst,
	StationNameLlansamlet:                        StationCodeLlansamlet,
	StationNameLlantwitMajor:                     StationCodeLlantwitMajor,
	StationNameLlanwrda:                          StationCodeLlanwrda,
	StationNameLlanwrtyd:                         StationCodeLlanwrtyd,
	StationNameLlwyngwril:                        StationCodeLlwyngwril,
	StationNameLlwynypia:                         StationCodeLlwynypia,
	StationNameLochAwe:                           StationCodeLochAwe,
	StationNameLochEilOutwardBound:               StationCodeLochEilOutwardBound,
	StationNameLochailort:                        StationCodeLochailort,
	StationNameLocheilside:                       StationCodeLocheilside,
	StationNameLochgelly:                         StationCodeLochgelly,
	StationNameLochluichart:                      StationCodeLochluichart,
	StationNameLochwinnoch:                       StationCodeLochwinnoch,
	StationNameLockerbie:                         StationCodeLockerbie,
	StationNameLockwood:                          StationCodeLockwood,
	StationNameLondonBlackfriars:                 StationCodeLondonBlackfriars,
	StationNameLondonBridge:                      StationCodeLondonBridge,
	StationNameLondonCannonStreet:                StationCodeLondonCannonStreet,
	StationNameLondonCharingCross:                StationCodeLondonCharingCross,
	StationNameLondonEuston:                      StationCodeLondonEuston,
	StationNameLondonFenchurchStreet:             StationCodeLondonFenchurchStreet,
	StationNameLondonFields:                      StationCodeLondonFields,
	StationNameLondonKingsCross:                  StationCodeLondonKingsCross,
	StationNameLondonLiverpoolStreet:             StationCodeLondonLiverpoolStreet,
	StationNameLondonMarylebone:                  StationCodeLondonMarylebone,
	StationNameLondonPaddington:                  StationCodeLondonPaddington,
	StationNameLondonRoadBrighton:                StationCodeLondonRoadBrighton,
	StationNameLondonRoadGuildford:               StationCodeLondonRoadGuildford,
	StationNameLondonStPancrasIntl:               StationCodeLondonStPancrasIntl,
	StationNameLondonVictoria:                    StationCodeLondonVictoria,
	StationNameLondonWaterloo:                    StationCodeLondonWaterloo,
	StationNameLondonWaterlooEast:                StationCodeLondonWaterlooEast,
	StationNameLongBuckby:                        StationCodeLongBuckby,
	StationNameLongEaton:                         StationCodeLongEaton,
	StationNameLongPreston:                       StationCodeLongPreston,
	StationNameLongbeck:                          StationCodeLongbeck,
	StationNameLongbridge:                        StationCodeLongbridge,
	StationNameLongcross:                         StationCodeLongcross,
	StationNameLongfield:                         StationCodeLongfield,
	StationNameLongniddry:                        StationCodeLongniddry,
	StationNameLongport:                          StationCodeLongport,
	StationNameLongton:                           StationCodeLongton,
	StationNameLooe:                              StationCodeLooe,
	StationNameLostock:                           StationCodeLostock,
	StationNameLostockGralam:                     StationCodeLostockGralam,
	StationNameLostockHall:                       StationCodeLostockHall,
	StationNameLostwithiel:                       StationCodeLostwithiel,
	StationNameLoughborough:                      StationCodeLoughborough,
	StationNameLoughboroughJunction:              StationCodeLoughboroughJunction,
	StationNameLowdham:                           StationCodeLowdham,
	StationNameLowerSydenham:                     StationCodeLowerSydenham,
	StationNameLowestoft:                         StationCodeLowestoft,
	StationNameLudlow:                            StationCodeLudlow,
	StationNameLuton:                             StationCodeLuton,
	StationNameLutonAirportParkway:               StationCodeLutonAirportParkway,
	StationNameLuxulyan:                          StationCodeLuxulyan,
	StationNameLydney:                            StationCodeLydney,
	StationNameLyeWestMidlands:                   StationCodeLyeWestMidlands,
	StationNameLymingtonPier:                     StationCodeLymingtonPier,
	StationNameLymingtonTown:                     StationCodeLymingtonTown,
	StationNameLympstoneCommando:                 StationCodeLympstoneCommando,
	StationNameLympstoneVillage:                  StationCodeLympstoneVillage,
	StationNameLytham:                            StationCodeLytham,
	StationNameMacclesfield:                      StationCodeMacclesfield,
	StationNameMachynlleth:                       StationCodeMachynlleth,
	StationNameMaesteg:                           StationCodeMaesteg,
	StationNameMaestegEwennyRoad:                 StationCodeMaestegEwennyRoad,
	StationNameMaghull:                           StationCodeMaghull,
	StationNameMaidenNewton:                      StationCodeMaidenNewton,
	StationNameMaidenhead:                        StationCodeMaidenhead,
	StationNameMaidstoneBarracks:                 StationCodeMaidstoneBarracks,
	StationNameMaidstoneEast:                     StationCodeMaidstoneEast,
	StationNameMaidstoneWest:                     StationCodeMaidstoneWest,
	StationNameMaldenManor:                       StationCodeMaldenManor,
	StationNameMallaig:                           StationCodeMallaig,
	StationNameMalton:                            StationCodeMalton,
	StationNameMalvernLink:                       StationCodeMalvernLink,
	StationNameManchesterAirport:                 StationCodeManchesterAirport,
	StationNameManchesterOxfordRoad:              StationCodeManchesterOxfordRoad,
	StationNameManchesterPiccadilly:              StationCodeManchesterPiccadilly,
	StationNameManchesterUnitedFootballGround:    StationCodeManchesterUnitedFootballGround,
	StationNameManchesterVictoria:                StationCodeManchesterVictoria,
	StationNameManea:                             StationCodeManea,
	StationNameManningtree:                       StationCodeManningtree,
	StationNameManorPark:                         StationCodeManorPark,
	StationNameManorRoad:                         StationCodeManorRoad,
	StationNameManorbier:                         StationCodeManorbier,
	StationNameManors:                            StationCodeManors,
	StationNameMansfield:                         StationCodeMansfield,
	StationNameMansfieldWoodhouse:                StationCodeMansfieldWoodhouse,
	StationNameMarch:                             StationCodeMarch,
	StationNameMardenKent:                        StationCodeMardenKent,
	StationNameMargate:                           StationCodeMargate,
	StationNameMarketHarborough:                  StationCodeMarketHarborough,
	StationNameMarketRasen:                       StationCodeMarketRasen,
	StationNameMarkinch:                          StationCodeMarkinch,
	StationNameMarksTey:                          StationCodeMarksTey,
	StationNameMarlow:                            StationCodeMarlow,
	StationNameMarple:                            StationCodeMarple,
	StationNameMarsdenYorks:                      StationCodeMarsdenYorks,
	StationNameMarske:                            StationCodeMarske,
	StationNameMarstonGreen:                      StationCodeMarstonGreen,
	StationNameMartinMill:                        StationCodeMartinMill,
	StationNameMartinsHeron:                      StationCodeMartinsHeron,
	StationNameMarton:                            StationCodeMarton,
	StationNameMaryhill:                          StationCodeMaryhill,
	StationNameMaryland:                          StationCodeMaryland,
	StationNameMaryport:                          StationCodeMaryport,
	StationNameMatlock:                           StationCodeMatlock,
	StationNameMatlockBath:                       StationCodeMatlockBath,
	StationNameMauldethRoad:                      StationCodeMauldethRoad,
	StationNameMaxwellPark:                       StationCodeMaxwellPark,
	StationNameMaybole:                           StationCodeMaybole,
	StationNameMazeHill:                          StationCodeMazeHill,
	StationNameMeadowhall:                        StationCodeMeadowhall,
	StationNameMeldreth:                          StationCodeMeldreth,
	StationNameMelksham:                          StationCodeMelksham,
	StationNameMeltonMowbray:                     StationCodeMeltonMowbray,
	StationNameMeltonSuffolk:                     StationCodeMeltonSuffolk,
	StationNameMenheniot:                         StationCodeMenheniot,
	StationNameMenston:                           StationCodeMenston,
	StationNameMeols:                             StationCodeMeols,
	StationNameMeolsCop:                          StationCodeMeolsCop,
	StationNameMeopham:                           StationCodeMeopham,
	StationNameMerryton:                          StationCodeMerryton,
	StationNameMerstham:                          StationCodeMerstham,
	StationNameMerthyrTydfil:                     StationCodeMerthyrTydfil,
	StationNameMerthyrVale:                       StationCodeMerthyrVale,
	StationNameMetheringham:                      StationCodeMetheringham,
	StationNameMetroCentre:                       StationCodeMetroCentre,
	StationNameMexborough:                        StationCodeMexborough,
	StationNameMicheldever:                       StationCodeMicheldever,
	StationNameMicklefield:                       StationCodeMicklefield,
	StationNameMiddlesbrough:                     StationCodeMiddlesbrough,
	StationNameMiddlewood:                        StationCodeMiddlewood,
	StationNameMidgham:                           StationCodeMidgham,
	StationNameMilfordHaven:                      StationCodeMilfordHaven,
	StationNameMilfordSurrey:                     StationCodeMilfordSurrey,
	StationNameMillHillBroadway:                  StationCodeMillHillBroadway,
	StationNameMillHillLancs:                     StationCodeMillHillLancs,
	StationNameMillbrookBeds:                     StationCodeMillbrookBeds,
	StationNameMillbrookHants:                    StationCodeMillbrookHants,
	StationNameMillikenPark:                      StationCodeMillikenPark,
	StationNameMillom:                            StationCodeMillom,
	StationNameMillsHillManchester:               StationCodeMillsHillManchester,
	StationNameMilngavie:                         StationCodeMilngavie,
	StationNameMiltonKeynesCentral:               StationCodeMiltonKeynesCentral,
	StationNameMinffordd:                         StationCodeMinffordd,
	StationNameMinster:                           StationCodeMinster,
	StationNameMirfield:                          StationCodeMirfield,
	StationNameMistley:                           StationCodeMistley,
	StationNameMitchamEastfields:                 StationCodeMitchamEastfields,
	StationNameMitchamJunction:                   StationCodeMitchamJunction,
	StationNameMobberley:                         StationCodeMobberley,
	StationNameMonifieth:                         StationCodeMonifieth,
	StationNameMonksRisborough:                   StationCodeMonksRisborough,
	StationNameMontpelier:                        StationCodeMontpelier,
	StationNameMontrose:                          StationCodeMontrose,
	StationNameMoorfields:                        StationCodeMoorfields,
	StationNameMoorgate:                          StationCodeMoorgate,
	StationNameMoorside:                          StationCodeMoorside,
	StationNameMoorthorpe:                        StationCodeMoorthorpe,
	StationNameMorar:                             StationCodeMorar,
	StationNameMorchardRoad:                      StationCodeMorchardRoad,
	StationNameMordenSouth:                       StationCodeMordenSouth,
	StationNameMorecambe:                         StationCodeMorecambe,
	StationNameMoretonDorset:                     StationCodeMoretonDorset,
	StationNameMoretonMerseyside:                 StationCodeMoretonMerseyside,
	StationNameMoretoninMarsh:                    StationCodeMoretoninMarsh,
	StationNameMorfaMawddach:                     StationCodeMorfaMawddach,
	StationNameMorley:                            StationCodeMorley,
	StationNameMorpeth:                           StationCodeMorpeth,
	StationNameMortimer:                          StationCodeMortimer,
	StationNameMortlake:                          StationCodeMortlake,
	StationNameMosesGate:                         StationCodeMosesGate,
	StationNameMossSide:                          StationCodeMossSide,
	StationNameMossleyHill:                       StationCodeMossleyHill,
	StationNameMossleyManchester:                 StationCodeMossleyManchester,
	StationNameMosspark:                          StationCodeMosspark,
	StationNameMoston:                            StationCodeMoston,
	StationNameMotherwell:                        StationCodeMotherwell,
	StationNameMotspurPark:                       StationCodeMotspurPark,
	StationNameMottingham:                        StationCodeMottingham,
	StationNameMottisfontAndDunbridge:            StationCodeMottisfontAndDunbridge,
	StationNameMouldsworth:                       StationCodeMouldsworth,
	StationNameMoulsecoomb:                       StationCodeMoulsecoomb,
	StationNameMountFlorida:                      StationCodeMountFlorida,
	StationNameMountVernon:                       StationCodeMountVernon,
	StationNameMountainAsh:                       StationCodeMountainAsh,
	StationNameMuirend:                           StationCodeMuirend,
	StationNameMuirofOrd:                         StationCodeMuirofOrd,
	StationNameMusselburgh:                       StationCodeMusselburgh,
	StationNameMytholmroyd:                       StationCodeMytholmroyd,
	StationNameNafferton:                         StationCodeNafferton,
	StationNameNailseaAndBackwell:                StationCodeNailseaAndBackwell,
	StationNameNairn:                             StationCodeNairn,
	StationNameNantwich:                          StationCodeNantwich,
	StationNameNarberth:                          StationCodeNarberth,
	StationNameNarborough:                        StationCodeNarborough,
	StationNameNavigationRoad:                    StationCodeNavigationRoad,
	StationNameNeath:                             StationCodeNeath,
	StationNameNeedhamMarket:                     StationCodeNeedhamMarket,
	StationNameNeilston:                          StationCodeNeilston,
	StationNameNelson:                            StationCodeNelson,
	StationNameNeston:                            StationCodeNeston,
	StationNameNetherfield:                       StationCodeNetherfield,
	StationNameNethertown:                        StationCodeNethertown,
	StationNameNetley:                            StationCodeNetley,
	StationNameNewBarnet:                         StationCodeNewBarnet,
	StationNameNewBeckenham:                      StationCodeNewBeckenham,
	StationNameNewBrighton:                       StationCodeNewBrighton,
	StationNameNewClee:                           StationCodeNewClee,
	StationNameNewCross:                          StationCodeNewCross,
	StationNameNewCrossGate:                      StationCodeNewCrossGate,
	StationNameNewCumnock:                        StationCodeNewCumnock,
	StationNameNewEltham:                         StationCodeNewEltham,
	StationNameNewHolland:                        StationCodeNewHolland,
	StationNameNewHythe:                          StationCodeNewHythe,
	StationNameNewLane:                           StationCodeNewLane,
	StationNameNewMalden:                         StationCodeNewMalden,
	StationNameNewMillsCentral:                   StationCodeNewMillsCentral,
	StationNameNewMillsNewtown:                   StationCodeNewMillsNewtown,
	StationNameNewMilton:                         StationCodeNewMilton,
	StationNameNewPudsey:                         StationCodeNewPudsey,
	StationNameNewSouthgate:                      StationCodeNewSouthgate,
	StationNameNewarkCastle:                      StationCodeNewarkCastle,
	StationNameNewarkNorthGate:                   StationCodeNewarkNorthGate,
	StationNameNewbridge:                         StationCodeNewbridge,
	StationNameNewbury:                           StationCodeNewbury,
	StationNameNewburyRacecourse:                 StationCodeNewburyRacecourse,
	StationNameNewcastle:                         StationCodeNewcastle,
	StationNameNewcourt:                          StationCodeNewcourt,
	StationNameNewcraighall:                      StationCodeNewcraighall,
	StationNameNewhavenHarbour:                   StationCodeNewhavenHarbour,
	StationNameNewhavenTown:                      StationCodeNewhavenTown,
	StationNameNewington:                         StationCodeNewington,
	StationNameNewmarket:                         StationCodeNewmarket,
	StationNameNewportEssex:                      StationCodeNewportEssex,
	StationNameNewportSouthWales:                 StationCodeNewportSouthWales,
	StationNameNewquay:                           StationCodeNewquay,
	StationNameNewstead:                          StationCodeNewstead,
	StationNameNewtonAbbot:                       StationCodeNewtonAbbot,
	StationNameNewtonAycliffe:                    StationCodeNewtonAycliffe,
	StationNameNewtonLanark:                      StationCodeNewtonLanark,
	StationNameNewtonStCyres:                     StationCodeNewtonStCyres,
	StationNameNewtonforHyde:                     StationCodeNewtonforHyde,
	StationNameNewtongrange:                      StationCodeNewtongrange,
	StationNameNewtonleWillows:                   StationCodeNewtonleWillows,
	StationNameNewtonmore:                        StationCodeNewtonmore,
	StationNameNewtononAyr:                       StationCodeNewtononAyr,
	StationNameNewtownPowys:                      StationCodeNewtownPowys,
	StationNameNinianPark:                        StationCodeNinianPark,
	StationNameNitshill:                          StationCodeNitshill,
	StationNameNorbiton:                          StationCodeNorbiton,
	StationNameNorbury:                           StationCodeNorbury,
	StationNameNormansBay:                        StationCodeNormansBay,
	StationNameNormanton:                         StationCodeNormanton,
	StationNameNorthBerwick:                      StationCodeNorthBerwick,
	StationNameNorthCamp:                         StationCodeNorthCamp,
	StationNameNorthDulwich:                      StationCodeNorthDulwich,
	StationNameNorthFambridge:                    StationCodeNorthFambridge,
	StationNameNorthLlanrwst:                     StationCodeNorthLlanrwst,
	StationNameNorthQueensferry:                  StationCodeNorthQueensferry,
	StationNameNorthRoadDarlington:               StationCodeNorthRoadDarlington,
	StationNameNorthSheen:                        StationCodeNorthSheen,
	StationNameNorthWalsham:                      StationCodeNorthWalsham,
	StationNameNorthWembley:                      StationCodeNorthWembley,
	StationNameNorthallerton:                     StationCodeNorthallerton,
	StationNameNorthampton:                       StationCodeNorthampton,
	StationNameNorthfield:                        StationCodeNorthfield,
	StationNameNorthfleet:                        StationCodeNorthfleet,
	StationNameNortholtPark:                      StationCodeNortholtPark,
	StationNameNorthumberlandPark:                StationCodeNorthumberlandPark,
	StationNameNorthwich:                         StationCodeNorthwich,
	StationNameNortonBridge:                      StationCodeNortonBridge,
	StationNameNorwich:                           StationCodeNorwich,
	StationNameNorwoodJunction:                   StationCodeNorwoodJunction,
	StationNameNottingham:                        StationCodeNottingham,
	StationNameNuneaton:                          StationCodeNuneaton,
	StationNameNunhead:                           StationCodeNunhead,
	StationNameNunthorpe:                         StationCodeNunthorpe,
	StationNameNutbourne:                         StationCodeNutbourne,
	StationNameNutfield:                          StationCodeNutfield,
	StationNameOakengates:                        StationCodeOakengates,
	StationNameOakham:                            StationCodeOakham,
	StationNameOakleighPark:                      StationCodeOakleighPark,
	StationNameOban:                              StationCodeOban,
	StationNameOckendon:                          StationCodeOckendon,
	StationNameOckley:                            StationCodeOckley,
	StationNameOkehampton:                        StationCodeOkehampton,
	StationNameOldHill:                           StationCodeOldHill,
	StationNameOldRoan:                           StationCodeOldRoan,
	StationNameOldStreet:                         StationCodeOldStreet,
	StationNameOldfieldPark:                      StationCodeOldfieldPark,
	StationNameOlton:                             StationCodeOlton,
	StationNameOre:                               StationCodeOre,
	StationNameOrmskirk:                          StationCodeOrmskirk,
	StationNameOrpington:                         StationCodeOrpington,
	StationNameOrrell:                            StationCodeOrrell,
	StationNameOrrellPark:                        StationCodeOrrellPark,
	StationNameOtford:                            StationCodeOtford,
	StationNameOultonBroadNorth:                  StationCodeOultonBroadNorth,
	StationNameOultonBroadSouth:                  StationCodeOultonBroadSouth,
	StationNameOutwood:                           StationCodeOutwood,
	StationNameOverpool:                          StationCodeOverpool,
	StationNameOverton:                           StationCodeOverton,
	StationNameOxenholmeLakeDistrict:             StationCodeOxenholmeLakeDistrict,
	StationNameOxford:                            StationCodeOxford,
	StationNameOxfordParkway:                     StationCodeOxfordParkway,
	StationNameOxshott:                           StationCodeOxshott,
	StationNameOxted:                             StationCodeOxted,
	StationNamePaddockWood:                       StationCodePaddockWood,
	StationNamePadgate:                           StationCodePadgate,
	StationNamePaignton:                          StationCodePaignton,
	StationNamePaisleyCanal:                      StationCodePaisleyCanal,
	StationNamePaisleyGilmourStreet:              StationCodePaisleyGilmourStreet,
	StationNamePaisleyStJames:                    StationCodePaisleyStJames,
	StationNamePalmersGreen:                      StationCodePalmersGreen,
	StationNamePangbourne:                        StationCodePangbourne,
	StationNamePannal:                            StationCodePannal,
	StationNamePantyffynnon:                      StationCodePantyffynnon,
	StationNamePar:                               StationCodePar,
	StationNameParbold:                           StationCodeParbold,
	StationNameParkStreet:                        StationCodeParkStreet,
	StationNameParkstoneDorset:                   StationCodeParkstoneDorset,
	StationNameParsonStreet:                      StationCodeParsonStreet,
	StationNamePartick:                           StationCodePartick,
	StationNameParton:                            StationCodeParton,
	StationNamePatchway:                          StationCodePatchway,
	StationNamePatricroft:                        StationCodePatricroft,
	StationNamePatterton:                         StationCodePatterton,
	StationNamePeartree:                          StationCodePeartree,
	StationNamePeckhamRye:                        StationCodePeckhamRye,
	StationNamePegswood:                          StationCodePegswood,
	StationNamePemberton:                         StationCodePemberton,
	StationNamePembreyAndBurryPort:               StationCodePembreyAndBurryPort,
	StationNamePembroke:                          StationCodePembroke,
	StationNamePembrokeDock:                      StationCodePembrokeDock,
	StationNamePenally:                           StationCodePenally,
	StationNamePenarth:                           StationCodePenarth,
	StationNamePencoed:                           StationCodePencoed,
	StationNamePengam:                            StationCodePengam,
	StationNamePengeEast:                         StationCodePengeEast,
	StationNamePengeWest:                         StationCodePengeWest,
	StationNamePenhelig:                          StationCodePenhelig,
	StationNamePenistone:                         StationCodePenistone,
	StationNamePenkridge:                         StationCodePenkridge,
	StationNamePenmaenmawr:                       StationCodePenmaenmawr,
	StationNamePenmere:                           StationCodePenmere,
	StationNamePenrhiwceiber:                     StationCodePenrhiwceiber,
	StationNamePenrhyndeudraeth:                  StationCodePenrhyndeudraeth,
	StationNamePenrithNorthLakes:                 StationCodePenrithNorthLakes,
	StationNamePenrynCornwall:                    StationCodePenrynCornwall,
	StationNamePensarnGwynedd:                    StationCodePensarnGwynedd,
	StationNamePenshurst:                         StationCodePenshurst,
	StationNamePentreBach:                        StationCodePentreBach,
	StationNamePenyBont:                          StationCodePenyBont,
	StationNamePenychain:                         StationCodePenychain,
	StationNamePenyffordd:                        StationCodePenyffordd,
	StationNamePenzance:                          StationCodePenzance,
	StationNamePerranwell:                        StationCodePerranwell,
	StationNamePerryBarr:                         StationCodePerryBarr,
	StationNamePershore:                          StationCodePershore,
	StationNamePerth:                             StationCodePerth,
	StationNamePeterborough:                      StationCodePeterborough,
	StationNamePetersfield:                       StationCodePetersfield,
	StationNamePettsWood:                         StationCodePettsWood,
	StationNamePevenseyAndWestham:                StationCodePevenseyAndWestham,
	StationNamePevenseyBay:                       StationCodePevenseyBay,
	StationNamePewsey:                            StationCodePewsey,
	StationNamePilning:                           StationCodePilning,
	StationNamePinhoe:                            StationCodePinhoe,
	StationNamePitlochry:                         StationCodePitlochry,
	StationNamePitsea:                            StationCodePitsea,
	StationNamePleasington:                       StationCodePleasington,
	StationNamePlockton:                          StationCodePlockton,
	StationNamePluckley:                          StationCodePluckley,
	StationNamePlumley:                           StationCodePlumley,
	StationNamePlumpton:                          StationCodePlumpton,
	StationNamePlumstead:                         StationCodePlumstead,
	StationNamePlymouth:                          StationCodePlymouth,
	StationNamePokesdown:                         StationCodePokesdown,
	StationNamePolegate:                          StationCodePolegate,
	StationNamePolesworth:                        StationCodePolesworth,
	StationNamePollokshawsEast:                   StationCodePollokshawsEast,
	StationNamePollokshawsWest:                   StationCodePollokshawsWest,
	StationNamePollokshieldsEast:                 StationCodePollokshieldsEast,
	StationNamePollokshieldsWest:                 StationCodePollokshieldsWest,
	StationNamePolmont:                           StationCodePolmont,
	StationNamePolsloeBridge:                     StationCodePolsloeBridge,
	StationNamePondersEnd:                        StationCodePondersEnd,
	StationNamePontarddulais:                     StationCodePontarddulais,
	StationNamePontefractBaghill:                 StationCodePontefractBaghill,
	StationNamePontefractMonkhill:                StationCodePontefractMonkhill,
	StationNamePontefractTanshelf:                StationCodePontefractTanshelf,
	StationNamePontlottyn:                        StationCodePontlottyn,
	StationNamePontyPant:                         StationCodePontyPant,
	StationNamePontyclun:                         StationCodePontyclun,
	StationNamePontypoolAndNewInn:                StationCodePontypoolAndNewInn,
	StationNamePontypridd:                        StationCodePontypridd,
	StationNamePoole:                             StationCodePoole,
	StationNamePoppleton:                         StationCodePoppleton,
	StationNamePortGlasgow:                       StationCodePortGlasgow,
	StationNamePortSunlight:                      StationCodePortSunlight,
	StationNamePortTalbotParkway:                 StationCodePortTalbotParkway,
	StationNamePortchester:                       StationCodePortchester,
	StationNamePorth:                             StationCodePorth,
	StationNamePorthmadog:                        StationCodePorthmadog,
	StationNamePortlethen:                        StationCodePortlethen,
	StationNamePortslade:                         StationCodePortslade,
	StationNamePortsmouthAndSouthsea:             StationCodePortsmouthAndSouthsea,
	StationNamePortsmouthArms:                    StationCodePortsmouthArms,
	StationNamePortsmouthHarbour:                 StationCodePortsmouthHarbour,
	StationNamePossilparkAndParkhouse:            StationCodePossilparkAndParkhouse,
	StationNamePottersBar:                        StationCodePottersBar,
	StationNamePoultonleFylde:                    StationCodePoultonleFylde,
	StationNamePoynton:                           StationCodePoynton,
	StationNamePrees:                             StationCodePrees,
	StationNamePrescot:                           StationCodePrescot,
	StationNamePrestatyn:                         StationCodePrestatyn,
	StationNamePrestbury:                         StationCodePrestbury,
	StationNamePrestonLancs:                      StationCodePrestonLancs,
	StationNamePrestonPark:                       StationCodePrestonPark,
	StationNamePrestonpans:                       StationCodePrestonpans,
	StationNamePrestwickInternationalAirport:     StationCodePrestwickInternationalAirport,
	StationNamePrestwickTown:                     StationCodePrestwickTown,
	StationNamePriesthillAndDarnley:              StationCodePriesthillAndDarnley,
	StationNamePrincesRisborough:                 StationCodePrincesRisborough,
	StationNamePrittlewell:                       StationCodePrittlewell,
	StationNamePrudhoe:                           StationCodePrudhoe,
	StationNamePulborough:                        StationCodePulborough,
	StationNamePurfleet:                          StationCodePurfleet,
	StationNamePurley:                            StationCodePurley,
	StationNamePurleyOaks:                        StationCodePurleyOaks,
	StationNamePutney:                            StationCodePutney,
	StationNamePwllheli:                          StationCodePwllheli,
	StationNamePyeCorner:                         StationCodePyeCorner,
	StationNamePyle:                              StationCodePyle,
	StationNameQuakersYard:                       StationCodeQuakersYard,
	StationNameQueenborough:                      StationCodeQueenborough,
	StationNameQueensParkGlasgow:                 StationCodeQueensParkGlasgow,
	StationNameQueensParkLondon:                  StationCodeQueensParkLondon,
	StationNameQueensRoadPeckham:                 StationCodeQueensRoadPeckham,
	StationNameQueenstownRoadBattersea:           StationCodeQueenstownRoadBattersea,
	StationNameQuintrellDowns:                    StationCodeQuintrellDowns,
	StationNameRadcliffeonTrent:                  StationCodeRadcliffeonTrent,
	StationNameRadlett:                           StationCodeRadlett,
	StationNameRadley:                            StationCodeRadley,
	StationNameRadyr:                             StationCodeRadyr,
	StationNameRainford:                          StationCodeRainford,
	StationNameRainhamEssex:                      StationCodeRainhamEssex,
	StationNameRainhamKent:                       StationCodeRainhamKent,
	StationNameRainhill:                          StationCodeRainhill,
	StationNameRamsgate:                          StationCodeRamsgate,
	StationNameRamsgreaveAndWilpshire:            StationCodeRamsgreaveAndWilpshire,
	StationNameRannoch:                           StationCodeRannoch,
	StationNameRauceby:                           StationCodeRauceby,
	StationNameRavenglassforEskdale:              StationCodeRavenglassforEskdale,
	StationNameRavensbourne:                      StationCodeRavensbourne,
	StationNameRavensthorpe:                      StationCodeRavensthorpe,
	StationNameRawcliffe:                         StationCodeRawcliffe,
	StationNameRayleigh:                          StationCodeRayleigh,
	StationNameRaynesPark:                        StationCodeRaynesPark,
	StationNameReading:                           StationCodeReading,
	StationNameReadingWest:                       StationCodeReadingWest,
	StationNameRectoryRoad:                       StationCodeRectoryRoad,
	StationNameRedbridge:                         StationCodeRedbridge,
	StationNameRedcarBritishSteel:                StationCodeRedcarBritishSteel,
	StationNameRedcarCentral:                     StationCodeRedcarCentral,
	StationNameRedcarEast:                        StationCodeRedcarEast,
	StationNameReddishNorth:                      StationCodeReddishNorth,
	StationNameReddishSouth:                      StationCodeReddishSouth,
	StationNameRedditch:                          StationCodeRedditch,
	StationNameRedhill:                           StationCodeRedhill,
	StationNameRedland:                           StationCodeRedland,
	StationNameRedruth:                           StationCodeRedruth,
	StationNameReedhamNorfolk:                    StationCodeReedhamNorfolk,
	StationNameReedhamSurrey:                     StationCodeReedhamSurrey,
	StationNameReigate:                           StationCodeReigate,
	StationNameRenton:                            StationCodeRenton,
	StationNameRetford:                           StationCodeRetford,
	StationNameRhiwbina:                          StationCodeRhiwbina,
	StationNameRhooseCardiffInternationalAirport: StationCodeRhooseCardiffInternationalAirport,
	StationNameRhosneigr:                         StationCodeRhosneigr,
	StationNameRhyl:                              StationCodeRhyl,
	StationNameRhymney:                           StationCodeRhymney,
	StationNameRibblehead:                        StationCodeRibblehead,
	StationNameRiceLane:                          StationCodeRiceLane,
	StationNameRichmondLondon:                    StationCodeRichmondLondon,
	StationNameRickmansworth:                     StationCodeRickmansworth,
	StationNameRiddlesdown:                       StationCodeRiddlesdown,
	StationNameRidgmont:                          StationCodeRidgmont,
	StationNameRidingMill:                        StationCodeRidingMill,
	StationNameRiscaAndPontymister:               StationCodeRiscaAndPontymister,
	StationNameRishton:                           StationCodeRishton,
	StationNameRobertsbridge:                     StationCodeRobertsbridge,
	StationNameRoby:                              StationCodeRoby,
	StationNameRochdale:                          StationCodeRochdale,
	StationNameRoche:                             StationCodeRoche,
	StationNameRochester:                         StationCodeRochester,
	StationNameRochford:                          StationCodeRochford,
	StationNameRockFerry:                         StationCodeRockFerry,
	StationNameRogart:                            StationCodeRogart,
	StationNameRogerstone:                        StationCodeRogerstone,
	StationNameRolleston:                         StationCodeRolleston,
	StationNameRomanBridge:                       StationCodeRomanBridge,
	StationNameRomford:                           StationCodeRomford,
	StationNameRomiley:                           StationCodeRomiley,
	StationNameRomsey:                            StationCodeRomsey,
	StationNameRoose:                             StationCodeRoose,
	StationNameRoseGrove:                         StationCodeRoseGrove,
	StationNameRoseHillMarple:                    StationCodeRoseHillMarple,
	StationNameRosyth:                            StationCodeRosyth,
	StationNameRotherhamCentral:                  StationCodeRotherhamCentral,
	StationNameRotherhithe:                       StationCodeRotherhithe,
	StationNameRoughtonRoad:                      StationCodeRoughtonRoad,
	StationNameRowlandsCastle:                    StationCodeRowlandsCastle,
	StationNameRowleyRegis:                       StationCodeRowleyRegis,
	StationNameRoyBridge:                         StationCodeRoyBridge,
	StationNameRoydon:                            StationCodeRoydon,
	StationNameRoyston:                           StationCodeRoyston,
	StationNameRuabon:                            StationCodeRuabon,
	StationNameRufford:                           StationCodeRufford,
	StationNameRugby:                             StationCodeRugby,
	StationNameRugeleyTown:                       StationCodeRugeleyTown,
	StationNameRugeleyTrentValley:                StationCodeRugeleyTrentValley,
	StationNameRuncorn:                           StationCodeRuncorn,
	StationNameRuncornEast:                       StationCodeRuncornEast,
	StationNameRuskington:                        StationCodeRuskington,
	StationNameRuswarp:                           StationCodeRuswarp,
	StationNameRutherglen:                        StationCodeRutherglen,
	StationNameRydeEsplanade:                     StationCodeRydeEsplanade,
	StationNameRydePierHead:                      StationCodeRydePierHead,
	StationNameRydeStJohnsRoad:                   StationCodeRydeStJohnsRoad,
	StationNameRyderBrow:                         StationCodeRyderBrow,
	StationNameRyeHouse:                          StationCodeRyeHouse,
	StationNameRyeSussex:                         StationCodeRyeSussex,
	StationNameSalfordCentral:                    StationCodeSalfordCentral,
	StationNameSalfordCrescent:                   StationCodeSalfordCrescent,
	StationNameSalfordsSurrey:                    StationCodeSalfordsSurrey,
	StationNameSalhouse:                          StationCodeSalhouse,
	StationNameSalisbury:                         StationCodeSalisbury,
	StationNameSaltaire:                          StationCodeSaltaire,
	StationNameSaltash:                           StationCodeSaltash,
	StationNameSaltburn:                          StationCodeSaltburn,
	StationNameSaltcoats:                         StationCodeSaltcoats,
	StationNameSaltmarshe:                        StationCodeSaltmarshe,
	StationNameSalwick:                           StationCodeSalwick,
	StationNameSampfordCourtenay:                 StationCodeSampfordCourtenay,
	StationNameSandalAndAgbrigg:                  StationCodeSandalAndAgbrigg,
	StationNameSandbach:                          StationCodeSandbach,
	StationNameSanderstead:                       StationCodeSanderstead,
	StationNameSandhills:                         StationCodeSandhills,
	StationNameSandhurstBerks:                    StationCodeSandhurstBerks,
	StationNameSandling:                          StationCodeSandling,
	StationNameSandown:                           StationCodeSandown,
	StationNameSandplace:                         StationCodeSandplace,
	StationNameSandwellAndDudley:                 StationCodeSandwellAndDudley,
	StationNameSandwich:                          StationCodeSandwich,
	StationNameSandy:                             StationCodeSandy,
	StationNameSankeyforPenketh:                  StationCodeSankeyforPenketh,
	StationNameSanquhar:                          StationCodeSanquhar,
	StationNameSarn:                              StationCodeSarn,
	StationNameSaundersfoot:                      StationCodeSaundersfoot,
	StationNameSaunderton:                        StationCodeSaunderton,
	StationNameSawbridgeworth:                    StationCodeSawbridgeworth,
	StationNameSaxilby:                           StationCodeSaxilby,
	StationNameSaxmundham:                        StationCodeSaxmundham,
	StationNameScarborough:                       StationCodeScarborough,
	StationNameScotscalder:                       StationCodeScotscalder,
	StationNameScotstounhill:                     StationCodeScotstounhill,
	StationNameScunthorpe:                        StationCodeScunthorpe,
	StationNameSeaMills:                          StationCodeSeaMills,
	StationNameSeafordSussex:                     StationCodeSeafordSussex,
	StationNameSeaforthAndLitherland:             StationCodeSeaforthAndLitherland,
	StationNameSeaham:                            StationCodeSeaham,
	StationNameSeamer:                            StationCodeSeamer,
	StationNameSeascale:                          StationCodeSeascale,
	StationNameSeatonCarew:                       StationCodeSeatonCarew,
	StationNameSeerGreenAndJordans:               StationCodeSeerGreenAndJordans,
	StationNameSelby:                             StationCodeSelby,
	StationNameSelhurst:                          StationCodeSelhurst,
	StationNameSellafield:                        StationCodeSellafield,
	StationNameSelling:                           StationCodeSelling,
	StationNameSellyOak:                          StationCodeSellyOak,
	StationNameSettle:                            StationCodeSettle,
	StationNameSevenKings:                        StationCodeSevenKings,
	StationNameSevenSisters:                      StationCodeSevenSisters,
	StationNameSevenoaks:                         StationCodeSevenoaks,
	StationNameSevernBeach:                       StationCodeSevernBeach,
	StationNameSevernTunnelJunction:              StationCodeSevernTunnelJunction,
	StationNameShadwell:                          StationCodeShadwell,
	StationNameShalfordSurrey:                    StationCodeShalfordSurrey,
	StationNameShanklin:                          StationCodeShanklin,
	StationNameShawfair:                          StationCodeShawfair,
	StationNameShawford:                          StationCodeShawford,
	StationNameShawlands:                         StationCodeShawlands,
	StationNameSheernessonSea:                    StationCodeSheernessonSea,
	StationNameSheffield:                         StationCodeSheffield,
	StationNameShelfordCambs:                     StationCodeShelfordCambs,
	StationNameShenfield:                         StationCodeShenfield,
	StationNameShenstone:                         StationCodeShenstone,
	StationNameShepherdsBush:                     StationCodeShepherdsBush,
	StationNameShepherdsWell:                     StationCodeShepherdsWell,
	StationNameShepley:                           StationCodeShepley,
	StationNameShepperton:                        StationCodeShepperton,
	StationNameShepreth:                          StationCodeShepreth,
	StationNameSherborne:                         StationCodeSherborne,
	StationNameSherburninElmet:                   StationCodeSherburninElmet,
	StationNameSheringham:                        StationCodeSheringham,
	StationNameShettleston:                       StationCodeShettleston,
	StationNameShieldmuir:                        StationCodeShieldmuir,
	StationNameShifnal:                           StationCodeShifnal,
	StationNameShildon:                           StationCodeShildon,
	StationNameShiplake:                          StationCodeShiplake,
	StationNameShipleyYorks:                      StationCodeShipleyYorks,
	StationNameShippeaHill:                       StationCodeShippeaHill,
	StationNameShipton:                           StationCodeShipton,
	StationNameShirebrook:                        StationCodeShirebrook,
	StationNameShirehampton:                      StationCodeShirehampton,
	StationNameShireoaks:                         StationCodeShireoaks,
	StationNameShirley:                           StationCodeShirley,
	StationNameShoeburyness:                      StationCodeShoeburyness,
	StationNameSholing:                           StationCodeSholing,
	StationNameShoreditchHighStreet:              StationCodeShoreditchHighStreet,
	StationNameShorehamKent:                      StationCodeShorehamKent,
	StationNameShorehambySea:                     StationCodeShorehambySea,
	StationNameShortlands:                        StationCodeShortlands,
	StationNameShotton:                           StationCodeShotton,
	StationNameShotts:                            StationCodeShotts,
	StationNameShrewsbury:                        StationCodeShrewsbury,
	StationNameSidcup:                            StationCodeSidcup,
	StationNameSileby:                            StationCodeSileby,
	StationNameSilecroft:                         StationCodeSilecroft,
	StationNameSilkstoneCommon:                   StationCodeSilkstoneCommon,
	StationNameSilverStreet:                      StationCodeSilverStreet,
	StationNameSilverdale:                        StationCodeSilverdale,
	StationNameSinger:                            StationCodeSinger,
	StationNameSittingbourne:                     StationCodeSittingbourne,
	StationNameSkegness:                          StationCodeSkegness,
	StationNameSkewen:                            StationCodeSkewen,
	StationNameSkipton:                           StationCodeSkipton,
	StationNameSladeGreen:                        StationCodeSladeGreen,
	StationNameSlaithwaite:                       StationCodeSlaithwaite,
	StationNameSlateford:                         StationCodeSlateford,
	StationNameSleaford:                          StationCodeSleaford,
	StationNameSleights:                          StationCodeSleights,
	StationNameSlough:                            StationCodeSlough,
	StationNameSmallHeath:                        StationCodeSmallHeath,
	StationNameSmallbrookJunction:                StationCodeSmallbrookJunction,
	StationNameSmethwickGaltonBridge:             StationCodeSmethwickGaltonBridge,
	StationNameSmethwickRolfeStreet:              StationCodeSmethwickRolfeStreet,
	StationNameSmithyBridge:                      StationCodeSmithyBridge,
	StationNameSnaith:                            StationCodeSnaith,
	StationNameSnodland:                          StationCodeSnodland,
	StationNameSnowdown:                          StationCodeSnowdown,
	StationNameSoleStreet:                        StationCodeSoleStreet,
	StationNameSolihull:                          StationCodeSolihull,
	StationNameSomerleyton:                       StationCodeSomerleyton,
	StationNameSouthActon:                        StationCodeSouthActon,
	StationNameSouthBank:                         StationCodeSouthBank,
	StationNameSouthBermondsey:                   StationCodeSouthBermondsey,
	StationNameSouthCroydon:                      StationCodeSouthCroydon,
	StationNameSouthElmsall:                      StationCodeSouthElmsall,
	StationNameSouthGreenford:                    StationCodeSouthGreenford,
	StationNameSouthGyle:                         StationCodeSouthGyle,
	StationNameSouthHampstead:                    StationCodeSouthHampstead,
	StationNameSouthKenton:                       StationCodeSouthKenton,
	StationNameSouthMerton:                       StationCodeSouthMerton,
	StationNameSouthMilford:                      StationCodeSouthMilford,
	StationNameSouthRuislip:                      StationCodeSouthRuislip,
	StationNameSouthTottenham:                    StationCodeSouthTottenham,
	StationNameSouthWigston:                      StationCodeSouthWigston,
	StationNameSouthWoodhamFerrers:               StationCodeSouthWoodhamFerrers,
	StationNameSouthall:                          StationCodeSouthall,
	StationNameSouthamptonAirportParkway:         StationCodeSouthamptonAirportParkway,
	StationNameSouthamptonCentral:                StationCodeSouthamptonCentral,
	StationNameSouthbourne:                       StationCodeSouthbourne,
	StationNameSouthbury:                         StationCodeSouthbury,
	StationNameSouthease:                         StationCodeSouthease,
	StationNameSouthendAirport:                   StationCodeSouthendAirport,
	StationNameSouthendCentral:                   StationCodeSouthendCentral,
	StationNameSouthendEast:                      StationCodeSouthendEast,
	StationNameSouthendVictoria:                  StationCodeSouthendVictoria,
	StationNameSouthminster:                      StationCodeSouthminster,
	StationNameSouthport:                         StationCodeSouthport,
	StationNameSouthwick:                         StationCodeSouthwick,
	StationNameSowerbyBridge:                     StationCodeSowerbyBridge,
	StationNameSpalding:                          StationCodeSpalding,
	StationNameSpeanBridge:                       StationCodeSpeanBridge,
	StationNameSpital:                            StationCodeSpital,
	StationNameSpondon:                           StationCodeSpondon,
	StationNameSpoonerRow:                        StationCodeSpoonerRow,
	StationNameSpringRoad:                        StationCodeSpringRoad,
	StationNameSpringburn:                        StationCodeSpringburn,
	StationNameSpringfield:                       StationCodeSpringfield,
	StationNameSquiresGate:                       StationCodeSquiresGate,
	StationNameStAlbansAbbey:                     StationCodeStAlbansAbbey,
	StationNameStAlbansCity:                      StationCodeStAlbansCity,
	StationNameStAndrewsRoad:                     StationCodeStAndrewsRoad,
	StationNameStAnnesonSea:                      StationCodeStAnnesonSea,
	StationNameStAustell:                         StationCodeStAustell,
	StationNameStBees:                            StationCodeStBees,
	StationNameStBudeauxFerryRoad:                StationCodeStBudeauxFerryRoad,
	StationNameStBudeauxVictoriaRoad:             StationCodeStBudeauxVictoriaRoad,
	StationNameStColumbRoad:                      StationCodeStColumbRoad,
	StationNameStDenys:                           StationCodeStDenys,
	StationNameStErth:                            StationCodeStErth,
	StationNameStGermans:                         StationCodeStGermans,
	StationNameStHelensCentral:                   StationCodeStHelensCentral,
	StationNameStHelensJunction:                  StationCodeStHelensJunction,
	StationNameStHelierSurrey:                    StationCodeStHelierSurrey,
	StationNameStIvesCornwall:                    StationCodeStIvesCornwall,
	StationNameStJamesParkExeter:                 StationCodeStJamesParkExeter,
	StationNameStJamesStreetWalthamstow:          StationCodeStJamesStreetWalthamstow,
	StationNameStJohnsLondon:                     StationCodeStJohnsLondon,
	StationNameStKeyneWishingWellHalt:            StationCodeStKeyneWishingWellHalt,
	StationNameStLeonardsWarriorSquare:           StationCodeStLeonardsWarriorSquare,
	StationNameStMargaretsHerts:                  StationCodeStMargaretsHerts,
	StationNameStMargaretsLondon:                 StationCodeStMargaretsLondon,
	StationNameStMaryCray:                        StationCodeStMaryCray,
	StationNameStMichaels:                        StationCodeStMichaels,
	StationNameStNeots:                           StationCodeStNeots,
	StationNameStafford:                          StationCodeStafford,
	StationNameStaines:                           StationCodeStaines,
	StationNameStallingborough:                   StationCodeStallingborough,
	StationNameStalybridge:                       StationCodeStalybridge,
	StationNameStamfordHill:                      StationCodeStamfordHill,
	StationNameStamfordLincs:                     StationCodeStamfordLincs,
	StationNameStanfordleHope:                    StationCodeStanfordleHope,
	StationNameStanlowAndThornton:                StationCodeStanlowAndThornton,
	StationNameStanstedAirport:                   StationCodeStanstedAirport,
	StationNameStanstedMountfitchet:              StationCodeStanstedMountfitchet,
	StationNameStaplehurst:                       StationCodeStaplehurst,
	StationNameStapletonRoad:                     StationCodeStapletonRoad,
	StationNameStarbeck:                          StationCodeStarbeck,
	StationNameStarcross:                         StationCodeStarcross,
	StationNameStaveleyCumbria:                   StationCodeStaveleyCumbria,
	StationNameStechford:                         StationCodeStechford,
	StationNameSteetonAndSilsden:                 StationCodeSteetonAndSilsden,
	StationNameStepps:                            StationCodeStepps,
	StationNameStevenage:                         StationCodeStevenage,
	StationNameStevenston:                        StationCodeStevenston,
	StationNameStewartby:                         StationCodeStewartby,
	StationNameStewarton:                         StationCodeStewarton,
	StationNameStirling:                          StationCodeStirling,
	StationNameStockport:                         StationCodeStockport,
	StationNameStocksfield:                       StationCodeStocksfield,
	StationNameStocksmoor:                        StationCodeStocksmoor,
	StationNameStockton:                          StationCodeStockton,
	StationNameStokeMandeville:                   StationCodeStokeMandeville,
	StationNameStokeNewington:                    StationCodeStokeNewington,
	StationNameStokeonTrent:                      StationCodeStokeonTrent,
	StationNameStoneCrossing:                     StationCodeStoneCrossing,
	StationNameStoneStaffs:                       StationCodeStoneStaffs,
	StationNameStonebridgePark:                   StationCodeStonebridgePark,
	StationNameStonegate:                         StationCodeStonegate,
	StationNameStonehaven:                        StationCodeStonehaven,
	StationNameStonehouse:                        StationCodeStonehouse,
	StationNameStoneleigh:                        StationCodeStoneleigh,
	StationNameStourbridgeJunction:               StationCodeStourbridgeJunction,
	StationNameStourbridgeTown:                   StationCodeStourbridgeTown,
	StationNameStow:                              StationCodeStow,
	StationNameStowmarket:                        StationCodeStowmarket,
	StationNameStranraer:                         StationCodeStranraer,
	StationNameStratfordInternational:            StationCodeStratfordInternational,
	StationNameStratfordLondon:                   StationCodeStratfordLondon,
	StationNameStratfordUponAvon:                 StationCodeStratfordUponAvon,
	StationNameStratfordUponAvonParkway:          StationCodeStratfordUponAvonParkway,
	StationNameStrathcarron:                      StationCodeStrathcarron,
	StationNameStrawberryHill:                    StationCodeStrawberryHill,
	StationNameStreathamCommon:                   StationCodeStreathamCommon,
	StationNameStreathamGreaterLondon:            StationCodeStreathamGreaterLondon,
	StationNameStreathamHill:                     StationCodeStreathamHill,
	StationNameStreethouse:                       StationCodeStreethouse,
	StationNameStrines:                           StationCodeStrines,
	StationNameStromeferry:                       StationCodeStromeferry,
	StationNameStrood:                            StationCodeStrood,
	StationNameStroudGloucs:                      StationCodeStroudGloucs,
	StationNameSturry:                            StationCodeSturry,
	StationNameStyal:                             StationCodeStyal,
	StationNameSudburyAndHarrowRoad:              StationCodeSudburyAndHarrowRoad,
	StationNameSudburyHillHarrow:                 StationCodeSudburyHillHarrow,
	StationNameSudburySuffolk:                    StationCodeSudburySuffolk,
	StationNameSugarLoaf:                         StationCodeSugarLoaf,
	StationNameSummerston:                        StationCodeSummerston,
	StationNameSunbury:                           StationCodeSunbury,
	StationNameSunderland:                        StationCodeSunderland,
	StationNameSundridgePark:                     StationCodeSundridgePark,
	StationNameSunningdale:                       StationCodeSunningdale,
	StationNameSunnymeads:                        StationCodeSunnymeads,
	StationNameSurbiton:                          StationCodeSurbiton,
	StationNameSurreyQuays:                       StationCodeSurreyQuays,
	StationNameSuttonColdfield:                   StationCodeSuttonColdfield,
	StationNameSuttonCommon:                      StationCodeSuttonCommon,
	StationNameSuttonParkway:                     StationCodeSuttonParkway,
	StationNameSuttonSurrey:                      StationCodeSuttonSurrey,
	StationNameSwale:                             StationCodeSwale,
	StationNameSwanley:                           StationCodeSwanley,
	StationNameSwanscombe:                        StationCodeSwanscombe,
	StationNameSwansea:                           StationCodeSwansea,
	StationNameSwanwick:                          StationCodeSwanwick,
	StationNameSway:                              StationCodeSway,
	StationNameSwaythling:                        StationCodeSwaythling,
	StationNameSwinderby:                         StationCodeSwinderby,
	StationNameSwindonWilts:                      StationCodeSwindonWilts,
	StationNameSwineshead:                        StationCodeSwineshead,
	StationNameSwintonManchester:                 StationCodeSwintonManchester,
	StationNameSwintonSouthYorks:                 StationCodeSwintonSouthYorks,
	StationNameSydenhamHill:                      StationCodeSydenhamHill,
	StationNameSydenhamLondon:                    StationCodeSydenhamLondon,
	StationNameSyonLane:                          StationCodeSyonLane,
	StationNameSyston:                            StationCodeSyston,
	StationNameTackley:                           StationCodeTackley,
	StationNameTadworth:                          StationCodeTadworth,
	StationNameTaffsWell:                         StationCodeTaffsWell,
	StationNameTain:                              StationCodeTain,
	StationNameTalsarnau:                         StationCodeTalsarnau,
	StationNameTalyCafn:                          StationCodeTalyCafn,
	StationNameTalybont:                          StationCodeTalybont,
	StationNameTameBridgeParkway:                 StationCodeTameBridgeParkway,
	StationNameTamworth:                          StationCodeTamworth,
	StationNameTaplow:                            StationCodeTaplow,
	StationNameTattenhamCorner:                   StationCodeTattenhamCorner,
	StationNameTaunton:                           StationCodeTaunton,
	StationNameTaynuilt:                          StationCodeTaynuilt,
	StationNameTeddington:                        StationCodeTeddington,
	StationNameTeessideAirport:                   StationCodeTeessideAirport,
	StationNameTeignmouth:                        StationCodeTeignmouth,
	StationNameTelfordCentral:                    StationCodeTelfordCentral,
	StationNameTemplecombe:                       StationCodeTemplecombe,
	StationNameTenby:                             StationCodeTenby,
	StationNameTeynham:                           StationCodeTeynham,
	StationNameThamesDitton:                      StationCodeThamesDitton,
	StationNameThatcham:                          StationCodeThatcham,
	StationNameThattoHeath:                       StationCodeThattoHeath,
	StationNameTheHawthorns:                      StationCodeTheHawthorns,
	StationNameTheLakesWarks:                     StationCodeTheLakesWarks,
	StationNameTheale:                            StationCodeTheale,
	StationNameTheobaldsGrove:                    StationCodeTheobaldsGrove,
	StationNameThetford:                          StationCodeThetford,
	StationNameThirsk:                            StationCodeThirsk,
	StationNameThornaby:                          StationCodeThornaby,
	StationNameThorneNorth:                       StationCodeThorneNorth,
	StationNameThorneSouth:                       StationCodeThorneSouth,
	StationNameThornford:                         StationCodeThornford,
	StationNameThornliebank:                      StationCodeThornliebank,
	StationNameThorntonAbbey:                     StationCodeThorntonAbbey,
	StationNameThorntonHeath:                     StationCodeThorntonHeath,
	StationNameThorntonhall:                      StationCodeThorntonhall,
	StationNameThorpeBay:                         StationCodeThorpeBay,
	StationNameThorpeCulvert:                     StationCodeThorpeCulvert,
	StationNameThorpeleSoken:                     StationCodeThorpeleSoken,
	StationNameThreeBridges:                      StationCodeThreeBridges,
	StationNameThreeOaks:                         StationCodeThreeOaks,
	StationNameThurgarton:                        StationCodeThurgarton,
	StationNameThurnscoe:                         StationCodeThurnscoe,
	StationNameThurso:                            StationCodeThurso,
	StationNameThurston:                          StationCodeThurston,
	StationNameTilburyTown:                       StationCodeTilburyTown,
	StationNameTileHill:                          StationCodeTileHill,
	StationNameTilehurst:                         StationCodeTilehurst,
	StationNameTipton:                            StationCodeTipton,
	StationNameTirPhil:                           StationCodeTirPhil,
	StationNameTisbury:                           StationCodeTisbury,
	StationNameTivertonParkway:                   StationCodeTivertonParkway,
	StationNameTodmorden:                         StationCodeTodmorden,
	StationNameTolworth:                          StationCodeTolworth,
	StationNameTonPentre:                         StationCodeTonPentre,
	StationNameTonbridge:                         StationCodeTonbridge,
	StationNameTondu:                             StationCodeTondu,
	StationNameTonfanau:                          StationCodeTonfanau,
	StationNameTonypandy:                         StationCodeTonypandy,
	StationNameTooting:                           StationCodeTooting,
	StationNameTopsham:                           StationCodeTopsham,
	StationNameTorquay:                           StationCodeTorquay,
	StationNameTorre:                             StationCodeTorre,
	StationNameTotnes:                            StationCodeTotnes,
	StationNameTottenhamHale:                     StationCodeTottenhamHale,
	StationNameTotton:                            StationCodeTotton,
	StationNameTownGreen:                         StationCodeTownGreen,
	StationNameTraffordPark:                      StationCodeTraffordPark,
	StationNameTrefforest:                        StationCodeTrefforest,
	StationNameTrefforestEstate:                  StationCodeTrefforestEstate,
	StationNameTrehafod:                          StationCodeTrehafod,
	StationNameTreherbert:                        StationCodeTreherbert,
	StationNameTreorchy:                          StationCodeTreorchy,
	StationNameTrimley:                           StationCodeTrimley,
	StationNameTring:                             StationCodeTring,
	StationNameTroedyrhiw:                        StationCodeTroedyrhiw,
	StationNameTroon:                             StationCodeTroon,
	StationNameTrowbridge:                        StationCodeTrowbridge,
	StationNameTruro:                             StationCodeTruro,
	StationNameTulloch:                           StationCodeTulloch,
	StationNameTulseHill:                         StationCodeTulseHill,
	StationNameTunbridgeWells:                    StationCodeTunbridgeWells,
	StationNameTurkeyStreet:                      StationCodeTurkeyStreet,
	StationNameTutburyAndHatton:                  StationCodeTutburyAndHatton,
	StationNameTweedbank:                         StationCodeTweedbank,
	StationNameTwickenham:                        StationCodeTwickenham,
	StationNameTwyford:                           StationCodeTwyford,
	StationNameTyCroes:                           StationCodeTyCroes,
	StationNameTyGlas:                            StationCodeTyGlas,
	StationNameTygwyn:                            StationCodeTygwyn,
	StationNameTyndrumLower:                      StationCodeTyndrumLower,
	StationNameTyseley:                           StationCodeTyseley,
	StationNameTywyn:                             StationCodeTywyn,
	StationNameUckfield:                          StationCodeUckfield,
	StationNameUddingston:                        StationCodeUddingston,
	StationNameUlceby:                            StationCodeUlceby,
	StationNameUlleskelf:                         StationCodeUlleskelf,
	StationNameUlverston:                         StationCodeUlverston,
	StationNameUmberleigh:                        StationCodeUmberleigh,
	StationNameUniversityBirmingham:              StationCodeUniversityBirmingham,
	StationNameUphall:                            StationCodeUphall,
	StationNameUpholland:                         StationCodeUpholland,
	StationNameUpminster:                         StationCodeUpminster,
	StationNameUpperHalliford:                    StationCodeUpperHalliford,
	StationNameUpperHolloway:                     StationCodeUpperHolloway,
	StationNameUpperTyndrum:                      StationCodeUpperTyndrum,
	StationNameUpperWarlingham:                   StationCodeUpperWarlingham,
	StationNameUptonMerseyside:                   StationCodeUptonMerseyside,
	StationNameUpwey:                             StationCodeUpwey,
	StationNameUrmston:                           StationCodeUrmston,
	StationNameUttoxeter:                         StationCodeUttoxeter,
	StationNameValley:                            StationCodeValley,
	StationNameVauxhall:                          StationCodeVauxhall,
	StationNameVirginiaWater:                     StationCodeVirginiaWater,
	StationNameWaddon:                            StationCodeWaddon,
	StationNameWadhurst:                          StationCodeWadhurst,
	StationNameWainfleet:                         StationCodeWainfleet,
	StationNameWakefieldKirkgate:                 StationCodeWakefieldKirkgate,
	StationNameWakefieldWestgate:                 StationCodeWakefieldWestgate,
	StationNameWalkden:                           StationCodeWalkden,
	StationNameWallaseyGroveRoad:                 StationCodeWallaseyGroveRoad,
	StationNameWallaseyVillage:                   StationCodeWallaseyVillage,
	StationNameWallington:                        StationCodeWallington,
	StationNameWallyford:                         StationCodeWallyford,
	StationNameWalmer:                            StationCodeWalmer,
	StationNameWalsall:                           StationCodeWalsall,
	StationNameWalsden:                           StationCodeWalsden,
	StationNameWalthamCross:                      StationCodeWalthamCross,
	StationNameWalthamstowCentral:                StationCodeWalthamstowCentral,
	StationNameWalthamstowQueensRoad:             StationCodeWalthamstowQueensRoad,
	StationNameWaltonMerseyside:                  StationCodeWaltonMerseyside,
	StationNameWaltononThames:                    StationCodeWaltononThames,
	StationNameWaltonontheNaze:                   StationCodeWaltonontheNaze,
	StationNameWanborough:                        StationCodeWanborough,
	StationNameWandsworthCommon:                  StationCodeWandsworthCommon,
	StationNameWandsworthRoad:                    StationCodeWandsworthRoad,
	StationNameWandsworthTown:                    StationCodeWandsworthTown,
	StationNameWansteadPark:                      StationCodeWansteadPark,
	StationNameWapping:                           StationCodeWapping,
	StationNameWarblington:                       StationCodeWarblington,
	StationNameWareHerts:                         StationCodeWareHerts,
	StationNameWarehamDorset:                     StationCodeWarehamDorset,
	StationNameWargrave:                          StationCodeWargrave,
	StationNameWarminster:                        StationCodeWarminster,
	StationNameWarnham:                           StationCodeWarnham,
	StationNameWarringtonBankQuay:                StationCodeWarringtonBankQuay,
	StationNameWarringtonCentral:                 StationCodeWarringtonCentral,
	StationNameWarwick:                           StationCodeWarwick,
	StationNameWarwickParkway:                    StationCodeWarwickParkway,
	StationNameWaterOrton:                        StationCodeWaterOrton,
	StationNameWaterbeach:                        StationCodeWaterbeach,
	StationNameWateringbury:                      StationCodeWateringbury,
	StationNameWaterlooMerseyside:                StationCodeWaterlooMerseyside,
	StationNameWatfordHighStreet:                 StationCodeWatfordHighStreet,
	StationNameWatfordJunction:                   StationCodeWatfordJunction,
	StationNameWatfordNorth:                      StationCodeWatfordNorth,
	StationNameWatlington:                        StationCodeWatlington,
	StationNameWattonatStone:                     StationCodeWattonatStone,
	StationNameWaunGronPark:                      StationCodeWaunGronPark,
	StationNameWavertreeTechnologyPark:           StationCodeWavertreeTechnologyPark,
	StationNameWedgwood:                          StationCodeWedgwood,
	StationNameWeeley:                            StationCodeWeeley,
	StationNameWeeton:                            StationCodeWeeton,
	StationNameWelhamGreen:                       StationCodeWelhamGreen,
	StationNameWelling:                           StationCodeWelling,
	StationNameWellingborough:                    StationCodeWellingborough,
	StationNameWellingtonShropshire:              StationCodeWellingtonShropshire,
	StationNameWelshpool:                         StationCodeWelshpool,
	StationNameWelwynGardenCity:                  StationCodeWelwynGardenCity,
	StationNameWelwynNorth:                       StationCodeWelwynNorth,
	StationNameWem:                               StationCodeWem,
	StationNameWembleyCentral:                    StationCodeWembleyCentral,
	StationNameWembleyStadium:                    StationCodeWembleyStadium,
	StationNameWemyssBay:                         StationCodeWemyssBay,
	StationNameWendover:                          StationCodeWendover,
	StationNameWennington:                        StationCodeWennington,
	StationNameWestAllerton:                      StationCodeWestAllerton,
	StationNameWestBrompton:                      StationCodeWestBrompton,
	StationNameWestByfleet:                       StationCodeWestByfleet,
	StationNameWestCalder:                        StationCodeWestCalder,
	StationNameWestCroydon:                       StationCodeWestCroydon,
	StationNameWestDrayton:                       StationCodeWestDrayton,
	StationNameWestDulwich:                       StationCodeWestDulwich,
	StationNameWestEaling:                        StationCodeWestEaling,
	StationNameWestHam:                           StationCodeWestHam,
	StationNameWestHampstead:                     StationCodeWestHampstead,
	StationNameWestHampsteadThameslink:           StationCodeWestHampsteadThameslink,
	StationNameWestHorndon:                       StationCodeWestHorndon,
	StationNameWestKilbride:                      StationCodeWestKilbride,
	StationNameWestKirby:                         StationCodeWestKirby,
	StationNameWestMalling:                       StationCodeWestMalling,
	StationNameWestNorwood:                       StationCodeWestNorwood,
	StationNameWestRuislip:                       StationCodeWestRuislip,
	StationNameWestRunton:                        StationCodeWestRunton,
	StationNameWestStLeonards:                    StationCodeWestStLeonards,
	StationNameWestSutton:                        StationCodeWestSutton,
	StationNameWestWickham:                       StationCodeWestWickham,
	StationNameWestWorthing:                      StationCodeWestWorthing,
	StationNameWestburyWilts:                     StationCodeWestburyWilts,
	StationNameWestcliff:                         StationCodeWestcliff,
	StationNameWestcombePark:                     StationCodeWestcombePark,
	StationNameWestenhanger:                      StationCodeWestenhanger,
	StationNameWesterHailes:                      StationCodeWesterHailes,
	StationNameWesterfield:                       StationCodeWesterfield,
	StationNameWesterton:                         StationCodeWesterton,
	StationNameWestgateOnSea:                     StationCodeWestgateOnSea,
	StationNameWesthoughton:                      StationCodeWesthoughton,
	StationNameWestonMilton:                      StationCodeWestonMilton,
	StationNameWestonsuperMare:                   StationCodeWestonsuperMare,
	StationNameWetheral:                          StationCodeWetheral,
	StationNameWeybridge:                         StationCodeWeybridge,
	StationNameWeymouth:                          StationCodeWeymouth,
	StationNameWhaleyBridge:                      StationCodeWhaleyBridge,
	StationNameWhalleyLancs:                      StationCodeWhalleyLancs,
	StationNameWhatstandwell:                     StationCodeWhatstandwell,
	StationNameWhifflet:                          StationCodeWhifflet,
	StationNameWhimple:                           StationCodeWhimple,
	StationNameWhinhill:                          StationCodeWhinhill,
	StationNameWhiston:                           StationCodeWhiston,
	StationNameWhitby:                            StationCodeWhitby,
	StationNameWhitchurchCardiff:                 StationCodeWhitchurchCardiff,
	StationNameWhitchurchHants:                   StationCodeWhitchurchHants,
	StationNameWhitchurchShropshire:              StationCodeWhitchurchShropshire,
	StationNameWhiteHartLane:                     StationCodeWhiteHartLane,
	StationNameWhiteNotley:                       StationCodeWhiteNotley,
	StationNameWhitechapel:                       StationCodeWhitechapel,
	StationNameWhitecraigs:                       StationCodeWhitecraigs,
	StationNameWhitehaven:                        StationCodeWhitehaven,
	StationNameWhitland:                          StationCodeWhitland,
	StationNameWhitleyBridge:                     StationCodeWhitleyBridge,
	StationNameWhitlocksEnd:                      StationCodeWhitlocksEnd,
	StationNameWhitstable:                        StationCodeWhitstable,
	StationNameWhittlesea:                        StationCodeWhittlesea,
	StationNameWhittlesfordParkway:               StationCodeWhittlesfordParkway,
	StationNameWhittonLondon:                     StationCodeWhittonLondon,
	StationNameWhitwellDerbyshire:                StationCodeWhitwellDerbyshire,
	StationNameWhyteleafe:                        StationCodeWhyteleafe,
	StationNameWhyteleafeSouth:                   StationCodeWhyteleafeSouth,
	StationNameWick:                              StationCodeWick,
	StationNameWickford:                          StationCodeWickford,
	StationNameWickhamMarket:                     StationCodeWickhamMarket,
	StationNameWiddrington:                       StationCodeWiddrington,
	StationNameWidnes:                            StationCodeWidnes,
	StationNameWidneyManor:                       StationCodeWidneyManor,
	StationNameWiganNorthWestern:                 StationCodeWiganNorthWestern,
	StationNameWiganWallgate:                     StationCodeWiganWallgate,
	StationNameWigton:                            StationCodeWigton,
	StationNameWildmill:                          StationCodeWildmill,
	StationNameWillesdenJunction:                 StationCodeWillesdenJunction,
	StationNameWilliamwood:                       StationCodeWilliamwood,
	StationNameWillington:                        StationCodeWillington,
	StationNameWilmcote:                          StationCodeWilmcote,
	StationNameWilmslow:                          StationCodeWilmslow,
	StationNameWilnecoteStaffs:                   StationCodeWilnecoteStaffs,
	StationNameWimbledon:                         StationCodeWimbledon,
	StationNameWimbledonChase:                    StationCodeWimbledonChase,
	StationNameWinchelsea:                        StationCodeWinchelsea,
	StationNameWinchester:                        StationCodeWinchester,
	StationNameWinchfield:                        StationCodeWinchfield,
	StationNameWinchmoreHill:                     StationCodeWinchmoreHill,
	StationNameWindermere:                        StationCodeWindermere,
	StationNameWindsorAndEtonCentral:             StationCodeWindsorAndEtonCentral,
	StationNameWindsorAndEtonRiverside:           StationCodeWindsorAndEtonRiverside,
	StationNameWinnersh:                          StationCodeWinnersh,
	StationNameWinnershTriangle:                  StationCodeWinnershTriangle,
	StationNameWinsford:                          StationCodeWinsford,
	StationNameWishaw:                            StationCodeWishaw,
	StationNameWitham:                            StationCodeWitham,
	StationNameWitley:                            StationCodeWitley,
	StationNameWittonWestMidlands:                StationCodeWittonWestMidlands,
	StationNameWivelsfield:                       StationCodeWivelsfield,
	StationNameWivenhoe:                          StationCodeWivenhoe,
	StationNameWoburnSands:                       StationCodeWoburnSands,
	StationNameWoking:                            StationCodeWoking,
	StationNameWokingham:                         StationCodeWokingham,
	StationNameWoldingham:                        StationCodeWoldingham,
	StationNameWolverhampton:                     StationCodeWolverhampton,
	StationNameWolverton:                         StationCodeWolverton,
	StationNameWombwell:                          StationCodeWombwell,
	StationNameWoodEnd:                           StationCodeWoodEnd,
	StationNameWoodStreet:                        StationCodeWoodStreet,
	StationNameWoodbridge:                        StationCodeWoodbridge,
	StationNameWoodgrangePark:                    StationCodeWoodgrangePark,
	StationNameWoodhall:                          StationCodeWoodhall,
	StationNameWoodhouse:                         StationCodeWoodhouse,
	StationNameWoodlesford:                       StationCodeWoodlesford,
	StationNameWoodley:                           StationCodeWoodley,
	StationNameWoodmansterne:                     StationCodeWoodmansterne,
	StationNameWoodsmoor:                         StationCodeWoodsmoor,
	StationNameWool:                              StationCodeWool,
	StationNameWoolston:                          StationCodeWoolston,
	StationNameWoolwichArsenal:                   StationCodeWoolwichArsenal,
	StationNameWoolwichDockyard:                  StationCodeWoolwichDockyard,
	StationNameWoottonWawen:                      StationCodeWoottonWawen,
	StationNameWorcesterForegateStreet:           StationCodeWorcesterForegateStreet,
	StationNameWorcesterPark:                     StationCodeWorcesterPark,
	StationNameWorcesterShrubHill:                StationCodeWorcesterShrubHill,
	StationNameWorkington:                        StationCodeWorkington,
	StationNameWorksop:                           StationCodeWorksop,
	StationNameWorle:                             StationCodeWorle,
	StationNameWorplesdon:                        StationCodeWorplesdon,
	StationNameWorstead:                          StationCodeWorstead,
	StationNameWorthing:                          StationCodeWorthing,
	StationNameWrabness:                          StationCodeWrabness,
	StationNameWraysbury:                         StationCodeWraysbury,
	StationNameWrenbury:                          StationCodeWrenbury,
	StationNameWressle:                           StationCodeWressle,
	StationNameWrexhamCentral:                    StationCodeWrexhamCentral,
	StationNameWrexhamGeneral:                    StationCodeWrexhamGeneral,
	StationNameWye:                               StationCodeWye,
	StationNameWylam:                             StationCodeWylam,
	StationNameWyldeGreen:                        StationCodeWyldeGreen,
	StationNameWymondham:                         StationCodeWymondham,
	StationNameWythall:                           StationCodeWythall,
	StationNameYalding:                           StationCodeYalding,
	StationNameYardleyWood:                       StationCodeYardleyWood,
	StationNameYarm:                              StationCodeYarm,
	StationNameYate:                              StationCodeYate,
	StationNameYatton:                            StationCodeYatton,
	StationNameYeoford:                           StationCodeYeoford,
	StationNameYeovilJunction:                    StationCodeYeovilJunction,
	StationNameYeovilPenMill:                     StationCodeYeovilPenMill,
	StationNameYetminster:                        StationCodeYetminster,
	StationNameYnyswen:                           StationCodeYnyswen,
	StationNameYoker:                             StationCodeYoker,
	StationNameYork:                              StationCodeYork,
	StationNameYorton:                            StationCodeYorton,
	StationNameYstradMynach:                      StationCodeYstradMynach,
	StationNameYstradRhondda:                     StationCodeYstradRhondda,
}

const (
	StationNameAbbeyWood                         = "Abbey Wood"
	StationNameAber                              = "Aber"
	StationNameAbercynon                         = "Abercynon"
	StationNameAberdare                          = "Aberdare"
	StationNameAberdeen                          = "Aberdeen"
	StationNameAberdour                          = "Aberdour"
	StationNameAberdovey                         = "Aberdovey"
	StationNameAbererch                          = "Abererch"
	StationNameAbergavenny                       = "Abergavenny"
	StationNameAbergeleAndPensarn                = "Abergele & Pensarn"
	StationNameAberystwyth                       = "Aberystwyth"
	StationNameAccrington                        = "Accrington"
	StationNameAchanalt                          = "Achanalt"
	StationNameAchnasheen                        = "Achnasheen"
	StationNameAchnashellach                     = "Achnashellach"
	StationNameAcklington                        = "Acklington"
	StationNameAcle                              = "Acle"
	StationNameAcocksGreen                       = "Acocks Green"
	StationNameActonBridgeCheshire               = "Acton Bridge (Cheshire)"
	StationNameActonCentral                      = "Acton Central"
	StationNameActonMainLine                     = "Acton Main Line"
	StationNameAdderleyPark                      = "Adderley Park"
	StationNameAddiewell                         = "Addiewell"
	StationNameAddlestone                        = "Addlestone"
	StationNameAdisham                           = "Adisham"
	StationNameAdlingtonCheshire                 = "Adlington (Cheshire)"
	StationNameAdlingtonLancs                    = "Adlington (Lancs)"
	StationNameAdwick                            = "Adwick"
	StationNameAigburth                          = "Aigburth"
	StationNameAinsdale                          = "Ainsdale"
	StationNameAintree                           = "Aintree"
	StationNameAirbles                           = "Airbles"
	StationNameAirdrie                           = "Airdrie"
	StationNameAlbanyPark                        = "Albany Park"
	StationNameAlbrighton                        = "Albrighton"
	StationNameAlderleyEdge                      = "Alderley Edge"
	StationNameAldermaston                       = "Aldermaston"
	StationNameAldershot                         = "Aldershot"
	StationNameAldrington                        = "Aldrington"
	StationNameAlexandraPalace                   = "Alexandra Palace"
	StationNameAlexandraParade                   = "Alexandra Parade"
	StationNameAlexandria                        = "Alexandria"
	StationNameAlfreton                          = "Alfreton"
	StationNameAllensWest                        = "Allens West"
	StationNameAlloa                             = "Alloa"
	StationNameAlness                            = "Alness"
	StationNameAlnmouth                          = "Alnmouth"
	StationNameAlresfordEssex                    = "Alresford (Essex)"
	StationNameAlsager                           = "Alsager"
	StationNameAlthorneEssex                     = "Althorne (Essex)"
	StationNameAlthorpe                          = "Althorpe"
	StationNameAltnabreac                        = "Altnabreac"
	StationNameAlton                             = "Alton"
	StationNameAltrincham                        = "Altrincham"
	StationNameAlvechurch                        = "Alvechurch"
	StationNameAmbergate                         = "Ambergate"
	StationNameAmberley                          = "Amberley"
	StationNameAmersham                          = "Amersham"
	StationNameAmmanford                         = "Ammanford"
	StationNameAncaster                          = "Ancaster"
	StationNameAnderston                         = "Anderston"
	StationNameAndover                           = "Andover"
	StationNameAnerley                           = "Anerley"
	StationNameAngelRoad                         = "Angel Road"
	StationNameAngmering                         = "Angmering"
	StationNameAnnan                             = "Annan"
	StationNameAnniesland                        = "Anniesland"
	StationNameAnsdellAndFairhaven               = "Ansdell & Fairhaven"
	StationNameApperleyBridge                    = "Apperley Bridge"
	StationNameAppleby                           = "Appleby"
	StationNameAppledoreKent                     = "Appledore (Kent)"
	StationNameAppleford                         = "Appleford"
	StationNameAppleyBridge                      = "Appley Bridge"
	StationNameApsley                            = "Apsley"
	StationNameArbroath                          = "Arbroath"
	StationNameArdgay                            = "Ardgay"
	StationNameArdlui                            = "Ardlui"
	StationNameArdrossanHarbour                  = "Ardrossan Harbour"
	StationNameArdrossanSouthBeach               = "Ardrossan South Beach"
	StationNameArdrossanTown                     = "Ardrossan Town"
	StationNameArdwick                           = "Ardwick"
	StationNameArgyleStreet                      = "Argyle Street"
	StationNameArisaig                           = "Arisaig"
	StationNameArlesey                           = "Arlesey"
	StationNameArmadaleWestLothian               = "Armadale (West Lothian)"
	StationNameArmathwaite                       = "Armathwaite"
	StationNameArnside                           = "Arnside"
	StationNameArram                             = "Arram"
	StationNameArrocharAndTarbet                 = "Arrochar & Tarbet"
	StationNameArundel                           = "Arundel"
	StationNameAscotBerks                        = "Ascot (Berks)"
	StationNameAscottunderWychwood               = "Ascott-under-Wychwood"
	StationNameAsh                               = "Ash"
	StationNameAshVale                           = "Ash Vale"
	StationNameAshburys                          = "Ashburys"
	StationNameAshchurchforTewkesbury            = "Ashchurch for Tewkesbury"
	StationNameAshfield                          = "Ashfield"
	StationNameAshfordInternational              = "Ashford International"
	StationNameAshfordInternationalEurostar      = "Ashford International (Eurostar)"
	StationNameAshfordSurrey                     = "Ashford (Surrey)"
	StationNameAshley                            = "Ashley"
	StationNameAshtead                           = "Ashtead"
	StationNameAshtonunderLyne                   = "Ashton-under-Lyne"
	StationNameAshurstKent                       = "Ashurst (Kent)"
	StationNameAshurstNewForest                  = "Ashurst New Forest"
	StationNameAshwellAndMorden                  = "Ashwell & Morden"
	StationNameAskam                             = "Askam"
	StationNameAslockton                         = "Aslockton"
	StationNameAspatria                          = "Aspatria"
	StationNameAspleyGuise                       = "Aspley Guise"
	StationNameAston                             = "Aston"
	StationNameAtherstone                        = "Atherstone"
	StationNameAtherton                          = "Atherton"
	StationNameAttadale                          = "Attadale"
	StationNameAttenborough                      = "Attenborough"
	StationNameAttleborough                      = "Attleborough"
	StationNameAuchinleck                        = "Auchinleck"
	StationNameAudleyEnd                         = "Audley End"
	StationNameAughtonPark                       = "Aughton Park"
	StationNameAviemore                          = "Aviemore"
	StationNameAvoncliff                         = "Avoncliff"
	StationNameAvonmouth                         = "Avonmouth"
	StationNameAxminster                         = "Axminster"
	StationNameAylesbury                         = "Aylesbury"
	StationNameAylesburyValeParkway              = "Aylesbury Vale Parkway"
	StationNameAylesford                         = "Aylesford"
	StationNameAylesham                          = "Aylesham"
	StationNameAyr                               = "Ayr"
	StationNameBache                             = "Bache"
	StationNameBaglan                            = "Baglan"
	StationNameBagshot                           = "Bagshot"
	StationNameBaildon                           = "Baildon"
	StationNameBaillieston                       = "Baillieston"
	StationNameBalcombe                          = "Balcombe"
	StationNameBaldock                           = "Baldock"
	StationNameBalham                            = "Balham"
	StationNameBalloch                           = "Balloch"
	StationNameBalmossie                         = "Balmossie"
	StationNameBamberBridge                      = "Bamber Bridge"
	StationNameBamford                           = "Bamford"
	StationNameBanavie                           = "Banavie"
	StationNameBanbury                           = "Banbury"
	StationNameBangorGwynedd                     = "Bangor (Gwynedd)"
	StationNameBankHall                          = "Bank Hall"
	StationNameBanstead                          = "Banstead"
	StationNameBarassie                          = "Barassie"
	StationNameBarbican                          = "Barbican"
	StationNameBardonMill                        = "Bardon Mill"
	StationNameBareLane                          = "Bare Lane"
	StationNameBargeddie                         = "Bargeddie"
	StationNameBargoed                           = "Bargoed"
	StationNameBarking                           = "Barking"
	StationNameBarlaston                         = "Barlaston"
	StationNameBarming                           = "Barming"
	StationNameBarmouth                          = "Barmouth"
	StationNameBarnehurst                        = "Barnehurst"
	StationNameBarnes                            = "Barnes"
	StationNameBarnesBridge                      = "Barnes Bridge"
	StationNameBarnetby                          = "Barnetby"
	StationNameBarnham                           = "Barnham"
	StationNameBarnhill                          = "Barnhill"
	StationNameBarnsley                          = "Barnsley"
	StationNameBarnstaple                        = "Barnstaple"
	StationNameBarntGreen                        = "Barnt Green"
	StationNameBarrhead                          = "Barrhead"
	StationNameBarrhill                          = "Barrhill"
	StationNameBarrowHaven                       = "Barrow Haven"
	StationNameBarrowUponSoar                    = "Barrow-Upon-Soar"
	StationNameBarrowinFurness                   = "Barrow-in-Furness"
	StationNameBarry                             = "Barry"
	StationNameBarryDocks                        = "Barry Docks"
	StationNameBarryIsland                       = "Barry Island"
	StationNameBarryLinks                        = "Barry Links"
	StationNameBartononHumber                    = "Barton-on-Humber"
	StationNameBasildon                          = "Basildon"
	StationNameBasingstoke                       = "Basingstoke"
	StationNameBatAndBall                        = "Bat & Ball"
	StationNameBathSpa                           = "Bath Spa"
	StationNameBathgate                          = "Bathgate"
	StationNameBatley                            = "Batley"
	StationNameBattersby                         = "Battersby"
	StationNameBatterseaPark                     = "Battersea Park"
	StationNameBattle                            = "Battle"
	StationNameBattlesbridge                     = "Battlesbridge"
	StationNameBayford                           = "Bayford"
	StationNameBeaconsfield                      = "Beaconsfield"
	StationNameBearley                           = "Bearley"
	StationNameBearsden                          = "Bearsden"
	StationNameBearsted                          = "Bearsted"
	StationNameBeasdale                          = "Beasdale"
	StationNameBeaulieuRoad                      = "Beaulieu Road"
	StationNameBeauly                            = "Beauly"
	StationNameBebington                         = "Bebington"
	StationNameBeccles                           = "Beccles"
	StationNameBeckenhamHill                     = "Beckenham Hill"
	StationNameBeckenhamJunction                 = "Beckenham Junction"
	StationNameBedford                           = "Bedford"
	StationNameBedfordStJohns                    = "Bedford St Johns"
	StationNameBedhampton                        = "Bedhampton"
	StationNameBedminster                        = "Bedminster"
	StationNameBedworth                          = "Bedworth"
	StationNameBedwyn                            = "Bedwyn"
	StationNameBeeston                           = "Beeston"
	StationNameBekesbourne                       = "Bekesbourne"
	StationNameBelleVue                          = "Belle Vue"
	StationNameBellgrove                         = "Bellgrove"
	StationNameBellingham                        = "Bellingham"
	StationNameBellshill                         = "Bellshill"
	StationNameBelmont                           = "Belmont"
	StationNameBelper                            = "Belper"
	StationNameBeltring                          = "Beltring"
	StationNameBelvedere                         = "Belvedere"
	StationNameBempton                           = "Bempton"
	StationNameBenRhydding                       = "Ben Rhydding"
	StationNameBenfleet                          = "Benfleet"
	StationNameBentham                           = "Bentham"
	StationNameBentleyHants                      = "Bentley (Hants)"
	StationNameBentleySouthYorks                 = "Bentley (South Yorks)"
	StationNameBereAlston                        = "Bere Alston"
	StationNameBereFerrers                       = "Bere Ferrers"
	StationNameBerkhamsted                       = "Berkhamsted"
	StationNameBerkswell                         = "Berkswell"
	StationNameBermudaPark                       = "Bermuda Park"
	StationNameBerneyArms                        = "Berney Arms"
	StationNameBerryBrow                         = "Berry Brow"
	StationNameBerrylands                        = "Berrylands"
	StationNameBerwickSussex                     = "Berwick (Sussex)"
	StationNameBerwickUponTweed                  = "Berwick-upon-Tweed"
	StationNameBescarLane                        = "Bescar Lane"
	StationNameBescotStadium                     = "Bescot Stadium"
	StationNameBetchworth                        = "Betchworth"
	StationNameBethnalGreen                      = "Bethnal Green"
	StationNameBetwsyCoed                        = "Betws-y-Coed"
	StationNameBeverley                          = "Beverley"
	StationNameBexhill                           = "Bexhill"
	StationNameBexley                            = "Bexley"
	StationNameBexleyheath                       = "Bexleyheath"
	StationNameBicesterNorth                     = "Bicester North"
	StationNameBicesterVillage                   = "Bicester Village"
	StationNameBickley                           = "Bickley"
	StationNameBidston                           = "Bidston"
	StationNameBiggleswade                       = "Biggleswade"
	StationNameBilbrook                          = "Bilbrook"
	StationNameBillericay                        = "Billericay"
	StationNameBillinghamCleveland               = "Billingham (Cleveland)"
	StationNameBillingshurst                     = "Billingshurst"
	StationNameBingham                           = "Bingham"
	StationNameBingley                           = "Bingley"
	StationNameBirchgrove                        = "Birchgrove"
	StationNameBirchingtonOnSea                  = "Birchington-on-Sea"
	StationNameBirchwood                         = "Birchwood"
	StationNameBirkbeck                          = "Birkbeck"
	StationNameBirkdale                          = "Birkdale"
	StationNameBirkenheadCentral                 = "Birkenhead Central"
	StationNameBirkenheadHamiltonSquare          = "Birkenhead Hamilton Square"
	StationNameBirkenheadNorth                   = "Birkenhead North"
	StationNameBirkenheadPark                    = "Birkenhead Park"
	StationNameBirminghamInternational           = "Birmingham International"
	StationNameBirminghamMoorStreet              = "Birmingham Moor Street"
	StationNameBirminghamNewStreet               = "Birmingham New Street"
	StationNameBirminghamSnowHill                = "Birmingham Snow Hill"
	StationNameBishopAuckland                    = "Bishop Auckland"
	StationNameBishopbriggs                      = "Bishopbriggs"
	StationNameBishopsStortford                  = "Bishops Stortford"
	StationNameBishopstoneSussex                 = "Bishopstone (Sussex)"
	StationNameBishoptonStrathclyde              = "Bishopton (Strathclyde)"
	StationNameBitterne                          = "Bitterne"
	StationNameBlackburn                         = "Blackburn"
	StationNameBlackheath                        = "Blackheath"
	StationNameBlackhorseRoad                    = "Blackhorse Road"
	StationNameBlackpoolNorth                    = "Blackpool North"
	StationNameBlackpoolPleasureBeach            = "Blackpool Pleasure Beach"
	StationNameBlackpoolSouth                    = "Blackpool South"
	StationNameBlackridge                        = "Blackridge"
	StationNameBlackrod                          = "Blackrod"
	StationNameBlackwater                        = "Blackwater"
	StationNameBlaenauFfestiniog                 = "Blaenau Ffestiniog"
	StationNameBlairAtholl                       = "Blair Atholl"
	StationNameBlairhill                         = "Blairhill"
	StationNameBlakeStreet                       = "Blake Street"
	StationNameBlakedown                         = "Blakedown"
	StationNameBlantyre                          = "Blantyre"
	StationNameBlaydon                           = "Blaydon"
	StationNameBleasby                           = "Bleasby"
	StationNameBletchley                         = "Bletchley"
	StationNameBloxwich                          = "Bloxwich"
	StationNameBloxwichNorth                     = "Bloxwich North"
	StationNameBlundellsandsAndCrosby            = "Blundellsands & Crosby"
	StationNameBlytheBridge                      = "Blythe Bridge"
	StationNameBodminParkway                     = "Bodmin Parkway"
	StationNameBodorgan                          = "Bodorgan"
	StationNameBognorRegis                       = "Bognor Regis"
	StationNameBogston                           = "Bogston"
	StationNameBolton                            = "Bolton"
	StationNameBoltonUponDearne                  = "Bolton-Upon-Dearne"
	StationNameBookham                           = "Bookham"
	StationNameBootleCumbria                     = "Bootle (Cumbria)"
	StationNameBootleNewStrand                   = "Bootle New Strand"
	StationNameBootleOrielRoad                   = "Bootle Oriel Road"
	StationNameBordesley                         = "Bordesley"
	StationNameBoroughGreenAndWrotham            = "Borough Green & Wrotham"
	StationNameBorth                             = "Borth"
	StationNameBosham                            = "Bosham"
	StationNameBoston                            = "Boston"
	StationNameBotley                            = "Botley"
	StationNameBottesford                        = "Bottesford"
	StationNameBourneEnd                         = "Bourne End"
	StationNameBournemouth                       = "Bournemouth"
	StationNameBournville                        = "Bournville"
	StationNameBowBrickhill                      = "Bow Brickhill"
	StationNameBowesPark                         = "Bowes Park"
	StationNameBowling                           = "Bowling"
	StationNameBoxHillAndWesthumble              = "Box Hill & Westhumble"
	StationNameBracknell                         = "Bracknell"
	StationNameBradfordForsterSquare             = "Bradford Forster Square"
	StationNameBradfordInterchange               = "Bradford Interchange"
	StationNameBradfordonAvon                    = "Bradford-on-Avon"
	StationNameBrading                           = "Brading"
	StationNameBraintree                         = "Braintree"
	StationNameBraintreeFreeport                 = "Braintree Freeport"
	StationNameBramhall                          = "Bramhall"
	StationNameBramleyHants                      = "Bramley (Hants)"
	StationNameBramleyWYorks                     = "Bramley (W Yorks)"
	StationNameBramptonCumbria                   = "Brampton (Cumbria)"
	StationNameBramptonSuffolk                   = "Brampton (Suffolk)"
	StationNameBranchton                         = "Branchton"
	StationNameBrandon                           = "Brandon"
	StationNameBranksome                         = "Branksome"
	StationNameBraystonesCumbria                 = "Braystones (Cumbria)"
	StationNameBredbury                          = "Bredbury"
	StationNameBreich                            = "Breich"
	StationNameBrentford                         = "Brentford"
	StationNameBrentwood                         = "Brentwood"
	StationNameBricketWood                       = "Bricket Wood"
	StationNameBridgend                          = "Bridgend"
	StationNameBridgeofAllan                     = "Bridge of Allan"
	StationNameBridgeofOrchy                     = "Bridge of Orchy"
	StationNameBridgeton                         = "Bridgeton"
	StationNameBridgwater                        = "Bridgwater"
	StationNameBridlington                       = "Bridlington"
	StationNameBrierfield                        = "Brierfield"
	StationNameBrigg                             = "Brigg"
	StationNameBrighouse                         = "Brighouse"
	StationNameBrightonEastSussex                = "Brighton (East Sussex)"
	StationNameBrimsdown                         = "Brimsdown"
	StationNameBrinnington                       = "Brinnington"
	StationNameBristolParkway                    = "Bristol Parkway"
	StationNameBristolTempleMeads                = "Bristol Temple Meads"
	StationNameBrithdir                          = "Brithdir"
	StationNameBritonFerry                       = "Briton Ferry"
	StationNameBrixton                           = "Brixton"
	StationNameBroadGreen                        = "Broad Green"
	StationNameBroadbottom                       = "Broadbottom"
	StationNameBroadstairs                       = "Broadstairs"
	StationNameBrockenhurst                      = "Brockenhurst"
	StationNameBrockholes                        = "Brockholes"
	StationNameBrockley                          = "Brockley"
	StationNameBromborough                       = "Bromborough"
	StationNameBromboroughRake                   = "Bromborough Rake"
	StationNameBromleyCrossLancs                 = "Bromley Cross (Lancs)"
	StationNameBromleyNorth                      = "Bromley North"
	StationNameBromleySouth                      = "Bromley South"
	StationNameBromsgrove                        = "Bromsgrove"
	StationNameBrondesbury                       = "Brondesbury"
	StationNameBrondesburyPark                   = "Brondesbury Park"
	StationNameBrookmansPark                     = "Brookmans Park"
	StationNameBrookwood                         = "Brookwood"
	StationNameBroome                            = "Broome"
	StationNameBroomfleet                        = "Broomfleet"
	StationNameBrora                             = "Brora"
	StationNameBrough                            = "Brough"
	StationNameBroughtyFerry                     = "Broughty Ferry"
	StationNameBroxbourne                        = "Broxbourne"
	StationNameBruceGrove                        = "Bruce Grove"
	StationNameBrundall                          = "Brundall"
	StationNameBrundallGardens                   = "Brundall Gardens"
	StationNameBrunstane                         = "Brunstane"
	StationNameBrunswick                         = "Brunswick"
	StationNameBruton                            = "Bruton"
	StationNameBryn                              = "Bryn"
	StationNameBuckenhamNorfolk                  = "Buckenham (Norfolk)"
	StationNameBuckley                           = "Buckley"
	StationNameBucknell                          = "Bucknell"
	StationNameBuckshawParkway                   = "Buckshaw Parkway"
	StationNameBugle                             = "Bugle"
	StationNameBuilthRoad                        = "Builth Road"
	StationNameBulwell                           = "Bulwell"
	StationNameBures                             = "Bures"
	StationNameBurgessHill                       = "Burgess Hill"
	StationNameBurleyPark                        = "Burley Park"
	StationNameBurleyinWharfedale                = "Burley-in-Wharfedale"
	StationNameBurnage                           = "Burnage"
	StationNameBurnesideCumbria                  = "Burneside (Cumbria)"
	StationNameBurnhamBucks                      = "Burnham (Bucks)"
	StationNameBurnhamonCrouch                   = "Burnham-on-Crouch"
	StationNameBurnleyBarracks                   = "Burnley Barracks"
	StationNameBurnleyCentral                    = "Burnley Central"
	StationNameBurnleyManchesterRoad             = "Burnley Manchester Road"
	StationNameBurnsideStrathclyde               = "Burnside (Strathclyde)"
	StationNameBurntisland                       = "Burntisland"
	StationNameBurscoughBridge                   = "Burscough Bridge"
	StationNameBurscoughJunction                 = "Burscough Junction"
	StationNameBursledon                         = "Bursledon"
	StationNameBurtonJoyce                       = "Burton Joyce"
	StationNameBurtononTrent                     = "Burton-on-Trent"
	StationNameBuryStEdmunds                     = "Bury St Edmunds"
	StationNameBusby                             = "Busby"
	StationNameBushHillPark                      = "Bush Hill Park"
	StationNameBushey                            = "Bushey"
	StationNameButlersLane                       = "Butlers Lane"
	StationNameBuxted                            = "Buxted"
	StationNameBuxton                            = "Buxton"
	StationNameByfleetAndNewHaw                  = "Byfleet & New Haw"
	StationNameBynea                             = "Bynea"
	StationNameCadoxton                          = "Cadoxton"
	StationNameCaergwrle                         = "Caergwrle"
	StationNameCaerphilly                        = "Caerphilly"
	StationNameCaersws                           = "Caersws"
	StationNameCaldercruix                       = "Caldercruix"
	StationNameCaldicot                          = "Caldicot"
	StationNameCaledonianRdAndBarnsbury          = "Caledonian Rd & Barnsbury"
	StationNameCalstock                          = "Calstock"
	StationNameCamAndDursley                     = "Cam & Dursley"
	StationNameCamberley                         = "Camberley"
	StationNameCamborne                          = "Camborne"
	StationNameCambridge                         = "Cambridge"
	StationNameCambridgeHeath                    = "Cambridge Heath"
	StationNameCambuslang                        = "Cambuslang"
	StationNameCamdenRoad                        = "Camden Road"
	StationNameCamelon                           = "Camelon"
	StationNameCanadaWater                       = "Canada Water"
	StationNameCanley                            = "Canley"
	StationNameCannock                           = "Cannock"
	StationNameCanonbury                         = "Canonbury"
	StationNameCanterburyEast                    = "Canterbury East"
	StationNameCanterburyWest                    = "Canterbury West"
	StationNameCantley                           = "Cantley"
	StationNameCapenhurst                        = "Capenhurst"
	StationNameCarbisBay                         = "Carbis Bay"
	StationNameCardenden                         = "Cardenden"
	StationNameCardiffBay                        = "Cardiff Bay"
	StationNameCardiffCentral                    = "Cardiff Central"
	StationNameCardiffQueenStreet                = "Cardiff Queen Street"
	StationNameCardonald                         = "Cardonald"
	StationNameCardross                          = "Cardross"
	StationNameCarfin                            = "Carfin"
	StationNameCarkAndCartmel                    = "Cark & Cartmel"
	StationNameCarlisle                          = "Carlisle"
	StationNameCarlton                           = "Carlton"
	StationNameCarluke                           = "Carluke"
	StationNameCarmarthen                        = "Carmarthen"
	StationNameCarmyle                           = "Carmyle"
	StationNameCarnforth                         = "Carnforth"
	StationNameCarnoustie                        = "Carnoustie"
	StationNameCarntyne                          = "Carntyne"
	StationNameCarpendersPark                    = "Carpenders Park"
	StationNameCarrbridge                        = "Carrbridge"
	StationNameCarshalton                        = "Carshalton"
	StationNameCarshaltonBeeches                 = "Carshalton Beeches"
	StationNameCarstairs                         = "Carstairs"
	StationNameCartsdyke                         = "Cartsdyke"
	StationNameCastleBarPark                     = "Castle Bar Park"
	StationNameCastleCary                        = "Castle Cary"
	StationNameCastleford                        = "Castleford"
	StationNameCastletonManchester               = "Castleton (Manchester)"
	StationNameCastletonMoor                     = "Castleton Moor"
	StationNameCaterham                          = "Caterham"
	StationNameCatford                           = "Catford"
	StationNameCatfordBridge                     = "Catford Bridge"
	StationNameCathays                           = "Cathays"
	StationNameCathcart                          = "Cathcart"
	StationNameCattal                            = "Cattal"
	StationNameCauseland                         = "Causeland"
	StationNameCefnyBedd                         = "Cefn-y-Bedd"
	StationNameChadwellHeath                     = "Chadwell Heath"
	StationNameChaffordHundredLakeside           = "Chafford Hundred Lakeside"
	StationNameChalfontAndLatimer                = "Chalfont & Latimer"
	StationNameChalkwell                         = "Chalkwell"
	StationNameChandlersFord                     = "Chandlers Ford"
	StationNameChapelenleFrith                   = "Chapel-en-le-Frith"
	StationNameChapeltonDevon                    = "Chapelton (Devon)"
	StationNameChapeltownSouthYorks              = "Chapeltown (South Yorks)"
	StationNameChappelAndWakesColne              = "Chappel & Wakes Colne"
	StationNameCharingCrossGlasgow               = "Charing Cross (Glasgow)"
	StationNameCharingKent                       = "Charing (Kent)"
	StationNameCharlbury                         = "Charlbury"
	StationNameCharlton                          = "Charlton"
	StationNameChartham                          = "Chartham"
	StationNameChassenRoad                       = "Chassen Road"
	StationNameChatelherault                     = "Chatelherault"
	StationNameChatham                           = "Chatham"
	StationNameChathill                          = "Chathill"
	StationNameCheadleHulme                      = "Cheadle Hulme"
	StationNameCheam                             = "Cheam"
	StationNameCheddington                       = "Cheddington"
	StationNameChelfordCheshire                  = "Chelford (Cheshire)"
	StationNameChelmsford                        = "Chelmsford"
	StationNameChelsfield                        = "Chelsfield"
	StationNameCheltenhamSpa                     = "Cheltenham Spa"
	StationNameChepstow                          = "Chepstow"
	StationNameCherryTree                        = "Cherry Tree"
	StationNameChertsey                          = "Chertsey"
	StationNameCheshunt                          = "Cheshunt"
	StationNameChessingtonNorth                  = "Chessington North"
	StationNameChessingtonSouth                  = "Chessington South"
	StationNameChester                           = "Chester"
	StationNameChesterRoad                       = "Chester Road"
	StationNameChesterfield                      = "Chesterfield"
	StationNameChesterleStreet                   = "Chester-le-Street"
	StationNameChestfieldAndSwalecliffe          = "Chestfield & Swalecliffe"
	StationNameChetnole                          = "Chetnole"
	StationNameChichester                        = "Chichester"
	StationNameChilham                           = "Chilham"
	StationNameChilworth                         = "Chilworth"
	StationNameChingford                         = "Chingford"
	StationNameChinley                           = "Chinley"
	StationNameChippenham                        = "Chippenham"
	StationNameChipstead                         = "Chipstead"
	StationNameChirk                             = "Chirk"
	StationNameChislehurst                       = "Chislehurst"
	StationNameChiswick                          = "Chiswick"
	StationNameCholsey                           = "Cholsey"
	StationNameChorley                           = "Chorley"
	StationNameChorleywood                       = "Chorleywood"
	StationNameChristchurch                      = "Christchurch"
	StationNameChristsHospital                   = "Christs Hospital"
	StationNameChurchAndOswaldtwistle            = "Church & Oswaldtwistle"
	StationNameChurchFenton                      = "Church Fenton"
	StationNameChurchStretton                    = "Church Stretton"
	StationNameCilmeri                           = "Cilmeri"
	StationNameCityThameslink                    = "City Thameslink"
	StationNameClactononSea                      = "Clacton-on-Sea"
	StationNameClandon                           = "Clandon"
	StationNameClaphamHighStreet                 = "Clapham High Street"
	StationNameClaphamJunction                   = "Clapham Junction"
	StationNameClaphamNorthYorkshire             = "Clapham (North Yorkshire)"
	StationNameClapton                           = "Clapton"
	StationNameClarbestonRoad                    = "Clarbeston Road"
	StationNameClarkston                         = "Clarkston"
	StationNameClaverdon                         = "Claverdon"
	StationNameClaygate                          = "Claygate"
	StationNameCleethorpes                       = "Cleethorpes"
	StationNameCleland                           = "Cleland"
	StationNameCliftonDown                       = "Clifton Down"
	StationNameCliftonManchester                 = "Clifton (Manchester)"
	StationNameClitheroe                         = "Clitheroe"
	StationNameClockHouse                        = "Clock House"
	StationNameClunderwen                        = "Clunderwen"
	StationNameClydebank                         = "Clydebank"
	StationNameCoatbridgeCentral                 = "Coatbridge Central"
	StationNameCoatbridgeSunnyside               = "Coatbridge Sunnyside"
	StationNameCoatdyke                          = "Coatdyke"
	StationNameCobhamAndStokedAbernon            = "Cobham & Stoke d'Abernon"
	StationNameCodsall                           = "Codsall"
	StationNameCogan                             = "Cogan"
	StationNameColchester                        = "Colchester"
	StationNameColchesterTown                    = "Colchester Town"
	StationNameColeshillParkway                  = "Coleshill Parkway"
	StationNameCollingham                        = "Collingham"
	StationNameCollington                        = "Collington"
	StationNameColne                             = "Colne"
	StationNameColwall                           = "Colwall"
	StationNameColwynBay                         = "Colwyn Bay"
	StationNameCombeOxon                         = "Combe (Oxon)"
	StationNameCommondale                        = "Commondale"
	StationNameCongleton                         = "Congleton"
	StationNameConisbrough                       = "Conisbrough"
	StationNameConnelFerry                       = "Connel Ferry"
	StationNameCononBridge                       = "Conon Bridge"
	StationNameCononley                          = "Cononley"
	StationNameConwayPark                        = "Conway Park"
	StationNameConwy                             = "Conwy"
	StationNameCoodenBeach                       = "Cooden Beach"
	StationNameCookham                           = "Cookham"
	StationNameCooksbridge                       = "Cooksbridge"
	StationNameCoombeJunctionHalt                = "Coombe Junction Halt"
	StationNameCopplestone                       = "Copplestone"
	StationNameCorbridge                         = "Corbridge"
	StationNameCorby                             = "Corby"
	StationNameCorkerhill                        = "Corkerhill"
	StationNameCorkickle                         = "Corkickle"
	StationNameCorpach                           = "Corpach"
	StationNameCorrour                           = "Corrour"
	StationNameCoryton                           = "Coryton"
	StationNameCoseley                           = "Coseley"
	StationNameCosford                           = "Cosford"
	StationNameCosham                            = "Cosham"
	StationNameCottingham                        = "Cottingham"
	StationNameCottingley                        = "Cottingley"
	StationNameCoulsdonSouth                     = "Coulsdon South"
	StationNameCoulsdonTown                      = "Coulsdon Town"
	StationNameCoventry                          = "Coventry"
	StationNameCoventryArena                     = "Coventry Arena"
	StationNameCowdenKent                        = "Cowden (Kent)"
	StationNameCowdenbeath                       = "Cowdenbeath"
	StationNameCradleyHeath                      = "Cradley Heath"
	StationNameCraigendoran                      = "Craigendoran"
	StationNameCramlington                       = "Cramlington"
	StationNameCranbrookDevon                    = "Cranbrook (Devon)"
	StationNameCravenArms                        = "Craven Arms"
	StationNameCrawley                           = "Crawley"
	StationNameCrayford                          = "Crayford"
	StationNameCrediton                          = "Crediton"
	StationNameCressingEssex                     = "Cressing (Essex)"
	StationNameCressington                       = "Cressington"
	StationNameCreswell                          = "Creswell"
	StationNameCrewe                             = "Crewe"
	StationNameCrewkerne                         = "Crewkerne"
	StationNameCrewsHill                         = "Crews Hill"
	StationNameCrianlarich                       = "Crianlarich"
	StationNameCriccieth                         = "Criccieth"
	StationNameCricklewood                       = "Cricklewood"
	StationNameCroftfoot                         = "Croftfoot"
	StationNameCroftonPark                       = "Crofton Park"
	StationNameCromer                            = "Cromer"
	StationNameCromford                          = "Cromford"
	StationNameCrookston                         = "Crookston"
	StationNameCrossGates                        = "Cross Gates"
	StationNameCrossflatts                       = "Crossflatts"
	StationNameCrosshill                         = "Crosshill"
	StationNameCrosskeys                         = "Crosskeys"
	StationNameCrossmyloof                       = "Crossmyloof"
	StationNameCroston                           = "Croston"
	StationNameCrouchHill                        = "Crouch Hill"
	StationNameCrowborough                       = "Crowborough"
	StationNameCrowhurst                         = "Crowhurst"
	StationNameCrowle                            = "Crowle"
	StationNameCrowthorne                        = "Crowthorne"
	StationNameCroy                              = "Croy"
	StationNameCrystalPalace                     = "Crystal Palace"
	StationNameCuddington                        = "Cuddington"
	StationNameCuffley                           = "Cuffley"
	StationNameCulham                            = "Culham"
	StationNameCulrain                           = "Culrain"
	StationNameCumbernauld                       = "Cumbernauld"
	StationNameCupar                             = "Cupar"
	StationNameCurriehill                        = "Curriehill"
	StationNameCuxton                            = "Cuxton"
	StationNameCwmbach                           = "Cwmbach"
	StationNameCwmbran                           = "Cwmbran"
	StationNameCynghordy                         = "Cynghordy"
	StationNameDagenhamDock                      = "Dagenham Dock"
	StationNameDaisyHill                         = "Daisy Hill"
	StationNameDalgetyBay                        = "Dalgety Bay"
	StationNameDalmally                          = "Dalmally"
	StationNameDalmarnock                        = "Dalmarnock"
	StationNameDalmeny                           = "Dalmeny"
	StationNameDalmuir                           = "Dalmuir"
	StationNameDalreoch                          = "Dalreoch"
	StationNameDalry                             = "Dalry"
	StationNameDalstonCumbria                    = "Dalston (Cumbria)"
	StationNameDalstonJunction                   = "Dalston Junction"
	StationNameDalstonKingsland                  = "Dalston Kingsland"
	StationNameDaltonCumbria                     = "Dalton (Cumbria)"
	StationNameDalwhinnie                        = "Dalwhinnie"
	StationNameDanby                             = "Danby"
	StationNameDanescourt                        = "Danescourt"
	StationNameDanzey                            = "Danzey"
	StationNameDarlington                        = "Darlington"
	StationNameDarnall                           = "Darnall"
	StationNameDarsham                           = "Darsham"
	StationNameDartford                          = "Dartford"
	StationNameDarton                            = "Darton"
	StationNameDarwen                            = "Darwen"
	StationNameDatchet                           = "Datchet"
	StationNameDavenport                         = "Davenport"
	StationNameDawlish                           = "Dawlish"
	StationNameDawlishWarren                     = "Dawlish Warren"
	StationNameDeal                              = "Deal"
	StationNameDeanWilts                         = "Dean (Wilts)"
	StationNameDeansgate                         = "Deansgate"
	StationNameDeganwy                           = "Deganwy"
	StationNameDeighton                          = "Deighton"
	StationNameDelamere                          = "Delamere"
	StationNameDenbyDale                         = "Denby Dale"
	StationNameDenham                            = "Denham"
	StationNameDenhamGolfClub                    = "Denham Golf Club"
	StationNameDenmarkHill                       = "Denmark Hill"
	StationNameDent                              = "Dent"
	StationNameDenton                            = "Denton"
	StationNameDeptford                          = "Deptford"
	StationNameDerby                             = "Derby"
	StationNameDerbyRoadIpswich                  = "Derby Road (Ipswich)"
	StationNameDevonportDevon                    = "Devonport (Devon)"
	StationNameDevonportDockyard                 = "Devonport Dockyard"
	StationNameDewsbury                          = "Dewsbury"
	StationNameDidcotParkway                     = "Didcot Parkway"
	StationNameDigbyAndSowton                    = "Digby & Sowton"
	StationNameDiltonMarsh                       = "Dilton Marsh"
	StationNameDinasPowys                        = "Dinas Powys"
	StationNameDinasRhondda                      = "Dinas (Rhondda)"
	StationNameDingleRoad                        = "Dingle Road"
	StationNameDingwall                          = "Dingwall"
	StationNameDinsdale                          = "Dinsdale"
	StationNameDinting                           = "Dinting"
	StationNameDisley                            = "Disley"
	StationNameDiss                              = "Diss"
	StationNameDodworth                          = "Dodworth"
	StationNameDolau                             = "Dolau"
	StationNameDoleham                           = "Doleham"
	StationNameDolgarrog                         = "Dolgarrog"
	StationNameDolwyddelan                       = "Dolwyddelan"
	StationNameDoncaster                         = "Doncaster"
	StationNameDorchesterSouth                   = "Dorchester South"
	StationNameDorchesterWest                    = "Dorchester West"
	StationNameDoreAndTotley                     = "Dore & Totley"
	StationNameDorkingDeepdene                   = "Dorking Deepdene"
	StationNameDorkingMain                       = "Dorking (Main)"
	StationNameDorkingWest                       = "Dorking West"
	StationNameDormans                           = "Dormans"
	StationNameDorridge                          = "Dorridge"
	StationNameDoveHoles                         = "Dove Holes"
	StationNameDoverPriory                       = "Dover Priory"
	StationNameDovercourt                        = "Dovercourt"
	StationNameDoveyJunction                     = "Dovey Junction"
	StationNameDownhamMarket                     = "Downham Market"
	StationNameDraytonGreen                      = "Drayton Green"
	StationNameDraytonPark                       = "Drayton Park"
	StationNameDrem                              = "Drem"
	StationNameDriffield                         = "Driffield"
	StationNameDrigg                             = "Drigg"
	StationNameDroitwichSpa                      = "Droitwich Spa"
	StationNameDronfield                         = "Dronfield"
	StationNameDrumchapel                        = "Drumchapel"
	StationNameDrumfrochar                       = "Drumfrochar"
	StationNameDrumgelloch                       = "Drumgelloch"
	StationNameDrumry                            = "Drumry"
	StationNameDublinFerryport                   = "Dublin Ferryport"
	StationNameDublinPortStena                   = "Dublin Port - Stena"
	StationNameDuddeston                         = "Duddeston"
	StationNameDudleyPort                        = "Dudley Port"
	StationNameDuffield                          = "Duffield"
	StationNameDuirinish                         = "Duirinish"
	StationNameDukeStreet                        = "Duke Street"
	StationNameDullingham                        = "Dullingham"
	StationNameDumbartonCentral                  = "Dumbarton Central"
	StationNameDumbartonEast                     = "Dumbarton East"
	StationNameDumbreck                          = "Dumbreck"
	StationNameDumfries                          = "Dumfries"
	StationNameDumptonPark                       = "Dumpton Park"
	StationNameDunbar                            = "Dunbar"
	StationNameDunblane                          = "Dunblane"
	StationNameDuncraig                          = "Duncraig"
	StationNameDundee                            = "Dundee"
	StationNameDunfermlineQueenMargaret          = "Dunfermline Queen Margaret"
	StationNameDunfermlineTown                   = "Dunfermline Town"
	StationNameDunkeldAndBirnam                  = "Dunkeld & Birnam"
	StationNameDunlop                            = "Dunlop"
	StationNameDunrobinCastle                    = "Dunrobin Castle"
	StationNameDunston                           = "Dunston"
	StationNameDuntonGreen                       = "Dunton Green"
	StationNameDurham                            = "Durham"
	StationNameDurringtononSea                   = "Durrington-on-Sea"
	StationNameDyce                              = "Dyce"
	StationNameDyffrynArdudwy                    = "Dyffryn Ardudwy"
	StationNameEaglescliffe                      = "Eaglescliffe"
	StationNameEalingBroadway                    = "Ealing Broadway"
	StationNameEarlestown                        = "Earlestown"
	StationNameEarley                            = "Earley"
	StationNameEarlsfield                        = "Earlsfield"
	StationNameEarlswoodSurrey                   = "Earlswood (Surrey)"
	StationNameEarlswoodWestMidlands             = "Earlswood (West Midlands)"
	StationNameEastCroydon                       = "East Croydon"
	StationNameEastDidsbury                      = "East Didsbury"
	StationNameEastDulwich                       = "East Dulwich"
	StationNameEastFarleigh                      = "East Farleigh"
	StationNameEastGarforth                      = "East Garforth"
	StationNameEastGrinstead                     = "East Grinstead"
	StationNameEastKilbride                      = "East Kilbride"
	StationNameEastMalling                       = "East Malling"
	StationNameEastMidlandsParkway               = "East Midlands Parkway"
	StationNameEastTilbury                       = "East Tilbury"
	StationNameEastWorthing                      = "East Worthing"
	StationNameEastbourne                        = "Eastbourne"
	StationNameEastbrook                         = "Eastbrook"
	StationNameEasterhouse                       = "Easterhouse"
	StationNameEasthamRake                       = "Eastham Rake"
	StationNameEastleigh                         = "Eastleigh"
	StationNameEastrington                       = "Eastrington"
	StationNameEbbsfleetInternational            = "Ebbsfleet International"
	StationNameEbbwValeParkway                   = "Ebbw Vale Parkway"
	StationNameEbbwValeTown                      = "Ebbw Vale Town"
	StationNameEcclesManchester                  = "Eccles (Manchester)"
	StationNameEcclesRoad                        = "Eccles Road"
	StationNameEcclestonPark                     = "Eccleston Park"
	StationNameEdale                             = "Edale"
	StationNameEdenPark                          = "Eden Park"
	StationNameEdenbridge                        = "Edenbridge"
	StationNameEdenbridgeTown                    = "Edenbridge Town"
	StationNameEdgeHill                          = "Edge Hill"
	StationNameEdinburgh                         = "Edinburgh"
	StationNameEdinburghGateway                  = "Edinburgh Gateway"
	StationNameEdinburghPark                     = "Edinburgh Park"
	StationNameEdmontonGreen                     = "Edmonton Green"
	StationNameEffinghamJunction                 = "Effingham Junction"
	StationNameEggesford                         = "Eggesford"
	StationNameEgham                             = "Egham"
	StationNameEgton                             = "Egton"
	StationNameElephantAndCastle                 = "Elephant & Castle"
	StationNameElephantAndCastleUnderground      = "Elephant & Castle (Underground)"
	StationNameElgin                             = "Elgin"
	StationNameEllesmerePort                     = "Ellesmere Port"
	StationNameElmersEnd                         = "Elmers End"
	StationNameElmsteadWoods                     = "Elmstead Woods"
	StationNameElmswell                          = "Elmswell"
	StationNameElsecar                           = "Elsecar"
	StationNameElsenhamEssex                     = "Elsenham (Essex)"
	StationNameElstreeAndBorehamwood             = "Elstree & Borehamwood"
	StationNameEltham                            = "Eltham"
	StationNameEltonAndOrston                    = "Elton & Orston"
	StationNameEly                               = "Ely"
	StationNameEmersonPark                       = "Emerson Park"
	StationNameEmsworth                          = "Emsworth"
	StationNameEnerglynAndChurchillPark          = "Energlyn & Churchill Park"
	StationNameEnfieldChase                      = "Enfield Chase"
	StationNameEnfieldLock                       = "Enfield Lock"
	StationNameEnfieldTown                       = "Enfield Town"
	StationNameEntwistle                         = "Entwistle"
	StationNameEpsomDowns                        = "Epsom Downs"
	StationNameEpsomSurrey                       = "Epsom (Surrey)"
	StationNameErdington                         = "Erdington"
	StationNameEridge                            = "Eridge"
	StationNameErith                             = "Erith"
	StationNameEsher                             = "Esher"
	StationNameEskbank                           = "Eskbank"
	StationNameEssexRoad                         = "Essex Road"
	StationNameEtchingham                        = "Etchingham"
	StationNameEuxtonBalshawLane                 = "Euxton Balshaw Lane"
	StationNameEvesham                           = "Evesham"
	StationNameEwellEast                         = "Ewell East"
	StationNameEwellWest                         = "Ewell West"
	StationNameExeterCentral                     = "Exeter Central"
	StationNameExeterStDavids                    = "Exeter St David's"
	StationNameExeterStThomas                    = "Exeter St Thomas"
	StationNameExhibitionCentreGlasgow           = "Exhibition Centre (Glasgow)"
	StationNameExmouth                           = "Exmouth"
	StationNameExton                             = "Exton"
	StationNameEynsford                          = "Eynsford"
	StationNameFairbourne                        = "Fairbourne"
	StationNameFairfield                         = "Fairfield"
	StationNameFairlie                           = "Fairlie"
	StationNameFairwater                         = "Fairwater"
	StationNameFalconwood                        = "Falconwood"
	StationNameFalkirkGrahamston                 = "Falkirk Grahamston"
	StationNameFalkirkHigh                       = "Falkirk High"
	StationNameFallsofCruachan                   = "Falls of Cruachan"
	StationNameFalmer                            = "Falmer"
	StationNameFalmouthDocks                     = "Falmouth Docks"
	StationNameFalmouthTown                      = "Falmouth Town"
	StationNameFareham                           = "Fareham"
	StationNameFarnboroughMain                   = "Farnborough (Main)"
	StationNameFarnboroughNorth                  = "Farnborough North"
	StationNameFarncombe                         = "Farncombe"
	StationNameFarnham                           = "Farnham"
	StationNameFarninghamRoad                    = "Farningham Road"
	StationNameFarnworth                         = "Farnworth"
	StationNameFarringdon                        = "Farringdon"
	StationNameFauldhouse                        = "Fauldhouse"
	StationNameFaversham                         = "Faversham"
	StationNameFaygate                           = "Faygate"
	StationNameFazakerley                        = "Fazakerley"
	StationNameFearn                             = "Fearn"
	StationNameFeatherstone                      = "Featherstone"
	StationNameFelixstowe                        = "Felixstowe"
	StationNameFeltham                           = "Feltham"
	StationNameFeniton                           = "Feniton"
	StationNameFennyStratford                    = "Fenny Stratford"
	StationNameFernhill                          = "Fernhill"
	StationNameFerriby                           = "Ferriby"
	StationNameFerryside                         = "Ferryside"
	StationNameFfairfach                         = "Ffairfach"
	StationNameFiley                             = "Filey"
	StationNameFiltonAbbeyWood                   = "Filton Abbey Wood"
	StationNameFinchleyRoadAndFrognal            = "Finchley Road & Frognal"
	StationNameFinsburyPark                      = "Finsbury Park"
	StationNameFinstock                          = "Finstock"
	StationNameFishbourneSussex                  = "Fishbourne (Sussex)"
	StationNameFishersgate                       = "Fishersgate"
	StationNameFishguardAndGoodwick              = "Fishguard & Goodwick"
	StationNameFishguardHarbour                  = "Fishguard Harbour"
	StationNameFiskerton                         = "Fiskerton"
	StationNameFitzwilliam                       = "Fitzwilliam"
	StationNameFiveWays                          = "Five Ways"
	StationNameFleet                             = "Fleet"
	StationNameFlimby                            = "Flimby"
	StationNameFlint                             = "Flint"
	StationNameFlitwick                          = "Flitwick"
	StationNameFlixton                           = "Flixton"
	StationNameFloweryField                      = "Flowery Field"
	StationNameFolkestoneCentral                 = "Folkestone Central"
	StationNameFolkestoneWest                    = "Folkestone West"
	StationNameFord                              = "Ford"
	StationNameForestGate                        = "Forest Gate"
	StationNameForestHill                        = "Forest Hill"
	StationNameFormby                            = "Formby"
	StationNameForres                            = "Forres"
	StationNameForsinard                         = "Forsinard"
	StationNameFortMatilda                       = "Fort Matilda"
	StationNameFortWilliam                       = "Fort William"
	StationNameFourOaks                          = "Four Oaks"
	StationNameFoxfield                          = "Foxfield"
	StationNameFoxton                            = "Foxton"
	StationNameFrant                             = "Frant"
	StationNameFratton                           = "Fratton"
	StationNameFreshfield                        = "Freshfield"
	StationNameFreshford                         = "Freshford"
	StationNameFrimley                           = "Frimley"
	StationNameFrintononSea                      = "Frinton-on-Sea"
	StationNameFrizinghall                       = "Frizinghall"
	StationNameFrodsham                          = "Frodsham"
	StationNameFrome                             = "Frome"
	StationNameFulwell                           = "Fulwell"
	StationNameFurnessVale                       = "Furness Vale"
	StationNameFurzePlatt                        = "Furze Platt"
	StationNameGainsboroughCentral               = "Gainsborough Central"
	StationNameGainsboroughLeaRoad               = "Gainsborough Lea Road"
	StationNameGalashiels                        = "Galashiels"
	StationNameGarelochhead                      = "Garelochhead"
	StationNameGarforth                          = "Garforth"
	StationNameGargrave                          = "Gargrave"
	StationNameGarrowhill                        = "Garrowhill"
	StationNameGarscadden                        = "Garscadden"
	StationNameGarsdale                          = "Garsdale"
	StationNameGarstonHertfordshire              = "Garston (Hertfordshire)"
	StationNameGarswood                          = "Garswood"
	StationNameGartcosh                          = "Gartcosh"
	StationNameGarthMidGlamorgan                 = "Garth (Mid Glamorgan)"
	StationNameGarthPowys                        = "Garth (Powys)"
	StationNameGarve                             = "Garve"
	StationNameGathurst                          = "Gathurst"
	StationNameGatley                            = "Gatley"
	StationNameGatwickAirport                    = "Gatwick Airport"
	StationNameGeorgemasJunction                 = "Georgemas Junction"
	StationNameGerrardsCross                     = "Gerrards Cross"
	StationNameGideaPark                         = "Gidea Park"
	StationNameGiffnock                          = "Giffnock"
	StationNameGiggleswick                       = "Giggleswick"
	StationNameGilberdyke                        = "Gilberdyke"
	StationNameGilfachFargoed                    = "Gilfach Fargoed"
	StationNameGillinghamDorset                  = "Gillingham (Dorset)"
	StationNameGillinghamKent                    = "Gillingham (Kent)"
	StationNameGilshochill                       = "Gilshochill"
	StationNameGipsyHill                         = "Gipsy Hill"
	StationNameGirvan                            = "Girvan"
	StationNameGlaisdale                         = "Glaisdale"
	StationNameGlanConwy                         = "Glan Conwy"
	StationNameGlasgowCentral                    = "Glasgow Central"
	StationNameGlasgowQueenStreet                = "Glasgow Queen Street"
	StationNameGlasshoughton                     = "Glasshoughton"
	StationNameGlazebrook                        = "Glazebrook"
	StationNameGleneagles                        = "Gleneagles"
	StationNameGlenfinnan                        = "Glenfinnan"
	StationNameGlengarnock                       = "Glengarnock"
	StationNameGlenrotheswithThornton            = "Glenrothes with Thornton"
	StationNameGlossop                           = "Glossop"
	StationNameGloucester                        = "Gloucester"
	StationNameGlynde                            = "Glynde"
	StationNameGobowen                           = "Gobowen"
	StationNameGodalming                         = "Godalming"
	StationNameGodley                            = "Godley"
	StationNameGodstone                          = "Godstone"
	StationNameGoldthorpe                        = "Goldthorpe"
	StationNameGolfStreet                        = "Golf Street"
	StationNameGolspie                           = "Golspie"
	StationNameGomshall                          = "Gomshall"
	StationNameGoodmayes                         = "Goodmayes"
	StationNameGoole                             = "Goole"
	StationNameGoostrey                          = "Goostrey"
	StationNameGordonHill                        = "Gordon Hill"
	StationNameGorebridge                        = "Gorebridge"
	StationNameGoringAndStreatley                = "Goring & Streatley"
	StationNameGoringbySea                       = "Goring-by-Sea"
	StationNameGorton                            = "Gorton"
	StationNameGospelOak                         = "Gospel Oak"
	StationNameGourock                           = "Gourock"
	StationNameGowerton                          = "Gowerton"
	StationNameGoxhill                           = "Goxhill"
	StationNameGrangeOverSands                   = "Grange-Over-Sands"
	StationNameGrangePark                        = "Grange Park"
	StationNameGrangetownCardiff                 = "Grangetown (Cardiff)"
	StationNameGrantham                          = "Grantham"
	StationNameGrateley                          = "Grateley"
	StationNameGravellyHill                      = "Gravelly Hill"
	StationNameGravesend                         = "Gravesend"
	StationNameGrays                             = "Grays"
	StationNameGreatAyton                        = "Great Ayton"
	StationNameGreatBentley                      = "Great Bentley"
	StationNameGreatChesterford                  = "Great Chesterford"
	StationNameGreatCoates                       = "Great Coates"
	StationNameGreatMalvern                      = "Great Malvern"
	StationNameGreatMissenden                    = "Great Missenden"
	StationNameGreatYarmouth                     = "Great Yarmouth"
	StationNameGreenLane                         = "Green Lane"
	StationNameGreenRoad                         = "Green Road"
	StationNameGreenbank                         = "Greenbank"
	StationNameGreenfaulds                       = "Greenfaulds"
	StationNameGreenfield                        = "Greenfield"
	StationNameGreenford                         = "Greenford"
	StationNameGreenhitheForBluewater            = "Greenhithe for Bluewater"
	StationNameGreenockCentral                   = "Greenock Central"
	StationNameGreenockWest                      = "Greenock West"
	StationNameGreenwich                         = "Greenwich"
	StationNameGretnaGreen                       = "Gretna Green"
	StationNameGrimsbyDocks                      = "Grimsby Docks"
	StationNameGrimsbyTown                       = "Grimsby Town"
	StationNameGrindleford                       = "Grindleford"
	StationNameGrosmont                          = "Grosmont"
	StationNameGrovePark                         = "Grove Park"
	StationNameGuideBridge                       = "Guide Bridge"
	StationNameGuildford                         = "Guildford"
	StationNameGuiseley                          = "Guiseley"
	StationNameGunnersbury                       = "Gunnersbury"
	StationNameGunnislake                        = "Gunnislake"
	StationNameGunton                            = "Gunton"
	StationNameGwersyllt                         = "Gwersyllt"
	StationNameGypsyLane                         = "Gypsy Lane"
	StationNameHabrough                          = "Habrough"
	StationNameHackbridge                        = "Hackbridge"
	StationNameHackneyCentral                    = "Hackney Central"
	StationNameHackneyDowns                      = "Hackney Downs"
	StationNameHackneyWick                       = "Hackney Wick"
	StationNameHaddenhamAndThameParkway          = "Haddenham & Thame Parkway"
	StationNameHaddiscoe                         = "Haddiscoe"
	StationNameHadfield                          = "Hadfield"
	StationNameHadleyWood                        = "Hadley Wood"
	StationNameHagFold                           = "Hag Fold"
	StationNameHaggerston                        = "Haggerston"
	StationNameHagley                            = "Hagley"
	StationNameHairmyres                         = "Hairmyres"
	StationNameHaleManchester                    = "Hale (Manchester)"
	StationNameHalesworth                        = "Halesworth"
	StationNameHalewood                          = "Halewood"
	StationNameHalifax                           = "Halifax"
	StationNameHallGreen                         = "Hall Green"
	StationNameHallRoad                          = "Hall Road"
	StationNameHalling                           = "Halling"
	StationNameHallithWood                       = "Hall-i'-th'-Wood"
	StationNameHaltwhistle                       = "Haltwhistle"
	StationNameHamStreet                         = "Ham Street"
	StationNameHamble                            = "Hamble"
	StationNameHamiltonCentral                   = "Hamilton Central"
	StationNameHamiltonWest                      = "Hamilton West"
	StationNameHammerton                         = "Hammerton"
	StationNameHampdenParkSussex                 = "Hampden Park (Sussex)"
	StationNameHampsteadHeath                    = "Hampstead Heath"
	StationNameHamptonCourt                      = "Hampton Court"
	StationNameHamptonLondon                     = "Hampton (London)"
	StationNameHamptonWick                       = "Hampton Wick"
	StationNameHamptoninArden                    = "Hampton-in-Arden"
	StationNameHamsteadBirmingham                = "Hamstead (Birmingham)"
	StationNameHamworthy                         = "Hamworthy"
	StationNameHanborough                        = "Hanborough"
	StationNameHandforth                         = "Handforth"
	StationNameHanwell                           = "Hanwell"
	StationNameHapton                            = "Hapton"
	StationNameHarlech                           = "Harlech"
	StationNameHarlesden                         = "Harlesden"
	StationNameHarlingRoad                       = "Harling Road"
	StationNameHarlingtonBeds                    = "Harlington (Beds)"
	StationNameHarlowMill                        = "Harlow Mill"
	StationNameHarlowTown                        = "Harlow Town"
	StationNameHaroldWood                        = "Harold Wood"
	StationNameHarpenden                         = "Harpenden"
	StationNameHarrietsham                       = "Harrietsham"
	StationNameHarringay                         = "Harringay"
	StationNameHarringayGreenLanes               = "Harringay Green Lanes"
	StationNameHarrington                        = "Harrington"
	StationNameHarrogate                         = "Harrogate"
	StationNameHarrowAndWealdstone               = "Harrow & Wealdstone"
	StationNameHarrowontheHill                   = "Harrow-on-the-Hill"
	StationNameHartfordCheshire                  = "Hartford (Cheshire)"
	StationNameHartlebury                        = "Hartlebury"
	StationNameHartlepool                        = "Hartlepool"
	StationNameHartwood                          = "Hartwood"
	StationNameHarwichInternational              = "Harwich International"
	StationNameHarwichTown                       = "Harwich Town"
	StationNameHaslemere                         = "Haslemere"
	StationNameHassocks                          = "Hassocks"
	StationNameHastings                          = "Hastings"
	StationNameHatchEnd                          = "Hatch End"
	StationNameHatfieldAndStainforth             = "Hatfield & Stainforth"
	StationNameHatfieldHerts                     = "Hatfield (Herts)"
	StationNameHatfieldPeverel                   = "Hatfield Peverel"
	StationNameHathersage                        = "Hathersage"
	StationNameHattersley                        = "Hattersley"
	StationNameHatton                            = "Hatton"
	StationNameHavant                            = "Havant"
	StationNameHavenhouse                        = "Havenhouse"
	StationNameHaverfordwest                     = "Haverfordwest"
	StationNameHawarden                          = "Hawarden"
	StationNameHawardenBridge                    = "Hawarden Bridge"
	StationNameHawkhead                          = "Hawkhead"
	StationNameHaydonBridge                      = "Haydon Bridge"
	StationNameHaydonsRoad                       = "Haydons Road"
	StationNameHayesAndHarlington                = "Hayes & Harlington"
	StationNameHayesKent                         = "Hayes (Kent)"
	StationNameHayle                             = "Hayle"
	StationNameHaymarket                         = "Haymarket"
	StationNameHaywardsHeath                     = "Haywards Heath"
	StationNameHazelGrove                        = "Hazel Grove"
	StationNameHeadcorn                          = "Headcorn"
	StationNameHeadingley                        = "Headingley"
	StationNameHeadstoneLane                     = "Headstone Lane"
	StationNameHealdGreen                        = "Heald Green"
	StationNameHealing                           = "Healing"
	StationNameHeathHighLevel                    = "Heath High Level"
	StationNameHeathLowLevel                     = "Heath Low Level"
	StationNameHeathrowAirportTerminal4          = "Heathrow Airport Terminal 4"
	StationNameHeathrowAirportTerminal5          = "Heathrow Airport Terminal 5"
	StationNameHeathrowAirportTerminals12and3    = "Heathrow Airport Terminals 1, 2 and 3"
	StationNameHeatonChapel                      = "Heaton Chapel"
	StationNameHebdenBridge                      = "Hebden Bridge"
	StationNameHeckington                        = "Heckington"
	StationNameHedgeEnd                          = "Hedge End"
	StationNameHednesford                        = "Hednesford"
	StationNameHeighington                       = "Heighington"
	StationNameHelensburghCentral                = "Helensburgh Central"
	StationNameHelensburghUpper                  = "Helensburgh Upper"
	StationNameHellifield                        = "Hellifield"
	StationNameHelmsdale                         = "Helmsdale"
	StationNameHelsby                            = "Helsby"
	StationNameHemelHempstead                    = "Hemel Hempstead"
	StationNameHendon                            = "Hendon"
	StationNameHengoed                           = "Hengoed"
	StationNameHenleyinArden                     = "Henley-in-Arden"
	StationNameHenleyonThames                    = "Henley-on-Thames"
	StationNameHensall                           = "Hensall"
	StationNameHereford                          = "Hereford"
	StationNameHerneBay                          = "Herne Bay"
	StationNameHerneHill                         = "Herne Hill"
	StationNameHersham                           = "Hersham"
	StationNameHertfordEast                      = "Hertford East"
	StationNameHertfordNorth                     = "Hertford North"
	StationNameHessle                            = "Hessle"
	StationNameHeswall                           = "Heswall"
	StationNameHever                             = "Hever"
	StationNameHeworth                           = "Heworth"
	StationNameHexham                            = "Hexham"
	StationNameHeyford                           = "Heyford"
	StationNameHeyshamPort                       = "Heysham Port"
	StationNameHighBrooms                        = "High Brooms"
	StationNameHighStreetGlasgow                 = "High Street (Glasgow)"
	StationNameHighStreetKensingtonUnderground   = "High Street Kensington Underground"
	StationNameHighWycombe                       = "High Wycombe"
	StationNameHigham                            = "Higham"
	StationNameHighamsPark                       = "Highams Park"
	StationNameHighbridgeAndBurnham              = "Highbridge & Burnham"
	StationNameHighburyAndIslington              = "Highbury & Islington"
	StationNameHightown                          = "Hightown"
	StationNameHildenborough                     = "Hildenborough"
	StationNameHillfoot                          = "Hillfoot"
	StationNameHillingtonEast                    = "Hillington East"
	StationNameHillingtonWest                    = "Hillington West"
	StationNameHillside                          = "Hillside"
	StationNameHilsea                            = "Hilsea"
	StationNameHinchleyWood                      = "Hinchley Wood"
	StationNameHinckleyLeics                     = "Hinckley (Leics)"
	StationNameHindley                           = "Hindley"
	StationNameHintonAdmiral                     = "Hinton Admiral"
	StationNameHitchin                           = "Hitchin"
	StationNameHitherGreen                       = "Hither Green"
	StationNameHockley                           = "Hockley"
	StationNameHollingbourne                     = "Hollingbourne"
	StationNameHolmesChapel                      = "Holmes Chapel"
	StationNameHolmwood                          = "Holmwood"
	StationNameHoltonHeath                       = "Holton Heath"
	StationNameHolyhead                          = "Holyhead"
	StationNameHolytown                          = "Holytown"
	StationNameHomerton                          = "Homerton"
	StationNameHoneybourne                       = "Honeybourne"
	StationNameHoniton                           = "Honiton"
	StationNameHonley                            = "Honley"
	StationNameHonorOakPark                      = "Honor Oak Park"
	StationNameHook                              = "Hook"
	StationNameHooton                            = "Hooton"
	StationNameHopeDerbyshire                    = "Hope (Derbyshire)"
	StationNameHopeFlintshire                    = "Hope (Flintshire)"
	StationNameHoptonHeath                       = "Hopton Heath"
	StationNameHorley                            = "Horley"
	StationNameHornbeamPark                      = "Hornbeam Park"
	StationNameHornsey                           = "Hornsey"
	StationNameHorsforth                         = "Horsforth"
	StationNameHorsham                           = "Horsham"
	StationNameHorsley                           = "Horsley"
	StationNameHortoninRibblesdale               = "Horton-in-Ribblesdale"
	StationNameHorwichParkway                    = "Horwich Parkway"
	StationNameHoscar                            = "Hoscar"
	StationNameHoughGreen                        = "Hough Green"
	StationNameHounslow                          = "Hounslow"
	StationNameHove                              = "Hove"
	StationNameHovetonAndWroxham                 = "Hoveton & Wroxham"
	StationNameHowWoodHerts                      = "How Wood (Herts)"
	StationNameHowden                            = "Howden"
	StationNameHowwoodRenfrewshire               = "Howwood (Renfrewshire)"
	StationNameHoxton                            = "Hoxton"
	StationNameHoylake                           = "Hoylake"
	StationNameHubbertsBridge                    = "Hubberts Bridge"
	StationNameHucknall                          = "Hucknall"
	StationNameHuddersfield                      = "Huddersfield"
	StationNameHull                              = "Hull"
	StationNameHumphreyPark                      = "Humphrey Park"
	StationNameHuncoat                           = "Huncoat"
	StationNameHungerford                        = "Hungerford"
	StationNameHunmanby                          = "Hunmanby"
	StationNameHuntingdon                        = "Huntingdon"
	StationNameHuntly                            = "Huntly"
	StationNameHuntsCross                        = "Hunts Cross"
	StationNameHurstGreen                        = "Hurst Green"
	StationNameHuttonCranswick                   = "Hutton Cranswick"
	StationNameHuyton                            = "Huyton"
	StationNameHydeCentral                       = "Hyde Central"
	StationNameHydeNorth                         = "Hyde North"
	StationNameHykeham                           = "Hykeham"
	StationNameHyndland                          = "Hyndland"
	StationNameHytheEssex                        = "Hythe (Essex)"
	StationNameIBMHalt                           = "IBM Halt"
	StationNameIfield                            = "Ifield"
	StationNameIlford                            = "Ilford"
	StationNameIlkley                            = "Ilkley"
	StationNameImperialWharf                     = "Imperial Wharf"
	StationNameInceAndElton                      = "Ince & Elton"
	StationNameInceManchester                    = "Ince (Manchester)"
	StationNameIngatestone                       = "Ingatestone"
	StationNameInsch                             = "Insch"
	StationNameInvergordon                       = "Invergordon"
	StationNameInvergowrie                       = "Invergowrie"
	StationNameInverkeithing                     = "Inverkeithing"
	StationNameInverkip                          = "Inverkip"
	StationNameInverness                         = "Inverness"
	StationNameInvershin                         = "Invershin"
	StationNameInverurie                         = "Inverurie"
	StationNameIpswich                           = "Ipswich"
	StationNameIrlam                             = "Irlam"
	StationNameIrvine                            = "Irvine"
	StationNameIsleworth                         = "Isleworth"
	StationNameIslip                             = "Islip"
	StationNameIver                              = "Iver"
	StationNameIvybridge                         = "Ivybridge"
	StationNameJamesCook                         = "James Cook"
	StationNameJewelleryQuarter                  = "Jewellery Quarter"
	StationNameJohnstonPembs                     = "Johnston (Pembs)"
	StationNameJohnstoneStrathclyde              = "Johnstone (Strathclyde)"
	StationNameJordanhill                        = "Jordanhill"
	StationNameKearsleyManchester                = "Kearsley (Manchester)"
	StationNameKearsney                          = "Kearsney"
	StationNameKeighley                          = "Keighley"
	StationNameKeith                             = "Keith"
	StationNameKelvedon                          = "Kelvedon"
	StationNameKelvindale                        = "Kelvindale"
	StationNameKemble                            = "Kemble"
	StationNameKempstonHardwick                  = "Kempston Hardwick"
	StationNameKemptonParkRacecourse             = "Kempton Park Racecourse"
	StationNameKemsing                           = "Kemsing"
	StationNameKemsley                           = "Kemsley"
	StationNameKendal                            = "Kendal"
	StationNameKenley                            = "Kenley"
	StationNameKennett                           = "Kennett"
	StationNameKennishead                        = "Kennishead"
	StationNameKensalGreen                       = "Kensal Green"
	StationNameKensalRise                        = "Kensal Rise"
	StationNameKensingtonOlympia                 = "Kensington Olympia"
	StationNameKentHouse                         = "Kent House"
	StationNameKentishTown                       = "Kentish Town"
	StationNameKentishTownWest                   = "Kentish Town West"
	StationNameKenton                            = "Kenton"
	StationNameKentsBank                         = "Kents Bank"
	StationNameKettering                         = "Kettering"
	StationNameKewBridge                         = "Kew Bridge"
	StationNameKewGardens                        = "Kew Gardens"
	StationNameKeyham                            = "Keyham"
	StationNameKeynsham                          = "Keynsham"
	StationNameKidbrooke                         = "Kidbrooke"
	StationNameKidderminster                     = "Kidderminster"
	StationNameKidsgrove                         = "Kidsgrove"
	StationNameKidwelly                          = "Kidwelly"
	StationNameKilburnHighRoad                   = "Kilburn High Road"
	StationNameKildale                           = "Kildale"
	StationNameKildonan                          = "Kildonan"
	StationNameKilgetty                          = "Kilgetty"
	StationNameKilmarnock                        = "Kilmarnock"
	StationNameKilmaurs                          = "Kilmaurs"
	StationNameKilpatrick                        = "Kilpatrick"
	StationNameKilwinning                        = "Kilwinning"
	StationNameKinbrace                          = "Kinbrace"
	StationNameKingham                           = "Kingham"
	StationNameKinghorn                          = "Kinghorn"
	StationNameKingsLangley                      = "Kings Langley"
	StationNameKingsLynn                         = "Kings Lynn"
	StationNameKingsNorton                       = "Kings Norton"
	StationNameKingsNympton                      = "Kings Nympton"
	StationNameKingsPark                         = "Kings Park"
	StationNameKingsSutton                       = "Kings Sutton"
	StationNameKingsknowe                        = "Kingsknowe"
	StationNameKingston                          = "Kingston"
	StationNameKingswood                         = "Kingswood"
	StationNameKingussie                         = "Kingussie"
	StationNameKintbury                          = "Kintbury"
	StationNameKirbyCross                        = "Kirby Cross"
	StationNameKirkSandall                       = "Kirk Sandall"
	StationNameKirkbyMerseyside                  = "Kirkby (Merseyside)"
	StationNameKirkbyStephen                     = "Kirkby Stephen"
	StationNameKirkbyinAshfield                  = "Kirkby-in-Ashfield"
	StationNameKirkbyinFurness                   = "Kirkby-in-Furness"
	StationNameKirkcaldy                         = "Kirkcaldy"
	StationNameKirkconnel                        = "Kirkconnel"
	StationNameKirkdale                          = "Kirkdale"
	StationNameKirkhamAndWesham                  = "Kirkham & Wesham"
	StationNameKirkhill                          = "Kirkhill"
	StationNameKirknewton                        = "Kirknewton"
	StationNameKirkstallForge                    = "Kirkstall Forge"
	StationNameKirkwood                          = "Kirkwood"
	StationNameKirtonLindsey                     = "Kirton Lindsey"
	StationNameKivetonBridge                     = "Kiveton Bridge"
	StationNameKivetonPark                       = "Kiveton Park"
	StationNameKnaresborough                     = "Knaresborough"
	StationNameKnebworth                         = "Knebworth"
	StationNameKnighton                          = "Knighton"
	StationNameKnockholt                         = "Knockholt"
	StationNameKnottingley                       = "Knottingley"
	StationNameKnucklas                          = "Knucklas"
	StationNameKnutsford                         = "Knutsford"
	StationNameKyleofLochalsh                    = "Kyle of Lochalsh"
	StationNameLadybank                          = "Ladybank"
	StationNameLadywell                          = "Ladywell"
	StationNameLaindon                           = "Laindon"
	StationNameLairg                             = "Lairg"
	StationNameLake                              = "Lake"
	StationNameLakenheath                        = "Lakenheath"
	StationNameLamphey                           = "Lamphey"
	StationNameLanark                            = "Lanark"
	StationNameLancaster                         = "Lancaster"
	StationNameLancing                           = "Lancing"
	StationNameLandywood                         = "Landywood"
	StationNameLangbank                          = "Langbank"
	StationNameLangho                            = "Langho"
	StationNameLangleyBerks                      = "Langley (Berks)"
	StationNameLangleyGreen                      = "Langley Green"
	StationNameLangleyMill                       = "Langley Mill"
	StationNameLangside                          = "Langside"
	StationNameLangwathby                        = "Langwathby"
	StationNameLangwithWhaleyThorns              = "Langwith-Whaley Thorns"
	StationNameLapford                           = "Lapford"
	StationNameLapworth                          = "Lapworth"
	StationNameLarbert                           = "Larbert"
	StationNameLargs                             = "Largs"
	StationNameLarkhall                          = "Larkhall"
	StationNameLaurencekirk                      = "Laurencekirk"
	StationNameLawrenceHill                      = "Lawrence Hill"
	StationNameLaytonLancs                       = "Layton (Lancs)"
	StationNameLazonbyAndKirkoswald              = "Lazonby & Kirkoswald"
	StationNameLeaBridge                         = "Lea Bridge"
	StationNameLeaGreen                          = "Lea Green"
	StationNameLeaHall                           = "Lea Hall"
	StationNameLeagrave                          = "Leagrave"
	StationNameLealholm                          = "Lealholm"
	StationNameLeamingtonSpa                     = "Leamington Spa"
	StationNameLeasowe                           = "Leasowe"
	StationNameLeatherhead                       = "Leatherhead"
	StationNameLedbury                           = "Ledbury"
	StationNameLeeLondon                         = "Lee (London)"
	StationNameLeeds                             = "Leeds"
	StationNameLeicester                         = "Leicester"
	StationNameLeighKent                         = "Leigh (Kent)"
	StationNameLeighonSea                        = "Leigh-on-Sea"
	StationNameLeightonBuzzard                   = "Leighton Buzzard"
	StationNameLelant                            = "Lelant"
	StationNameLelantSaltings                    = "Lelant Saltings"
	StationNameLenham                            = "Lenham"
	StationNameLenzie                            = "Lenzie"
	StationNameLeominster                        = "Leominster"
	StationNameLetchworthGardenCity              = "Letchworth Garden City"
	StationNameLeucharsforStAndrews              = "Leuchars (for St. Andrews)"
	StationNameLevenshulme                       = "Levenshulme"
	StationNameLewes                             = "Lewes"
	StationNameLewisham                          = "Lewisham"
	StationNameLeyland                           = "Leyland"
	StationNameLeytonMidlandRoad                 = "Leyton Midland Road"
	StationNameLeytonstoneHighRoad               = "Leytonstone High Road"
	StationNameLichfieldCity                     = "Lichfield City"
	StationNameLichfieldTrentValley              = "Lichfield Trent Valley"
	StationNameLidlington                        = "Lidlington"
	StationNameLimehouse                         = "Limehouse"
	StationNameLincolnCentral                    = "Lincoln Central"
	StationNameLingfield                         = "Lingfield"
	StationNameLingwood                          = "Lingwood"
	StationNameLinlithgow                        = "Linlithgow"
	StationNameLiphook                           = "Liphook"
	StationNameLiskeard                          = "Liskeard"
	StationNameLiss                              = "Liss"
	StationNameLisvaneAndThornhill               = "Lisvane & Thornhill"
	StationNameLittleKimble                      = "Little Kimble"
	StationNameLittleSutton                      = "Little Sutton"
	StationNameLittleborough                     = "Littleborough"
	StationNameLittlehampton                     = "Littlehampton"
	StationNameLittlehaven                       = "Littlehaven"
	StationNameLittleport                        = "Littleport"
	StationNameLiverpoolCentral                  = "Liverpool Central"
	StationNameLiverpoolJamesStreet              = "Liverpool James Street"
	StationNameLiverpoolLimeStreet               = "Liverpool Lime Street"
	StationNameLiverpoolSouthParkway             = "Liverpool South Parkway"
	StationNameLivingstonNorth                   = "Livingston North"
	StationNameLivingstonSouth                   = "Livingston South"
	StationNameLlanaber                          = "Llanaber"
	StationNameLlanbedr                          = "Llanbedr"
	StationNameLlanbisterRoad                    = "Llanbister Road"
	StationNameLlanbradach                       = "Llanbradach"
	StationNameLlandaf                           = "Llandaf"
	StationNameLlandanwg                         = "Llandanwg"
	StationNameLlandecwyn                        = "Llandecwyn"
	StationNameLlandeilo                         = "Llandeilo"
	StationNameLlandovery                        = "Llandovery"
	StationNameLlandrindod                       = "Llandrindod"
	StationNameLlandudno                         = "Llandudno"
	StationNameLlandudnoJunction                 = "Llandudno Junction"
	StationNameLlandybie                         = "Llandybie"
	StationNameLlanelli                          = "Llanelli"
	StationNameLlanfairfechan                    = "Llanfairfechan"
	StationNameLlanfairpwll                      = "Llanfairpwll"
	StationNameLlangadog                         = "Llangadog"
	StationNameLlangammarch                      = "Llangammarch"
	StationNameLlangennech                       = "Llangennech"
	StationNameLlangynllo                        = "Llangynllo"
	StationNameLlanharan                         = "Llanharan"
	StationNameLlanhilleth                       = "Llanhilleth"
	StationNameLlanishen                         = "Llanishen"
	StationNameLlanrwst                          = "Llanrwst"
	StationNameLlansamlet                        = "Llansamlet"
	StationNameLlantwitMajor                     = "Llantwit Major"
	StationNameLlanwrda                          = "Llanwrda"
	StationNameLlanwrtyd                         = "Llanwrtyd"
	StationNameLlwyngwril                        = "Llwyngwril"
	StationNameLlwynypia                         = "Llwynypia"
	StationNameLochAwe                           = "Loch Awe"
	StationNameLochEilOutwardBound               = "Loch Eil Outward Bound"
	StationNameLochailort                        = "Lochailort"
	StationNameLocheilside                       = "Locheilside"
	StationNameLochgelly                         = "Lochgelly"
	StationNameLochluichart                      = "Lochluichart"
	StationNameLochwinnoch                       = "Lochwinnoch"
	StationNameLockerbie                         = "Lockerbie"
	StationNameLockwood                          = "Lockwood"
	StationNameLondonBlackfriars                 = "London Blackfriars"
	StationNameLondonBridge                      = "London Bridge"
	StationNameLondonCannonStreet                = "London Cannon Street"
	StationNameLondonCharingCross                = "London Charing Cross"
	StationNameLondonEuston                      = "London Euston"
	StationNameLondonFenchurchStreet             = "London Fenchurch Street"
	StationNameLondonFields                      = "London Fields"
	StationNameLondonKingsCross                  = "London Kings Cross"
	StationNameLondonLiverpoolStreet             = "London Liverpool Street"
	StationNameLondonMarylebone                  = "London Marylebone"
	StationNameLondonPaddington                  = "London Paddington"
	StationNameLondonRoadBrighton                = "London Road (Brighton)"
	StationNameLondonRoadGuildford               = "London Road (Guildford)"
	StationNameLondonStPancrasIntl               = "London St Pancras (Intl)"
	StationNameLondonVictoria                    = "London Victoria"
	StationNameLondonWaterloo                    = "London Waterloo"
	StationNameLondonWaterlooEast                = "London Waterloo East"
	StationNameLongBuckby                        = "Long Buckby"
	StationNameLongEaton                         = "Long Eaton"
	StationNameLongPreston                       = "Long Preston"
	StationNameLongbeck                          = "Longbeck"
	StationNameLongbridge                        = "Longbridge"
	StationNameLongcross                         = "Longcross"
	StationNameLongfield                         = "Longfield"
	StationNameLongniddry                        = "Longniddry"
	StationNameLongport                          = "Longport"
	StationNameLongton                           = "Longton"
	StationNameLooe                              = "Looe"
	StationNameLostock                           = "Lostock"
	StationNameLostockGralam                     = "Lostock Gralam"
	StationNameLostockHall                       = "Lostock Hall"
	StationNameLostwithiel                       = "Lostwithiel"
	StationNameLoughborough                      = "Loughborough"
	StationNameLoughboroughJunction              = "Loughborough Junction"
	StationNameLowdham                           = "Lowdham"
	StationNameLowerSydenham                     = "Lower Sydenham"
	StationNameLowestoft                         = "Lowestoft"
	StationNameLudlow                            = "Ludlow"
	StationNameLuton                             = "Luton"
	StationNameLutonAirportParkway               = "Luton Airport Parkway"
	StationNameLuxulyan                          = "Luxulyan"
	StationNameLydney                            = "Lydney"
	StationNameLyeWestMidlands                   = "Lye (West Midlands)"
	StationNameLymingtonPier                     = "Lymington Pier"
	StationNameLymingtonTown                     = "Lymington Town"
	StationNameLympstoneCommando                 = "Lympstone Commando"
	StationNameLympstoneVillage                  = "Lympstone Village"
	StationNameLytham                            = "Lytham"
	StationNameMacclesfield                      = "Macclesfield"
	StationNameMachynlleth                       = "Machynlleth"
	StationNameMaesteg                           = "Maesteg"
	StationNameMaestegEwennyRoad                 = "Maesteg (Ewenny Road)"
	StationNameMaghull                           = "Maghull"
	StationNameMaidenNewton                      = "Maiden Newton"
	StationNameMaidenhead                        = "Maidenhead"
	StationNameMaidstoneBarracks                 = "Maidstone Barracks"
	StationNameMaidstoneEast                     = "Maidstone East"
	StationNameMaidstoneWest                     = "Maidstone West"
	StationNameMaldenManor                       = "Malden Manor"
	StationNameMallaig                           = "Mallaig"
	StationNameMalton                            = "Malton"
	StationNameMalvernLink                       = "Malvern Link"
	StationNameManchesterAirport                 = "Manchester Airport"
	StationNameManchesterOxfordRoad              = "Manchester Oxford Road"
	StationNameManchesterPiccadilly              = "Manchester Piccadilly"
	StationNameManchesterUnitedFootballGround    = "Manchester United Football Ground"
	StationNameManchesterVictoria                = "Manchester Victoria"
	StationNameManea                             = "Manea"
	StationNameManningtree                       = "Manningtree"
	StationNameManorPark                         = "Manor Park"
	StationNameManorRoad                         = "Manor Road"
	StationNameManorbier                         = "Manorbier"
	StationNameManors                            = "Manors"
	StationNameMansfield                         = "Mansfield"
	StationNameMansfieldWoodhouse                = "Mansfield Woodhouse"
	StationNameMarch                             = "March"
	StationNameMardenKent                        = "Marden (Kent)"
	StationNameMargate                           = "Margate"
	StationNameMarketHarborough                  = "Market Harborough"
	StationNameMarketRasen                       = "Market Rasen"
	StationNameMarkinch                          = "Markinch"
	StationNameMarksTey                          = "Marks Tey"
	StationNameMarlow                            = "Marlow"
	StationNameMarple                            = "Marple"
	StationNameMarsdenYorks                      = "Marsden (Yorks)"
	StationNameMarske                            = "Marske"
	StationNameMarstonGreen                      = "Marston Green"
	StationNameMartinMill                        = "Martin Mill"
	StationNameMartinsHeron                      = "Martins Heron"
	StationNameMarton                            = "Marton"
	StationNameMaryhill                          = "Maryhill"
	StationNameMaryland                          = "Maryland"
	StationNameMaryport                          = "Maryport"
	StationNameMatlock                           = "Matlock"
	StationNameMatlockBath                       = "Matlock Bath"
	StationNameMauldethRoad                      = "Mauldeth Road"
	StationNameMaxwellPark                       = "Maxwell Park"
	StationNameMaybole                           = "Maybole"
	StationNameMazeHill                          = "Maze Hill"
	StationNameMeadowhall                        = "Meadowhall"
	StationNameMeldreth                          = "Meldreth"
	StationNameMelksham                          = "Melksham"
	StationNameMeltonMowbray                     = "Melton Mowbray"
	StationNameMeltonSuffolk                     = "Melton (Suffolk)"
	StationNameMenheniot                         = "Menheniot"
	StationNameMenston                           = "Menston"
	StationNameMeols                             = "Meols"
	StationNameMeolsCop                          = "Meols Cop"
	StationNameMeopham                           = "Meopham"
	StationNameMerryton                          = "Merryton"
	StationNameMerstham                          = "Merstham"
	StationNameMerthyrTydfil                     = "Merthyr Tydfil"
	StationNameMerthyrVale                       = "Merthyr Vale"
	StationNameMetheringham                      = "Metheringham"
	StationNameMetroCentre                       = "MetroCentre"
	StationNameMexborough                        = "Mexborough"
	StationNameMicheldever                       = "Micheldever"
	StationNameMicklefield                       = "Micklefield"
	StationNameMiddlesbrough                     = "Middlesbrough"
	StationNameMiddlewood                        = "Middlewood"
	StationNameMidgham                           = "Midgham"
	StationNameMilfordHaven                      = "Milford Haven"
	StationNameMilfordSurrey                     = "Milford (Surrey)"
	StationNameMillHillBroadway                  = "Mill Hill Broadway"
	StationNameMillHillLancs                     = "Mill Hill (Lancs)"
	StationNameMillbrookBeds                     = "Millbrook (Beds)"
	StationNameMillbrookHants                    = "Millbrook (Hants)"
	StationNameMillikenPark                      = "Milliken Park"
	StationNameMillom                            = "Millom"
	StationNameMillsHillManchester               = "Mills Hill (Manchester)"
	StationNameMilngavie                         = "Milngavie"
	StationNameMiltonKeynesCentral               = "Milton Keynes Central"
	StationNameMinffordd                         = "Minffordd"
	StationNameMinster                           = "Minster"
	StationNameMirfield                          = "Mirfield"
	StationNameMistley                           = "Mistley"
	StationNameMitchamEastfields                 = "Mitcham Eastfields"
	StationNameMitchamJunction                   = "Mitcham Junction"
	StationNameMobberley                         = "Mobberley"
	StationNameMonifieth                         = "Monifieth"
	StationNameMonksRisborough                   = "Monks Risborough"
	StationNameMontpelier                        = "Montpelier"
	StationNameMontrose                          = "Montrose"
	StationNameMoorfields                        = "Moorfields"
	StationNameMoorgate                          = "Moorgate"
	StationNameMoorside                          = "Moorside"
	StationNameMoorthorpe                        = "Moorthorpe"
	StationNameMorar                             = "Morar"
	StationNameMorchardRoad                      = "Morchard Road"
	StationNameMordenSouth                       = "Morden South"
	StationNameMorecambe                         = "Morecambe"
	StationNameMoretonDorset                     = "Moreton (Dorset)"
	StationNameMoretonMerseyside                 = "Moreton (Merseyside)"
	StationNameMoretoninMarsh                    = "Moreton-in-Marsh"
	StationNameMorfaMawddach                     = "Morfa Mawddach"
	StationNameMorley                            = "Morley"
	StationNameMorpeth                           = "Morpeth"
	StationNameMortimer                          = "Mortimer"
	StationNameMortlake                          = "Mortlake"
	StationNameMosesGate                         = "Moses Gate"
	StationNameMossSide                          = "Moss Side"
	StationNameMossleyHill                       = "Mossley Hill"
	StationNameMossleyManchester                 = "Mossley (Manchester)"
	StationNameMosspark                          = "Mosspark"
	StationNameMoston                            = "Moston"
	StationNameMotherwell                        = "Motherwell"
	StationNameMotspurPark                       = "Motspur Park"
	StationNameMottingham                        = "Mottingham"
	StationNameMottisfontAndDunbridge            = "Mottisfont & Dunbridge"
	StationNameMouldsworth                       = "Mouldsworth"
	StationNameMoulsecoomb                       = "Moulsecoomb"
	StationNameMountFlorida                      = "Mount Florida"
	StationNameMountVernon                       = "Mount Vernon"
	StationNameMountainAsh                       = "Mountain Ash"
	StationNameMuirend                           = "Muirend"
	StationNameMuirofOrd                         = "Muir of Ord"
	StationNameMusselburgh                       = "Musselburgh"
	StationNameMytholmroyd                       = "Mytholmroyd"
	StationNameNafferton                         = "Nafferton"
	StationNameNailseaAndBackwell                = "Nailsea & Backwell"
	StationNameNairn                             = "Nairn"
	StationNameNantwich                          = "Nantwich"
	StationNameNarberth                          = "Narberth"
	StationNameNarborough                        = "Narborough"
	StationNameNavigationRoad                    = "Navigation Road"
	StationNameNeath                             = "Neath"
	StationNameNeedhamMarket                     = "Needham Market"
	StationNameNeilston                          = "Neilston"
	StationNameNelson                            = "Nelson"
	StationNameNeston                            = "Neston"
	StationNameNetherfield                       = "Netherfield"
	StationNameNethertown                        = "Nethertown"
	StationNameNetley                            = "Netley"
	StationNameNewBarnet                         = "New Barnet"
	StationNameNewBeckenham                      = "New Beckenham"
	StationNameNewBrighton                       = "New Brighton"
	StationNameNewClee                           = "New Clee"
	StationNameNewCross                          = "New Cross"
	StationNameNewCrossGate                      = "New Cross Gate"
	StationNameNewCumnock                        = "New Cumnock"
	StationNameNewEltham                         = "New Eltham"
	StationNameNewHolland                        = "New Holland"
	StationNameNewHythe                          = "New Hythe"
	StationNameNewLane                           = "New Lane"
	StationNameNewMalden                         = "New Malden"
	StationNameNewMillsCentral                   = "New Mills Central"
	StationNameNewMillsNewtown                   = "New Mills Newtown"
	StationNameNewMilton                         = "New Milton"
	StationNameNewPudsey                         = "New Pudsey"
	StationNameNewSouthgate                      = "New Southgate"
	StationNameNewarkCastle                      = "Newark Castle"
	StationNameNewarkNorthGate                   = "Newark North Gate"
	StationNameNewbridge                         = "Newbridge"
	StationNameNewbury                           = "Newbury"
	StationNameNewburyRacecourse                 = "Newbury Racecourse"
	StationNameNewcastle                         = "Newcastle"
	StationNameNewcourt                          = "Newcourt"
	StationNameNewcraighall                      = "Newcraighall"
	StationNameNewhavenHarbour                   = "Newhaven Harbour"
	StationNameNewhavenTown                      = "Newhaven Town"
	StationNameNewington                         = "Newington"
	StationNameNewmarket                         = "Newmarket"
	StationNameNewportEssex                      = "Newport (Essex)"
	StationNameNewportSouthWales                 = "Newport (South Wales)"
	StationNameNewquay                           = "Newquay"
	StationNameNewstead                          = "Newstead"
	StationNameNewtonAbbot                       = "Newton Abbot"
	StationNameNewtonAycliffe                    = "Newton Aycliffe"
	StationNameNewtonLanark                      = "Newton (Lanark)"
	StationNameNewtonStCyres                     = "Newton St Cyres"
	StationNameNewtonforHyde                     = "Newton for Hyde"
	StationNameNewtongrange                      = "Newtongrange"
	StationNameNewtonleWillows                   = "Newton-le-Willows"
	StationNameNewtonmore                        = "Newtonmore"
	StationNameNewtononAyr                       = "Newton-on-Ayr"
	StationNameNewtownPowys                      = "Newtown (Powys)"
	StationNameNinianPark                        = "Ninian Park"
	StationNameNitshill                          = "Nitshill"
	StationNameNorbiton                          = "Norbiton"
	StationNameNorbury                           = "Norbury"
	StationNameNormansBay                        = "Normans Bay"
	StationNameNormanton                         = "Normanton"
	StationNameNorthBerwick                      = "North Berwick"
	StationNameNorthCamp                         = "North Camp"
	StationNameNorthDulwich                      = "North Dulwich"
	StationNameNorthFambridge                    = "North Fambridge"
	StationNameNorthLlanrwst                     = "North Llanrwst"
	StationNameNorthQueensferry                  = "North Queensferry"
	StationNameNorthRoadDarlington               = "North Road (Darlington)"
	StationNameNorthSheen                        = "North Sheen"
	StationNameNorthWalsham                      = "North Walsham"
	StationNameNorthWembley                      = "North Wembley"
	StationNameNorthallerton                     = "Northallerton"
	StationNameNorthampton                       = "Northampton"
	StationNameNorthfield                        = "Northfield"
	StationNameNorthfleet                        = "Northfleet"
	StationNameNortholtPark                      = "Northolt Park"
	StationNameNorthumberlandPark                = "Northumberland Park"
	StationNameNorthwich                         = "Northwich"
	StationNameNortonBridge                      = "Norton Bridge"
	StationNameNorwich                           = "Norwich"
	StationNameNorwoodJunction                   = "Norwood Junction"
	StationNameNottingham                        = "Nottingham"
	StationNameNuneaton                          = "Nuneaton"
	StationNameNunhead                           = "Nunhead"
	StationNameNunthorpe                         = "Nunthorpe"
	StationNameNutbourne                         = "Nutbourne"
	StationNameNutfield                          = "Nutfield"
	StationNameOakengates                        = "Oakengates"
	StationNameOakham                            = "Oakham"
	StationNameOakleighPark                      = "Oakleigh Park"
	StationNameOban                              = "Oban"
	StationNameOckendon                          = "Ockendon"
	StationNameOckley                            = "Ockley"
	StationNameOkehampton                        = "Okehampton"
	StationNameOldHill                           = "Old Hill"
	StationNameOldRoan                           = "Old Roan"
	StationNameOldStreet                         = "Old Street"
	StationNameOldfieldPark                      = "Oldfield Park"
	StationNameOlton                             = "Olton"
	StationNameOre                               = "Ore"
	StationNameOrmskirk                          = "Ormskirk"
	StationNameOrpington                         = "Orpington"
	StationNameOrrell                            = "Orrell"
	StationNameOrrellPark                        = "Orrell Park"
	StationNameOtford                            = "Otford"
	StationNameOultonBroadNorth                  = "Oulton Broad North"
	StationNameOultonBroadSouth                  = "Oulton Broad South"
	StationNameOutwood                           = "Outwood"
	StationNameOverpool                          = "Overpool"
	StationNameOverton                           = "Overton"
	StationNameOxenholmeLakeDistrict             = "Oxenholme Lake District"
	StationNameOxford                            = "Oxford"
	StationNameOxfordParkway                     = "Oxford Parkway"
	StationNameOxshott                           = "Oxshott"
	StationNameOxted                             = "Oxted"
	StationNamePaddockWood                       = "Paddock Wood"
	StationNamePadgate                           = "Padgate"
	StationNamePaignton                          = "Paignton"
	StationNamePaisleyCanal                      = "Paisley Canal"
	StationNamePaisleyGilmourStreet              = "Paisley Gilmour Street"
	StationNamePaisleyStJames                    = "Paisley St James"
	StationNamePalmersGreen                      = "Palmers Green"
	StationNamePangbourne                        = "Pangbourne"
	StationNamePannal                            = "Pannal"
	StationNamePantyffynnon                      = "Pantyffynnon"
	StationNamePar                               = "Par"
	StationNameParbold                           = "Parbold"
	StationNameParkStreet                        = "Park Street"
	StationNameParkstoneDorset                   = "Parkstone (Dorset)"
	StationNameParsonStreet                      = "Parson Street"
	StationNamePartick                           = "Partick"
	StationNameParton                            = "Parton"
	StationNamePatchway                          = "Patchway"
	StationNamePatricroft                        = "Patricroft"
	StationNamePatterton                         = "Patterton"
	StationNamePeartree                          = "Peartree"
	StationNamePeckhamRye                        = "Peckham Rye"
	StationNamePegswood                          = "Pegswood"
	StationNamePemberton                         = "Pemberton"
	StationNamePembreyAndBurryPort               = "Pembrey & Burry Port"
	StationNamePembroke                          = "Pembroke"
	StationNamePembrokeDock                      = "Pembroke Dock"
	StationNamePenally                           = "Penally"
	StationNamePenarth                           = "Penarth"
	StationNamePencoed                           = "Pencoed"
	StationNamePengam                            = "Pengam"
	StationNamePengeEast                         = "Penge East"
	StationNamePengeWest                         = "Penge West"
	StationNamePenhelig                          = "Penhelig"
	StationNamePenistone                         = "Penistone"
	StationNamePenkridge                         = "Penkridge"
	StationNamePenmaenmawr                       = "Penmaenmawr"
	StationNamePenmere                           = "Penmere"
	StationNamePenrhiwceiber                     = "Penrhiwceiber"
	StationNamePenrhyndeudraeth                  = "Penrhyndeudraeth"
	StationNamePenrithNorthLakes                 = "Penrith (North Lakes)"
	StationNamePenrynCornwall                    = "Penryn (Cornwall)"
	StationNamePensarnGwynedd                    = "Pensarn (Gwynedd)"
	StationNamePenshurst                         = "Penshurst"
	StationNamePentreBach                        = "Pentre-Bach"
	StationNamePenyBont                          = "Pen-y-Bont"
	StationNamePenychain                         = "Penychain"
	StationNamePenyffordd                        = "Penyffordd"
	StationNamePenzance                          = "Penzance"
	StationNamePerranwell                        = "Perranwell"
	StationNamePerryBarr                         = "Perry Barr"
	StationNamePershore                          = "Pershore"
	StationNamePerth                             = "Perth"
	StationNamePeterborough                      = "Peterborough"
	StationNamePetersfield                       = "Petersfield"
	StationNamePettsWood                         = "Petts Wood"
	StationNamePevenseyAndWestham                = "Pevensey & Westham"
	StationNamePevenseyBay                       = "Pevensey Bay"
	StationNamePewsey                            = "Pewsey"
	StationNamePilning                           = "Pilning"
	StationNamePinhoe                            = "Pinhoe"
	StationNamePitlochry                         = "Pitlochry"
	StationNamePitsea                            = "Pitsea"
	StationNamePleasington                       = "Pleasington"
	StationNamePlockton                          = "Plockton"
	StationNamePluckley                          = "Pluckley"
	StationNamePlumley                           = "Plumley"
	StationNamePlumpton                          = "Plumpton"
	StationNamePlumstead                         = "Plumstead"
	StationNamePlymouth                          = "Plymouth"
	StationNamePokesdown                         = "Pokesdown"
	StationNamePolegate                          = "Polegate"
	StationNamePolesworth                        = "Polesworth"
	StationNamePollokshawsEast                   = "Pollokshaws East"
	StationNamePollokshawsWest                   = "Pollokshaws West"
	StationNamePollokshieldsEast                 = "Pollokshields East"
	StationNamePollokshieldsWest                 = "Pollokshields West"
	StationNamePolmont                           = "Polmont"
	StationNamePolsloeBridge                     = "Polsloe Bridge"
	StationNamePondersEnd                        = "Ponders End"
	StationNamePontarddulais                     = "Pontarddulais"
	StationNamePontefractBaghill                 = "Pontefract Baghill"
	StationNamePontefractMonkhill                = "Pontefract Monkhill"
	StationNamePontefractTanshelf                = "Pontefract Tanshelf"
	StationNamePontlottyn                        = "Pontlottyn"
	StationNamePontyPant                         = "Pont-y-Pant"
	StationNamePontyclun                         = "Pontyclun"
	StationNamePontypoolAndNewInn                = "Pontypool & New Inn"
	StationNamePontypridd                        = "Pontypridd"
	StationNamePoole                             = "Poole"
	StationNamePoppleton                         = "Poppleton"
	StationNamePortGlasgow                       = "Port Glasgow"
	StationNamePortSunlight                      = "Port Sunlight"
	StationNamePortTalbotParkway                 = "Port Talbot Parkway"
	StationNamePortchester                       = "Portchester"
	StationNamePorth                             = "Porth"
	StationNamePorthmadog                        = "Porthmadog"
	StationNamePortlethen                        = "Portlethen"
	StationNamePortslade                         = "Portslade"
	StationNamePortsmouthAndSouthsea             = "Portsmouth & Southsea"
	StationNamePortsmouthArms                    = "Portsmouth Arms"
	StationNamePortsmouthHarbour                 = "Portsmouth Harbour"
	StationNamePossilparkAndParkhouse            = "Possilpark & Parkhouse"
	StationNamePottersBar                        = "Potters Bar"
	StationNamePoultonleFylde                    = "Poulton-le-Fylde"
	StationNamePoynton                           = "Poynton"
	StationNamePrees                             = "Prees"
	StationNamePrescot                           = "Prescot"
	StationNamePrestatyn                         = "Prestatyn"
	StationNamePrestbury                         = "Prestbury"
	StationNamePrestonLancs                      = "Preston (Lancs)"
	StationNamePrestonPark                       = "Preston Park"
	StationNamePrestonpans                       = "Prestonpans"
	StationNamePrestwickInternationalAirport     = "Prestwick International Airport"
	StationNamePrestwickTown                     = "Prestwick Town"
	StationNamePriesthillAndDarnley              = "Priesthill & Darnley"
	StationNamePrincesRisborough                 = "Princes Risborough"
	StationNamePrittlewell                       = "Prittlewell"
	StationNamePrudhoe                           = "Prudhoe"
	StationNamePulborough                        = "Pulborough"
	StationNamePurfleet                          = "Purfleet"
	StationNamePurley                            = "Purley"
	StationNamePurleyOaks                        = "Purley Oaks"
	StationNamePutney                            = "Putney"
	StationNamePwllheli                          = "Pwllheli"
	StationNamePyeCorner                         = "Pye Corner"
	StationNamePyle                              = "Pyle"
	StationNameQuakersYard                       = "Quakers Yard"
	StationNameQueenborough                      = "Queenborough"
	StationNameQueensParkGlasgow                 = "Queens Park (Glasgow)"
	StationNameQueensParkLondon                  = "Queens Park (London)"
	StationNameQueensRoadPeckham                 = "Queens Road (Peckham)"
	StationNameQueenstownRoadBattersea           = "Queenstown Road (Battersea)"
	StationNameQuintrellDowns                    = "Quintrell Downs"
	StationNameRadcliffeonTrent                  = "Radcliffe-on-Trent"
	StationNameRadlett                           = "Radlett"
	StationNameRadley                            = "Radley"
	StationNameRadyr                             = "Radyr"
	StationNameRainford                          = "Rainford"
	StationNameRainhamEssex                      = "Rainham (Essex)"
	StationNameRainhamKent                       = "Rainham (Kent)"
	StationNameRainhill                          = "Rainhill"
	StationNameRamsgate                          = "Ramsgate"
	StationNameRamsgreaveAndWilpshire            = "Ramsgreave & Wilpshire"
	StationNameRannoch                           = "Rannoch"
	StationNameRauceby                           = "Rauceby"
	StationNameRavenglassforEskdale              = "Ravenglass for Eskdale"
	StationNameRavensbourne                      = "Ravensbourne"
	StationNameRavensthorpe                      = "Ravensthorpe"
	StationNameRawcliffe                         = "Rawcliffe"
	StationNameRayleigh                          = "Rayleigh"
	StationNameRaynesPark                        = "Raynes Park"
	StationNameReading                           = "Reading"
	StationNameReadingWest                       = "Reading West"
	StationNameRectoryRoad                       = "Rectory Road"
	StationNameRedbridge                         = "Redbridge"
	StationNameRedcarBritishSteel                = "Redcar British Steel"
	StationNameRedcarCentral                     = "Redcar Central"
	StationNameRedcarEast                        = "Redcar East"
	StationNameReddishNorth                      = "Reddish North"
	StationNameReddishSouth                      = "Reddish South"
	StationNameRedditch                          = "Redditch"
	StationNameRedhill                           = "Redhill"
	StationNameRedland                           = "Redland"
	StationNameRedruth                           = "Redruth"
	StationNameReedhamNorfolk                    = "Reedham (Norfolk)"
	StationNameReedhamSurrey                     = "Reedham (Surrey)"
	StationNameReigate                           = "Reigate"
	StationNameRenton                            = "Renton"
	StationNameRetford                           = "Retford"
	StationNameRhiwbina                          = "Rhiwbina"
	StationNameRhooseCardiffInternationalAirport = "Rhoose Cardiff International Airport"
	StationNameRhosneigr                         = "Rhosneigr"
	StationNameRhyl                              = "Rhyl"
	StationNameRhymney                           = "Rhymney"
	StationNameRibblehead                        = "Ribblehead"
	StationNameRiceLane                          = "Rice Lane"
	StationNameRichmondLondon                    = "Richmond (London)"
	StationNameRickmansworth                     = "Rickmansworth"
	StationNameRiddlesdown                       = "Riddlesdown"
	StationNameRidgmont                          = "Ridgmont"
	StationNameRidingMill                        = "Riding Mill"
	StationNameRiscaAndPontymister               = "Risca & Pontymister"
	StationNameRishton                           = "Rishton"
	StationNameRobertsbridge                     = "Robertsbridge"
	StationNameRoby                              = "Roby"
	StationNameRochdale                          = "Rochdale"
	StationNameRoche                             = "Roche"
	StationNameRochester                         = "Rochester"
	StationNameRochford                          = "Rochford"
	StationNameRockFerry                         = "Rock Ferry"
	StationNameRogart                            = "Rogart"
	StationNameRogerstone                        = "Rogerstone"
	StationNameRolleston                         = "Rolleston"
	StationNameRomanBridge                       = "Roman Bridge"
	StationNameRomford                           = "Romford"
	StationNameRomiley                           = "Romiley"
	StationNameRomsey                            = "Romsey"
	StationNameRoose                             = "Roose"
	StationNameRoseGrove                         = "Rose Grove"
	StationNameRoseHillMarple                    = "Rose Hill Marple"
	StationNameRosyth                            = "Rosyth"
	StationNameRotherhamCentral                  = "Rotherham Central"
	StationNameRotherhithe                       = "Rotherhithe"
	StationNameRoughtonRoad                      = "Roughton Road"
	StationNameRowlandsCastle                    = "Rowlands Castle"
	StationNameRowleyRegis                       = "Rowley Regis"
	StationNameRoyBridge                         = "Roy Bridge"
	StationNameRoydon                            = "Roydon"
	StationNameRoyston                           = "Royston"
	StationNameRuabon                            = "Ruabon"
	StationNameRufford                           = "Rufford"
	StationNameRugby                             = "Rugby"
	StationNameRugeleyTown                       = "Rugeley Town"
	StationNameRugeleyTrentValley                = "Rugeley Trent Valley"
	StationNameRuncorn                           = "Runcorn"
	StationNameRuncornEast                       = "Runcorn East"
	StationNameRuskington                        = "Ruskington"
	StationNameRuswarp                           = "Ruswarp"
	StationNameRutherglen                        = "Rutherglen"
	StationNameRydeEsplanade                     = "Ryde Esplanade"
	StationNameRydePierHead                      = "Ryde Pier Head"
	StationNameRydeStJohnsRoad                   = "Ryde St Johns Road"
	StationNameRyderBrow                         = "Ryder Brow"
	StationNameRyeHouse                          = "Rye House"
	StationNameRyeSussex                         = "Rye (Sussex)"
	StationNameSalfordCentral                    = "Salford Central"
	StationNameSalfordCrescent                   = "Salford Crescent"
	StationNameSalfordsSurrey                    = "Salfords (Surrey)"
	StationNameSalhouse                          = "Salhouse"
	StationNameSalisbury                         = "Salisbury"
	StationNameSaltaire                          = "Saltaire"
	StationNameSaltash                           = "Saltash"
	StationNameSaltburn                          = "Saltburn"
	StationNameSaltcoats                         = "Saltcoats"
	StationNameSaltmarshe                        = "Saltmarshe"
	StationNameSalwick                           = "Salwick"
	StationNameSampfordCourtenay                 = "Sampford Courtenay"
	StationNameSandalAndAgbrigg                  = "Sandal & Agbrigg"
	StationNameSandbach                          = "Sandbach"
	StationNameSanderstead                       = "Sanderstead"
	StationNameSandhills                         = "Sandhills"
	StationNameSandhurstBerks                    = "Sandhurst (Berks)"
	StationNameSandling                          = "Sandling"
	StationNameSandown                           = "Sandown"
	StationNameSandplace                         = "Sandplace"
	StationNameSandwellAndDudley                 = "Sandwell & Dudley"
	StationNameSandwich                          = "Sandwich"
	StationNameSandy                             = "Sandy"
	StationNameSankeyforPenketh                  = "Sankey for Penketh"
	StationNameSanquhar                          = "Sanquhar"
	StationNameSarn                              = "Sarn"
	StationNameSaundersfoot                      = "Saundersfoot"
	StationNameSaunderton                        = "Saunderton"
	StationNameSawbridgeworth                    = "Sawbridgeworth"
	StationNameSaxilby                           = "Saxilby"
	StationNameSaxmundham                        = "Saxmundham"
	StationNameScarborough                       = "Scarborough"
	StationNameScotscalder                       = "Scotscalder"
	StationNameScotstounhill                     = "Scotstounhill"
	StationNameScunthorpe                        = "Scunthorpe"
	StationNameSeaMills                          = "Sea Mills"
	StationNameSeafordSussex                     = "Seaford (Sussex)"
	StationNameSeaforthAndLitherland             = "Seaforth & Litherland"
	StationNameSeaham                            = "Seaham"
	StationNameSeamer                            = "Seamer"
	StationNameSeascale                          = "Seascale"
	StationNameSeatonCarew                       = "Seaton Carew"
	StationNameSeerGreenAndJordans               = "Seer Green & Jordans"
	StationNameSelby                             = "Selby"
	StationNameSelhurst                          = "Selhurst"
	StationNameSellafield                        = "Sellafield"
	StationNameSelling                           = "Selling"
	StationNameSellyOak                          = "Selly Oak"
	StationNameSettle                            = "Settle"
	StationNameSevenKings                        = "Seven Kings"
	StationNameSevenSisters                      = "Seven Sisters"
	StationNameSevenoaks                         = "Sevenoaks"
	StationNameSevernBeach                       = "Severn Beach"
	StationNameSevernTunnelJunction              = "Severn Tunnel Junction"
	StationNameShadwell                          = "Shadwell"
	StationNameShalfordSurrey                    = "Shalford (Surrey)"
	StationNameShanklin                          = "Shanklin"
	StationNameShawfair                          = "Shawfair"
	StationNameShawford                          = "Shawford"
	StationNameShawlands                         = "Shawlands"
	StationNameSheernessonSea                    = "Sheerness-on-Sea"
	StationNameSheffield                         = "Sheffield"
	StationNameShelfordCambs                     = "Shelford (Cambs)"
	StationNameShenfield                         = "Shenfield"
	StationNameShenstone                         = "Shenstone"
	StationNameShepherdsBush                     = "Shepherd's Bush"
	StationNameShepherdsWell                     = "Shepherds Well"
	StationNameShepley                           = "Shepley"
	StationNameShepperton                        = "Shepperton"
	StationNameShepreth                          = "Shepreth"
	StationNameSherborne                         = "Sherborne"
	StationNameSherburninElmet                   = "Sherburn-in-Elmet"
	StationNameSheringham                        = "Sheringham"
	StationNameShettleston                       = "Shettleston"
	StationNameShieldmuir                        = "Shieldmuir"
	StationNameShifnal                           = "Shifnal"
	StationNameShildon                           = "Shildon"
	StationNameShiplake                          = "Shiplake"
	StationNameShipleyYorks                      = "Shipley (Yorks)"
	StationNameShippeaHill                       = "Shippea Hill"
	StationNameShipton                           = "Shipton"
	StationNameShirebrook                        = "Shirebrook"
	StationNameShirehampton                      = "Shirehampton"
	StationNameShireoaks                         = "Shireoaks"
	StationNameShirley                           = "Shirley"
	StationNameShoeburyness                      = "Shoeburyness"
	StationNameSholing                           = "Sholing"
	StationNameShoreditchHighStreet              = "Shoreditch High Street"
	StationNameShorehamKent                      = "Shoreham (Kent)"
	StationNameShorehambySea                     = "Shoreham-by-Sea"
	StationNameShortlands                        = "Shortlands"
	StationNameShotton                           = "Shotton"
	StationNameShotts                            = "Shotts"
	StationNameShrewsbury                        = "Shrewsbury"
	StationNameSidcup                            = "Sidcup"
	StationNameSileby                            = "Sileby"
	StationNameSilecroft                         = "Silecroft"
	StationNameSilkstoneCommon                   = "Silkstone Common"
	StationNameSilverStreet                      = "Silver Street"
	StationNameSilverdale                        = "Silverdale"
	StationNameSinger                            = "Singer"
	StationNameSittingbourne                     = "Sittingbourne"
	StationNameSkegness                          = "Skegness"
	StationNameSkewen                            = "Skewen"
	StationNameSkipton                           = "Skipton"
	StationNameSladeGreen                        = "Slade Green"
	StationNameSlaithwaite                       = "Slaithwaite"
	StationNameSlateford                         = "Slateford"
	StationNameSleaford                          = "Sleaford"
	StationNameSleights                          = "Sleights"
	StationNameSlough                            = "Slough"
	StationNameSmallHeath                        = "Small Heath"
	StationNameSmallbrookJunction                = "Smallbrook Junction"
	StationNameSmethwickGaltonBridge             = "Smethwick Galton Bridge"
	StationNameSmethwickRolfeStreet              = "Smethwick Rolfe Street"
	StationNameSmithyBridge                      = "Smithy Bridge"
	StationNameSnaith                            = "Snaith"
	StationNameSnodland                          = "Snodland"
	StationNameSnowdown                          = "Snowdown"
	StationNameSoleStreet                        = "Sole Street"
	StationNameSolihull                          = "Solihull"
	StationNameSomerleyton                       = "Somerleyton"
	StationNameSouthActon                        = "South Acton"
	StationNameSouthBank                         = "South Bank"
	StationNameSouthBermondsey                   = "South Bermondsey"
	StationNameSouthCroydon                      = "South Croydon"
	StationNameSouthElmsall                      = "South Elmsall"
	StationNameSouthGreenford                    = "South Greenford"
	StationNameSouthGyle                         = "South Gyle"
	StationNameSouthHampstead                    = "South Hampstead"
	StationNameSouthKenton                       = "South Kenton"
	StationNameSouthMerton                       = "South Merton"
	StationNameSouthMilford                      = "South Milford"
	StationNameSouthRuislip                      = "South Ruislip"
	StationNameSouthTottenham                    = "South Tottenham"
	StationNameSouthWigston                      = "South Wigston"
	StationNameSouthWoodhamFerrers               = "South Woodham Ferrers"
	StationNameSouthall                          = "Southall"
	StationNameSouthamptonAirportParkway         = "Southampton Airport Parkway"
	StationNameSouthamptonCentral                = "Southampton Central"
	StationNameSouthbourne                       = "Southbourne"
	StationNameSouthbury                         = "Southbury"
	StationNameSouthease                         = "Southease"
	StationNameSouthendAirport                   = "Southend Airport"
	StationNameSouthendCentral                   = "Southend Central"
	StationNameSouthendEast                      = "Southend East"
	StationNameSouthendVictoria                  = "Southend Victoria"
	StationNameSouthminster                      = "Southminster"
	StationNameSouthport                         = "Southport"
	StationNameSouthwick                         = "Southwick"
	StationNameSowerbyBridge                     = "Sowerby Bridge"
	StationNameSpalding                          = "Spalding"
	StationNameSpeanBridge                       = "Spean Bridge"
	StationNameSpital                            = "Spital"
	StationNameSpondon                           = "Spondon"
	StationNameSpoonerRow                        = "Spooner Row"
	StationNameSpringRoad                        = "Spring Road"
	StationNameSpringburn                        = "Springburn"
	StationNameSpringfield                       = "Springfield"
	StationNameSquiresGate                       = "Squires Gate"
	StationNameStAlbansAbbey                     = "St Albans Abbey"
	StationNameStAlbansCity                      = "St Albans City"
	StationNameStAndrewsRoad                     = "St Andrews Road"
	StationNameStAnnesonSea                      = "St Annes-on-Sea"
	StationNameStAustell                         = "St Austell"
	StationNameStBees                            = "St Bees"
	StationNameStBudeauxFerryRoad                = "St Budeaux Ferry Road"
	StationNameStBudeauxVictoriaRoad             = "St Budeaux Victoria Road"
	StationNameStColumbRoad                      = "St Columb Road"
	StationNameStDenys                           = "St Denys"
	StationNameStErth                            = "St Erth"
	StationNameStGermans                         = "St Germans"
	StationNameStHelensCentral                   = "St Helens Central"
	StationNameStHelensJunction                  = "St Helens Junction"
	StationNameStHelierSurrey                    = "St Helier (Surrey)"
	StationNameStIvesCornwall                    = "St Ives (Cornwall)"
	StationNameStJamesParkExeter                 = "St James Park (Exeter)"
	StationNameStJamesStreetWalthamstow          = "St James Street (Walthamstow)"
	StationNameStJohnsLondon                     = "St Johns (London)"
	StationNameStKeyneWishingWellHalt            = "St Keyne Wishing Well Halt"
	StationNameStLeonardsWarriorSquare           = "St Leonards Warrior Square"
	StationNameStMargaretsHerts                  = "St Margarets (Herts)"
	StationNameStMargaretsLondon                 = "St Margarets (London)"
	StationNameStMaryCray                        = "St Mary Cray"
	StationNameStMichaels                        = "St Michaels"
	StationNameStNeots                           = "St Neots"
	StationNameStafford                          = "Stafford"
	StationNameStaines                           = "Staines"
	StationNameStallingborough                   = "Stallingborough"
	StationNameStalybridge                       = "Stalybridge"
	StationNameStamfordHill                      = "Stamford Hill"
	StationNameStamfordLincs                     = "Stamford (Lincs)"
	StationNameStanfordleHope                    = "Stanford-le-Hope"
	StationNameStanlowAndThornton                = "Stanlow & Thornton"
	StationNameStanstedAirport                   = "Stansted Airport"
	StationNameStanstedMountfitchet              = "Stansted Mountfitchet"
	StationNameStaplehurst                       = "Staplehurst"
	StationNameStapletonRoad                     = "Stapleton Road"
	StationNameStarbeck                          = "Starbeck"
	StationNameStarcross                         = "Starcross"
	StationNameStaveleyCumbria                   = "Staveley (Cumbria)"
	StationNameStechford                         = "Stechford"
	StationNameSteetonAndSilsden                 = "Steeton & Silsden"
	StationNameStepps                            = "Stepps"
	StationNameStevenage                         = "Stevenage"
	StationNameStevenston                        = "Stevenston"
	StationNameStewartby                         = "Stewartby"
	StationNameStewarton                         = "Stewarton"
	StationNameStirling                          = "Stirling"
	StationNameStockport                         = "Stockport"
	StationNameStocksfield                       = "Stocksfield"
	StationNameStocksmoor                        = "Stocksmoor"
	StationNameStockton                          = "Stockton"
	StationNameStokeMandeville                   = "Stoke Mandeville"
	StationNameStokeNewington                    = "Stoke Newington"
	StationNameStokeonTrent                      = "Stoke-on-Trent"
	StationNameStoneCrossing                     = "Stone Crossing"
	StationNameStoneStaffs                       = "Stone (Staffs)"
	StationNameStonebridgePark                   = "Stonebridge Park"
	StationNameStonegate                         = "Stonegate"
	StationNameStonehaven                        = "Stonehaven"
	StationNameStonehouse                        = "Stonehouse"
	StationNameStoneleigh                        = "Stoneleigh"
	StationNameStourbridgeJunction               = "Stourbridge Junction"
	StationNameStourbridgeTown                   = "Stourbridge Town"
	StationNameStow                              = "Stow"
	StationNameStowmarket                        = "Stowmarket"
	StationNameStranraer                         = "Stranraer"
	StationNameStratfordInternational            = "Stratford International"
	StationNameStratfordLondon                   = "Stratford (London)"
	StationNameStratfordUponAvon                 = "Stratford-upon-Avon"
	StationNameStratfordUponAvonParkway          = "Stratford-upon-Avon Parkway"
	StationNameStrathcarron                      = "Strathcarron"
	StationNameStrawberryHill                    = "Strawberry Hill"
	StationNameStreathamCommon                   = "Streatham Common"
	StationNameStreathamGreaterLondon            = "Streatham (Greater London)"
	StationNameStreathamHill                     = "Streatham Hill"
	StationNameStreethouse                       = "Streethouse"
	StationNameStrines                           = "Strines"
	StationNameStromeferry                       = "Stromeferry"
	StationNameStrood                            = "Strood"
	StationNameStroudGloucs                      = "Stroud (Gloucs)"
	StationNameSturry                            = "Sturry"
	StationNameStyal                             = "Styal"
	StationNameSudburyAndHarrowRoad              = "Sudbury & Harrow Road"
	StationNameSudburyHillHarrow                 = "Sudbury Hill Harrow"
	StationNameSudburySuffolk                    = "Sudbury (Suffolk)"
	StationNameSugarLoaf                         = "Sugar Loaf"
	StationNameSummerston                        = "Summerston"
	StationNameSunbury                           = "Sunbury"
	StationNameSunderland                        = "Sunderland"
	StationNameSundridgePark                     = "Sundridge Park"
	StationNameSunningdale                       = "Sunningdale"
	StationNameSunnymeads                        = "Sunnymeads"
	StationNameSurbiton                          = "Surbiton"
	StationNameSurreyQuays                       = "Surrey Quays"
	StationNameSuttonColdfield                   = "Sutton Coldfield"
	StationNameSuttonCommon                      = "Sutton Common"
	StationNameSuttonParkway                     = "Sutton Parkway"
	StationNameSuttonSurrey                      = "Sutton (Surrey)"
	StationNameSwale                             = "Swale"
	StationNameSwanley                           = "Swanley"
	StationNameSwanscombe                        = "Swanscombe"
	StationNameSwansea                           = "Swansea"
	StationNameSwanwick                          = "Swanwick"
	StationNameSway                              = "Sway"
	StationNameSwaythling                        = "Swaythling"
	StationNameSwinderby                         = "Swinderby"
	StationNameSwindonWilts                      = "Swindon (Wilts)"
	StationNameSwineshead                        = "Swineshead"
	StationNameSwintonManchester                 = "Swinton (Manchester)"
	StationNameSwintonSouthYorks                 = "Swinton (South Yorks)"
	StationNameSydenhamHill                      = "Sydenham Hill"
	StationNameSydenhamLondon                    = "Sydenham (London)"
	StationNameSyonLane                          = "Syon Lane"
	StationNameSyston                            = "Syston"
	StationNameTackley                           = "Tackley"
	StationNameTadworth                          = "Tadworth"
	StationNameTaffsWell                         = "Taffs Well"
	StationNameTain                              = "Tain"
	StationNameTalsarnau                         = "Talsarnau"
	StationNameTalyCafn                          = "Tal-y-Cafn"
	StationNameTalybont                          = "Talybont"
	StationNameTameBridgeParkway                 = "Tame Bridge Parkway"
	StationNameTamworth                          = "Tamworth"
	StationNameTaplow                            = "Taplow"
	StationNameTattenhamCorner                   = "Tattenham Corner"
	StationNameTaunton                           = "Taunton"
	StationNameTaynuilt                          = "Taynuilt"
	StationNameTeddington                        = "Teddington"
	StationNameTeessideAirport                   = "Tees-side Airport"
	StationNameTeignmouth                        = "Teignmouth"
	StationNameTelfordCentral                    = "Telford Central"
	StationNameTemplecombe                       = "Templecombe"
	StationNameTenby                             = "Tenby"
	StationNameTeynham                           = "Teynham"
	StationNameThamesDitton                      = "Thames Ditton"
	StationNameThatcham                          = "Thatcham"
	StationNameThattoHeath                       = "Thatto Heath"
	StationNameTheHawthorns                      = "The Hawthorns"
	StationNameTheLakesWarks                     = "The Lakes (Warks)"
	StationNameTheale                            = "Theale"
	StationNameTheobaldsGrove                    = "Theobalds Grove"
	StationNameThetford                          = "Thetford"
	StationNameThirsk                            = "Thirsk"
	StationNameThornaby                          = "Thornaby"
	StationNameThorneNorth                       = "Thorne North"
	StationNameThorneSouth                       = "Thorne South"
	StationNameThornford                         = "Thornford"
	StationNameThornliebank                      = "Thornliebank"
	StationNameThorntonAbbey                     = "Thornton Abbey"
	StationNameThorntonHeath                     = "Thornton Heath"
	StationNameThorntonhall                      = "Thorntonhall"
	StationNameThorpeBay                         = "Thorpe Bay"
	StationNameThorpeCulvert                     = "Thorpe Culvert"
	StationNameThorpeleSoken                     = "Thorpe-le-Soken"
	StationNameThreeBridges                      = "Three Bridges"
	StationNameThreeOaks                         = "Three Oaks"
	StationNameThurgarton                        = "Thurgarton"
	StationNameThurnscoe                         = "Thurnscoe"
	StationNameThurso                            = "Thurso"
	StationNameThurston                          = "Thurston"
	StationNameTilburyTown                       = "Tilbury Town"
	StationNameTileHill                          = "Tile Hill"
	StationNameTilehurst                         = "Tilehurst"
	StationNameTipton                            = "Tipton"
	StationNameTirPhil                           = "Tir-Phil"
	StationNameTisbury                           = "Tisbury"
	StationNameTivertonParkway                   = "Tiverton Parkway"
	StationNameTodmorden                         = "Todmorden"
	StationNameTolworth                          = "Tolworth"
	StationNameTonPentre                         = "Ton Pentre"
	StationNameTonbridge                         = "Tonbridge"
	StationNameTondu                             = "Tondu"
	StationNameTonfanau                          = "Tonfanau"
	StationNameTonypandy                         = "Tonypandy"
	StationNameTooting                           = "Tooting"
	StationNameTopsham                           = "Topsham"
	StationNameTorquay                           = "Torquay"
	StationNameTorre                             = "Torre"
	StationNameTotnes                            = "Totnes"
	StationNameTottenhamHale                     = "Tottenham Hale"
	StationNameTotton                            = "Totton"
	StationNameTownGreen                         = "Town Green"
	StationNameTraffordPark                      = "Trafford Park"
	StationNameTrefforest                        = "Trefforest"
	StationNameTrefforestEstate                  = "Trefforest Estate"
	StationNameTrehafod                          = "Trehafod"
	StationNameTreherbert                        = "Treherbert"
	StationNameTreorchy                          = "Treorchy"
	StationNameTrimley                           = "Trimley"
	StationNameTring                             = "Tring"
	StationNameTroedyrhiw                        = "Troed-y-rhiw"
	StationNameTroon                             = "Troon"
	StationNameTrowbridge                        = "Trowbridge"
	StationNameTruro                             = "Truro"
	StationNameTulloch                           = "Tulloch"
	StationNameTulseHill                         = "Tulse Hill"
	StationNameTunbridgeWells                    = "Tunbridge Wells"
	StationNameTurkeyStreet                      = "Turkey Street"
	StationNameTutburyAndHatton                  = "Tutbury & Hatton"
	StationNameTweedbank                         = "Tweedbank"
	StationNameTwickenham                        = "Twickenham"
	StationNameTwyford                           = "Twyford"
	StationNameTyCroes                           = "Ty Croes"
	StationNameTyGlas                            = "Ty Glas"
	StationNameTygwyn                            = "Tygwyn"
	StationNameTyndrumLower                      = "Tyndrum Lower"
	StationNameTyseley                           = "Tyseley"
	StationNameTywyn                             = "Tywyn"
	StationNameUckfield                          = "Uckfield"
	StationNameUddingston                        = "Uddingston"
	StationNameUlceby                            = "Ulceby"
	StationNameUlleskelf                         = "Ulleskelf"
	StationNameUlverston                         = "Ulverston"
	StationNameUmberleigh                        = "Umberleigh"
	StationNameUniversityBirmingham              = "University (Birmingham)"
	StationNameUphall                            = "Uphall"
	StationNameUpholland                         = "Upholland"
	StationNameUpminster                         = "Upminster"
	StationNameUpperHalliford                    = "Upper Halliford"
	StationNameUpperHolloway                     = "Upper Holloway"
	StationNameUpperTyndrum                      = "Upper Tyndrum"
	StationNameUpperWarlingham                   = "Upper Warlingham"
	StationNameUptonMerseyside                   = "Upton (Merseyside)"
	StationNameUpwey                             = "Upwey"
	StationNameUrmston                           = "Urmston"
	StationNameUttoxeter                         = "Uttoxeter"
	StationNameValley                            = "Valley"
	StationNameVauxhall                          = "Vauxhall"
	StationNameVirginiaWater                     = "Virginia Water"
	StationNameWaddon                            = "Waddon"
	StationNameWadhurst                          = "Wadhurst"
	StationNameWainfleet                         = "Wainfleet"
	StationNameWakefieldKirkgate                 = "Wakefield Kirkgate"
	StationNameWakefieldWestgate                 = "Wakefield Westgate"
	StationNameWalkden                           = "Walkden"
	StationNameWallaseyGroveRoad                 = "Wallasey Grove Road"
	StationNameWallaseyVillage                   = "Wallasey Village"
	StationNameWallington                        = "Wallington"
	StationNameWallyford                         = "Wallyford"
	StationNameWalmer                            = "Walmer"
	StationNameWalsall                           = "Walsall"
	StationNameWalsden                           = "Walsden"
	StationNameWalthamCross                      = "Waltham Cross"
	StationNameWalthamstowCentral                = "Walthamstow Central"
	StationNameWalthamstowQueensRoad             = "Walthamstow Queen's Road"
	StationNameWaltonMerseyside                  = "Walton (Merseyside)"
	StationNameWaltononThames                    = "Walton-on-Thames"
	StationNameWaltonontheNaze                   = "Walton-on-the-Naze"
	StationNameWanborough                        = "Wanborough"
	StationNameWandsworthCommon                  = "Wandsworth Common"
	StationNameWandsworthRoad                    = "Wandsworth Road"
	StationNameWandsworthTown                    = "Wandsworth Town"
	StationNameWansteadPark                      = "Wanstead Park"
	StationNameWapping                           = "Wapping"
	StationNameWarblington                       = "Warblington"
	StationNameWareHerts                         = "Ware (Herts)"
	StationNameWarehamDorset                     = "Wareham (Dorset)"
	StationNameWargrave                          = "Wargrave"
	StationNameWarminster                        = "Warminster"
	StationNameWarnham                           = "Warnham"
	StationNameWarringtonBankQuay                = "Warrington Bank Quay"
	StationNameWarringtonCentral                 = "Warrington Central"
	StationNameWarwick                           = "Warwick"
	StationNameWarwickParkway                    = "Warwick Parkway"
	StationNameWaterOrton                        = "Water Orton"
	StationNameWaterbeach                        = "Waterbeach"
	StationNameWateringbury                      = "Wateringbury"
	StationNameWaterlooMerseyside                = "Waterloo (Merseyside)"
	StationNameWatfordHighStreet                 = "Watford High Street"
	StationNameWatfordJunction                   = "Watford Junction"
	StationNameWatfordNorth                      = "Watford North"
	StationNameWatlington                        = "Watlington"
	StationNameWattonatStone                     = "Watton-at-Stone"
	StationNameWaunGronPark                      = "Waun-Gron Park"
	StationNameWavertreeTechnologyPark           = "Wavertree Technology Park"
	StationNameWedgwood                          = "Wedgwood"
	StationNameWeeley                            = "Weeley"
	StationNameWeeton                            = "Weeton"
	StationNameWelhamGreen                       = "Welham Green"
	StationNameWelling                           = "Welling"
	StationNameWellingborough                    = "Wellingborough"
	StationNameWellingtonShropshire              = "Wellington (Shropshire)"
	StationNameWelshpool                         = "Welshpool"
	StationNameWelwynGardenCity                  = "Welwyn Garden City"
	StationNameWelwynNorth                       = "Welwyn North"
	StationNameWem                               = "Wem"
	StationNameWembleyCentral                    = "Wembley Central"
	StationNameWembleyStadium                    = "Wembley Stadium"
	StationNameWemyssBay                         = "Wemyss Bay"
	StationNameWendover                          = "Wendover"
	StationNameWennington                        = "Wennington"
	StationNameWestAllerton                      = "West Allerton"
	StationNameWestBrompton                      = "West Brompton"
	StationNameWestByfleet                       = "West Byfleet"
	StationNameWestCalder                        = "West Calder"
	StationNameWestCroydon                       = "West Croydon"
	StationNameWestDrayton                       = "West Drayton"
	StationNameWestDulwich                       = "West Dulwich"
	StationNameWestEaling                        = "West Ealing"
	StationNameWestHam                           = "West Ham"
	StationNameWestHampstead                     = "West Hampstead"
	StationNameWestHampsteadThameslink           = "West Hampstead Thameslink"
	StationNameWestHorndon                       = "West Horndon"
	StationNameWestKilbride                      = "West Kilbride"
	StationNameWestKirby                         = "West Kirby"
	StationNameWestMalling                       = "West Malling"
	StationNameWestNorwood                       = "West Norwood"
	StationNameWestRuislip                       = "West Ruislip"
	StationNameWestRunton                        = "West Runton"
	StationNameWestStLeonards                    = "West St Leonards"
	StationNameWestSutton                        = "West Sutton"
	StationNameWestWickham                       = "West Wickham"
	StationNameWestWorthing                      = "West Worthing"
	StationNameWestburyWilts                     = "Westbury (Wilts)"
	StationNameWestcliff                         = "Westcliff"
	StationNameWestcombePark                     = "Westcombe Park"
	StationNameWestenhanger                      = "Westenhanger"
	StationNameWesterHailes                      = "Wester Hailes"
	StationNameWesterfield                       = "Westerfield"
	StationNameWesterton                         = "Westerton"
	StationNameWestgateOnSea                     = "Westgate-on-Sea"
	StationNameWesthoughton                      = "Westhoughton"
	StationNameWestonMilton                      = "Weston Milton"
	StationNameWestonsuperMare                   = "Weston-super-Mare"
	StationNameWetheral                          = "Wetheral"
	StationNameWeybridge                         = "Weybridge"
	StationNameWeymouth                          = "Weymouth"
	StationNameWhaleyBridge                      = "Whaley Bridge"
	StationNameWhalleyLancs                      = "Whalley (Lancs)"
	StationNameWhatstandwell                     = "Whatstandwell"
	StationNameWhifflet                          = "Whifflet"
	StationNameWhimple                           = "Whimple"
	StationNameWhinhill                          = "Whinhill"
	StationNameWhiston                           = "Whiston"
	StationNameWhitby                            = "Whitby"
	StationNameWhitchurchCardiff                 = "Whitchurch (Cardiff)"
	StationNameWhitchurchHants                   = "Whitchurch (Hants)"
	StationNameWhitchurchShropshire              = "Whitchurch (Shropshire)"
	StationNameWhiteHartLane                     = "White Hart Lane"
	StationNameWhiteNotley                       = "White Notley"
	StationNameWhitechapel                       = "Whitechapel"
	StationNameWhitecraigs                       = "Whitecraigs"
	StationNameWhitehaven                        = "Whitehaven"
	StationNameWhitland                          = "Whitland"
	StationNameWhitleyBridge                     = "Whitley Bridge"
	StationNameWhitlocksEnd                      = "Whitlocks End"
	StationNameWhitstable                        = "Whitstable"
	StationNameWhittlesea                        = "Whittlesea"
	StationNameWhittlesfordParkway               = "Whittlesford Parkway"
	StationNameWhittonLondon                     = "Whitton (London)"
	StationNameWhitwellDerbyshire                = "Whitwell (Derbyshire)"
	StationNameWhyteleafe                        = "Whyteleafe"
	StationNameWhyteleafeSouth                   = "Whyteleafe South"
	StationNameWick                              = "Wick"
	StationNameWickford                          = "Wickford"
	StationNameWickhamMarket                     = "Wickham Market"
	StationNameWiddrington                       = "Widdrington"
	StationNameWidnes                            = "Widnes"
	StationNameWidneyManor                       = "Widney Manor"
	StationNameWiganNorthWestern                 = "Wigan North Western"
	StationNameWiganWallgate                     = "Wigan Wallgate"
	StationNameWigton                            = "Wigton"
	StationNameWildmill                          = "Wildmill"
	StationNameWillesdenJunction                 = "Willesden Junction"
	StationNameWilliamwood                       = "Williamwood"
	StationNameWillington                        = "Willington"
	StationNameWilmcote                          = "Wilmcote"
	StationNameWilmslow                          = "Wilmslow"
	StationNameWilnecoteStaffs                   = "Wilnecote (Staffs)"
	StationNameWimbledon                         = "Wimbledon"
	StationNameWimbledonChase                    = "Wimbledon Chase"
	StationNameWinchelsea                        = "Winchelsea"
	StationNameWinchester                        = "Winchester"
	StationNameWinchfield                        = "Winchfield"
	StationNameWinchmoreHill                     = "Winchmore Hill"
	StationNameWindermere                        = "Windermere"
	StationNameWindsorAndEtonCentral             = "Windsor & Eton Central"
	StationNameWindsorAndEtonRiverside           = "Windsor & Eton Riverside"
	StationNameWinnersh                          = "Winnersh"
	StationNameWinnershTriangle                  = "Winnersh Triangle"
	StationNameWinsford                          = "Winsford"
	StationNameWishaw                            = "Wishaw"
	StationNameWitham                            = "Witham"
	StationNameWitley                            = "Witley"
	StationNameWittonWestMidlands                = "Witton (West Midlands)"
	StationNameWivelsfield                       = "Wivelsfield"
	StationNameWivenhoe                          = "Wivenhoe"
	StationNameWoburnSands                       = "Woburn Sands"
	StationNameWoking                            = "Woking"
	StationNameWokingham                         = "Wokingham"
	StationNameWoldingham                        = "Woldingham"
	StationNameWolverhampton                     = "Wolverhampton"
	StationNameWolverton                         = "Wolverton"
	StationNameWombwell                          = "Wombwell"
	StationNameWoodEnd                           = "Wood End"
	StationNameWoodStreet                        = "Wood Street"
	StationNameWoodbridge                        = "Woodbridge"
	StationNameWoodgrangePark                    = "Woodgrange Park"
	StationNameWoodhall                          = "Woodhall"
	StationNameWoodhouse                         = "Woodhouse"
	StationNameWoodlesford                       = "Woodlesford"
	StationNameWoodley                           = "Woodley"
	StationNameWoodmansterne                     = "Woodmansterne"
	StationNameWoodsmoor                         = "Woodsmoor"
	StationNameWool                              = "Wool"
	StationNameWoolston                          = "Woolston"
	StationNameWoolwichArsenal                   = "Woolwich Arsenal"
	StationNameWoolwichDockyard                  = "Woolwich Dockyard"
	StationNameWoottonWawen                      = "Wootton Wawen"
	StationNameWorcesterForegateStreet           = "Worcester Foregate Street"
	StationNameWorcesterPark                     = "Worcester Park"
	StationNameWorcesterShrubHill                = "Worcester Shrub Hill"
	StationNameWorkington                        = "Workington"
	StationNameWorksop                           = "Worksop"
	StationNameWorle                             = "Worle"
	StationNameWorplesdon                        = "Worplesdon"
	StationNameWorstead                          = "Worstead"
	StationNameWorthing                          = "Worthing"
	StationNameWrabness                          = "Wrabness"
	StationNameWraysbury                         = "Wraysbury"
	StationNameWrenbury                          = "Wrenbury"
	StationNameWressle                           = "Wressle"
	StationNameWrexhamCentral                    = "Wrexham Central"
	StationNameWrexhamGeneral                    = "Wrexham General"
	StationNameWye                               = "Wye"
	StationNameWylam                             = "Wylam"
	StationNameWyldeGreen                        = "Wylde Green"
	StationNameWymondham                         = "Wymondham"
	StationNameWythall                           = "Wythall"
	StationNameYalding                           = "Yalding"
	StationNameYardleyWood                       = "Yardley Wood"
	StationNameYarm                              = "Yarm"
	StationNameYate                              = "Yate"
	StationNameYatton                            = "Yatton"
	StationNameYeoford                           = "Yeoford"
	StationNameYeovilJunction                    = "Yeovil Junction"
	StationNameYeovilPenMill                     = "Yeovil Pen Mill"
	StationNameYetminster                        = "Yetminster"
	StationNameYnyswen                           = "Ynyswen"
	StationNameYoker                             = "Yoker"
	StationNameYork                              = "York"
	StationNameYorton                            = "Yorton"
	StationNameYstradMynach                      = "Ystrad Mynach"
	StationNameYstradRhondda                     = "Ystrad Rhondda"
)

// CRSCode Computer Reservation System Code (CRS code) refers to
// a code used to identify railway stations in the United Kingdom.
// Each train station in the UK has a unique three-letter CRS code,
// which is used for ticketing, timetabling, and reservation purposes.
// For example, the CRS code for London Paddington is PAD. These codes
// are essential for the efficient operation of the rail network.
type CRSCode string

const (
	StationCodeAbbeyWood                         CRSCode = "ABW"
	StationCodeAber                              CRSCode = "ABE"
	StationCodeAbercynon                         CRSCode = "ACY"
	StationCodeAberdare                          CRSCode = "ABA"
	StationCodeAberdeen                          CRSCode = "ABD"
	StationCodeAberdour                          CRSCode = "AUR"
	StationCodeAberdovey                         CRSCode = "AVY"
	StationCodeAbererch                          CRSCode = "ABH"
	StationCodeAbergavenny                       CRSCode = "AGV"
	StationCodeAbergeleAndPensarn                CRSCode = "AGL"
	StationCodeAberystwyth                       CRSCode = "AYW"
	StationCodeAccrington                        CRSCode = "ACR"
	StationCodeAchanalt                          CRSCode = "AAT"
	StationCodeAchnasheen                        CRSCode = "ACN"
	StationCodeAchnashellach                     CRSCode = "ACH"
	StationCodeAcklington                        CRSCode = "ACK"
	StationCodeAcle                              CRSCode = "ACL"
	StationCodeAcocksGreen                       CRSCode = "ACG"
	StationCodeActonBridgeCheshire               CRSCode = "ACB"
	StationCodeActonCentral                      CRSCode = "ACC"
	StationCodeActonMainLine                     CRSCode = "AML"
	StationCodeAdderleyPark                      CRSCode = "ADD"
	StationCodeAddiewell                         CRSCode = "ADW"
	StationCodeAddlestone                        CRSCode = "ASN"
	StationCodeAdisham                           CRSCode = "ADM"
	StationCodeAdlingtonCheshire                 CRSCode = "ADC"
	StationCodeAdlingtonLancs                    CRSCode = "ADL"
	StationCodeAdwick                            CRSCode = "AWK"
	StationCodeAigburth                          CRSCode = "AIG"
	StationCodeAinsdale                          CRSCode = "ANS"
	StationCodeAintree                           CRSCode = "AIN"
	StationCodeAirbles                           CRSCode = "AIR"
	StationCodeAirdrie                           CRSCode = "ADR"
	StationCodeAlbanyPark                        CRSCode = "AYP"
	StationCodeAlbrighton                        CRSCode = "ALB"
	StationCodeAlderleyEdge                      CRSCode = "ALD"
	StationCodeAldermaston                       CRSCode = "AMT"
	StationCodeAldershot                         CRSCode = "AHT"
	StationCodeAldrington                        CRSCode = "AGT"
	StationCodeAlexandraPalace                   CRSCode = "AAP"
	StationCodeAlexandraParade                   CRSCode = "AXP"
	StationCodeAlexandria                        CRSCode = "ALX"
	StationCodeAlfreton                          CRSCode = "ALF"
	StationCodeAllensWest                        CRSCode = "ALW"
	StationCodeAlloa                             CRSCode = "ALO"
	StationCodeAlness                            CRSCode = "ASS"
	StationCodeAlnmouth                          CRSCode = "ALM"
	StationCodeAlresfordEssex                    CRSCode = "ALR"
	StationCodeAlsager                           CRSCode = "ASG"
	StationCodeAlthorneEssex                     CRSCode = "ALN"
	StationCodeAlthorpe                          CRSCode = "ALP"
	StationCodeAltnabreac                        CRSCode = "ABC"
	StationCodeAlton                             CRSCode = "AON"
	StationCodeAltrincham                        CRSCode = "ALT"
	StationCodeAlvechurch                        CRSCode = "ALV"
	StationCodeAmbergate                         CRSCode = "AMB"
	StationCodeAmberley                          CRSCode = "AMY"
	StationCodeAmersham                          CRSCode = "AMR"
	StationCodeAmmanford                         CRSCode = "AMF"
	StationCodeAncaster                          CRSCode = "ANC"
	StationCodeAnderston                         CRSCode = "AND"
	StationCodeAndover                           CRSCode = "ADV"
	StationCodeAnerley                           CRSCode = "ANZ"
	StationCodeAngelRoad                         CRSCode = "AGR"
	StationCodeAngmering                         CRSCode = "ANG"
	StationCodeAnnan                             CRSCode = "ANN"
	StationCodeAnniesland                        CRSCode = "ANL"
	StationCodeAnsdellAndFairhaven               CRSCode = "AFV"
	StationCodeApperleyBridge                    CRSCode = "APY"
	StationCodeAppleby                           CRSCode = "APP"
	StationCodeAppledoreKent                     CRSCode = "APD"
	StationCodeAppleford                         CRSCode = "APF"
	StationCodeAppleyBridge                      CRSCode = "APB"
	StationCodeApsley                            CRSCode = "APS"
	StationCodeArbroath                          CRSCode = "ARB"
	StationCodeArdgay                            CRSCode = "ARD"
	StationCodeArdlui                            CRSCode = "AUI"
	StationCodeArdrossanHarbour                  CRSCode = "ADS"
	StationCodeArdrossanSouthBeach               CRSCode = "ASB"
	StationCodeArdrossanTown                     CRSCode = "ADN"
	StationCodeArdwick                           CRSCode = "ADK"
	StationCodeArgyleStreet                      CRSCode = "AGS"
	StationCodeArisaig                           CRSCode = "ARG"
	StationCodeArlesey                           CRSCode = "ARL"
	StationCodeArmadaleWestLothian               CRSCode = "ARM"
	StationCodeArmathwaite                       CRSCode = "AWT"
	StationCodeArnside                           CRSCode = "ARN"
	StationCodeArram                             CRSCode = "ARR"
	StationCodeArrocharAndTarbet                 CRSCode = "ART"
	StationCodeArundel                           CRSCode = "ARU"
	StationCodeAscotBerks                        CRSCode = "ACT"
	StationCodeAscottunderWychwood               CRSCode = "AUW"
	StationCodeAsh                               CRSCode = "ASH"
	StationCodeAshVale                           CRSCode = "AHV"
	StationCodeAshburys                          CRSCode = "ABY"
	StationCodeAshchurchforTewkesbury            CRSCode = "ASC"
	StationCodeAshfield                          CRSCode = "ASF"
	StationCodeAshfordInternational              CRSCode = "AFK"
	StationCodeAshfordInternationalEurostar      CRSCode = "ASI"
	StationCodeAshfordSurrey                     CRSCode = "AFS"
	StationCodeAshley                            CRSCode = "ASY"
	StationCodeAshtead                           CRSCode = "AHD"
	StationCodeAshtonunderLyne                   CRSCode = "AHN"
	StationCodeAshurstKent                       CRSCode = "AHS"
	StationCodeAshurstNewForest                  CRSCode = "ANF"
	StationCodeAshwellAndMorden                  CRSCode = "AWM"
	StationCodeAskam                             CRSCode = "ASK"
	StationCodeAslockton                         CRSCode = "ALK"
	StationCodeAspatria                          CRSCode = "ASP"
	StationCodeAspleyGuise                       CRSCode = "APG"
	StationCodeAston                             CRSCode = "AST"
	StationCodeAtherstone                        CRSCode = "ATH"
	StationCodeAtherton                          CRSCode = "ATN"
	StationCodeAttadale                          CRSCode = "ATT"
	StationCodeAttenborough                      CRSCode = "ATB"
	StationCodeAttleborough                      CRSCode = "ATL"
	StationCodeAuchinleck                        CRSCode = "AUK"
	StationCodeAudleyEnd                         CRSCode = "AUD"
	StationCodeAughtonPark                       CRSCode = "AUG"
	StationCodeAviemore                          CRSCode = "AVM"
	StationCodeAvoncliff                         CRSCode = "AVF"
	StationCodeAvonmouth                         CRSCode = "AVN"
	StationCodeAxminster                         CRSCode = "AXM"
	StationCodeAylesbury                         CRSCode = "AYS"
	StationCodeAylesburyValeParkway              CRSCode = "AVP"
	StationCodeAylesford                         CRSCode = "AYL"
	StationCodeAylesham                          CRSCode = "AYH"
	StationCodeAyr                               CRSCode = "AYR"
	StationCodeBache                             CRSCode = "BAC"
	StationCodeBaglan                            CRSCode = "BAJ"
	StationCodeBagshot                           CRSCode = "BAG"
	StationCodeBaildon                           CRSCode = "BLD"
	StationCodeBaillieston                       CRSCode = "BIO"
	StationCodeBalcombe                          CRSCode = "BAB"
	StationCodeBaldock                           CRSCode = "BDK"
	StationCodeBalham                            CRSCode = "BAL"
	StationCodeBalloch                           CRSCode = "BHC"
	StationCodeBalmossie                         CRSCode = "BSI"
	StationCodeBamberBridge                      CRSCode = "BMB"
	StationCodeBamford                           CRSCode = "BAM"
	StationCodeBanavie                           CRSCode = "BNV"
	StationCodeBanbury                           CRSCode = "BAN"
	StationCodeBangorGwynedd                     CRSCode = "BNG"
	StationCodeBankHall                          CRSCode = "BAH"
	StationCodeBanstead                          CRSCode = "BAD"
	StationCodeBarassie                          CRSCode = "BSS"
	StationCodeBarbican                          CRSCode = "ZBB"
	StationCodeBardonMill                        CRSCode = "BLL"
	StationCodeBareLane                          CRSCode = "BAR"
	StationCodeBargeddie                         CRSCode = "BGI"
	StationCodeBargoed                           CRSCode = "BGD"
	StationCodeBarking                           CRSCode = "BKG"
	StationCodeBarlaston                         CRSCode = "BRT"
	StationCodeBarming                           CRSCode = "BMG"
	StationCodeBarmouth                          CRSCode = "BRM"
	StationCodeBarnehurst                        CRSCode = "BNH"
	StationCodeBarnes                            CRSCode = "BNS"
	StationCodeBarnesBridge                      CRSCode = "BNI"
	StationCodeBarnetby                          CRSCode = "BTB"
	StationCodeBarnham                           CRSCode = "BAA"
	StationCodeBarnhill                          CRSCode = "BNL"
	StationCodeBarnsley                          CRSCode = "BNY"
	StationCodeBarnstaple                        CRSCode = "BNP"
	StationCodeBarntGreen                        CRSCode = "BTG"
	StationCodeBarrhead                          CRSCode = "BRR"
	StationCodeBarrhill                          CRSCode = "BRL"
	StationCodeBarrowHaven                       CRSCode = "BAV"
	StationCodeBarrowUponSoar                    CRSCode = "BWS"
	StationCodeBarrowinFurness                   CRSCode = "BIF"
	StationCodeBarry                             CRSCode = "BRY"
	StationCodeBarryDocks                        CRSCode = "BYD"
	StationCodeBarryIsland                       CRSCode = "BYI"
	StationCodeBarryLinks                        CRSCode = "BYL"
	StationCodeBartononHumber                    CRSCode = "BAU"
	StationCodeBasildon                          CRSCode = "BSO"
	StationCodeBasingstoke                       CRSCode = "BSK"
	StationCodeBatAndBall                        CRSCode = "BBL"
	StationCodeBathSpa                           CRSCode = "BTH"
	StationCodeBathgate                          CRSCode = "BHG"
	StationCodeBatley                            CRSCode = "BTL"
	StationCodeBattersby                         CRSCode = "BTT"
	StationCodeBatterseaPark                     CRSCode = "BAK"
	StationCodeBattle                            CRSCode = "BAT"
	StationCodeBattlesbridge                     CRSCode = "BLB"
	StationCodeBayford                           CRSCode = "BAY"
	StationCodeBeaconsfield                      CRSCode = "BCF"
	StationCodeBearley                           CRSCode = "BER"
	StationCodeBearsden                          CRSCode = "BRN"
	StationCodeBearsted                          CRSCode = "BSD"
	StationCodeBeasdale                          CRSCode = "BSL"
	StationCodeBeaulieuRoad                      CRSCode = "BEU"
	StationCodeBeauly                            CRSCode = "BEL"
	StationCodeBebington                         CRSCode = "BEB"
	StationCodeBeccles                           CRSCode = "BCC"
	StationCodeBeckenhamHill                     CRSCode = "BEC"
	StationCodeBeckenhamJunction                 CRSCode = "BKJ"
	StationCodeBedford                           CRSCode = "BDM"
	StationCodeBedfordStJohns                    CRSCode = "BSJ"
	StationCodeBedhampton                        CRSCode = "BDH"
	StationCodeBedminster                        CRSCode = "BMT"
	StationCodeBedworth                          CRSCode = "BEH"
	StationCodeBedwyn                            CRSCode = "BDW"
	StationCodeBeeston                           CRSCode = "BEE"
	StationCodeBekesbourne                       CRSCode = "BKS"
	StationCodeBelleVue                          CRSCode = "BLV"
	StationCodeBellgrove                         CRSCode = "BLG"
	StationCodeBellingham                        CRSCode = "BGM"
	StationCodeBellshill                         CRSCode = "BLH"
	StationCodeBelmont                           CRSCode = "BLM"
	StationCodeBelper                            CRSCode = "BLP"
	StationCodeBeltring                          CRSCode = "BEG"
	StationCodeBelvedere                         CRSCode = "BVD"
	StationCodeBempton                           CRSCode = "BEM"
	StationCodeBenRhydding                       CRSCode = "BEY"
	StationCodeBenfleet                          CRSCode = "BEF"
	StationCodeBentham                           CRSCode = "BEN"
	StationCodeBentleyHants                      CRSCode = "BTY"
	StationCodeBentleySouthYorks                 CRSCode = "BYK"
	StationCodeBereAlston                        CRSCode = "BAS"
	StationCodeBereFerrers                       CRSCode = "BFE"
	StationCodeBerkhamsted                       CRSCode = "BKM"
	StationCodeBerkswell                         CRSCode = "BKW"
	StationCodeBermudaPark                       CRSCode = "BEP"
	StationCodeBerneyArms                        CRSCode = "BYA"
	StationCodeBerryBrow                         CRSCode = "BBW"
	StationCodeBerrylands                        CRSCode = "BRS"
	StationCodeBerwickSussex                     CRSCode = "BRK"
	StationCodeBerwickUponTweed                  CRSCode = "BWK"
	StationCodeBescarLane                        CRSCode = "BES"
	StationCodeBescotStadium                     CRSCode = "BSC"
	StationCodeBetchworth                        CRSCode = "BTO"
	StationCodeBethnalGreen                      CRSCode = "BET"
	StationCodeBetwsyCoed                        CRSCode = "BYC"
	StationCodeBeverley                          CRSCode = "BEV"
	StationCodeBexhill                           CRSCode = "BEX"
	StationCodeBexley                            CRSCode = "BXY"
	StationCodeBexleyheath                       CRSCode = "BXH"
	StationCodeBicesterNorth                     CRSCode = "BCS"
	StationCodeBicesterVillage                   CRSCode = "BIT"
	StationCodeBickley                           CRSCode = "BKL"
	StationCodeBidston                           CRSCode = "BID"
	StationCodeBiggleswade                       CRSCode = "BIW"
	StationCodeBilbrook                          CRSCode = "BBK"
	StationCodeBillericay                        CRSCode = "BIC"
	StationCodeBillinghamCleveland               CRSCode = "BIL"
	StationCodeBillingshurst                     CRSCode = "BIG"
	StationCodeBingham                           CRSCode = "BIN"
	StationCodeBingley                           CRSCode = "BIY"
	StationCodeBirchgrove                        CRSCode = "BCG"
	StationCodeBirchingtonOnSea                  CRSCode = "BCH"
	StationCodeBirchwood                         CRSCode = "BWD"
	StationCodeBirkbeck                          CRSCode = "BIK"
	StationCodeBirkdale                          CRSCode = "BDL"
	StationCodeBirkenheadCentral                 CRSCode = "BKC"
	StationCodeBirkenheadHamiltonSquare          CRSCode = "BKQ"
	StationCodeBirkenheadNorth                   CRSCode = "BKN"
	StationCodeBirkenheadPark                    CRSCode = "BKP"
	StationCodeBirminghamInternational           CRSCode = "BHI"
	StationCodeBirminghamMoorStreet              CRSCode = "BMO"
	StationCodeBirminghamNewStreet               CRSCode = "BHM"
	StationCodeBirminghamSnowHill                CRSCode = "BSW"
	StationCodeBishopAuckland                    CRSCode = "BIA"
	StationCodeBishopbriggs                      CRSCode = "BBG"
	StationCodeBishopsStortford                  CRSCode = "BIS"
	StationCodeBishopstoneSussex                 CRSCode = "BIP"
	StationCodeBishoptonStrathclyde              CRSCode = "BPT"
	StationCodeBitterne                          CRSCode = "BTE"
	StationCodeBlackburn                         CRSCode = "BBN"
	StationCodeBlackheath                        CRSCode = "BKH"
	StationCodeBlackhorseRoad                    CRSCode = "BHO"
	StationCodeBlackpoolNorth                    CRSCode = "BPN"
	StationCodeBlackpoolPleasureBeach            CRSCode = "BPB"
	StationCodeBlackpoolSouth                    CRSCode = "BPS"
	StationCodeBlackridge                        CRSCode = "BKR"
	StationCodeBlackrod                          CRSCode = "BLK"
	StationCodeBlackwater                        CRSCode = "BAW"
	StationCodeBlaenauFfestiniog                 CRSCode = "BFF"
	StationCodeBlairAtholl                       CRSCode = "BLA"
	StationCodeBlairhill                         CRSCode = "BAI"
	StationCodeBlakeStreet                       CRSCode = "BKT"
	StationCodeBlakedown                         CRSCode = "BKD"
	StationCodeBlantyre                          CRSCode = "BLT"
	StationCodeBlaydon                           CRSCode = "BLO"
	StationCodeBleasby                           CRSCode = "BSB"
	StationCodeBletchley                         CRSCode = "BLY"
	StationCodeBloxwich                          CRSCode = "BLX"
	StationCodeBloxwichNorth                     CRSCode = "BWN"
	StationCodeBlundellsandsAndCrosby            CRSCode = "BLN"
	StationCodeBlytheBridge                      CRSCode = "BYB"
	StationCodeBodminParkway                     CRSCode = "BOD"
	StationCodeBodorgan                          CRSCode = "BOR"
	StationCodeBognorRegis                       CRSCode = "BOG"
	StationCodeBogston                           CRSCode = "BGS"
	StationCodeBolton                            CRSCode = "BON"
	StationCodeBoltonUponDearne                  CRSCode = "BTD"
	StationCodeBookham                           CRSCode = "BKA"
	StationCodeBootleCumbria                     CRSCode = "BOC"
	StationCodeBootleNewStrand                   CRSCode = "BNW"
	StationCodeBootleOrielRoad                   CRSCode = "BOT"
	StationCodeBordesley                         CRSCode = "BBS"
	StationCodeBoroughGreenAndWrotham            CRSCode = "BRG"
	StationCodeBorth                             CRSCode = "BRH"
	StationCodeBosham                            CRSCode = "BOH"
	StationCodeBoston                            CRSCode = "BSN"
	StationCodeBotley                            CRSCode = "BOE"
	StationCodeBottesford                        CRSCode = "BTF"
	StationCodeBourneEnd                         CRSCode = "BNE"
	StationCodeBournemouth                       CRSCode = "BMH"
	StationCodeBournville                        CRSCode = "BRV"
	StationCodeBowBrickhill                      CRSCode = "BWB"
	StationCodeBowesPark                         CRSCode = "BOP"
	StationCodeBowling                           CRSCode = "BWG"
	StationCodeBoxHillAndWesthumble              CRSCode = "BXW"
	StationCodeBracknell                         CRSCode = "BCE"
	StationCodeBradfordForsterSquare             CRSCode = "BDQ"
	StationCodeBradfordInterchange               CRSCode = "BDI"
	StationCodeBradfordonAvon                    CRSCode = "BOA"
	StationCodeBrading                           CRSCode = "BDN"
	StationCodeBraintree                         CRSCode = "BTR"
	StationCodeBraintreeFreeport                 CRSCode = "BTP"
	StationCodeBramhall                          CRSCode = "BML"
	StationCodeBramleyHants                      CRSCode = "BMY"
	StationCodeBramleyWYorks                     CRSCode = "BLE"
	StationCodeBramptonCumbria                   CRSCode = "BMP"
	StationCodeBramptonSuffolk                   CRSCode = "BRP"
	StationCodeBranchton                         CRSCode = "BCN"
	StationCodeBrandon                           CRSCode = "BND"
	StationCodeBranksome                         CRSCode = "BSM"
	StationCodeBraystonesCumbria                 CRSCode = "BYS"
	StationCodeBredbury                          CRSCode = "BDY"
	StationCodeBreich                            CRSCode = "BRC"
	StationCodeBrentford                         CRSCode = "BFD"
	StationCodeBrentwood                         CRSCode = "BRE"
	StationCodeBricketWood                       CRSCode = "BWO"
	StationCodeBridgend                          CRSCode = "BGN"
	StationCodeBridgeofAllan                     CRSCode = "BEA"
	StationCodeBridgeofOrchy                     CRSCode = "BRO"
	StationCodeBridgeton                         CRSCode = "BDG"
	StationCodeBridgwater                        CRSCode = "BWT"
	StationCodeBridlington                       CRSCode = "BDT"
	StationCodeBrierfield                        CRSCode = "BRF"
	StationCodeBrigg                             CRSCode = "BGG"
	StationCodeBrighouse                         CRSCode = "BGH"
	StationCodeBrightonEastSussex                CRSCode = "BTN"
	StationCodeBrimsdown                         CRSCode = "BMD"
	StationCodeBrinnington                       CRSCode = "BNT"
	StationCodeBristolParkway                    CRSCode = "BPW"
	StationCodeBristolTempleMeads                CRSCode = "BRI"
	StationCodeBrithdir                          CRSCode = "BHD"
	StationCodeBritonFerry                       CRSCode = "BNF"
	StationCodeBrixton                           CRSCode = "BRX"
	StationCodeBroadGreen                        CRSCode = "BGE"
	StationCodeBroadbottom                       CRSCode = "BDB"
	StationCodeBroadstairs                       CRSCode = "BSR"
	StationCodeBrockenhurst                      CRSCode = "BCU"
	StationCodeBrockholes                        CRSCode = "BHS"
	StationCodeBrockley                          CRSCode = "BCY"
	StationCodeBromborough                       CRSCode = "BOM"
	StationCodeBromboroughRake                   CRSCode = "BMR"
	StationCodeBromleyCrossLancs                 CRSCode = "BMC"
	StationCodeBromleyNorth                      CRSCode = "BMN"
	StationCodeBromleySouth                      CRSCode = "BMS"
	StationCodeBromsgrove                        CRSCode = "BMV"
	StationCodeBrondesbury                       CRSCode = "BSY"
	StationCodeBrondesburyPark                   CRSCode = "BSP"
	StationCodeBrookmansPark                     CRSCode = "BPK"
	StationCodeBrookwood                         CRSCode = "BKO"
	StationCodeBroome                            CRSCode = "BME"
	StationCodeBroomfleet                        CRSCode = "BMF"
	StationCodeBrora                             CRSCode = "BRA"
	StationCodeBrough                            CRSCode = "BUH"
	StationCodeBroughtyFerry                     CRSCode = "BYF"
	StationCodeBroxbourne                        CRSCode = "BXB"
	StationCodeBruceGrove                        CRSCode = "BCV"
	StationCodeBrundall                          CRSCode = "BDA"
	StationCodeBrundallGardens                   CRSCode = "BGA"
	StationCodeBrunstane                         CRSCode = "BSU"
	StationCodeBrunswick                         CRSCode = "BRW"
	StationCodeBruton                            CRSCode = "BRU"
	StationCodeBryn                              CRSCode = "BYN"
	StationCodeBuckenhamNorfolk                  CRSCode = "BUC"
	StationCodeBuckley                           CRSCode = "BCK"
	StationCodeBucknell                          CRSCode = "BUK"
	StationCodeBuckshawParkway                   CRSCode = "BSV"
	StationCodeBugle                             CRSCode = "BGL"
	StationCodeBuilthRoad                        CRSCode = "BHR"
	StationCodeBulwell                           CRSCode = "BLW"
	StationCodeBures                             CRSCode = "BUE"
	StationCodeBurgessHill                       CRSCode = "BUG"
	StationCodeBurleyPark                        CRSCode = "BUY"
	StationCodeBurleyinWharfedale                CRSCode = "BUW"
	StationCodeBurnage                           CRSCode = "BNA"
	StationCodeBurnesideCumbria                  CRSCode = "BUD"
	StationCodeBurnhamBucks                      CRSCode = "BNM"
	StationCodeBurnhamonCrouch                   CRSCode = "BUU"
	StationCodeBurnleyBarracks                   CRSCode = "BUB"
	StationCodeBurnleyCentral                    CRSCode = "BNC"
	StationCodeBurnleyManchesterRoad             CRSCode = "BYM"
	StationCodeBurnsideStrathclyde               CRSCode = "BUI"
	StationCodeBurntisland                       CRSCode = "BTS"
	StationCodeBurscoughBridge                   CRSCode = "BCB"
	StationCodeBurscoughJunction                 CRSCode = "BCJ"
	StationCodeBursledon                         CRSCode = "BUO"
	StationCodeBurtonJoyce                       CRSCode = "BUJ"
	StationCodeBurtononTrent                     CRSCode = "BUT"
	StationCodeBuryStEdmunds                     CRSCode = "BSE"
	StationCodeBusby                             CRSCode = "BUS"
	StationCodeBushHillPark                      CRSCode = "BHK"
	StationCodeBushey                            CRSCode = "BSH"
	StationCodeButlersLane                       CRSCode = "BUL"
	StationCodeBuxted                            CRSCode = "BXD"
	StationCodeBuxton                            CRSCode = "BUX"
	StationCodeByfleetAndNewHaw                  CRSCode = "BFN"
	StationCodeBynea                             CRSCode = "BYE"
	StationCodeCadoxton                          CRSCode = "CAD"
	StationCodeCaergwrle                         CRSCode = "CGW"
	StationCodeCaerphilly                        CRSCode = "CPH"
	StationCodeCaersws                           CRSCode = "CWS"
	StationCodeCaldercruix                       CRSCode = "CAC"
	StationCodeCaldicot                          CRSCode = "CDT"
	StationCodeCaledonianRdAndBarnsbury          CRSCode = "CIR"
	StationCodeCalstock                          CRSCode = "CSK"
	StationCodeCamAndDursley                     CRSCode = "CDU"
	StationCodeCamberley                         CRSCode = "CAM"
	StationCodeCamborne                          CRSCode = "CBN"
	StationCodeCambridge                         CRSCode = "CBG"
	StationCodeCambridgeHeath                    CRSCode = "CBH"
	StationCodeCambuslang                        CRSCode = "CBL"
	StationCodeCamdenRoad                        CRSCode = "CMD"
	StationCodeCamelon                           CRSCode = "CMO"
	StationCodeCanadaWater                       CRSCode = "ZCW"
	StationCodeCanley                            CRSCode = "CNL"
	StationCodeCannock                           CRSCode = "CAO"
	StationCodeCanonbury                         CRSCode = "CNN"
	StationCodeCanterburyEast                    CRSCode = "CBE"
	StationCodeCanterburyWest                    CRSCode = "CBW"
	StationCodeCantley                           CRSCode = "CNY"
	StationCodeCapenhurst                        CRSCode = "CPU"
	StationCodeCarbisBay                         CRSCode = "CBB"
	StationCodeCardenden                         CRSCode = "CDD"
	StationCodeCardiffBay                        CRSCode = "CDB"
	StationCodeCardiffCentral                    CRSCode = "CDF"
	StationCodeCardiffQueenStreet                CRSCode = "CDQ"
	StationCodeCardonald                         CRSCode = "CDO"
	StationCodeCardross                          CRSCode = "CDR"
	StationCodeCarfin                            CRSCode = "CRF"
	StationCodeCarkAndCartmel                    CRSCode = "CAK"
	StationCodeCarlisle                          CRSCode = "CAR"
	StationCodeCarlton                           CRSCode = "CTO"
	StationCodeCarluke                           CRSCode = "CLU"
	StationCodeCarmarthen                        CRSCode = "CMN"
	StationCodeCarmyle                           CRSCode = "CML"
	StationCodeCarnforth                         CRSCode = "CNF"
	StationCodeCarnoustie                        CRSCode = "CAN"
	StationCodeCarntyne                          CRSCode = "CAY"
	StationCodeCarpendersPark                    CRSCode = "CPK"
	StationCodeCarrbridge                        CRSCode = "CAG"
	StationCodeCarshalton                        CRSCode = "CSH"
	StationCodeCarshaltonBeeches                 CRSCode = "CSB"
	StationCodeCarstairs                         CRSCode = "CRS"
	StationCodeCartsdyke                         CRSCode = "CDY"
	StationCodeCastleBarPark                     CRSCode = "CBP"
	StationCodeCastleCary                        CRSCode = "CLC"
	StationCodeCastleford                        CRSCode = "CFD"
	StationCodeCastletonManchester               CRSCode = "CAS"
	StationCodeCastletonMoor                     CRSCode = "CSM"
	StationCodeCaterham                          CRSCode = "CAT"
	StationCodeCatford                           CRSCode = "CTF"
	StationCodeCatfordBridge                     CRSCode = "CFB"
	StationCodeCathays                           CRSCode = "CYS"
	StationCodeCathcart                          CRSCode = "CCT"
	StationCodeCattal                            CRSCode = "CTL"
	StationCodeCauseland                         CRSCode = "CAU"
	StationCodeCefnyBedd                         CRSCode = "CYB"
	StationCodeChadwellHeath                     CRSCode = "CTH"
	StationCodeChaffordHundredLakeside           CRSCode = "CFH"
	StationCodeChalfontAndLatimer                CRSCode = "CFO"
	StationCodeChalkwell                         CRSCode = "CHW"
	StationCodeChandlersFord                     CRSCode = "CFR"
	StationCodeChapelenleFrith                   CRSCode = "CEF"
	StationCodeChapeltonDevon                    CRSCode = "CPN"
	StationCodeChapeltownSouthYorks              CRSCode = "CLN"
	StationCodeChappelAndWakesColne              CRSCode = "CWC"
	StationCodeCharingCrossGlasgow               CRSCode = "CHC"
	StationCodeCharingKent                       CRSCode = "CHG"
	StationCodeCharlbury                         CRSCode = "CBY"
	StationCodeCharlton                          CRSCode = "CTN"
	StationCodeChartham                          CRSCode = "CRT"
	StationCodeChassenRoad                       CRSCode = "CSR"
	StationCodeChatelherault                     CRSCode = "CTE"
	StationCodeChatham                           CRSCode = "CTM"
	StationCodeChathill                          CRSCode = "CHT"
	StationCodeCheadleHulme                      CRSCode = "CHU"
	StationCodeCheam                             CRSCode = "CHE"
	StationCodeCheddington                       CRSCode = "CED"
	StationCodeChelfordCheshire                  CRSCode = "CEL"
	StationCodeChelmsford                        CRSCode = "CHM"
	StationCodeChelsfield                        CRSCode = "CLD"
	StationCodeCheltenhamSpa                     CRSCode = "CNM"
	StationCodeChepstow                          CRSCode = "CPW"
	StationCodeCherryTree                        CRSCode = "CYT"
	StationCodeChertsey                          CRSCode = "CHY"
	StationCodeCheshunt                          CRSCode = "CHN"
	StationCodeChessingtonNorth                  CRSCode = "CSN"
	StationCodeChessingtonSouth                  CRSCode = "CSS"
	StationCodeChester                           CRSCode = "CTR"
	StationCodeChesterRoad                       CRSCode = "CRD"
	StationCodeChesterfield                      CRSCode = "CHD"
	StationCodeChesterleStreet                   CRSCode = "CLS"
	StationCodeChestfieldAndSwalecliffe          CRSCode = "CSW"
	StationCodeChetnole                          CRSCode = "CNO"
	StationCodeChichester                        CRSCode = "CCH"
	StationCodeChilham                           CRSCode = "CIL"
	StationCodeChilworth                         CRSCode = "CHL"
	StationCodeChingford                         CRSCode = "CHI"
	StationCodeChinley                           CRSCode = "CLY"
	StationCodeChippenham                        CRSCode = "CPM"
	StationCodeChipstead                         CRSCode = "CHP"
	StationCodeChirk                             CRSCode = "CRK"
	StationCodeChislehurst                       CRSCode = "CIT"
	StationCodeChiswick                          CRSCode = "CHK"
	StationCodeCholsey                           CRSCode = "CHO"
	StationCodeChorley                           CRSCode = "CRL"
	StationCodeChorleywood                       CRSCode = "CLW"
	StationCodeChristchurch                      CRSCode = "CHR"
	StationCodeChristsHospital                   CRSCode = "CHH"
	StationCodeChurchAndOswaldtwistle            CRSCode = "CTW"
	StationCodeChurchFenton                      CRSCode = "CHF"
	StationCodeChurchStretton                    CRSCode = "CTT"
	StationCodeCilmeri                           CRSCode = "CIM"
	StationCodeCityThameslink                    CRSCode = "CTK"
	StationCodeClactononSea                      CRSCode = "CLT"
	StationCodeClandon                           CRSCode = "CLA"
	StationCodeClaphamHighStreet                 CRSCode = "CLP"
	StationCodeClaphamJunction                   CRSCode = "CLJ"
	StationCodeClaphamNorthYorkshire             CRSCode = "CPY"
	StationCodeClapton                           CRSCode = "CPT"
	StationCodeClarbestonRoad                    CRSCode = "CLR"
	StationCodeClarkston                         CRSCode = "CKS"
	StationCodeClaverdon                         CRSCode = "CLV"
	StationCodeClaygate                          CRSCode = "CLG"
	StationCodeCleethorpes                       CRSCode = "CLE"
	StationCodeCleland                           CRSCode = "CEA"
	StationCodeCliftonDown                       CRSCode = "CFN"
	StationCodeCliftonManchester                 CRSCode = "CLI"
	StationCodeClitheroe                         CRSCode = "CLH"
	StationCodeClockHouse                        CRSCode = "CLK"
	StationCodeClunderwen                        CRSCode = "CUW"
	StationCodeClydebank                         CRSCode = "CYK"
	StationCodeCoatbridgeCentral                 CRSCode = "CBC"
	StationCodeCoatbridgeSunnyside               CRSCode = "CBS"
	StationCodeCoatdyke                          CRSCode = "COA"
	StationCodeCobhamAndStokedAbernon            CRSCode = "CSD"
	StationCodeCodsall                           CRSCode = "CSL"
	StationCodeCogan                             CRSCode = "CGN"
	StationCodeColchester                        CRSCode = "COL"
	StationCodeColchesterTown                    CRSCode = "CET"
	StationCodeColeshillParkway                  CRSCode = "CEH"
	StationCodeCollingham                        CRSCode = "CLM"
	StationCodeCollington                        CRSCode = "CLL"
	StationCodeColne                             CRSCode = "CNE"
	StationCodeColwall                           CRSCode = "CWL"
	StationCodeColwynBay                         CRSCode = "CWB"
	StationCodeCombeOxon                         CRSCode = "CME"
	StationCodeCommondale                        CRSCode = "COM"
	StationCodeCongleton                         CRSCode = "CNG"
	StationCodeConisbrough                       CRSCode = "CNS"
	StationCodeConnelFerry                       CRSCode = "CON"
	StationCodeCononBridge                       CRSCode = "CBD"
	StationCodeCononley                          CRSCode = "CEY"
	StationCodeConwayPark                        CRSCode = "CNP"
	StationCodeConwy                             CRSCode = "CNW"
	StationCodeCoodenBeach                       CRSCode = "COB"
	StationCodeCookham                           CRSCode = "COO"
	StationCodeCooksbridge                       CRSCode = "CBR"
	StationCodeCoombeJunctionHalt                CRSCode = "COE"
	StationCodeCopplestone                       CRSCode = "COP"
	StationCodeCorbridge                         CRSCode = "CRB"
	StationCodeCorby                             CRSCode = "COR"
	StationCodeCorkerhill                        CRSCode = "CKH"
	StationCodeCorkickle                         CRSCode = "CKL"
	StationCodeCorpach                           CRSCode = "CPA"
	StationCodeCorrour                           CRSCode = "CRR"
	StationCodeCoryton                           CRSCode = "COY"
	StationCodeCoseley                           CRSCode = "CSY"
	StationCodeCosford                           CRSCode = "COS"
	StationCodeCosham                            CRSCode = "CSA"
	StationCodeCottingham                        CRSCode = "CGM"
	StationCodeCottingley                        CRSCode = "COT"
	StationCodeCoulsdonSouth                     CRSCode = "CDS"
	StationCodeCoulsdonTown                      CRSCode = "CDN"
	StationCodeCoventry                          CRSCode = "COV"
	StationCodeCoventryArena                     CRSCode = "CAA"
	StationCodeCowdenKent                        CRSCode = "CWN"
	StationCodeCowdenbeath                       CRSCode = "COW"
	StationCodeCradleyHeath                      CRSCode = "CRA"
	StationCodeCraigendoran                      CRSCode = "CGD"
	StationCodeCramlington                       CRSCode = "CRM"
	StationCodeCranbrookDevon                    CRSCode = "CBK"
	StationCodeCravenArms                        CRSCode = "CRV"
	StationCodeCrawley                           CRSCode = "CRW"
	StationCodeCrayford                          CRSCode = "CRY"
	StationCodeCrediton                          CRSCode = "CDI"
	StationCodeCressingEssex                     CRSCode = "CES"
	StationCodeCressington                       CRSCode = "CSG"
	StationCodeCreswell                          CRSCode = "CWD"
	StationCodeCrewe                             CRSCode = "CRE"
	StationCodeCrewkerne                         CRSCode = "CKN"
	StationCodeCrewsHill                         CRSCode = "CWH"
	StationCodeCrianlarich                       CRSCode = "CNR"
	StationCodeCriccieth                         CRSCode = "CCC"
	StationCodeCricklewood                       CRSCode = "CRI"
	StationCodeCroftfoot                         CRSCode = "CFF"
	StationCodeCroftonPark                       CRSCode = "CFT"
	StationCodeCromer                            CRSCode = "CMR"
	StationCodeCromford                          CRSCode = "CMF"
	StationCodeCrookston                         CRSCode = "CKT"
	StationCodeCrossGates                        CRSCode = "CRG"
	StationCodeCrossflatts                       CRSCode = "CFL"
	StationCodeCrosshill                         CRSCode = "COI"
	StationCodeCrosskeys                         CRSCode = "CKY"
	StationCodeCrossmyloof                       CRSCode = "CMY"
	StationCodeCroston                           CRSCode = "CSO"
	StationCodeCrouchHill                        CRSCode = "CRH"
	StationCodeCrowborough                       CRSCode = "COH"
	StationCodeCrowhurst                         CRSCode = "CWU"
	StationCodeCrowle                            CRSCode = "CWE"
	StationCodeCrowthorne                        CRSCode = "CRN"
	StationCodeCroy                              CRSCode = "CRO"
	StationCodeCrystalPalace                     CRSCode = "CYP"
	StationCodeCuddington                        CRSCode = "CUD"
	StationCodeCuffley                           CRSCode = "CUF"
	StationCodeCulham                            CRSCode = "CUM"
	StationCodeCulrain                           CRSCode = "CUA"
	StationCodeCumbernauld                       CRSCode = "CUB"
	StationCodeCupar                             CRSCode = "CUP"
	StationCodeCurriehill                        CRSCode = "CUH"
	StationCodeCuxton                            CRSCode = "CUX"
	StationCodeCwmbach                           CRSCode = "CMH"
	StationCodeCwmbran                           CRSCode = "CWM"
	StationCodeCynghordy                         CRSCode = "CYN"
	StationCodeDagenhamDock                      CRSCode = "DDK"
	StationCodeDaisyHill                         CRSCode = "DSY"
	StationCodeDalgetyBay                        CRSCode = "DAG"
	StationCodeDalmally                          CRSCode = "DAL"
	StationCodeDalmarnock                        CRSCode = "DAK"
	StationCodeDalmeny                           CRSCode = "DAM"
	StationCodeDalmuir                           CRSCode = "DMR"
	StationCodeDalreoch                          CRSCode = "DLR"
	StationCodeDalry                             CRSCode = "DLY"
	StationCodeDalstonCumbria                    CRSCode = "DLS"
	StationCodeDalstonJunction                   CRSCode = "DLJ"
	StationCodeDalstonKingsland                  CRSCode = "DLK"
	StationCodeDaltonCumbria                     CRSCode = "DLT"
	StationCodeDalwhinnie                        CRSCode = "DLW"
	StationCodeDanby                             CRSCode = "DNY"
	StationCodeDanescourt                        CRSCode = "DCT"
	StationCodeDanzey                            CRSCode = "DZY"
	StationCodeDarlington                        CRSCode = "DAR"
	StationCodeDarnall                           CRSCode = "DAN"
	StationCodeDarsham                           CRSCode = "DSM"
	StationCodeDartford                          CRSCode = "DFD"
	StationCodeDarton                            CRSCode = "DRT"
	StationCodeDarwen                            CRSCode = "DWN"
	StationCodeDatchet                           CRSCode = "DAT"
	StationCodeDavenport                         CRSCode = "DVN"
	StationCodeDawlish                           CRSCode = "DWL"
	StationCodeDawlishWarren                     CRSCode = "DWW"
	StationCodeDeal                              CRSCode = "DEA"
	StationCodeDeanWilts                         CRSCode = "DEN"
	StationCodeDeansgate                         CRSCode = "DGT"
	StationCodeDeganwy                           CRSCode = "DGY"
	StationCodeDeighton                          CRSCode = "DHN"
	StationCodeDelamere                          CRSCode = "DLM"
	StationCodeDenbyDale                         CRSCode = "DBD"
	StationCodeDenham                            CRSCode = "DNM"
	StationCodeDenhamGolfClub                    CRSCode = "DGC"
	StationCodeDenmarkHill                       CRSCode = "DMK"
	StationCodeDent                              CRSCode = "DNT"
	StationCodeDenton                            CRSCode = "DTN"
	StationCodeDeptford                          CRSCode = "DEP"
	StationCodeDerby                             CRSCode = "DBY"
	StationCodeDerbyRoadIpswich                  CRSCode = "DBR"
	StationCodeDevonportDevon                    CRSCode = "DPT"
	StationCodeDevonportDockyard                 CRSCode = "DOC"
	StationCodeDewsbury                          CRSCode = "DEW"
	StationCodeDidcotParkway                     CRSCode = "DID"
	StationCodeDigbyAndSowton                    CRSCode = "DIG"
	StationCodeDiltonMarsh                       CRSCode = "DMH"
	StationCodeDinasPowys                        CRSCode = "DNS"
	StationCodeDinasRhondda                      CRSCode = "DMG"
	StationCodeDingleRoad                        CRSCode = "DGL"
	StationCodeDingwall                          CRSCode = "DIN"
	StationCodeDinsdale                          CRSCode = "DND"
	StationCodeDinting                           CRSCode = "DTG"
	StationCodeDisley                            CRSCode = "DSL"
	StationCodeDiss                              CRSCode = "DIS"
	StationCodeDodworth                          CRSCode = "DOD"
	StationCodeDolau                             CRSCode = "DOL"
	StationCodeDoleham                           CRSCode = "DLH"
	StationCodeDolgarrog                         CRSCode = "DLG"
	StationCodeDolwyddelan                       CRSCode = "DWD"
	StationCodeDoncaster                         CRSCode = "DON"
	StationCodeDorchesterSouth                   CRSCode = "DCH"
	StationCodeDorchesterWest                    CRSCode = "DCW"
	StationCodeDoreAndTotley                     CRSCode = "DOR"
	StationCodeDorkingDeepdene                   CRSCode = "DPD"
	StationCodeDorkingMain                       CRSCode = "DKG"
	StationCodeDorkingWest                       CRSCode = "DKT"
	StationCodeDormans                           CRSCode = "DMS"
	StationCodeDorridge                          CRSCode = "DDG"
	StationCodeDoveHoles                         CRSCode = "DVH"
	StationCodeDoverPriory                       CRSCode = "DVP"
	StationCodeDovercourt                        CRSCode = "DVC"
	StationCodeDoveyJunction                     CRSCode = "DVY"
	StationCodeDownhamMarket                     CRSCode = "DOW"
	StationCodeDraytonGreen                      CRSCode = "DRG"
	StationCodeDraytonPark                       CRSCode = "DYP"
	StationCodeDrem                              CRSCode = "DRM"
	StationCodeDriffield                         CRSCode = "DRF"
	StationCodeDrigg                             CRSCode = "DRI"
	StationCodeDroitwichSpa                      CRSCode = "DTW"
	StationCodeDronfield                         CRSCode = "DRO"
	StationCodeDrumchapel                        CRSCode = "DMC"
	StationCodeDrumfrochar                       CRSCode = "DFR"
	StationCodeDrumgelloch                       CRSCode = "DRU"
	StationCodeDrumry                            CRSCode = "DMY"
	StationCodeDublinFerryport                   CRSCode = "DFP"
	StationCodeDublinPortStena                   CRSCode = "DPS"
	StationCodeDuddeston                         CRSCode = "DUD"
	StationCodeDudleyPort                        CRSCode = "DDP"
	StationCodeDuffield                          CRSCode = "DFI"
	StationCodeDuirinish                         CRSCode = "DRN"
	StationCodeDukeStreet                        CRSCode = "DST"
	StationCodeDullingham                        CRSCode = "DUL"
	StationCodeDumbartonCentral                  CRSCode = "DBC"
	StationCodeDumbartonEast                     CRSCode = "DBE"
	StationCodeDumbreck                          CRSCode = "DUM"
	StationCodeDumfries                          CRSCode = "DMF"
	StationCodeDumptonPark                       CRSCode = "DMP"
	StationCodeDunbar                            CRSCode = "DUN"
	StationCodeDunblane                          CRSCode = "DBL"
	StationCodeDuncraig                          CRSCode = "DCG"
	StationCodeDundee                            CRSCode = "DEE"
	StationCodeDunfermlineQueenMargaret          CRSCode = "DFL"
	StationCodeDunfermlineTown                   CRSCode = "DFE"
	StationCodeDunkeldAndBirnam                  CRSCode = "DKD"
	StationCodeDunlop                            CRSCode = "DNL"
	StationCodeDunrobinCastle                    CRSCode = "DNO"
	StationCodeDunston                           CRSCode = "DOT"
	StationCodeDuntonGreen                       CRSCode = "DNG"
	StationCodeDurham                            CRSCode = "DHM"
	StationCodeDurringtononSea                   CRSCode = "DUR"
	StationCodeDyce                              CRSCode = "DYC"
	StationCodeDyffrynArdudwy                    CRSCode = "DYF"
	StationCodeEaglescliffe                      CRSCode = "EAG"
	StationCodeEalingBroadway                    CRSCode = "EAL"
	StationCodeEarlestown                        CRSCode = "ERL"
	StationCodeEarley                            CRSCode = "EAR"
	StationCodeEarlsfield                        CRSCode = "EAD"
	StationCodeEarlswoodSurrey                   CRSCode = "ELD"
	StationCodeEarlswoodWestMidlands             CRSCode = "EWD"
	StationCodeEastCroydon                       CRSCode = "ECR"
	StationCodeEastDidsbury                      CRSCode = "EDY"
	StationCodeEastDulwich                       CRSCode = "EDW"
	StationCodeEastFarleigh                      CRSCode = "EFL"
	StationCodeEastGarforth                      CRSCode = "EGF"
	StationCodeEastGrinstead                     CRSCode = "EGR"
	StationCodeEastKilbride                      CRSCode = "EKL"
	StationCodeEastMalling                       CRSCode = "EML"
	StationCodeEastMidlandsParkway               CRSCode = "EMD"
	StationCodeEastTilbury                       CRSCode = "ETL"
	StationCodeEastWorthing                      CRSCode = "EWR"
	StationCodeEastbourne                        CRSCode = "EBN"
	StationCodeEastbrook                         CRSCode = "EBK"
	StationCodeEasterhouse                       CRSCode = "EST"
	StationCodeEasthamRake                       CRSCode = "ERA"
	StationCodeEastleigh                         CRSCode = "ESL"
	StationCodeEastrington                       CRSCode = "EGN"
	StationCodeEbbsfleetInternational            CRSCode = "EBD"
	StationCodeEbbwValeParkway                   CRSCode = "EBV"
	StationCodeEbbwValeTown                      CRSCode = "EBB"
	StationCodeEcclesManchester                  CRSCode = "ECC"
	StationCodeEcclesRoad                        CRSCode = "ECS"
	StationCodeEcclestonPark                     CRSCode = "ECL"
	StationCodeEdale                             CRSCode = "EDL"
	StationCodeEdenPark                          CRSCode = "EDN"
	StationCodeEdenbridge                        CRSCode = "EBR"
	StationCodeEdenbridgeTown                    CRSCode = "EBT"
	StationCodeEdgeHill                          CRSCode = "EDG"
	StationCodeEdinburgh                         CRSCode = "EDB"
	StationCodeEdinburghGateway                  CRSCode = "EGY"
	StationCodeEdinburghPark                     CRSCode = "EDP"
	StationCodeEdmontonGreen                     CRSCode = "EDR"
	StationCodeEffinghamJunction                 CRSCode = "EFF"
	StationCodeEggesford                         CRSCode = "EGG"
	StationCodeEgham                             CRSCode = "EGH"
	StationCodeEgton                             CRSCode = "EGT"
	StationCodeElephantAndCastle                 CRSCode = "EPH"
	StationCodeElephantAndCastleUnderground      CRSCode = "ZEL"
	StationCodeElgin                             CRSCode = "ELG"
	StationCodeEllesmerePort                     CRSCode = "ELP"
	StationCodeElmersEnd                         CRSCode = "ELE"
	StationCodeElmsteadWoods                     CRSCode = "ESD"
	StationCodeElmswell                          CRSCode = "ESW"
	StationCodeElsecar                           CRSCode = "ELR"
	StationCodeElsenhamEssex                     CRSCode = "ESM"
	StationCodeElstreeAndBorehamwood             CRSCode = "ELS"
	StationCodeEltham                            CRSCode = "ELW"
	StationCodeEltonAndOrston                    CRSCode = "ELO"
	StationCodeEly                               CRSCode = "ELY"
	StationCodeEmersonPark                       CRSCode = "EMP"
	StationCodeEmsworth                          CRSCode = "EMS"
	StationCodeEnerglynAndChurchillPark          CRSCode = "ECP"
	StationCodeEnfieldChase                      CRSCode = "ENC"
	StationCodeEnfieldLock                       CRSCode = "ENL"
	StationCodeEnfieldTown                       CRSCode = "ENF"
	StationCodeEntwistle                         CRSCode = "ENT"
	StationCodeEpsomDowns                        CRSCode = "EPD"
	StationCodeEpsomSurrey                       CRSCode = "EPS"
	StationCodeErdington                         CRSCode = "ERD"
	StationCodeEridge                            CRSCode = "ERI"
	StationCodeErith                             CRSCode = "ERH"
	StationCodeEsher                             CRSCode = "ESH"
	StationCodeEskbank                           CRSCode = "EKB"
	StationCodeEssexRoad                         CRSCode = "EXR"
	StationCodeEtchingham                        CRSCode = "ETC"
	StationCodeEuxtonBalshawLane                 CRSCode = "EBA"
	StationCodeEvesham                           CRSCode = "EVE"
	StationCodeEwellEast                         CRSCode = "EWE"
	StationCodeEwellWest                         CRSCode = "EWW"
	StationCodeExeterCentral                     CRSCode = "EXC"
	StationCodeExeterStDavids                    CRSCode = "EXD"
	StationCodeExeterStThomas                    CRSCode = "EXT"
	StationCodeExhibitionCentreGlasgow           CRSCode = "EXG"
	StationCodeExmouth                           CRSCode = "EXM"
	StationCodeExton                             CRSCode = "EXN"
	StationCodeEynsford                          CRSCode = "EYN"
	StationCodeFairbourne                        CRSCode = "FRB"
	StationCodeFairfield                         CRSCode = "FRF"
	StationCodeFairlie                           CRSCode = "FRL"
	StationCodeFairwater                         CRSCode = "FRW"
	StationCodeFalconwood                        CRSCode = "FCN"
	StationCodeFalkirkGrahamston                 CRSCode = "FKG"
	StationCodeFalkirkHigh                       CRSCode = "FKK"
	StationCodeFallsofCruachan                   CRSCode = "FOC"
	StationCodeFalmer                            CRSCode = "FMR"
	StationCodeFalmouthDocks                     CRSCode = "FAL"
	StationCodeFalmouthTown                      CRSCode = "FMT"
	StationCodeFareham                           CRSCode = "FRM"
	StationCodeFarnboroughMain                   CRSCode = "FNB"
	StationCodeFarnboroughNorth                  CRSCode = "FNN"
	StationCodeFarncombe                         CRSCode = "FNC"
	StationCodeFarnham                           CRSCode = "FNH"
	StationCodeFarninghamRoad                    CRSCode = "FNR"
	StationCodeFarnworth                         CRSCode = "FNW"
	StationCodeFarringdon                        CRSCode = "ZFD"
	StationCodeFauldhouse                        CRSCode = "FLD"
	StationCodeFaversham                         CRSCode = "FAV"
	StationCodeFaygate                           CRSCode = "FGT"
	StationCodeFazakerley                        CRSCode = "FAZ"
	StationCodeFearn                             CRSCode = "FRN"
	StationCodeFeatherstone                      CRSCode = "FEA"
	StationCodeFelixstowe                        CRSCode = "FLX"
	StationCodeFeltham                           CRSCode = "FEL"
	StationCodeFeniton                           CRSCode = "FNT"
	StationCodeFennyStratford                    CRSCode = "FEN"
	StationCodeFernhill                          CRSCode = "FER"
	StationCodeFerriby                           CRSCode = "FRY"
	StationCodeFerryside                         CRSCode = "FYS"
	StationCodeFfairfach                         CRSCode = "FFA"
	StationCodeFiley                             CRSCode = "FIL"
	StationCodeFiltonAbbeyWood                   CRSCode = "FIT"
	StationCodeFinchleyRoadAndFrognal            CRSCode = "FNY"
	StationCodeFinsburyPark                      CRSCode = "FPK"
	StationCodeFinstock                          CRSCode = "FIN"
	StationCodeFishbourneSussex                  CRSCode = "FSB"
	StationCodeFishersgate                       CRSCode = "FSG"
	StationCodeFishguardAndGoodwick              CRSCode = "FGW"
	StationCodeFishguardHarbour                  CRSCode = "FGH"
	StationCodeFiskerton                         CRSCode = "FSK"
	StationCodeFitzwilliam                       CRSCode = "FZW"
	StationCodeFiveWays                          CRSCode = "FWY"
	StationCodeFleet                             CRSCode = "FLE"
	StationCodeFlimby                            CRSCode = "FLM"
	StationCodeFlint                             CRSCode = "FLN"
	StationCodeFlitwick                          CRSCode = "FLT"
	StationCodeFlixton                           CRSCode = "FLI"
	StationCodeFloweryField                      CRSCode = "FLF"
	StationCodeFolkestoneCentral                 CRSCode = "FKC"
	StationCodeFolkestoneWest                    CRSCode = "FKW"
	StationCodeFord                              CRSCode = "FOD"
	StationCodeForestGate                        CRSCode = "FOG"
	StationCodeForestHill                        CRSCode = "FOH"
	StationCodeFormby                            CRSCode = "FBY"
	StationCodeForres                            CRSCode = "FOR"
	StationCodeForsinard                         CRSCode = "FRS"
	StationCodeFortMatilda                       CRSCode = "FTM"
	StationCodeFortWilliam                       CRSCode = "FTW"
	StationCodeFourOaks                          CRSCode = "FOK"
	StationCodeFoxfield                          CRSCode = "FOX"
	StationCodeFoxton                            CRSCode = "FXN"
	StationCodeFrant                             CRSCode = "FRT"
	StationCodeFratton                           CRSCode = "FTN"
	StationCodeFreshfield                        CRSCode = "FRE"
	StationCodeFreshford                         CRSCode = "FFD"
	StationCodeFrimley                           CRSCode = "FML"
	StationCodeFrintononSea                      CRSCode = "FRI"
	StationCodeFrizinghall                       CRSCode = "FZH"
	StationCodeFrodsham                          CRSCode = "FRD"
	StationCodeFrome                             CRSCode = "FRO"
	StationCodeFulwell                           CRSCode = "FLW"
	StationCodeFurnessVale                       CRSCode = "FNV"
	StationCodeFurzePlatt                        CRSCode = "FZP"
	StationCodeGainsboroughCentral               CRSCode = "GNB"
	StationCodeGainsboroughLeaRoad               CRSCode = "GBL"
	StationCodeGalashiels                        CRSCode = "GAL"
	StationCodeGarelochhead                      CRSCode = "GCH"
	StationCodeGarforth                          CRSCode = "GRF"
	StationCodeGargrave                          CRSCode = "GGV"
	StationCodeGarrowhill                        CRSCode = "GAR"
	StationCodeGarscadden                        CRSCode = "GRS"
	StationCodeGarsdale                          CRSCode = "GSD"
	StationCodeGarstonHertfordshire              CRSCode = "GSN"
	StationCodeGarswood                          CRSCode = "GSW"
	StationCodeGartcosh                          CRSCode = "GRH"
	StationCodeGarthMidGlamorgan                 CRSCode = "GMG"
	StationCodeGarthPowys                        CRSCode = "GTH"
	StationCodeGarve                             CRSCode = "GVE"
	StationCodeGathurst                          CRSCode = "GST"
	StationCodeGatley                            CRSCode = "GTY"
	StationCodeGatwickAirport                    CRSCode = "GTW"
	StationCodeGeorgemasJunction                 CRSCode = "GGJ"
	StationCodeGerrardsCross                     CRSCode = "GER"
	StationCodeGideaPark                         CRSCode = "GDP"
	StationCodeGiffnock                          CRSCode = "GFN"
	StationCodeGiggleswick                       CRSCode = "GIG"
	StationCodeGilberdyke                        CRSCode = "GBD"
	StationCodeGilfachFargoed                    CRSCode = "GFF"
	StationCodeGillinghamDorset                  CRSCode = "GIL"
	StationCodeGillinghamKent                    CRSCode = "GLM"
	StationCodeGilshochill                       CRSCode = "GSC"
	StationCodeGipsyHill                         CRSCode = "GIP"
	StationCodeGirvan                            CRSCode = "GIR"
	StationCodeGlaisdale                         CRSCode = "GLS"
	StationCodeGlanConwy                         CRSCode = "GCW"
	StationCodeGlasgowCentral                    CRSCode = "GLC"
	StationCodeGlasgowQueenStreet                CRSCode = "GLQ"
	StationCodeGlasshoughton                     CRSCode = "GLH"
	StationCodeGlazebrook                        CRSCode = "GLZ"
	StationCodeGleneagles                        CRSCode = "GLE"
	StationCodeGlenfinnan                        CRSCode = "GLF"
	StationCodeGlengarnock                       CRSCode = "GLG"
	StationCodeGlenrotheswithThornton            CRSCode = "GLT"
	StationCodeGlossop                           CRSCode = "GLO"
	StationCodeGloucester                        CRSCode = "GCR"
	StationCodeGlynde                            CRSCode = "GLY"
	StationCodeGobowen                           CRSCode = "GOB"
	StationCodeGodalming                         CRSCode = "GOD"
	StationCodeGodley                            CRSCode = "GDL"
	StationCodeGodstone                          CRSCode = "GDN"
	StationCodeGoldthorpe                        CRSCode = "GOE"
	StationCodeGolfStreet                        CRSCode = "GOF"
	StationCodeGolspie                           CRSCode = "GOL"
	StationCodeGomshall                          CRSCode = "GOM"
	StationCodeGoodmayes                         CRSCode = "GMY"
	StationCodeGoole                             CRSCode = "GOO"
	StationCodeGoostrey                          CRSCode = "GTR"
	StationCodeGordonHill                        CRSCode = "GDH"
	StationCodeGorebridge                        CRSCode = "GBG"
	StationCodeGoringAndStreatley                CRSCode = "GOR"
	StationCodeGoringbySea                       CRSCode = "GBS"
	StationCodeGorton                            CRSCode = "GTO"
	StationCodeGospelOak                         CRSCode = "GPO"
	StationCodeGourock                           CRSCode = "GRK"
	StationCodeGowerton                          CRSCode = "GWN"
	StationCodeGoxhill                           CRSCode = "GOX"
	StationCodeGrangeOverSands                   CRSCode = "GOS"
	StationCodeGrangePark                        CRSCode = "GPK"
	StationCodeGrangetownCardiff                 CRSCode = "GTN"
	StationCodeGrantham                          CRSCode = "GRA"
	StationCodeGrateley                          CRSCode = "GRT"
	StationCodeGravellyHill                      CRSCode = "GVH"
	StationCodeGravesend                         CRSCode = "GRV"
	StationCodeGrays                             CRSCode = "GRY"
	StationCodeGreatAyton                        CRSCode = "GTA"
	StationCodeGreatBentley                      CRSCode = "GRB"
	StationCodeGreatChesterford                  CRSCode = "GRC"
	StationCodeGreatCoates                       CRSCode = "GCT"
	StationCodeGreatMalvern                      CRSCode = "GMV"
	StationCodeGreatMissenden                    CRSCode = "GMN"
	StationCodeGreatYarmouth                     CRSCode = "GYM"
	StationCodeGreenLane                         CRSCode = "GNL"
	StationCodeGreenRoad                         CRSCode = "GNR"
	StationCodeGreenbank                         CRSCode = "GBK"
	StationCodeGreenfaulds                       CRSCode = "GRL"
	StationCodeGreenfield                        CRSCode = "GNF"
	StationCodeGreenford                         CRSCode = "GFD"
	StationCodeGreenhitheForBluewater            CRSCode = "GNH"
	StationCodeGreenockCentral                   CRSCode = "GKC"
	StationCodeGreenockWest                      CRSCode = "GKW"
	StationCodeGreenwich                         CRSCode = "GNW"
	StationCodeGretnaGreen                       CRSCode = "GEA"
	StationCodeGrimsbyDocks                      CRSCode = "GMD"
	StationCodeGrimsbyTown                       CRSCode = "GMB"
	StationCodeGrindleford                       CRSCode = "GRN"
	StationCodeGrosmont                          CRSCode = "GMT"
	StationCodeGrovePark                         CRSCode = "GRP"
	StationCodeGuideBridge                       CRSCode = "GUI"
	StationCodeGuildford                         CRSCode = "GLD"
	StationCodeGuiseley                          CRSCode = "GSY"
	StationCodeGunnersbury                       CRSCode = "GUN"
	StationCodeGunnislake                        CRSCode = "GSL"
	StationCodeGunton                            CRSCode = "GNT"
	StationCodeGwersyllt                         CRSCode = "GWE"
	StationCodeGypsyLane                         CRSCode = "GYP"
	StationCodeHabrough                          CRSCode = "HAB"
	StationCodeHackbridge                        CRSCode = "HCB"
	StationCodeHackneyCentral                    CRSCode = "HKC"
	StationCodeHackneyDowns                      CRSCode = "HAC"
	StationCodeHackneyWick                       CRSCode = "HKW"
	StationCodeHaddenhamAndThameParkway          CRSCode = "HDM"
	StationCodeHaddiscoe                         CRSCode = "HAD"
	StationCodeHadfield                          CRSCode = "HDF"
	StationCodeHadleyWood                        CRSCode = "HDW"
	StationCodeHagFold                           CRSCode = "HGF"
	StationCodeHaggerston                        CRSCode = "HGG"
	StationCodeHagley                            CRSCode = "HAG"
	StationCodeHairmyres                         CRSCode = "HMY"
	StationCodeHaleManchester                    CRSCode = "HAL"
	StationCodeHalesworth                        CRSCode = "HAS"
	StationCodeHalewood                          CRSCode = "HED"
	StationCodeHalifax                           CRSCode = "HFX"
	StationCodeHallGreen                         CRSCode = "HLG"
	StationCodeHallRoad                          CRSCode = "HLR"
	StationCodeHalling                           CRSCode = "HAI"
	StationCodeHallithWood                       CRSCode = "HID"
	StationCodeHaltwhistle                       CRSCode = "HWH"
	StationCodeHamStreet                         CRSCode = "HMT"
	StationCodeHamble                            CRSCode = "HME"
	StationCodeHamiltonCentral                   CRSCode = "HNC"
	StationCodeHamiltonWest                      CRSCode = "HNW"
	StationCodeHammerton                         CRSCode = "HMM"
	StationCodeHampdenParkSussex                 CRSCode = "HMD"
	StationCodeHampsteadHeath                    CRSCode = "HDH"
	StationCodeHamptonCourt                      CRSCode = "HMC"
	StationCodeHamptonLondon                     CRSCode = "HMP"
	StationCodeHamptonWick                       CRSCode = "HMW"
	StationCodeHamptoninArden                    CRSCode = "HIA"
	StationCodeHamsteadBirmingham                CRSCode = "HSD"
	StationCodeHamworthy                         CRSCode = "HAM"
	StationCodeHanborough                        CRSCode = "HND"
	StationCodeHandforth                         CRSCode = "HTH"
	StationCodeHanwell                           CRSCode = "HAN"
	StationCodeHapton                            CRSCode = "HPN"
	StationCodeHarlech                           CRSCode = "HRL"
	StationCodeHarlesden                         CRSCode = "HDN"
	StationCodeHarlingRoad                       CRSCode = "HRD"
	StationCodeHarlingtonBeds                    CRSCode = "HLN"
	StationCodeHarlowMill                        CRSCode = "HWM"
	StationCodeHarlowTown                        CRSCode = "HWN"
	StationCodeHaroldWood                        CRSCode = "HRO"
	StationCodeHarpenden                         CRSCode = "HPD"
	StationCodeHarrietsham                       CRSCode = "HRM"
	StationCodeHarringay                         CRSCode = "HGY"
	StationCodeHarringayGreenLanes               CRSCode = "HRY"
	StationCodeHarrington                        CRSCode = "HRR"
	StationCodeHarrogate                         CRSCode = "HGT"
	StationCodeHarrowAndWealdstone               CRSCode = "HRW"
	StationCodeHarrowontheHill                   CRSCode = "HOH"
	StationCodeHartfordCheshire                  CRSCode = "HTF"
	StationCodeHartlebury                        CRSCode = "HBY"
	StationCodeHartlepool                        CRSCode = "HPL"
	StationCodeHartwood                          CRSCode = "HTW"
	StationCodeHarwichInternational              CRSCode = "HPQ"
	StationCodeHarwichTown                       CRSCode = "HWC"
	StationCodeHaslemere                         CRSCode = "HSL"
	StationCodeHassocks                          CRSCode = "HSK"
	StationCodeHastings                          CRSCode = "HGS"
	StationCodeHatchEnd                          CRSCode = "HTE"
	StationCodeHatfieldAndStainforth             CRSCode = "HFS"
	StationCodeHatfieldHerts                     CRSCode = "HAT"
	StationCodeHatfieldPeverel                   CRSCode = "HAP"
	StationCodeHathersage                        CRSCode = "HSG"
	StationCodeHattersley                        CRSCode = "HTY"
	StationCodeHatton                            CRSCode = "HTN"
	StationCodeHavant                            CRSCode = "HAV"
	StationCodeHavenhouse                        CRSCode = "HVN"
	StationCodeHaverfordwest                     CRSCode = "HVF"
	StationCodeHawarden                          CRSCode = "HWD"
	StationCodeHawardenBridge                    CRSCode = "HWB"
	StationCodeHawkhead                          CRSCode = "HKH"
	StationCodeHaydonBridge                      CRSCode = "HDB"
	StationCodeHaydonsRoad                       CRSCode = "HYR"
	StationCodeHayesAndHarlington                CRSCode = "HAY"
	StationCodeHayesKent                         CRSCode = "HYS"
	StationCodeHayle                             CRSCode = "HYL"
	StationCodeHaymarket                         CRSCode = "HYM"
	StationCodeHaywardsHeath                     CRSCode = "HHE"
	StationCodeHazelGrove                        CRSCode = "HAZ"
	StationCodeHeadcorn                          CRSCode = "HCN"
	StationCodeHeadingley                        CRSCode = "HDY"
	StationCodeHeadstoneLane                     CRSCode = "HDL"
	StationCodeHealdGreen                        CRSCode = "HDG"
	StationCodeHealing                           CRSCode = "HLI"
	StationCodeHeathHighLevel                    CRSCode = "HHL"
	StationCodeHeathLowLevel                     CRSCode = "HLL"
	StationCodeHeathrowAirportTerminal4          CRSCode = "HAF"
	StationCodeHeathrowAirportTerminal5          CRSCode = "HWV"
	StationCodeHeathrowAirportTerminals12and3    CRSCode = "HXX"
	StationCodeHeatonChapel                      CRSCode = "HTC"
	StationCodeHebdenBridge                      CRSCode = "HBD"
	StationCodeHeckington                        CRSCode = "HEC"
	StationCodeHedgeEnd                          CRSCode = "HDE"
	StationCodeHednesford                        CRSCode = "HNF"
	StationCodeHeighington                       CRSCode = "HEI"
	StationCodeHelensburghCentral                CRSCode = "HLC"
	StationCodeHelensburghUpper                  CRSCode = "HLU"
	StationCodeHellifield                        CRSCode = "HLD"
	StationCodeHelmsdale                         CRSCode = "HMS"
	StationCodeHelsby                            CRSCode = "HSB"
	StationCodeHemelHempstead                    CRSCode = "HML"
	StationCodeHendon                            CRSCode = "HEN"
	StationCodeHengoed                           CRSCode = "HNG"
	StationCodeHenleyinArden                     CRSCode = "HNL"
	StationCodeHenleyonThames                    CRSCode = "HOT"
	StationCodeHensall                           CRSCode = "HEL"
	StationCodeHereford                          CRSCode = "HFD"
	StationCodeHerneBay                          CRSCode = "HNB"
	StationCodeHerneHill                         CRSCode = "HNH"
	StationCodeHersham                           CRSCode = "HER"
	StationCodeHertfordEast                      CRSCode = "HFE"
	StationCodeHertfordNorth                     CRSCode = "HFN"
	StationCodeHessle                            CRSCode = "HES"
	StationCodeHeswall                           CRSCode = "HSW"
	StationCodeHever                             CRSCode = "HEV"
	StationCodeHeworth                           CRSCode = "HEW"
	StationCodeHexham                            CRSCode = "HEX"
	StationCodeHeyford                           CRSCode = "HYD"
	StationCodeHeyshamPort                       CRSCode = "HHB"
	StationCodeHighBrooms                        CRSCode = "HIB"
	StationCodeHighStreetGlasgow                 CRSCode = "HST"
	StationCodeHighStreetKensingtonUnderground   CRSCode = "ZHS"
	StationCodeHighWycombe                       CRSCode = "HWY"
	StationCodeHigham                            CRSCode = "HGM"
	StationCodeHighamsPark                       CRSCode = "HIP"
	StationCodeHighbridgeAndBurnham              CRSCode = "HIG"
	StationCodeHighburyAndIslington              CRSCode = "HHY"
	StationCodeHightown                          CRSCode = "HTO"
	StationCodeHildenborough                     CRSCode = "HLB"
	StationCodeHillfoot                          CRSCode = "HLF"
	StationCodeHillingtonEast                    CRSCode = "HLE"
	StationCodeHillingtonWest                    CRSCode = "HLW"
	StationCodeHillside                          CRSCode = "HIL"
	StationCodeHilsea                            CRSCode = "HLS"
	StationCodeHinchleyWood                      CRSCode = "HYW"
	StationCodeHinckleyLeics                     CRSCode = "HNK"
	StationCodeHindley                           CRSCode = "HIN"
	StationCodeHintonAdmiral                     CRSCode = "HNA"
	StationCodeHitchin                           CRSCode = "HIT"
	StationCodeHitherGreen                       CRSCode = "HGR"
	StationCodeHockley                           CRSCode = "HOC"
	StationCodeHollingbourne                     CRSCode = "HBN"
	StationCodeHolmesChapel                      CRSCode = "HCH"
	StationCodeHolmwood                          CRSCode = "HLM"
	StationCodeHoltonHeath                       CRSCode = "HOL"
	StationCodeHolyhead                          CRSCode = "HHD"
	StationCodeHolytown                          CRSCode = "HLY"
	StationCodeHomerton                          CRSCode = "HMN"
	StationCodeHoneybourne                       CRSCode = "HYB"
	StationCodeHoniton                           CRSCode = "HON"
	StationCodeHonley                            CRSCode = "HOY"
	StationCodeHonorOakPark                      CRSCode = "HPA"
	StationCodeHook                              CRSCode = "HOK"
	StationCodeHooton                            CRSCode = "HOO"
	StationCodeHopeDerbyshire                    CRSCode = "HOP"
	StationCodeHopeFlintshire                    CRSCode = "HPE"
	StationCodeHoptonHeath                       CRSCode = "HPT"
	StationCodeHorley                            CRSCode = "HOR"
	StationCodeHornbeamPark                      CRSCode = "HBP"
	StationCodeHornsey                           CRSCode = "HRN"
	StationCodeHorsforth                         CRSCode = "HRS"
	StationCodeHorsham                           CRSCode = "HRH"
	StationCodeHorsley                           CRSCode = "HSY"
	StationCodeHortoninRibblesdale               CRSCode = "HIR"
	StationCodeHorwichParkway                    CRSCode = "HWI"
	StationCodeHoscar                            CRSCode = "HSC"
	StationCodeHoughGreen                        CRSCode = "HGN"
	StationCodeHounslow                          CRSCode = "HOU"
	StationCodeHove                              CRSCode = "HOV"
	StationCodeHovetonAndWroxham                 CRSCode = "HXM"
	StationCodeHowWoodHerts                      CRSCode = "HWW"
	StationCodeHowden                            CRSCode = "HOW"
	StationCodeHowwoodRenfrewshire               CRSCode = "HOZ"
	StationCodeHoxton                            CRSCode = "HOX"
	StationCodeHoylake                           CRSCode = "HYK"
	StationCodeHubbertsBridge                    CRSCode = "HBB"
	StationCodeHucknall                          CRSCode = "HKN"
	StationCodeHuddersfield                      CRSCode = "HUD"
	StationCodeHull                              CRSCode = "HUL"
	StationCodeHumphreyPark                      CRSCode = "HUP"
	StationCodeHuncoat                           CRSCode = "HCT"
	StationCodeHungerford                        CRSCode = "HGD"
	StationCodeHunmanby                          CRSCode = "HUB"
	StationCodeHuntingdon                        CRSCode = "HUN"
	StationCodeHuntly                            CRSCode = "HNT"
	StationCodeHuntsCross                        CRSCode = "HNX"
	StationCodeHurstGreen                        CRSCode = "HUR"
	StationCodeHuttonCranswick                   CRSCode = "HUT"
	StationCodeHuyton                            CRSCode = "HUY"
	StationCodeHydeCentral                       CRSCode = "HYC"
	StationCodeHydeNorth                         CRSCode = "HYT"
	StationCodeHykeham                           CRSCode = "HKM"
	StationCodeHyndland                          CRSCode = "HYN"
	StationCodeHytheEssex                        CRSCode = "HYH"
	StationCodeIBMHalt                           CRSCode = "IBM"
	StationCodeIfield                            CRSCode = "IFI"
	StationCodeIlford                            CRSCode = "IFD"
	StationCodeIlkley                            CRSCode = "ILK"
	StationCodeImperialWharf                     CRSCode = "IMW"
	StationCodeInceAndElton                      CRSCode = "INE"
	StationCodeInceManchester                    CRSCode = "INC"
	StationCodeIngatestone                       CRSCode = "INT"
	StationCodeInsch                             CRSCode = "INS"
	StationCodeInvergordon                       CRSCode = "IGD"
	StationCodeInvergowrie                       CRSCode = "ING"
	StationCodeInverkeithing                     CRSCode = "INK"
	StationCodeInverkip                          CRSCode = "INP"
	StationCodeInverness                         CRSCode = "INV"
	StationCodeInvershin                         CRSCode = "INH"
	StationCodeInverurie                         CRSCode = "INR"
	StationCodeIpswich                           CRSCode = "IPS"
	StationCodeIrlam                             CRSCode = "IRL"
	StationCodeIrvine                            CRSCode = "IRV"
	StationCodeIsleworth                         CRSCode = "ISL"
	StationCodeIslip                             CRSCode = "ISP"
	StationCodeIver                              CRSCode = "IVR"
	StationCodeIvybridge                         CRSCode = "IVY"
	StationCodeJamesCook                         CRSCode = "JCH"
	StationCodeJewelleryQuarter                  CRSCode = "JEQ"
	StationCodeJohnstonPembs                     CRSCode = "JOH"
	StationCodeJohnstoneStrathclyde              CRSCode = "JHN"
	StationCodeJordanhill                        CRSCode = "JOR"
	StationCodeKearsleyManchester                CRSCode = "KSL"
	StationCodeKearsney                          CRSCode = "KSN"
	StationCodeKeighley                          CRSCode = "KEI"
	StationCodeKeith                             CRSCode = "KEH"
	StationCodeKelvedon                          CRSCode = "KEL"
	StationCodeKelvindale                        CRSCode = "KVD"
	StationCodeKemble                            CRSCode = "KEM"
	StationCodeKempstonHardwick                  CRSCode = "KMH"
	StationCodeKemptonParkRacecourse             CRSCode = "KMP"
	StationCodeKemsing                           CRSCode = "KMS"
	StationCodeKemsley                           CRSCode = "KML"
	StationCodeKendal                            CRSCode = "KEN"
	StationCodeKenley                            CRSCode = "KLY"
	StationCodeKennett                           CRSCode = "KNE"
	StationCodeKennishead                        CRSCode = "KNS"
	StationCodeKensalGreen                       CRSCode = "KNL"
	StationCodeKensalRise                        CRSCode = "KNR"
	StationCodeKensingtonOlympia                 CRSCode = "KPA"
	StationCodeKentHouse                         CRSCode = "KTH"
	StationCodeKentishTown                       CRSCode = "KTN"
	StationCodeKentishTownWest                   CRSCode = "KTW"
	StationCodeKenton                            CRSCode = "KNT"
	StationCodeKentsBank                         CRSCode = "KBK"
	StationCodeKettering                         CRSCode = "KET"
	StationCodeKewBridge                         CRSCode = "KWB"
	StationCodeKewGardens                        CRSCode = "KWG"
	StationCodeKeyham                            CRSCode = "KEY"
	StationCodeKeynsham                          CRSCode = "KYN"
	StationCodeKidbrooke                         CRSCode = "KDB"
	StationCodeKidderminster                     CRSCode = "KID"
	StationCodeKidsgrove                         CRSCode = "KDG"
	StationCodeKidwelly                          CRSCode = "KWL"
	StationCodeKilburnHighRoad                   CRSCode = "KBN"
	StationCodeKildale                           CRSCode = "KLD"
	StationCodeKildonan                          CRSCode = "KIL"
	StationCodeKilgetty                          CRSCode = "KGT"
	StationCodeKilmarnock                        CRSCode = "KMK"
	StationCodeKilmaurs                          CRSCode = "KLM"
	StationCodeKilpatrick                        CRSCode = "KPT"
	StationCodeKilwinning                        CRSCode = "KWN"
	StationCodeKinbrace                          CRSCode = "KBC"
	StationCodeKingham                           CRSCode = "KGM"
	StationCodeKinghorn                          CRSCode = "KGH"
	StationCodeKingsLangley                      CRSCode = "KGL"
	StationCodeKingsLynn                         CRSCode = "KLN"
	StationCodeKingsNorton                       CRSCode = "KNN"
	StationCodeKingsNympton                      CRSCode = "KGN"
	StationCodeKingsPark                         CRSCode = "KGP"
	StationCodeKingsSutton                       CRSCode = "KGS"
	StationCodeKingsknowe                        CRSCode = "KGE"
	StationCodeKingston                          CRSCode = "KNG"
	StationCodeKingswood                         CRSCode = "KND"
	StationCodeKingussie                         CRSCode = "KIN"
	StationCodeKintbury                          CRSCode = "KIT"
	StationCodeKirbyCross                        CRSCode = "KBX"
	StationCodeKirkSandall                       CRSCode = "KKS"
	StationCodeKirkbyMerseyside                  CRSCode = "KIR"
	StationCodeKirkbyStephen                     CRSCode = "KSW"
	StationCodeKirkbyinAshfield                  CRSCode = "KKB"
	StationCodeKirkbyinFurness                   CRSCode = "KBF"
	StationCodeKirkcaldy                         CRSCode = "KDY"
	StationCodeKirkconnel                        CRSCode = "KRK"
	StationCodeKirkdale                          CRSCode = "KKD"
	StationCodeKirkhamAndWesham                  CRSCode = "KKM"
	StationCodeKirkhill                          CRSCode = "KKH"
	StationCodeKirknewton                        CRSCode = "KKN"
	StationCodeKirkstallForge                    CRSCode = "KLF"
	StationCodeKirkwood                          CRSCode = "KWD"
	StationCodeKirtonLindsey                     CRSCode = "KTL"
	StationCodeKivetonBridge                     CRSCode = "KIV"
	StationCodeKivetonPark                       CRSCode = "KVP"
	StationCodeKnaresborough                     CRSCode = "KNA"
	StationCodeKnebworth                         CRSCode = "KBW"
	StationCodeKnighton                          CRSCode = "KNI"
	StationCodeKnockholt                         CRSCode = "KCK"
	StationCodeKnottingley                       CRSCode = "KNO"
	StationCodeKnucklas                          CRSCode = "KNU"
	StationCodeKnutsford                         CRSCode = "KNF"
	StationCodeKyleofLochalsh                    CRSCode = "KYL"
	StationCodeLadybank                          CRSCode = "LDY"
	StationCodeLadywell                          CRSCode = "LAD"
	StationCodeLaindon                           CRSCode = "LAI"
	StationCodeLairg                             CRSCode = "LRG"
	StationCodeLake                              CRSCode = "LKE"
	StationCodeLakenheath                        CRSCode = "LAK"
	StationCodeLamphey                           CRSCode = "LAM"
	StationCodeLanark                            CRSCode = "LNK"
	StationCodeLancaster                         CRSCode = "LAN"
	StationCodeLancing                           CRSCode = "LAC"
	StationCodeLandywood                         CRSCode = "LAW"
	StationCodeLangbank                          CRSCode = "LGB"
	StationCodeLangho                            CRSCode = "LHO"
	StationCodeLangleyBerks                      CRSCode = "LNY"
	StationCodeLangleyGreen                      CRSCode = "LGG"
	StationCodeLangleyMill                       CRSCode = "LGM"
	StationCodeLangside                          CRSCode = "LGS"
	StationCodeLangwathby                        CRSCode = "LGW"
	StationCodeLangwithWhaleyThorns              CRSCode = "LAG"
	StationCodeLapford                           CRSCode = "LAP"
	StationCodeLapworth                          CRSCode = "LPW"
	StationCodeLarbert                           CRSCode = "LBT"
	StationCodeLargs                             CRSCode = "LAR"
	StationCodeLarkhall                          CRSCode = "LRH"
	StationCodeLaurencekirk                      CRSCode = "LAU"
	StationCodeLawrenceHill                      CRSCode = "LWH"
	StationCodeLaytonLancs                       CRSCode = "LAY"
	StationCodeLazonbyAndKirkoswald              CRSCode = "LZB"
	StationCodeLeaBridge                         CRSCode = "LEB"
	StationCodeLeaGreen                          CRSCode = "LEG"
	StationCodeLeaHall                           CRSCode = "LEH"
	StationCodeLeagrave                          CRSCode = "LEA"
	StationCodeLealholm                          CRSCode = "LHM"
	StationCodeLeamingtonSpa                     CRSCode = "LMS"
	StationCodeLeasowe                           CRSCode = "LSW"
	StationCodeLeatherhead                       CRSCode = "LHD"
	StationCodeLedbury                           CRSCode = "LED"
	StationCodeLeeLondon                         CRSCode = "LEE"
	StationCodeLeeds                             CRSCode = "LDS"
	StationCodeLeicester                         CRSCode = "LEI"
	StationCodeLeighKent                         CRSCode = "LIH"
	StationCodeLeighonSea                        CRSCode = "LES"
	StationCodeLeightonBuzzard                   CRSCode = "LBZ"
	StationCodeLelant                            CRSCode = "LEL"
	StationCodeLelantSaltings                    CRSCode = "LTS"
	StationCodeLenham                            CRSCode = "LEN"
	StationCodeLenzie                            CRSCode = "LNZ"
	StationCodeLeominster                        CRSCode = "LEO"
	StationCodeLetchworthGardenCity              CRSCode = "LET"
	StationCodeLeucharsforStAndrews              CRSCode = "LEU"
	StationCodeLevenshulme                       CRSCode = "LVM"
	StationCodeLewes                             CRSCode = "LWS"
	StationCodeLewisham                          CRSCode = "LEW"
	StationCodeLeyland                           CRSCode = "LEY"
	StationCodeLeytonMidlandRoad                 CRSCode = "LEM"
	StationCodeLeytonstoneHighRoad               CRSCode = "LER"
	StationCodeLichfieldCity                     CRSCode = "LIC"
	StationCodeLichfieldTrentValley              CRSCode = "LTV"
	StationCodeLidlington                        CRSCode = "LID"
	StationCodeLimehouse                         CRSCode = "LHS"
	StationCodeLincolnCentral                    CRSCode = "LCN"
	StationCodeLingfield                         CRSCode = "LFD"
	StationCodeLingwood                          CRSCode = "LGD"
	StationCodeLinlithgow                        CRSCode = "LIN"
	StationCodeLiphook                           CRSCode = "LIP"
	StationCodeLiskeard                          CRSCode = "LSK"
	StationCodeLiss                              CRSCode = "LIS"
	StationCodeLisvaneAndThornhill               CRSCode = "LVT"
	StationCodeLittleKimble                      CRSCode = "LTK"
	StationCodeLittleSutton                      CRSCode = "LTT"
	StationCodeLittleborough                     CRSCode = "LTL"
	StationCodeLittlehampton                     CRSCode = "LIT"
	StationCodeLittlehaven                       CRSCode = "LVN"
	StationCodeLittleport                        CRSCode = "LTP"
	StationCodeLiverpoolCentral                  CRSCode = "LVC"
	StationCodeLiverpoolJamesStreet              CRSCode = "LVJ"
	StationCodeLiverpoolLimeStreet               CRSCode = "LIV"
	StationCodeLiverpoolSouthParkway             CRSCode = "LPY"
	StationCodeLivingstonNorth                   CRSCode = "LSN"
	StationCodeLivingstonSouth                   CRSCode = "LVG"
	StationCodeLlanaber                          CRSCode = "LLA"
	StationCodeLlanbedr                          CRSCode = "LBR"
	StationCodeLlanbisterRoad                    CRSCode = "LLT"
	StationCodeLlanbradach                       CRSCode = "LNB"
	StationCodeLlandaf                           CRSCode = "LLN"
	StationCodeLlandanwg                         CRSCode = "LDN"
	StationCodeLlandecwyn                        CRSCode = "LLC"
	StationCodeLlandeilo                         CRSCode = "LLL"
	StationCodeLlandovery                        CRSCode = "LLV"
	StationCodeLlandrindod                       CRSCode = "LLO"
	StationCodeLlandudno                         CRSCode = "LLD"
	StationCodeLlandudnoJunction                 CRSCode = "LLJ"
	StationCodeLlandybie                         CRSCode = "LLI"
	StationCodeLlanelli                          CRSCode = "LLE"
	StationCodeLlanfairfechan                    CRSCode = "LLF"
	StationCodeLlanfairpwll                      CRSCode = "LPG"
	StationCodeLlangadog                         CRSCode = "LLG"
	StationCodeLlangammarch                      CRSCode = "LLM"
	StationCodeLlangennech                       CRSCode = "LLH"
	StationCodeLlangynllo                        CRSCode = "LGO"
	StationCodeLlanharan                         CRSCode = "LLR"
	StationCodeLlanhilleth                       CRSCode = "LTH"
	StationCodeLlanishen                         CRSCode = "LLS"
	StationCodeLlanrwst                          CRSCode = "LWR"
	StationCodeLlansamlet                        CRSCode = "LAS"
	StationCodeLlantwitMajor                     CRSCode = "LWM"
	StationCodeLlanwrda                          CRSCode = "LNR"
	StationCodeLlanwrtyd                         CRSCode = "LNW"
	StationCodeLlwyngwril                        CRSCode = "LLW"
	StationCodeLlwynypia                         CRSCode = "LLY"
	StationCodeLochAwe                           CRSCode = "LHA"
	StationCodeLochEilOutwardBound               CRSCode = "LHE"
	StationCodeLochailort                        CRSCode = "LCL"
	StationCodeLocheilside                       CRSCode = "LCS"
	StationCodeLochgelly                         CRSCode = "LCG"
	StationCodeLochluichart                      CRSCode = "LCC"
	StationCodeLochwinnoch                       CRSCode = "LHW"
	StationCodeLockerbie                         CRSCode = "LOC"
	StationCodeLockwood                          CRSCode = "LCK"
	StationCodeLondonBlackfriars                 CRSCode = "BFR"
	StationCodeLondonBridge                      CRSCode = "LBG"
	StationCodeLondonCannonStreet                CRSCode = "CST"
	StationCodeLondonCharingCross                CRSCode = "CHX"
	StationCodeLondonEuston                      CRSCode = "EUS"
	StationCodeLondonFenchurchStreet             CRSCode = "FST"
	StationCodeLondonFields                      CRSCode = "LOF"
	StationCodeLondonKingsCross                  CRSCode = "KGX"
	StationCodeLondonLiverpoolStreet             CRSCode = "LST"
	StationCodeLondonMarylebone                  CRSCode = "MYB"
	StationCodeLondonPaddington                  CRSCode = "PAD"
	StationCodeLondonRoadBrighton                CRSCode = "LRB"
	StationCodeLondonRoadGuildford               CRSCode = "LRD"
	StationCodeLondonStPancrasIntl               CRSCode = "STP"
	StationCodeLondonVictoria                    CRSCode = "VIC"
	StationCodeLondonWaterloo                    CRSCode = "WAT"
	StationCodeLondonWaterlooEast                CRSCode = "WAE"
	StationCodeLongBuckby                        CRSCode = "LBK"
	StationCodeLongEaton                         CRSCode = "LGE"
	StationCodeLongPreston                       CRSCode = "LPR"
	StationCodeLongbeck                          CRSCode = "LGK"
	StationCodeLongbridge                        CRSCode = "LOB"
	StationCodeLongcross                         CRSCode = "LNG"
	StationCodeLongfield                         CRSCode = "LGF"
	StationCodeLongniddry                        CRSCode = "LND"
	StationCodeLongport                          CRSCode = "LPT"
	StationCodeLongton                           CRSCode = "LGN"
	StationCodeLooe                              CRSCode = "LOO"
	StationCodeLostock                           CRSCode = "LOT"
	StationCodeLostockGralam                     CRSCode = "LTG"
	StationCodeLostockHall                       CRSCode = "LOH"
	StationCodeLostwithiel                       CRSCode = "LOS"
	StationCodeLoughborough                      CRSCode = "LBO"
	StationCodeLoughboroughJunction              CRSCode = "LGJ"
	StationCodeLowdham                           CRSCode = "LOW"
	StationCodeLowerSydenham                     CRSCode = "LSY"
	StationCodeLowestoft                         CRSCode = "LWT"
	StationCodeLudlow                            CRSCode = "LUD"
	StationCodeLuton                             CRSCode = "LUT"
	StationCodeLutonAirportParkway               CRSCode = "LTN"
	StationCodeLuxulyan                          CRSCode = "LUX"
	StationCodeLydney                            CRSCode = "LYD"
	StationCodeLyeWestMidlands                   CRSCode = "LYE"
	StationCodeLymingtonPier                     CRSCode = "LYP"
	StationCodeLymingtonTown                     CRSCode = "LYT"
	StationCodeLympstoneCommando                 CRSCode = "LYC"
	StationCodeLympstoneVillage                  CRSCode = "LYM"
	StationCodeLytham                            CRSCode = "LTM"
	StationCodeMacclesfield                      CRSCode = "MAC"
	StationCodeMachynlleth                       CRSCode = "MCN"
	StationCodeMaesteg                           CRSCode = "MST"
	StationCodeMaestegEwennyRoad                 CRSCode = "MEW"
	StationCodeMaghull                           CRSCode = "MAG"
	StationCodeMaidenNewton                      CRSCode = "MDN"
	StationCodeMaidenhead                        CRSCode = "MAI"
	StationCodeMaidstoneBarracks                 CRSCode = "MDB"
	StationCodeMaidstoneEast                     CRSCode = "MDE"
	StationCodeMaidstoneWest                     CRSCode = "MDW"
	StationCodeMaldenManor                       CRSCode = "MAL"
	StationCodeMallaig                           CRSCode = "MLG"
	StationCodeMalton                            CRSCode = "MLT"
	StationCodeMalvernLink                       CRSCode = "MVL"
	StationCodeManchesterAirport                 CRSCode = "MIA"
	StationCodeManchesterOxfordRoad              CRSCode = "MCO"
	StationCodeManchesterPiccadilly              CRSCode = "MAN"
	StationCodeManchesterUnitedFootballGround    CRSCode = "MUF"
	StationCodeManchesterVictoria                CRSCode = "MCV"
	StationCodeManea                             CRSCode = "MNE"
	StationCodeManningtree                       CRSCode = "MNG"
	StationCodeManorPark                         CRSCode = "MNP"
	StationCodeManorRoad                         CRSCode = "MNR"
	StationCodeManorbier                         CRSCode = "MRB"
	StationCodeManors                            CRSCode = "MAS"
	StationCodeMansfield                         CRSCode = "MFT"
	StationCodeMansfieldWoodhouse                CRSCode = "MSW"
	StationCodeMarch                             CRSCode = "MCH"
	StationCodeMardenKent                        CRSCode = "MRN"
	StationCodeMargate                           CRSCode = "MAR"
	StationCodeMarketHarborough                  CRSCode = "MHR"
	StationCodeMarketRasen                       CRSCode = "MKR"
	StationCodeMarkinch                          CRSCode = "MNC"
	StationCodeMarksTey                          CRSCode = "MKT"
	StationCodeMarlow                            CRSCode = "MLW"
	StationCodeMarple                            CRSCode = "MPL"
	StationCodeMarsdenYorks                      CRSCode = "MSN"
	StationCodeMarske                            CRSCode = "MSK"
	StationCodeMarstonGreen                      CRSCode = "MGN"
	StationCodeMartinMill                        CRSCode = "MTM"
	StationCodeMartinsHeron                      CRSCode = "MAO"
	StationCodeMarton                            CRSCode = "MTO"
	StationCodeMaryhill                          CRSCode = "MYH"
	StationCodeMaryland                          CRSCode = "MYL"
	StationCodeMaryport                          CRSCode = "MRY"
	StationCodeMatlock                           CRSCode = "MAT"
	StationCodeMatlockBath                       CRSCode = "MTB"
	StationCodeMauldethRoad                      CRSCode = "MAU"
	StationCodeMaxwellPark                       CRSCode = "MAX"
	StationCodeMaybole                           CRSCode = "MAY"
	StationCodeMazeHill                          CRSCode = "MZH"
	StationCodeMeadowhall                        CRSCode = "MHS"
	StationCodeMeldreth                          CRSCode = "MEL"
	StationCodeMelksham                          CRSCode = "MKM"
	StationCodeMeltonMowbray                     CRSCode = "MMO"
	StationCodeMeltonSuffolk                     CRSCode = "MES"
	StationCodeMenheniot                         CRSCode = "MEN"
	StationCodeMenston                           CRSCode = "MNN"
	StationCodeMeols                             CRSCode = "MEO"
	StationCodeMeolsCop                          CRSCode = "MEC"
	StationCodeMeopham                           CRSCode = "MEP"
	StationCodeMerryton                          CRSCode = "MEY"
	StationCodeMerstham                          CRSCode = "MHM"
	StationCodeMerthyrTydfil                     CRSCode = "MER"
	StationCodeMerthyrVale                       CRSCode = "MEV"
	StationCodeMetheringham                      CRSCode = "MGM"
	StationCodeMetroCentre                       CRSCode = "MCE"
	StationCodeMexborough                        CRSCode = "MEX"
	StationCodeMicheldever                       CRSCode = "MIC"
	StationCodeMicklefield                       CRSCode = "MIK"
	StationCodeMiddlesbrough                     CRSCode = "MBR"
	StationCodeMiddlewood                        CRSCode = "MDL"
	StationCodeMidgham                           CRSCode = "MDG"
	StationCodeMilfordHaven                      CRSCode = "MFH"
	StationCodeMilfordSurrey                     CRSCode = "MLF"
	StationCodeMillHillBroadway                  CRSCode = "MIL"
	StationCodeMillHillLancs                     CRSCode = "MLH"
	StationCodeMillbrookBeds                     CRSCode = "MLB"
	StationCodeMillbrookHants                    CRSCode = "MBK"
	StationCodeMillikenPark                      CRSCode = "MIN"
	StationCodeMillom                            CRSCode = "MLM"
	StationCodeMillsHillManchester               CRSCode = "MIH"
	StationCodeMilngavie                         CRSCode = "MLN"
	StationCodeMiltonKeynesCentral               CRSCode = "MKC"
	StationCodeMinffordd                         CRSCode = "MFF"
	StationCodeMinster                           CRSCode = "MSR"
	StationCodeMirfield                          CRSCode = "MIR"
	StationCodeMistley                           CRSCode = "MIS"
	StationCodeMitchamEastfields                 CRSCode = "MTC"
	StationCodeMitchamJunction                   CRSCode = "MIJ"
	StationCodeMobberley                         CRSCode = "MOB"
	StationCodeMonifieth                         CRSCode = "MON"
	StationCodeMonksRisborough                   CRSCode = "MRS"
	StationCodeMontpelier                        CRSCode = "MTP"
	StationCodeMontrose                          CRSCode = "MTS"
	StationCodeMoorfields                        CRSCode = "MRF"
	StationCodeMoorgate                          CRSCode = "MOG"
	StationCodeMoorside                          CRSCode = "MSD"
	StationCodeMoorthorpe                        CRSCode = "MRP"
	StationCodeMorar                             CRSCode = "MRR"
	StationCodeMorchardRoad                      CRSCode = "MRD"
	StationCodeMordenSouth                       CRSCode = "MDS"
	StationCodeMorecambe                         CRSCode = "MCM"
	StationCodeMoretonDorset                     CRSCode = "MTN"
	StationCodeMoretonMerseyside                 CRSCode = "MRT"
	StationCodeMoretoninMarsh                    CRSCode = "MIM"
	StationCodeMorfaMawddach                     CRSCode = "MFA"
	StationCodeMorley                            CRSCode = "MLY"
	StationCodeMorpeth                           CRSCode = "MPT"
	StationCodeMortimer                          CRSCode = "MOR"
	StationCodeMortlake                          CRSCode = "MTL"
	StationCodeMosesGate                         CRSCode = "MSS"
	StationCodeMossSide                          CRSCode = "MOS"
	StationCodeMossleyHill                       CRSCode = "MSH"
	StationCodeMossleyManchester                 CRSCode = "MSL"
	StationCodeMosspark                          CRSCode = "MPK"
	StationCodeMoston                            CRSCode = "MSO"
	StationCodeMotherwell                        CRSCode = "MTH"
	StationCodeMotspurPark                       CRSCode = "MOT"
	StationCodeMottingham                        CRSCode = "MTG"
	StationCodeMottisfontAndDunbridge            CRSCode = "DBG"
	StationCodeMouldsworth                       CRSCode = "MLD"
	StationCodeMoulsecoomb                       CRSCode = "MCB"
	StationCodeMountFlorida                      CRSCode = "MFL"
	StationCodeMountVernon                       CRSCode = "MTV"
	StationCodeMountainAsh                       CRSCode = "MTA"
	StationCodeMuirend                           CRSCode = "MUI"
	StationCodeMuirofOrd                         CRSCode = "MOO"
	StationCodeMusselburgh                       CRSCode = "MUB"
	StationCodeMytholmroyd                       CRSCode = "MYT"
	StationCodeNafferton                         CRSCode = "NFN"
	StationCodeNailseaAndBackwell                CRSCode = "NLS"
	StationCodeNairn                             CRSCode = "NRN"
	StationCodeNantwich                          CRSCode = "NAN"
	StationCodeNarberth                          CRSCode = "NAR"
	StationCodeNarborough                        CRSCode = "NBR"
	StationCodeNavigationRoad                    CRSCode = "NVR"
	StationCodeNeath                             CRSCode = "NTH"
	StationCodeNeedhamMarket                     CRSCode = "NMT"
	StationCodeNeilston                          CRSCode = "NEI"
	StationCodeNelson                            CRSCode = "NEL"
	StationCodeNeston                            CRSCode = "NES"
	StationCodeNetherfield                       CRSCode = "NET"
	StationCodeNethertown                        CRSCode = "NRT"
	StationCodeNetley                            CRSCode = "NTL"
	StationCodeNewBarnet                         CRSCode = "NBA"
	StationCodeNewBeckenham                      CRSCode = "NBC"
	StationCodeNewBrighton                       CRSCode = "NBN"
	StationCodeNewClee                           CRSCode = "NCE"
	StationCodeNewCross                          CRSCode = "NWX"
	StationCodeNewCrossGate                      CRSCode = "NXG"
	StationCodeNewCumnock                        CRSCode = "NCK"
	StationCodeNewEltham                         CRSCode = "NEH"
	StationCodeNewHolland                        CRSCode = "NHL"
	StationCodeNewHythe                          CRSCode = "NHE"
	StationCodeNewLane                           CRSCode = "NLN"
	StationCodeNewMalden                         CRSCode = "NEM"
	StationCodeNewMillsCentral                   CRSCode = "NMC"
	StationCodeNewMillsNewtown                   CRSCode = "NMN"
	StationCodeNewMilton                         CRSCode = "NWM"
	StationCodeNewPudsey                         CRSCode = "NPD"
	StationCodeNewSouthgate                      CRSCode = "NSG"
	StationCodeNewarkCastle                      CRSCode = "NCT"
	StationCodeNewarkNorthGate                   CRSCode = "NNG"
	StationCodeNewbridge                         CRSCode = "NBE"
	StationCodeNewbury                           CRSCode = "NBY"
	StationCodeNewburyRacecourse                 CRSCode = "NRC"
	StationCodeNewcastle                         CRSCode = "NCL"
	StationCodeNewcourt                          CRSCode = "NCO"
	StationCodeNewcraighall                      CRSCode = "NEW"
	StationCodeNewhavenHarbour                   CRSCode = "NVH"
	StationCodeNewhavenTown                      CRSCode = "NVN"
	StationCodeNewington                         CRSCode = "NGT"
	StationCodeNewmarket                         CRSCode = "NMK"
	StationCodeNewportEssex                      CRSCode = "NWE"
	StationCodeNewportSouthWales                 CRSCode = "NWP"
	StationCodeNewquay                           CRSCode = "NQY"
	StationCodeNewstead                          CRSCode = "NSD"
	StationCodeNewtonAbbot                       CRSCode = "NTA"
	StationCodeNewtonAycliffe                    CRSCode = "NAY"
	StationCodeNewtonLanark                      CRSCode = "NTN"
	StationCodeNewtonStCyres                     CRSCode = "NTC"
	StationCodeNewtonforHyde                     CRSCode = "NWN"
	StationCodeNewtongrange                      CRSCode = "NEG"
	StationCodeNewtonleWillows                   CRSCode = "NLW"
	StationCodeNewtonmore                        CRSCode = "NWR"
	StationCodeNewtononAyr                       CRSCode = "NOA"
	StationCodeNewtownPowys                      CRSCode = "NWT"
	StationCodeNinianPark                        CRSCode = "NNP"
	StationCodeNitshill                          CRSCode = "NIT"
	StationCodeNorbiton                          CRSCode = "NBT"
	StationCodeNorbury                           CRSCode = "NRB"
	StationCodeNormansBay                        CRSCode = "NSB"
	StationCodeNormanton                         CRSCode = "NOR"
	StationCodeNorthBerwick                      CRSCode = "NBW"
	StationCodeNorthCamp                         CRSCode = "NCM"
	StationCodeNorthDulwich                      CRSCode = "NDL"
	StationCodeNorthFambridge                    CRSCode = "NFA"
	StationCodeNorthLlanrwst                     CRSCode = "NLR"
	StationCodeNorthQueensferry                  CRSCode = "NQU"
	StationCodeNorthRoadDarlington               CRSCode = "NRD"
	StationCodeNorthSheen                        CRSCode = "NSH"
	StationCodeNorthWalsham                      CRSCode = "NWA"
	StationCodeNorthWembley                      CRSCode = "NWB"
	StationCodeNorthallerton                     CRSCode = "NTR"
	StationCodeNorthampton                       CRSCode = "NMP"
	StationCodeNorthfield                        CRSCode = "NFD"
	StationCodeNorthfleet                        CRSCode = "NFL"
	StationCodeNortholtPark                      CRSCode = "NLT"
	StationCodeNorthumberlandPark                CRSCode = "NUM"
	StationCodeNorthwich                         CRSCode = "NWI"
	StationCodeNortonBridge                      CRSCode = "NTB"
	StationCodeNorwich                           CRSCode = "NRW"
	StationCodeNorwoodJunction                   CRSCode = "NWD"
	StationCodeNottingham                        CRSCode = "NOT"
	StationCodeNuneaton                          CRSCode = "NUN"
	StationCodeNunhead                           CRSCode = "NHD"
	StationCodeNunthorpe                         CRSCode = "NNT"
	StationCodeNutbourne                         CRSCode = "NUT"
	StationCodeNutfield                          CRSCode = "NUF"
	StationCodeOakengates                        CRSCode = "OKN"
	StationCodeOakham                            CRSCode = "OKM"
	StationCodeOakleighPark                      CRSCode = "OKL"
	StationCodeOban                              CRSCode = "OBN"
	StationCodeOckendon                          CRSCode = "OCK"
	StationCodeOckley                            CRSCode = "OLY"
	StationCodeOkehampton                        CRSCode = "OKE"
	StationCodeOldHill                           CRSCode = "OHL"
	StationCodeOldRoan                           CRSCode = "ORN"
	StationCodeOldStreet                         CRSCode = "OLD"
	StationCodeOldfieldPark                      CRSCode = "OLF"
	StationCodeOlton                             CRSCode = "OLT"
	StationCodeOre                               CRSCode = "ORE"
	StationCodeOrmskirk                          CRSCode = "OMS"
	StationCodeOrpington                         CRSCode = "ORP"
	StationCodeOrrell                            CRSCode = "ORR"
	StationCodeOrrellPark                        CRSCode = "OPK"
	StationCodeOtford                            CRSCode = "OTF"
	StationCodeOultonBroadNorth                  CRSCode = "OUN"
	StationCodeOultonBroadSouth                  CRSCode = "OUS"
	StationCodeOutwood                           CRSCode = "OUT"
	StationCodeOverpool                          CRSCode = "OVE"
	StationCodeOverton                           CRSCode = "OVR"
	StationCodeOxenholmeLakeDistrict             CRSCode = "OXN"
	StationCodeOxford                            CRSCode = "OXF"
	StationCodeOxfordParkway                     CRSCode = "OXP"
	StationCodeOxshott                           CRSCode = "OXS"
	StationCodeOxted                             CRSCode = "OXT"
	StationCodePaddockWood                       CRSCode = "PDW"
	StationCodePadgate                           CRSCode = "PDG"
	StationCodePaignton                          CRSCode = "PGN"
	StationCodePaisleyCanal                      CRSCode = "PCN"
	StationCodePaisleyGilmourStreet              CRSCode = "PYG"
	StationCodePaisleyStJames                    CRSCode = "PYJ"
	StationCodePalmersGreen                      CRSCode = "PAL"
	StationCodePangbourne                        CRSCode = "PAN"
	StationCodePannal                            CRSCode = "PNL"
	StationCodePantyffynnon                      CRSCode = "PTF"
	StationCodePar                               CRSCode = "PAR"
	StationCodeParbold                           CRSCode = "PBL"
	StationCodeParkStreet                        CRSCode = "PKT"
	StationCodeParkstoneDorset                   CRSCode = "PKS"
	StationCodeParsonStreet                      CRSCode = "PSN"
	StationCodePartick                           CRSCode = "PTK"
	StationCodeParton                            CRSCode = "PRN"
	StationCodePatchway                          CRSCode = "PWY"
	StationCodePatricroft                        CRSCode = "PAT"
	StationCodePatterton                         CRSCode = "PTT"
	StationCodePeartree                          CRSCode = "PEA"
	StationCodePeckhamRye                        CRSCode = "PMR"
	StationCodePegswood                          CRSCode = "PEG"
	StationCodePemberton                         CRSCode = "PEM"
	StationCodePembreyAndBurryPort               CRSCode = "PBY"
	StationCodePembroke                          CRSCode = "PMB"
	StationCodePembrokeDock                      CRSCode = "PMD"
	StationCodePenally                           CRSCode = "PNA"
	StationCodePenarth                           CRSCode = "PEN"
	StationCodePencoed                           CRSCode = "PCD"
	StationCodePengam                            CRSCode = "PGM"
	StationCodePengeEast                         CRSCode = "PNE"
	StationCodePengeWest                         CRSCode = "PNW"
	StationCodePenhelig                          CRSCode = "PHG"
	StationCodePenistone                         CRSCode = "PNS"
	StationCodePenkridge                         CRSCode = "PKG"
	StationCodePenmaenmawr                       CRSCode = "PMW"
	StationCodePenmere                           CRSCode = "PNM"
	StationCodePenrhiwceiber                     CRSCode = "PER"
	StationCodePenrhyndeudraeth                  CRSCode = "PRH"
	StationCodePenrithNorthLakes                 CRSCode = "PNR"
	StationCodePenrynCornwall                    CRSCode = "PYN"
	StationCodePensarnGwynedd                    CRSCode = "PES"
	StationCodePenshurst                         CRSCode = "PHR"
	StationCodePentreBach                        CRSCode = "PTB"
	StationCodePenyBont                          CRSCode = "PNY"
	StationCodePenychain                         CRSCode = "PNC"
	StationCodePenyffordd                        CRSCode = "PNF"
	StationCodePenzance                          CRSCode = "PNZ"
	StationCodePerranwell                        CRSCode = "PRW"
	StationCodePerryBarr                         CRSCode = "PRY"
	StationCodePershore                          CRSCode = "PSH"
	StationCodePerth                             CRSCode = "PTH"
	StationCodePeterborough                      CRSCode = "PBO"
	StationCodePetersfield                       CRSCode = "PTR"
	StationCodePettsWood                         CRSCode = "PET"
	StationCodePevenseyAndWestham                CRSCode = "PEV"
	StationCodePevenseyBay                       CRSCode = "PEB"
	StationCodePewsey                            CRSCode = "PEW"
	StationCodePilning                           CRSCode = "PIL"
	StationCodePinhoe                            CRSCode = "PIN"
	StationCodePitlochry                         CRSCode = "PIT"
	StationCodePitsea                            CRSCode = "PSE"
	StationCodePleasington                       CRSCode = "PLS"
	StationCodePlockton                          CRSCode = "PLK"
	StationCodePluckley                          CRSCode = "PLC"
	StationCodePlumley                           CRSCode = "PLM"
	StationCodePlumpton                          CRSCode = "PMP"
	StationCodePlumstead                         CRSCode = "PLU"
	StationCodePlymouth                          CRSCode = "PLY"
	StationCodePokesdown                         CRSCode = "POK"
	StationCodePolegate                          CRSCode = "PLG"
	StationCodePolesworth                        CRSCode = "PSW"
	StationCodePollokshawsEast                   CRSCode = "PWE"
	StationCodePollokshawsWest                   CRSCode = "PWW"
	StationCodePollokshieldsEast                 CRSCode = "PLE"
	StationCodePollokshieldsWest                 CRSCode = "PLW"
	StationCodePolmont                           CRSCode = "PMT"
	StationCodePolsloeBridge                     CRSCode = "POL"
	StationCodePondersEnd                        CRSCode = "PON"
	StationCodePontarddulais                     CRSCode = "PTD"
	StationCodePontefractBaghill                 CRSCode = "PFR"
	StationCodePontefractMonkhill                CRSCode = "PFM"
	StationCodePontefractTanshelf                CRSCode = "POT"
	StationCodePontlottyn                        CRSCode = "PLT"
	StationCodePontyPant                         CRSCode = "PYP"
	StationCodePontyclun                         CRSCode = "PYC"
	StationCodePontypoolAndNewInn                CRSCode = "PPL"
	StationCodePontypridd                        CRSCode = "PPD"
	StationCodePoole                             CRSCode = "POO"
	StationCodePoppleton                         CRSCode = "POP"
	StationCodePortGlasgow                       CRSCode = "PTG"
	StationCodePortSunlight                      CRSCode = "PSL"
	StationCodePortTalbotParkway                 CRSCode = "PTA"
	StationCodePortchester                       CRSCode = "PTC"
	StationCodePorth                             CRSCode = "POR"
	StationCodePorthmadog                        CRSCode = "PTM"
	StationCodePortlethen                        CRSCode = "PLN"
	StationCodePortslade                         CRSCode = "PLD"
	StationCodePortsmouthAndSouthsea             CRSCode = "PMS"
	StationCodePortsmouthArms                    CRSCode = "PMA"
	StationCodePortsmouthHarbour                 CRSCode = "PMH"
	StationCodePossilparkAndParkhouse            CRSCode = "PPK"
	StationCodePottersBar                        CRSCode = "PBR"
	StationCodePoultonleFylde                    CRSCode = "PFY"
	StationCodePoynton                           CRSCode = "PYT"
	StationCodePrees                             CRSCode = "PRS"
	StationCodePrescot                           CRSCode = "PSC"
	StationCodePrestatyn                         CRSCode = "PRT"
	StationCodePrestbury                         CRSCode = "PRB"
	StationCodePrestonLancs                      CRSCode = "PRE"
	StationCodePrestonPark                       CRSCode = "PRP"
	StationCodePrestonpans                       CRSCode = "PST"
	StationCodePrestwickInternationalAirport     CRSCode = "PRA"
	StationCodePrestwickTown                     CRSCode = "PTW"
	StationCodePriesthillAndDarnley              CRSCode = "PTL"
	StationCodePrincesRisborough                 CRSCode = "PRR"
	StationCodePrittlewell                       CRSCode = "PRL"
	StationCodePrudhoe                           CRSCode = "PRU"
	StationCodePulborough                        CRSCode = "PUL"
	StationCodePurfleet                          CRSCode = "PFL"
	StationCodePurley                            CRSCode = "PUR"
	StationCodePurleyOaks                        CRSCode = "PUO"
	StationCodePutney                            CRSCode = "PUT"
	StationCodePwllheli                          CRSCode = "PWL"
	StationCodePyeCorner                         CRSCode = "PYE"
	StationCodePyle                              CRSCode = "PYL"
	StationCodeQuakersYard                       CRSCode = "QYD"
	StationCodeQueenborough                      CRSCode = "QBR"
	StationCodeQueensParkGlasgow                 CRSCode = "QPK"
	StationCodeQueensParkLondon                  CRSCode = "QPW"
	StationCodeQueensRoadPeckham                 CRSCode = "QRP"
	StationCodeQueenstownRoadBattersea           CRSCode = "QRB"
	StationCodeQuintrellDowns                    CRSCode = "QUI"
	StationCodeRadcliffeonTrent                  CRSCode = "RDF"
	StationCodeRadlett                           CRSCode = "RDT"
	StationCodeRadley                            CRSCode = "RAD"
	StationCodeRadyr                             CRSCode = "RDR"
	StationCodeRainford                          CRSCode = "RNF"
	StationCodeRainhamEssex                      CRSCode = "RNM"
	StationCodeRainhamKent                       CRSCode = "RAI"
	StationCodeRainhill                          CRSCode = "RNH"
	StationCodeRamsgate                          CRSCode = "RAM"
	StationCodeRamsgreaveAndWilpshire            CRSCode = "RGW"
	StationCodeRannoch                           CRSCode = "RAN"
	StationCodeRauceby                           CRSCode = "RAU"
	StationCodeRavenglassforEskdale              CRSCode = "RAV"
	StationCodeRavensbourne                      CRSCode = "RVB"
	StationCodeRavensthorpe                      CRSCode = "RVN"
	StationCodeRawcliffe                         CRSCode = "RWC"
	StationCodeRayleigh                          CRSCode = "RLG"
	StationCodeRaynesPark                        CRSCode = "RAY"
	StationCodeReading                           CRSCode = "RDG"
	StationCodeReadingWest                       CRSCode = "RDW"
	StationCodeRectoryRoad                       CRSCode = "REC"
	StationCodeRedbridge                         CRSCode = "RDB"
	StationCodeRedcarBritishSteel                CRSCode = "RBS"
	StationCodeRedcarCentral                     CRSCode = "RCC"
	StationCodeRedcarEast                        CRSCode = "RCE"
	StationCodeReddishNorth                      CRSCode = "RDN"
	StationCodeReddishSouth                      CRSCode = "RDS"
	StationCodeRedditch                          CRSCode = "RDC"
	StationCodeRedhill                           CRSCode = "RDH"
	StationCodeRedland                           CRSCode = "RDA"
	StationCodeRedruth                           CRSCode = "RED"
	StationCodeReedhamNorfolk                    CRSCode = "REE"
	StationCodeReedhamSurrey                     CRSCode = "RHM"
	StationCodeReigate                           CRSCode = "REI"
	StationCodeRenton                            CRSCode = "RTN"
	StationCodeRetford                           CRSCode = "RET"
	StationCodeRhiwbina                          CRSCode = "RHI"
	StationCodeRhooseCardiffInternationalAirport CRSCode = "RIA"
	StationCodeRhosneigr                         CRSCode = "RHO"
	StationCodeRhyl                              CRSCode = "RHL"
	StationCodeRhymney                           CRSCode = "RHY"
	StationCodeRibblehead                        CRSCode = "RHD"
	StationCodeRiceLane                          CRSCode = "RIL"
	StationCodeRichmondLondon                    CRSCode = "RMD"
	StationCodeRickmansworth                     CRSCode = "RIC"
	StationCodeRiddlesdown                       CRSCode = "RDD"
	StationCodeRidgmont                          CRSCode = "RID"
	StationCodeRidingMill                        CRSCode = "RDM"
	StationCodeRiscaAndPontymister               CRSCode = "RCA"
	StationCodeRishton                           CRSCode = "RIS"
	StationCodeRobertsbridge                     CRSCode = "RBR"
	StationCodeRoby                              CRSCode = "ROB"
	StationCodeRochdale                          CRSCode = "RCD"
	StationCodeRoche                             CRSCode = "ROC"
	StationCodeRochester                         CRSCode = "RTR"
	StationCodeRochford                          CRSCode = "RFD"
	StationCodeRockFerry                         CRSCode = "RFY"
	StationCodeRogart                            CRSCode = "ROG"
	StationCodeRogerstone                        CRSCode = "ROR"
	StationCodeRolleston                         CRSCode = "ROL"
	StationCodeRomanBridge                       CRSCode = "RMB"
	StationCodeRomford                           CRSCode = "RMF"
	StationCodeRomiley                           CRSCode = "RML"
	StationCodeRomsey                            CRSCode = "ROM"
	StationCodeRoose                             CRSCode = "ROO"
	StationCodeRoseGrove                         CRSCode = "RSG"
	StationCodeRoseHillMarple                    CRSCode = "RSH"
	StationCodeRosyth                            CRSCode = "ROS"
	StationCodeRotherhamCentral                  CRSCode = "RMC"
	StationCodeRotherhithe                       CRSCode = "ROE"
	StationCodeRoughtonRoad                      CRSCode = "RNR"
	StationCodeRowlandsCastle                    CRSCode = "RLN"
	StationCodeRowleyRegis                       CRSCode = "ROW"
	StationCodeRoyBridge                         CRSCode = "RYB"
	StationCodeRoydon                            CRSCode = "RYN"
	StationCodeRoyston                           CRSCode = "RYS"
	StationCodeRuabon                            CRSCode = "RUA"
	StationCodeRufford                           CRSCode = "RUF"
	StationCodeRugby                             CRSCode = "RUG"
	StationCodeRugeleyTown                       CRSCode = "RGT"
	StationCodeRugeleyTrentValley                CRSCode = "RGL"
	StationCodeRuncorn                           CRSCode = "RUN"
	StationCodeRuncornEast                       CRSCode = "RUE"
	StationCodeRuskington                        CRSCode = "RKT"
	StationCodeRuswarp                           CRSCode = "RUS"
	StationCodeRutherglen                        CRSCode = "RUT"
	StationCodeRydeEsplanade                     CRSCode = "RYD"
	StationCodeRydePierHead                      CRSCode = "RYP"
	StationCodeRydeStJohnsRoad                   CRSCode = "RYR"
	StationCodeRyderBrow                         CRSCode = "RRB"
	StationCodeRyeHouse                          CRSCode = "RYH"
	StationCodeRyeSussex                         CRSCode = "RYE"
	StationCodeSalfordCentral                    CRSCode = "SFD"
	StationCodeSalfordCrescent                   CRSCode = "SLD"
	StationCodeSalfordsSurrey                    CRSCode = "SAF"
	StationCodeSalhouse                          CRSCode = "SAH"
	StationCodeSalisbury                         CRSCode = "SAL"
	StationCodeSaltaire                          CRSCode = "SAE"
	StationCodeSaltash                           CRSCode = "STS"
	StationCodeSaltburn                          CRSCode = "SLB"
	StationCodeSaltcoats                         CRSCode = "SLT"
	StationCodeSaltmarshe                        CRSCode = "SAM"
	StationCodeSalwick                           CRSCode = "SLW"
	StationCodeSampfordCourtenay                 CRSCode = "SMC"
	StationCodeSandalAndAgbrigg                  CRSCode = "SNA"
	StationCodeSandbach                          CRSCode = "SDB"
	StationCodeSanderstead                       CRSCode = "SNR"
	StationCodeSandhills                         CRSCode = "SDL"
	StationCodeSandhurstBerks                    CRSCode = "SND"
	StationCodeSandling                          CRSCode = "SDG"
	StationCodeSandown                           CRSCode = "SAN"
	StationCodeSandplace                         CRSCode = "SDP"
	StationCodeSandwellAndDudley                 CRSCode = "SAD"
	StationCodeSandwich                          CRSCode = "SDW"
	StationCodeSandy                             CRSCode = "SDY"
	StationCodeSankeyforPenketh                  CRSCode = "SNK"
	StationCodeSanquhar                          CRSCode = "SQH"
	StationCodeSarn                              CRSCode = "SRR"
	StationCodeSaundersfoot                      CRSCode = "SDF"
	StationCodeSaunderton                        CRSCode = "SDR"
	StationCodeSawbridgeworth                    CRSCode = "SAW"
	StationCodeSaxilby                           CRSCode = "SXY"
	StationCodeSaxmundham                        CRSCode = "SAX"
	StationCodeScarborough                       CRSCode = "SCA"
	StationCodeScotscalder                       CRSCode = "SCT"
	StationCodeScotstounhill                     CRSCode = "SCH"
	StationCodeScunthorpe                        CRSCode = "SCU"
	StationCodeSeaMills                          CRSCode = "SML"
	StationCodeSeafordSussex                     CRSCode = "SEF"
	StationCodeSeaforthAndLitherland             CRSCode = "SFL"
	StationCodeSeaham                            CRSCode = "SEA"
	StationCodeSeamer                            CRSCode = "SEM"
	StationCodeSeascale                          CRSCode = "SSC"
	StationCodeSeatonCarew                       CRSCode = "SEC"
	StationCodeSeerGreenAndJordans               CRSCode = "SRG"
	StationCodeSelby                             CRSCode = "SBY"
	StationCodeSelhurst                          CRSCode = "SRS"
	StationCodeSellafield                        CRSCode = "SEL"
	StationCodeSelling                           CRSCode = "SEG"
	StationCodeSellyOak                          CRSCode = "SLY"
	StationCodeSettle                            CRSCode = "SET"
	StationCodeSevenKings                        CRSCode = "SVK"
	StationCodeSevenSisters                      CRSCode = "SVS"
	StationCodeSevenoaks                         CRSCode = "SEV"
	StationCodeSevernBeach                       CRSCode = "SVB"
	StationCodeSevernTunnelJunction              CRSCode = "STJ"
	StationCodeShadwell                          CRSCode = "SDE"
	StationCodeShalfordSurrey                    CRSCode = "SFR"
	StationCodeShanklin                          CRSCode = "SHN"
	StationCodeShawfair                          CRSCode = "SFI"
	StationCodeShawford                          CRSCode = "SHW"
	StationCodeShawlands                         CRSCode = "SHL"
	StationCodeSheernessonSea                    CRSCode = "SSS"
	StationCodeSheffield                         CRSCode = "SHF"
	StationCodeShelfordCambs                     CRSCode = "SED"
	StationCodeShenfield                         CRSCode = "SNF"
	StationCodeShenstone                         CRSCode = "SEN"
	StationCodeShepherdsBush                     CRSCode = "SPB"
	StationCodeShepherdsWell                     CRSCode = "SPH"
	StationCodeShepley                           CRSCode = "SPY"
	StationCodeShepperton                        CRSCode = "SHP"
	StationCodeShepreth                          CRSCode = "STH"
	StationCodeSherborne                         CRSCode = "SHE"
	StationCodeSherburninElmet                   CRSCode = "SIE"
	StationCodeSheringham                        CRSCode = "SHM"
	StationCodeShettleston                       CRSCode = "SLS"
	StationCodeShieldmuir                        CRSCode = "SDM"
	StationCodeShifnal                           CRSCode = "SFN"
	StationCodeShildon                           CRSCode = "SHD"
	StationCodeShiplake                          CRSCode = "SHI"
	StationCodeShipleyYorks                      CRSCode = "SHY"
	StationCodeShippeaHill                       CRSCode = "SPP"
	StationCodeShipton                           CRSCode = "SIP"
	StationCodeShirebrook                        CRSCode = "SHB"
	StationCodeShirehampton                      CRSCode = "SHH"
	StationCodeShireoaks                         CRSCode = "SRO"
	StationCodeShirley                           CRSCode = "SRL"
	StationCodeShoeburyness                      CRSCode = "SRY"
	StationCodeSholing                           CRSCode = "SHO"
	StationCodeShoreditchHighStreet              CRSCode = "SDC"
	StationCodeShorehamKent                      CRSCode = "SEH"
	StationCodeShorehambySea                     CRSCode = "SSE"
	StationCodeShortlands                        CRSCode = "SRT"
	StationCodeShotton                           CRSCode = "SHT"
	StationCodeShotts                            CRSCode = "SHS"
	StationCodeShrewsbury                        CRSCode = "SHR"
	StationCodeSidcup                            CRSCode = "SID"
	StationCodeSileby                            CRSCode = "SIL"
	StationCodeSilecroft                         CRSCode = "SIC"
	StationCodeSilkstoneCommon                   CRSCode = "SLK"
	StationCodeSilverStreet                      CRSCode = "SLV"
	StationCodeSilverdale                        CRSCode = "SVR"
	StationCodeSinger                            CRSCode = "SIN"
	StationCodeSittingbourne                     CRSCode = "SIT"
	StationCodeSkegness                          CRSCode = "SKG"
	StationCodeSkewen                            CRSCode = "SKE"
	StationCodeSkipton                           CRSCode = "SKI"
	StationCodeSladeGreen                        CRSCode = "SGR"
	StationCodeSlaithwaite                       CRSCode = "SWT"
	StationCodeSlateford                         CRSCode = "SLA"
	StationCodeSleaford                          CRSCode = "SLR"
	StationCodeSleights                          CRSCode = "SLH"
	StationCodeSlough                            CRSCode = "SLO"
	StationCodeSmallHeath                        CRSCode = "SMA"
	StationCodeSmallbrookJunction                CRSCode = "SAB"
	StationCodeSmethwickGaltonBridge             CRSCode = "SGB"
	StationCodeSmethwickRolfeStreet              CRSCode = "SMR"
	StationCodeSmithyBridge                      CRSCode = "SMB"
	StationCodeSnaith                            CRSCode = "SNI"
	StationCodeSnodland                          CRSCode = "SDA"
	StationCodeSnowdown                          CRSCode = "SWO"
	StationCodeSoleStreet                        CRSCode = "SOR"
	StationCodeSolihull                          CRSCode = "SOL"
	StationCodeSomerleyton                       CRSCode = "SYT"
	StationCodeSouthActon                        CRSCode = "SAT"
	StationCodeSouthBank                         CRSCode = "SBK"
	StationCodeSouthBermondsey                   CRSCode = "SBM"
	StationCodeSouthCroydon                      CRSCode = "SCY"
	StationCodeSouthElmsall                      CRSCode = "SES"
	StationCodeSouthGreenford                    CRSCode = "SGN"
	StationCodeSouthGyle                         CRSCode = "SGL"
	StationCodeSouthHampstead                    CRSCode = "SOH"
	StationCodeSouthKenton                       CRSCode = "SOK"
	StationCodeSouthMerton                       CRSCode = "SMO"
	StationCodeSouthMilford                      CRSCode = "SOM"
	StationCodeSouthRuislip                      CRSCode = "SRU"
	StationCodeSouthTottenham                    CRSCode = "STO"
	StationCodeSouthWigston                      CRSCode = "SWS"
	StationCodeSouthWoodhamFerrers               CRSCode = "SOF"
	StationCodeSouthall                          CRSCode = "STL"
	StationCodeSouthamptonAirportParkway         CRSCode = "SOA"
	StationCodeSouthamptonCentral                CRSCode = "SOU"
	StationCodeSouthbourne                       CRSCode = "SOB"
	StationCodeSouthbury                         CRSCode = "SBU"
	StationCodeSouthease                         CRSCode = "SEE"
	StationCodeSouthendAirport                   CRSCode = "SIA"
	StationCodeSouthendCentral                   CRSCode = "SOC"
	StationCodeSouthendEast                      CRSCode = "SOE"
	StationCodeSouthendVictoria                  CRSCode = "SOV"
	StationCodeSouthminster                      CRSCode = "SMN"
	StationCodeSouthport                         CRSCode = "SOP"
	StationCodeSouthwick                         CRSCode = "SWK"
	StationCodeSowerbyBridge                     CRSCode = "SOW"
	StationCodeSpalding                          CRSCode = "SPA"
	StationCodeSpeanBridge                       CRSCode = "SBR"
	StationCodeSpital                            CRSCode = "SPI"
	StationCodeSpondon                           CRSCode = "SPO"
	StationCodeSpoonerRow                        CRSCode = "SPN"
	StationCodeSpringRoad                        CRSCode = "SRI"
	StationCodeSpringburn                        CRSCode = "SPR"
	StationCodeSpringfield                       CRSCode = "SPF"
	StationCodeSquiresGate                       CRSCode = "SQU"
	StationCodeStAlbansAbbey                     CRSCode = "SAA"
	StationCodeStAlbansCity                      CRSCode = "SAC"
	StationCodeStAndrewsRoad                     CRSCode = "SAR"
	StationCodeStAnnesonSea                      CRSCode = "SAS"
	StationCodeStAustell                         CRSCode = "SAU"
	StationCodeStBees                            CRSCode = "SBS"
	StationCodeStBudeauxFerryRoad                CRSCode = "SBF"
	StationCodeStBudeauxVictoriaRoad             CRSCode = "SBV"
	StationCodeStColumbRoad                      CRSCode = "SCR"
	StationCodeStDenys                           CRSCode = "SDN"
	StationCodeStErth                            CRSCode = "SER"
	StationCodeStGermans                         CRSCode = "SGM"
	StationCodeStHelensCentral                   CRSCode = "SNH"
	StationCodeStHelensJunction                  CRSCode = "SHJ"
	StationCodeStHelierSurrey                    CRSCode = "SIH"
	StationCodeStIvesCornwall                    CRSCode = "SIV"
	StationCodeStJamesParkExeter                 CRSCode = "SJP"
	StationCodeStJamesStreetWalthamstow          CRSCode = "SJS"
	StationCodeStJohnsLondon                     CRSCode = "SAJ"
	StationCodeStKeyneWishingWellHalt            CRSCode = "SKN"
	StationCodeStLeonardsWarriorSquare           CRSCode = "SLQ"
	StationCodeStMargaretsHerts                  CRSCode = "SMT"
	StationCodeStMargaretsLondon                 CRSCode = "SMG"
	StationCodeStMaryCray                        CRSCode = "SMY"
	StationCodeStMichaels                        CRSCode = "STM"
	StationCodeStNeots                           CRSCode = "SNO"
	StationCodeStafford                          CRSCode = "STA"
	StationCodeStaines                           CRSCode = "SNS"
	StationCodeStallingborough                   CRSCode = "SLL"
	StationCodeStalybridge                       CRSCode = "SYB"
	StationCodeStamfordHill                      CRSCode = "SMH"
	StationCodeStamfordLincs                     CRSCode = "SMD"
	StationCodeStanfordleHope                    CRSCode = "SFO"
	StationCodeStanlowAndThornton                CRSCode = "SNT"
	StationCodeStanstedAirport                   CRSCode = "SSD"
	StationCodeStanstedMountfitchet              CRSCode = "SST"
	StationCodeStaplehurst                       CRSCode = "SPU"
	StationCodeStapletonRoad                     CRSCode = "SRD"
	StationCodeStarbeck                          CRSCode = "SBE"
	StationCodeStarcross                         CRSCode = "SCS"
	StationCodeStaveleyCumbria                   CRSCode = "SVL"
	StationCodeStechford                         CRSCode = "SCF"
	StationCodeSteetonAndSilsden                 CRSCode = "SON"
	StationCodeStepps                            CRSCode = "SPS"
	StationCodeStevenage                         CRSCode = "SVG"
	StationCodeStevenston                        CRSCode = "STV"
	StationCodeStewartby                         CRSCode = "SWR"
	StationCodeStewarton                         CRSCode = "STT"
	StationCodeStirling                          CRSCode = "STG"
	StationCodeStockport                         CRSCode = "SPT"
	StationCodeStocksfield                       CRSCode = "SKS"
	StationCodeStocksmoor                        CRSCode = "SSM"
	StationCodeStockton                          CRSCode = "STK"
	StationCodeStokeMandeville                   CRSCode = "SKM"
	StationCodeStokeNewington                    CRSCode = "SKW"
	StationCodeStokeonTrent                      CRSCode = "SOT"
	StationCodeStoneCrossing                     CRSCode = "SCG"
	StationCodeStoneStaffs                       CRSCode = "SNE"
	StationCodeStonebridgePark                   CRSCode = "SBP"
	StationCodeStonegate                         CRSCode = "SOG"
	StationCodeStonehaven                        CRSCode = "STN"
	StationCodeStonehouse                        CRSCode = "SHU"
	StationCodeStoneleigh                        CRSCode = "SNL"
	StationCodeStourbridgeJunction               CRSCode = "SBJ"
	StationCodeStourbridgeTown                   CRSCode = "SBT"
	StationCodeStow                              CRSCode = "SOI"
	StationCodeStowmarket                        CRSCode = "SMK"
	StationCodeStranraer                         CRSCode = "STR"
	StationCodeStratfordInternational            CRSCode = "SFA"
	StationCodeStratfordLondon                   CRSCode = "SRA"
	StationCodeStratfordUponAvon                 CRSCode = "SAV"
	StationCodeStratfordUponAvonParkway          CRSCode = "STY"
	StationCodeStrathcarron                      CRSCode = "STC"
	StationCodeStrawberryHill                    CRSCode = "STW"
	StationCodeStreathamCommon                   CRSCode = "SRC"
	StationCodeStreathamGreaterLondon            CRSCode = "STE"
	StationCodeStreathamHill                     CRSCode = "SRH"
	StationCodeStreethouse                       CRSCode = "SHC"
	StationCodeStrines                           CRSCode = "SRN"
	StationCodeStromeferry                       CRSCode = "STF"
	StationCodeStrood                            CRSCode = "SOO"
	StationCodeStroudGloucs                      CRSCode = "STD"
	StationCodeSturry                            CRSCode = "STU"
	StationCodeStyal                             CRSCode = "SYA"
	StationCodeSudburyAndHarrowRoad              CRSCode = "SUD"
	StationCodeSudburyHillHarrow                 CRSCode = "SDH"
	StationCodeSudburySuffolk                    CRSCode = "SUY"
	StationCodeSugarLoaf                         CRSCode = "SUG"
	StationCodeSummerston                        CRSCode = "SUM"
	StationCodeSunbury                           CRSCode = "SUU"
	StationCodeSunderland                        CRSCode = "SUN"
	StationCodeSundridgePark                     CRSCode = "SUP"
	StationCodeSunningdale                       CRSCode = "SNG"
	StationCodeSunnymeads                        CRSCode = "SNY"
	StationCodeSurbiton                          CRSCode = "SUR"
	StationCodeSurreyQuays                       CRSCode = "SQE"
	StationCodeSuttonColdfield                   CRSCode = "SUT"
	StationCodeSuttonCommon                      CRSCode = "SUC"
	StationCodeSuttonParkway                     CRSCode = "SPK"
	StationCodeSuttonSurrey                      CRSCode = "SUO"
	StationCodeSwale                             CRSCode = "SWL"
	StationCodeSwanley                           CRSCode = "SAY"
	StationCodeSwanscombe                        CRSCode = "SWM"
	StationCodeSwansea                           CRSCode = "SWA"
	StationCodeSwanwick                          CRSCode = "SNW"
	StationCodeSway                              CRSCode = "SWY"
	StationCodeSwaythling                        CRSCode = "SWG"
	StationCodeSwinderby                         CRSCode = "SWD"
	StationCodeSwindonWilts                      CRSCode = "SWI"
	StationCodeSwineshead                        CRSCode = "SWE"
	StationCodeSwintonManchester                 CRSCode = "SNN"
	StationCodeSwintonSouthYorks                 CRSCode = "SWN"
	StationCodeSydenhamHill                      CRSCode = "SYH"
	StationCodeSydenhamLondon                    CRSCode = "SYD"
	StationCodeSyonLane                          CRSCode = "SYL"
	StationCodeSyston                            CRSCode = "SYS"
	StationCodeTackley                           CRSCode = "TAC"
	StationCodeTadworth                          CRSCode = "TAD"
	StationCodeTaffsWell                         CRSCode = "TAF"
	StationCodeTain                              CRSCode = "TAI"
	StationCodeTalsarnau                         CRSCode = "TAL"
	StationCodeTalyCafn                          CRSCode = "TLC"
	StationCodeTalybont                          CRSCode = "TLB"
	StationCodeTameBridgeParkway                 CRSCode = "TAB"
	StationCodeTamworth                          CRSCode = "TAM"
	StationCodeTaplow                            CRSCode = "TAP"
	StationCodeTattenhamCorner                   CRSCode = "TAT"
	StationCodeTaunton                           CRSCode = "TAU"
	StationCodeTaynuilt                          CRSCode = "TAY"
	StationCodeTeddington                        CRSCode = "TED"
	StationCodeTeessideAirport                   CRSCode = "TEA"
	StationCodeTeignmouth                        CRSCode = "TGM"
	StationCodeTelfordCentral                    CRSCode = "TFC"
	StationCodeTemplecombe                       CRSCode = "TMC"
	StationCodeTenby                             CRSCode = "TEN"
	StationCodeTeynham                           CRSCode = "TEY"
	StationCodeThamesDitton                      CRSCode = "THD"
	StationCodeThatcham                          CRSCode = "THA"
	StationCodeThattoHeath                       CRSCode = "THH"
	StationCodeTheHawthorns                      CRSCode = "THW"
	StationCodeTheLakesWarks                     CRSCode = "TLK"
	StationCodeTheale                            CRSCode = "THE"
	StationCodeTheobaldsGrove                    CRSCode = "TEO"
	StationCodeThetford                          CRSCode = "TTF"
	StationCodeThirsk                            CRSCode = "THI"
	StationCodeThornaby                          CRSCode = "TBY"
	StationCodeThorneNorth                       CRSCode = "TNN"
	StationCodeThorneSouth                       CRSCode = "TNS"
	StationCodeThornford                         CRSCode = "THO"
	StationCodeThornliebank                      CRSCode = "THB"
	StationCodeThorntonAbbey                     CRSCode = "TNA"
	StationCodeThorntonHeath                     CRSCode = "TTH"
	StationCodeThorntonhall                      CRSCode = "THT"
	StationCodeThorpeBay                         CRSCode = "TPB"
	StationCodeThorpeCulvert                     CRSCode = "TPC"
	StationCodeThorpeleSoken                     CRSCode = "TLS"
	StationCodeThreeBridges                      CRSCode = "TBD"
	StationCodeThreeOaks                         CRSCode = "TOK"
	StationCodeThurgarton                        CRSCode = "THU"
	StationCodeThurnscoe                         CRSCode = "THC"
	StationCodeThurso                            CRSCode = "THS"
	StationCodeThurston                          CRSCode = "TRS"
	StationCodeTilburyTown                       CRSCode = "TIL"
	StationCodeTileHill                          CRSCode = "THL"
	StationCodeTilehurst                         CRSCode = "TLH"
	StationCodeTipton                            CRSCode = "TIP"
	StationCodeTirPhil                           CRSCode = "TIR"
	StationCodeTisbury                           CRSCode = "TIS"
	StationCodeTivertonParkway                   CRSCode = "TVP"
	StationCodeTodmorden                         CRSCode = "TOD"
	StationCodeTolworth                          CRSCode = "TOL"
	StationCodeTonPentre                         CRSCode = "TPN"
	StationCodeTonbridge                         CRSCode = "TON"
	StationCodeTondu                             CRSCode = "TDU"
	StationCodeTonfanau                          CRSCode = "TNF"
	StationCodeTonypandy                         CRSCode = "TNP"
	StationCodeTooting                           CRSCode = "TOO"
	StationCodeTopsham                           CRSCode = "TOP"
	StationCodeTorquay                           CRSCode = "TQY"
	StationCodeTorre                             CRSCode = "TRR"
	StationCodeTotnes                            CRSCode = "TOT"
	StationCodeTottenhamHale                     CRSCode = "TOM"
	StationCodeTotton                            CRSCode = "TTN"
	StationCodeTownGreen                         CRSCode = "TWN"
	StationCodeTraffordPark                      CRSCode = "TRA"
	StationCodeTrefforest                        CRSCode = "TRF"
	StationCodeTrefforestEstate                  CRSCode = "TRE"
	StationCodeTrehafod                          CRSCode = "TRH"
	StationCodeTreherbert                        CRSCode = "TRB"
	StationCodeTreorchy                          CRSCode = "TRY"
	StationCodeTrimley                           CRSCode = "TRM"
	StationCodeTring                             CRSCode = "TRI"
	StationCodeTroedyrhiw                        CRSCode = "TRD"
	StationCodeTroon                             CRSCode = "TRN"
	StationCodeTrowbridge                        CRSCode = "TRO"
	StationCodeTruro                             CRSCode = "TRU"
	StationCodeTulloch                           CRSCode = "TUL"
	StationCodeTulseHill                         CRSCode = "TUH"
	StationCodeTunbridgeWells                    CRSCode = "TBW"
	StationCodeTurkeyStreet                      CRSCode = "TUR"
	StationCodeTutburyAndHatton                  CRSCode = "TUT"
	StationCodeTweedbank                         CRSCode = "TWB"
	StationCodeTwickenham                        CRSCode = "TWI"
	StationCodeTwyford                           CRSCode = "TWY"
	StationCodeTyCroes                           CRSCode = "TYC"
	StationCodeTyGlas                            CRSCode = "TGS"
	StationCodeTygwyn                            CRSCode = "TYG"
	StationCodeTyndrumLower                      CRSCode = "TYL"
	StationCodeTyseley                           CRSCode = "TYS"
	StationCodeTywyn                             CRSCode = "TYW"
	StationCodeUckfield                          CRSCode = "UCK"
	StationCodeUddingston                        CRSCode = "UDD"
	StationCodeUlceby                            CRSCode = "ULC"
	StationCodeUlleskelf                         CRSCode = "ULL"
	StationCodeUlverston                         CRSCode = "ULV"
	StationCodeUmberleigh                        CRSCode = "UMB"
	StationCodeUniversityBirmingham              CRSCode = "UNI"
	StationCodeUphall                            CRSCode = "UHA"
	StationCodeUpholland                         CRSCode = "UPL"
	StationCodeUpminster                         CRSCode = "UPM"
	StationCodeUpperHalliford                    CRSCode = "UPH"
	StationCodeUpperHolloway                     CRSCode = "UHL"
	StationCodeUpperTyndrum                      CRSCode = "UTY"
	StationCodeUpperWarlingham                   CRSCode = "UWL"
	StationCodeUptonMerseyside                   CRSCode = "UPT"
	StationCodeUpwey                             CRSCode = "UPW"
	StationCodeUrmston                           CRSCode = "URM"
	StationCodeUttoxeter                         CRSCode = "UTT"
	StationCodeValley                            CRSCode = "VAL"
	StationCodeVauxhall                          CRSCode = "VXH"
	StationCodeVirginiaWater                     CRSCode = "VIR"
	StationCodeWaddon                            CRSCode = "WDO"
	StationCodeWadhurst                          CRSCode = "WAD"
	StationCodeWainfleet                         CRSCode = "WFL"
	StationCodeWakefieldKirkgate                 CRSCode = "WKK"
	StationCodeWakefieldWestgate                 CRSCode = "WKF"
	StationCodeWalkden                           CRSCode = "WKD"
	StationCodeWallaseyGroveRoad                 CRSCode = "WLG"
	StationCodeWallaseyVillage                   CRSCode = "WLV"
	StationCodeWallington                        CRSCode = "WLT"
	StationCodeWallyford                         CRSCode = "WAF"
	StationCodeWalmer                            CRSCode = "WAM"
	StationCodeWalsall                           CRSCode = "WSL"
	StationCodeWalsden                           CRSCode = "WDN"
	StationCodeWalthamCross                      CRSCode = "WLC"
	StationCodeWalthamstowCentral                CRSCode = "WHC"
	StationCodeWalthamstowQueensRoad             CRSCode = "WMW"
	StationCodeWaltonMerseyside                  CRSCode = "WAO"
	StationCodeWaltononThames                    CRSCode = "WAL"
	StationCodeWaltonontheNaze                   CRSCode = "WON"
	StationCodeWanborough                        CRSCode = "WAN"
	StationCodeWandsworthCommon                  CRSCode = "WSW"
	StationCodeWandsworthRoad                    CRSCode = "WWR"
	StationCodeWandsworthTown                    CRSCode = "WNT"
	StationCodeWansteadPark                      CRSCode = "WNP"
	StationCodeWapping                           CRSCode = "WPE"
	StationCodeWarblington                       CRSCode = "WBL"
	StationCodeWareHerts                         CRSCode = "WAR"
	StationCodeWarehamDorset                     CRSCode = "WRM"
	StationCodeWargrave                          CRSCode = "WGV"
	StationCodeWarminster                        CRSCode = "WMN"
	StationCodeWarnham                           CRSCode = "WNH"
	StationCodeWarringtonBankQuay                CRSCode = "WBQ"
	StationCodeWarringtonCentral                 CRSCode = "WAC"
	StationCodeWarwick                           CRSCode = "WRW"
	StationCodeWarwickParkway                    CRSCode = "WRP"
	StationCodeWaterOrton                        CRSCode = "WTO"
	StationCodeWaterbeach                        CRSCode = "WBC"
	StationCodeWateringbury                      CRSCode = "WTR"
	StationCodeWaterlooMerseyside                CRSCode = "WLO"
	StationCodeWatfordHighStreet                 CRSCode = "WFH"
	StationCodeWatfordJunction                   CRSCode = "WFJ"
	StationCodeWatfordNorth                      CRSCode = "WFN"
	StationCodeWatlington                        CRSCode = "WTG"
	StationCodeWattonatStone                     CRSCode = "WAS"
	StationCodeWaunGronPark                      CRSCode = "WNG"
	StationCodeWavertreeTechnologyPark           CRSCode = "WAV"
	StationCodeWedgwood                          CRSCode = "WED"
	StationCodeWeeley                            CRSCode = "WEE"
	StationCodeWeeton                            CRSCode = "WET"
	StationCodeWelhamGreen                       CRSCode = "WMG"
	StationCodeWelling                           CRSCode = "WLI"
	StationCodeWellingborough                    CRSCode = "WEL"
	StationCodeWellingtonShropshire              CRSCode = "WLN"
	StationCodeWelshpool                         CRSCode = "WLP"
	StationCodeWelwynGardenCity                  CRSCode = "WGC"
	StationCodeWelwynNorth                       CRSCode = "WLW"
	StationCodeWem                               CRSCode = "WEM"
	StationCodeWembleyCentral                    CRSCode = "WMB"
	StationCodeWembleyStadium                    CRSCode = "WCX"
	StationCodeWemyssBay                         CRSCode = "WMS"
	StationCodeWendover                          CRSCode = "WND"
	StationCodeWennington                        CRSCode = "WNN"
	StationCodeWestAllerton                      CRSCode = "WSA"
	StationCodeWestBrompton                      CRSCode = "WBP"
	StationCodeWestByfleet                       CRSCode = "WBY"
	StationCodeWestCalder                        CRSCode = "WCL"
	StationCodeWestCroydon                       CRSCode = "WCY"
	StationCodeWestDrayton                       CRSCode = "WDT"
	StationCodeWestDulwich                       CRSCode = "WDU"
	StationCodeWestEaling                        CRSCode = "WEA"
	StationCodeWestHam                           CRSCode = "WEH"
	StationCodeWestHampstead                     CRSCode = "WHD"
	StationCodeWestHampsteadThameslink           CRSCode = "WHP"
	StationCodeWestHorndon                       CRSCode = "WHR"
	StationCodeWestKilbride                      CRSCode = "WKB"
	StationCodeWestKirby                         CRSCode = "WKI"
	StationCodeWestMalling                       CRSCode = "WMA"
	StationCodeWestNorwood                       CRSCode = "WNW"
	StationCodeWestRuislip                       CRSCode = "WRU"
	StationCodeWestRunton                        CRSCode = "WRN"
	StationCodeWestStLeonards                    CRSCode = "WLD"
	StationCodeWestSutton                        CRSCode = "WSU"
	StationCodeWestWickham                       CRSCode = "WWI"
	StationCodeWestWorthing                      CRSCode = "WWO"
	StationCodeWestburyWilts                     CRSCode = "WSB"
	StationCodeWestcliff                         CRSCode = "WCF"
	StationCodeWestcombePark                     CRSCode = "WCB"
	StationCodeWestenhanger                      CRSCode = "WHA"
	StationCodeWesterHailes                      CRSCode = "WTA"
	StationCodeWesterfield                       CRSCode = "WFI"
	StationCodeWesterton                         CRSCode = "WES"
	StationCodeWestgateOnSea                     CRSCode = "WGA"
	StationCodeWesthoughton                      CRSCode = "WHG"
	StationCodeWestonMilton                      CRSCode = "WNM"
	StationCodeWestonsuperMare                   CRSCode = "WSM"
	StationCodeWetheral                          CRSCode = "WRL"
	StationCodeWeybridge                         CRSCode = "WYB"
	StationCodeWeymouth                          CRSCode = "WEY"
	StationCodeWhaleyBridge                      CRSCode = "WBR"
	StationCodeWhalleyLancs                      CRSCode = "WHE"
	StationCodeWhatstandwell                     CRSCode = "WTS"
	StationCodeWhifflet                          CRSCode = "WFF"
	StationCodeWhimple                           CRSCode = "WHM"
	StationCodeWhinhill                          CRSCode = "WNL"
	StationCodeWhiston                           CRSCode = "WHN"
	StationCodeWhitby                            CRSCode = "WTB"
	StationCodeWhitchurchCardiff                 CRSCode = "WHT"
	StationCodeWhitchurchHants                   CRSCode = "WCH"
	StationCodeWhitchurchShropshire              CRSCode = "WTC"
	StationCodeWhiteHartLane                     CRSCode = "WHL"
	StationCodeWhiteNotley                       CRSCode = "WNY"
	StationCodeWhitechapel                       CRSCode = "ZLW"
	StationCodeWhitecraigs                       CRSCode = "WCR"
	StationCodeWhitehaven                        CRSCode = "WTH"
	StationCodeWhitland                          CRSCode = "WTL"
	StationCodeWhitleyBridge                     CRSCode = "WBD"
	StationCodeWhitlocksEnd                      CRSCode = "WTE"
	StationCodeWhitstable                        CRSCode = "WHI"
	StationCodeWhittlesea                        CRSCode = "WLE"
	StationCodeWhittlesfordParkway               CRSCode = "WLF"
	StationCodeWhittonLondon                     CRSCode = "WTN"
	StationCodeWhitwellDerbyshire                CRSCode = "WWL"
	StationCodeWhyteleafe                        CRSCode = "WHY"
	StationCodeWhyteleafeSouth                   CRSCode = "WHS"
	StationCodeWick                              CRSCode = "WCK"
	StationCodeWickford                          CRSCode = "WIC"
	StationCodeWickhamMarket                     CRSCode = "WCM"
	StationCodeWiddrington                       CRSCode = "WDD"
	StationCodeWidnes                            CRSCode = "WID"
	StationCodeWidneyManor                       CRSCode = "WMR"
	StationCodeWiganNorthWestern                 CRSCode = "WGN"
	StationCodeWiganWallgate                     CRSCode = "WGW"
	StationCodeWigton                            CRSCode = "WGT"
	StationCodeWildmill                          CRSCode = "WMI"
	StationCodeWillesdenJunction                 CRSCode = "WIJ"
	StationCodeWilliamwood                       CRSCode = "WLM"
	StationCodeWillington                        CRSCode = "WIL"
	StationCodeWilmcote                          CRSCode = "WMC"
	StationCodeWilmslow                          CRSCode = "WML"
	StationCodeWilnecoteStaffs                   CRSCode = "WNE"
	StationCodeWimbledon                         CRSCode = "WIM"
	StationCodeWimbledonChase                    CRSCode = "WBO"
	StationCodeWinchelsea                        CRSCode = "WSE"
	StationCodeWinchester                        CRSCode = "WIN"
	StationCodeWinchfield                        CRSCode = "WNF"
	StationCodeWinchmoreHill                     CRSCode = "WIH"
	StationCodeWindermere                        CRSCode = "WDM"
	StationCodeWindsorAndEtonCentral             CRSCode = "WNC"
	StationCodeWindsorAndEtonRiverside           CRSCode = "WNR"
	StationCodeWinnersh                          CRSCode = "WNS"
	StationCodeWinnershTriangle                  CRSCode = "WTI"
	StationCodeWinsford                          CRSCode = "WSF"
	StationCodeWishaw                            CRSCode = "WSH"
	StationCodeWitham                            CRSCode = "WTM"
	StationCodeWitley                            CRSCode = "WTY"
	StationCodeWittonWestMidlands                CRSCode = "WTT"
	StationCodeWivelsfield                       CRSCode = "WVF"
	StationCodeWivenhoe                          CRSCode = "WIV"
	StationCodeWoburnSands                       CRSCode = "WOB"
	StationCodeWoking                            CRSCode = "WOK"
	StationCodeWokingham                         CRSCode = "WKM"
	StationCodeWoldingham                        CRSCode = "WOH"
	StationCodeWolverhampton                     CRSCode = "WVH"
	StationCodeWolverton                         CRSCode = "WOL"
	StationCodeWombwell                          CRSCode = "WOM"
	StationCodeWoodEnd                           CRSCode = "WDE"
	StationCodeWoodStreet                        CRSCode = "WST"
	StationCodeWoodbridge                        CRSCode = "WDB"
	StationCodeWoodgrangePark                    CRSCode = "WGR"
	StationCodeWoodhall                          CRSCode = "WDL"
	StationCodeWoodhouse                         CRSCode = "WDH"
	StationCodeWoodlesford                       CRSCode = "WDS"
	StationCodeWoodley                           CRSCode = "WLY"
	StationCodeWoodmansterne                     CRSCode = "WME"
	StationCodeWoodsmoor                         CRSCode = "WSR"
	StationCodeWool                              CRSCode = "WOO"
	StationCodeWoolston                          CRSCode = "WLS"
	StationCodeWoolwichArsenal                   CRSCode = "WWA"
	StationCodeWoolwichDockyard                  CRSCode = "WWD"
	StationCodeWoottonWawen                      CRSCode = "WWW"
	StationCodeWorcesterForegateStreet           CRSCode = "WOF"
	StationCodeWorcesterPark                     CRSCode = "WCP"
	StationCodeWorcesterShrubHill                CRSCode = "WOS"
	StationCodeWorkington                        CRSCode = "WKG"
	StationCodeWorksop                           CRSCode = "WRK"
	StationCodeWorle                             CRSCode = "WOR"
	StationCodeWorplesdon                        CRSCode = "WPL"
	StationCodeWorstead                          CRSCode = "WRT"
	StationCodeWorthing                          CRSCode = "WRH"
	StationCodeWrabness                          CRSCode = "WRB"
	StationCodeWraysbury                         CRSCode = "WRY"
	StationCodeWrenbury                          CRSCode = "WRE"
	StationCodeWressle                           CRSCode = "WRS"
	StationCodeWrexhamCentral                    CRSCode = "WXC"
	StationCodeWrexhamGeneral                    CRSCode = "WRX"
	StationCodeWye                               CRSCode = "WYE"
	StationCodeWylam                             CRSCode = "WYM"
	StationCodeWyldeGreen                        CRSCode = "WYL"
	StationCodeWymondham                         CRSCode = "WMD"
	StationCodeWythall                           CRSCode = "WYT"
	StationCodeYalding                           CRSCode = "YAL"
	StationCodeYardleyWood                       CRSCode = "YRD"
	StationCodeYarm                              CRSCode = "YRM"
	StationCodeYate                              CRSCode = "YAE"
	StationCodeYatton                            CRSCode = "YAT"
	StationCodeYeoford                           CRSCode = "YEO"
	StationCodeYeovilJunction                    CRSCode = "YVJ"
	StationCodeYeovilPenMill                     CRSCode = "YVP"
	StationCodeYetminster                        CRSCode = "YET"
	StationCodeYnyswen                           CRSCode = "YNW"
	StationCodeYoker                             CRSCode = "YOK"
	StationCodeYork                              CRSCode = "YRK"
	StationCodeYorton                            CRSCode = "YRT"
	StationCodeYstradMynach                      CRSCode = "YSM"
	StationCodeYstradRhondda                     CRSCode = "YSR"
)

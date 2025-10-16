<map version="freeplane 1.7.0">
<!--To view this file, download free mind mapping software Freeplane from http://freeplane.sourceforge.net -->
<node TEXT="Parsers Intro" LOCALIZED_STYLE_REF="default" FOLDED="false" ID="ID_1359429857" CREATED="1346072531070" MODIFIED="1760614531521"><hook NAME="MapStyle" zoom="1.5">
    <properties show_icon_for_attributes="true" show_note_icons="true" fit_to_viewport="false"/>

<map_styles>
<stylenode LOCALIZED_TEXT="styles.root_node" STYLE="oval" UNIFORM_SHAPE="true" VGAP_QUANTITY="24.0 pt">
<font SIZE="24"/>
<stylenode LOCALIZED_TEXT="styles.predefined" POSITION="right" STYLE="bubble">
<stylenode LOCALIZED_TEXT="default" MAX_WIDTH="600.0 px" COLOR="#000000" STYLE="as_parent">
<font NAME="SansSerif" SIZE="10" BOLD="false" ITALIC="false"/>
</stylenode>
<stylenode LOCALIZED_TEXT="defaultstyle.details"/>
<stylenode LOCALIZED_TEXT="defaultstyle.attributes">
<font SIZE="9"/>
</stylenode>
<stylenode LOCALIZED_TEXT="defaultstyle.note"/>
<stylenode LOCALIZED_TEXT="defaultstyle.floating">
<edge STYLE="hide_edge"/>
<cloud COLOR="#f0f0f0" SHAPE="ROUND_RECT"/>
</stylenode>
</stylenode>
<stylenode LOCALIZED_TEXT="styles.user-defined" POSITION="right" STYLE="bubble">
<stylenode LOCALIZED_TEXT="styles.topic" COLOR="#18898b" STYLE="fork">
<font NAME="Liberation Sans" SIZE="10" BOLD="true"/>
</stylenode>
<stylenode LOCALIZED_TEXT="styles.subtopic" COLOR="#cc3300" STYLE="fork">
<font NAME="Liberation Sans" SIZE="10" BOLD="true"/>
</stylenode>
<stylenode LOCALIZED_TEXT="styles.subsubtopic" COLOR="#669900">
<font NAME="Liberation Sans" SIZE="10" BOLD="true"/>
</stylenode>
<stylenode LOCALIZED_TEXT="styles.important">
<icon BUILTIN="yes"/>
</stylenode>
<stylenode TEXT="lsyh" COLOR="#990000">
<font NAME="SansSerif" SIZE="10" BOLD="true"/>
<edge COLOR="#808080"/>
</stylenode>
<stylenode TEXT="home" COLOR="#215800">
<font SIZE="10" BOLD="true"/>
</stylenode>
<stylenode TEXT="ohome" COLOR="#990033">
<font NAME="SansSerif" SIZE="10" BOLD="true"/>
</stylenode>
<stylenode TEXT="code" COLOR="#3631f6">
<font NAME="Monospaced" SIZE="10" BOLD="true" STRIKETHROUGH="false"/>
<edge COLOR="#808080"/>
</stylenode>
</stylenode>
<stylenode LOCALIZED_TEXT="styles.AutomaticLayout" POSITION="right" STYLE="bubble">
<stylenode LOCALIZED_TEXT="AutomaticLayout.level.root" COLOR="#000000">
<font SIZE="18"/>
</stylenode>
<stylenode LOCALIZED_TEXT="AutomaticLayout.level,1" COLOR="#0033ff">
<font NAME="Monospaced" BOLD="true"/>
</stylenode>
<stylenode LOCALIZED_TEXT="AutomaticLayout.level,2" COLOR="#00b439">
<font BOLD="true"/>
</stylenode>
<stylenode LOCALIZED_TEXT="AutomaticLayout.level,3" COLOR="#990000">
<font BOLD="true"/>
</stylenode>
<stylenode LOCALIZED_TEXT="AutomaticLayout.level,4" COLOR="#111111">
<font NAME="SansSerif" SIZE="10"/>
</stylenode>
</stylenode>
</stylenode>
</map_styles>
</hook>
<node TEXT="formula feature demo" POSITION="right" ID="ID_553710072" CREATED="1497525630303" MODIFIED="1760624883135">
<node TEXT="A feladat" FOLDED="true" ID="ID_898072602" CREATED="1760614533573" MODIFIED="1760614535555">
<node TEXT="formula editor &#xe9;s runner" ID="ID_1761447870" CREATED="1760614580597" MODIFIED="1760614593672"/>
<node TEXT="formul&#xe1;k:" ID="ID_1948781331" CREATED="1760614594269" MODIFIED="1760614600679">
<node TEXT="pl.:" ID="ID_1966468266" CREATED="1760614848243" MODIFIED="1760614850791">
<node TEXT="((10. * 2) + ((50.0 / pi) - 2.))" STYLE_REF="code" ID="ID_1682660564" CREATED="1760614953577" MODIFIED="1760615900572"/>
</node>
<node TEXT="addit&#xed;v (+-) &#xe9;s multiplikat&#xed;v (*/) m&#x171;veletek, a precedencia figyelembev&#xe9;tel&#xe9;vel" ID="ID_943564820" CREATED="1760614600924" MODIFIED="1760614642536"/>
<node TEXT="k&#xe9;ttag&#xfa;, &#xe9;s t&#xf6;bbtag&#xfa; m&#x171;veletek" ID="ID_1425069901" CREATED="1760614643052" MODIFIED="1760614652807"/>
<node TEXT="z&#xe1;r&#xf3;jeles kifejez&#xe9;sek" ID="ID_1785817749" CREATED="1760614653011" MODIFIED="1760614664095"/>
<node TEXT="kifejez&#xe9;sek egym&#xe1;sba &#xe1;gyazhat&#xf3;s&#xe1;ga, tetsz&#x151;leges m&#xe9;rt&#xe9;kig" ID="ID_1472693250" CREATED="1760614664347" MODIFIED="1760614674807"/>
<node TEXT="Go programba &#xe1;gyazva fusson mind a h&#xe1;rom (parser, compile, run) f&#xe1;zis" ID="ID_1745059281" CREATED="1760615203491" MODIFIED="1760620307055"/>
<node TEXT="hozz&#xe1;f&#xe9;r&#xe9;s, a futtat&#xf3; k&#xf6;rnyezetben defini&#xe1;lt v&#xe1;ltoz&#xf3;k &#xe9;s konstansok &#xe9;rt&#xe9;keihez, szimb&#xf3;likus nevekkel" ID="ID_712994830" CREATED="1760614675044" MODIFIED="1760614716808"/>
<node TEXT="a futtat&#xf3; k&#xf6;rnyezetben defini&#xe1;lt f&#xfc;ggv&#xe9;nyek h&#xed;v&#xe1;sa" ID="ID_1292260496" CREATED="1760614717076" MODIFIED="1760614735558"/>
<node TEXT="eredm&#xe9;ny &#xe1;tad&#xe1;sa a h&#xed;v&#xf3; programnak" ID="ID_301089581" CREATED="1760614735803" MODIFIED="1760614747040"/>
<node TEXT="elemz&#xe9;s (parse), ford&#xed;t&#xe1;s (compile) &#xe9;s futtat&#xe1;s (run) f&#xfc;ggetlenek egym&#xe1;st&#xf3;l" ID="ID_922310338" CREATED="1760614747220" MODIFIED="1760614799039"/>
<node TEXT="Az elemz&#xe9;s eset&#xe9;ben, felhaszn&#xe1;l&#xf3;bar&#xe1;t hiba&#xfc;zenetek szolg&#xe1;ltat&#xe1;sa, szintaxis hiba eset&#xe9;n" ID="ID_1024305004" CREATED="1760614820723" MODIFIED="1760614841207"/>
<node TEXT="A futtat&#xe1;s eset&#xe9;ben nagy performancia" ID="ID_1808429199" CREATED="1760614799268" MODIFIED="1760614819695"/>
</node>
</node>
<node TEXT="Lehets&#xe9;ges megold&#xe1;sok" FOLDED="true" ID="ID_481590454" CREATED="1760614536516" MODIFIED="1760615031482">
<node TEXT="Be&#xe1;gyazott interpreter, pl.:&#xa;JavaScript, Lua" ID="ID_499813817" CREATED="1760615031690" MODIFIED="1760615078390">
<node TEXT="" ID="ID_1579940693" CREATED="1760618445073" MODIFIED="1760618445073">
<node TEXT="el&#x151;ny&#xf6;k:" ID="ID_823345653" CREATED="1760618140480" MODIFIED="1760618143595">
<node TEXT="off-the-shelf" ID="ID_1594236249" CREATED="1760618143799" MODIFIED="1760618151113"/>
<node TEXT="lehet el&#xe9;g gyors, de nem minden esetben" ID="ID_1917438475" CREATED="1760618151527" MODIFIED="1760618165509"/>
</node>
<node TEXT="h&#xe1;tr&#xe1;nyok" ID="ID_677868862" CREATED="1760618166783" MODIFIED="1760618169584">
<node TEXT="hiba&#xfc;zenetek gener&#xe1;l&#xe1;sa" ID="ID_1394942650" CREATED="1760618169815" MODIFIED="1760618181455"/>
<node TEXT="egy komplett programoz&#xe1;si nyelvet tud/kell haszn&#xe1;lni, megtanulni" ID="ID_68839370" CREATED="1760618182359" MODIFIED="1760618212313"/>
</node>
</node>
</node>
<node TEXT="Saj&#xe1;t Parser, compiler" ID="ID_1096231839" CREATED="1760615079116" MODIFIED="1760615087110">
<node TEXT="Gener&#xe1;lt Parser" ID="ID_910887344" CREATED="1760615104841" MODIFIED="1760615129910">
<node TEXT="goyacc" ID="ID_1216520508" CREATED="1760618340458" MODIFIED="1760618342036"/>
<node TEXT="gocc" ID="ID_601557704" CREATED="1760618220949" MODIFIED="1760618241872"/>
<node TEXT="" ID="ID_748185966" CREATED="1760618443017" MODIFIED="1760618443017">
<node TEXT="el&#x151;ny&#xf6;k:" ID="ID_620870256" CREATED="1760618355946" MODIFIED="1760618360301">
<node TEXT="gyors, hat&#xe9;kony" ID="ID_847985917" CREATED="1760618389721" MODIFIED="1760618395923"/>
</node>
<node TEXT="h&#xe1;tr&#xe1;nyok:" ID="ID_1293841307" CREATED="1760618362682" MODIFIED="1760618366068">
<node TEXT="nagy betanul&#xe1;si befektet&#xe9;s" ID="ID_773855711" CREATED="1760618366921" MODIFIED="1760618387547"/>
<node TEXT="hibakezel&#xe9;st relat&#xed;ve neh&#xe9;z implement&#xe1;lni" ID="ID_356848749" CREATED="1760618403249" MODIFIED="1760620527236"/>
<node TEXT="compilert, &#xe9;s g&#xe9;pi k&#xf3;dot mindenk&#xe9;ppen implement&#xe1;lni kell hozz&#xe1;" ID="ID_1815838936" CREATED="1760620529842" MODIFIED="1760620545156"/>
<node TEXT="overkill" ID="ID_131599009" CREATED="1760618427833" MODIFIED="1760618432849"/>
</node>
</node>
</node>
<node TEXT="Hand-Written Parser" ID="ID_34055816" CREATED="1760615130540" MODIFIED="1760615140846">
<node TEXT="egyedi megold&#xe1;s" ID="ID_1585818975" CREATED="1760615141956" MODIFIED="1760615145391">
<node TEXT="el&#x151;ny&#xf6;k:" ID="ID_1653575688" CREATED="1760618355946" MODIFIED="1760618360301">
<node TEXT="gyors, hat&#xe9;kony" ID="ID_1464233362" CREATED="1760618389721" MODIFIED="1760618395923"/>
<node TEXT="nagyon flexibilis" ID="ID_1067777691" CREATED="1760618457808" MODIFIED="1760618462636"/>
<node TEXT="hibakezel&#xe9;st relat&#xed;ve k&#xf6;nny&#x171; implement&#xe1;lni" ID="ID_58718325" CREATED="1760618403249" MODIFIED="1760620577683"/>
</node>
<node TEXT="h&#xe1;tr&#xe1;nyok:" ID="ID_1985129639" CREATED="1760618362682" MODIFIED="1760618366068">
<node TEXT="nagyobb fejleszt&#xe9;si ig&#xe9;ny" ID="ID_770671233" CREATED="1760618366921" MODIFIED="1760618480448"/>
</node>
</node>
<node TEXT="parser Combinator" ID="ID_1457856029" CREATED="1760615145563" MODIFIED="1760615149158">
<node TEXT="el&#x151;ny&#xf6;k:" ID="ID_1110425071" CREATED="1760618355946" MODIFIED="1760618360301">
<node TEXT="gyors, hat&#xe9;kony" ID="ID_103043889" CREATED="1760618389721" MODIFIED="1760618395923"/>
<node TEXT="nagyon flexibilis" ID="ID_649549578" CREATED="1760618457808" MODIFIED="1760618462636"/>
<node TEXT="nem t&#xfa;l nagy tanul&#xe1;si befektet&#xe9;s" ID="ID_290719710" CREATED="1760618543821" MODIFIED="1760620618282"/>
<node TEXT="hibakezel&#xe9;st relat&#xed;ve k&#xf6;nny&#x171; implement&#xe1;lni" ID="ID_354657823" CREATED="1760618403249" MODIFIED="1760620628676"/>
<node TEXT="relat&#xed;ve kicsi fejleszt&#xe9;si ig&#xe9;ny" ID="ID_228934947" CREATED="1760618366921" MODIFIED="1760618529656"/>
</node>
<node TEXT="h&#xe1;tr&#xe1;nyok:" ID="ID_445678666" CREATED="1760618362682" MODIFIED="1760618366068">
<node TEXT="meg kell tanulni" ID="ID_1581309217" CREATED="1760620640209" MODIFIED="1760620644701"/>
<node TEXT="n&#xe9;ha nem trivi&#xe1;lis a haszn&#xe1;lata" ID="ID_1541268606" CREATED="1760620648808" MODIFIED="1760620656101"/>
</node>
</node>
</node>
</node>
</node>
<node TEXT="Egy p&#xe9;lda implement&#xe1;ci&#xf3;:" FOLDED="true" ID="ID_1153746595" CREATED="1760615088291" MODIFIED="1760615096359">
<node TEXT="gocalc" ID="ID_401983340" CREATED="1760615096635" MODIFIED="1760615098512">
<node TEXT="elemi sz&#xe1;m&#xed;t&#xe1;si m&#x171;veletek" ID="ID_1676734725" CREATED="1760615156803" MODIFIED="1760615167198"/>
<node TEXT="szintakszis:" FOLDED="true" ID="ID_1395027575" CREATED="1760616118945" MODIFIED="1760616123486">
<node TEXT="lang-syntax.png" ID="ID_293241533" CREATED="1760616153712" MODIFIED="1760616169623" LINK="lang-syntax.png">
<icon BUILTIN="video"/>
<node TEXT="lang-syntax.png" ID="ID_308938773" CREATED="1760616124769" MODIFIED="1760616142036">
<hook URI="lang-syntax.png" SIZE="1.0" NAME="ExternalObject"/>
</node>
</node>
</node>
<node TEXT="parser:" ID="ID_1660868589" CREATED="1760615181987" MODIFIED="1760615189527">
<node TEXT="parc: Parser Combinator go-ban implement&#xe1;lva" ID="ID_869948875" CREATED="1760615189730" MODIFIED="1760615262359"/>
</node>
<node TEXT="interpreter:" ID="ID_1676202431" CREATED="1760615490348" MODIFIED="1760615493448">
<node TEXT="K&#xf6;zvetlen&#xfc;l a parser &#xe1;ltal el&#x151;&#xe1;ll&#xed;tott AST-t futtatja" ID="ID_150376379" CREATED="1760615493659" MODIFIED="1760615519907"/>
</node>
<node TEXT="compiler:" ID="ID_1925127689" CREATED="1760615265731" MODIFIED="1760615270048">
<node TEXT="A parser &#xe1;ltal el&#x151;&#xe1;ll&#xed;tott AST-b&#x151;l egy utas&#xed;t&#xe1;ssort gener&#xe1;l." ID="ID_1507622447" CREATED="1760615522300" MODIFIED="1760615540159"/>
<node TEXT="A program egy line&#xe1;ris utas&#xed;t&#xe1;s array" ID="ID_385462404" CREATED="1760615270339" MODIFIED="1760615478104">
<node ID="ID_512312000" CREATED="1760615479220" MODIFIED="1760615479220"><richcontent TYPE="NODE">

<html>
  <head>
    
  </head>
  <body>
    <p>
      egyszer&#369;, stack-machine alap&#250; m&#369;k&#246;d&#233;s.
    </p>
  </body>
</html>
</richcontent>
</node>
<node TEXT="elemi utas&#xed;t&#xe1;sok, be&#xe9;p&#xed;tett f&#xfc;ggv&#xe9;nyek form&#xe1;j&#xe1;ban" ID="ID_1400577383" CREATED="1760615337460" MODIFIED="1760615354759"/>
<node TEXT="nil" STYLE_REF="code" ID="ID_366222779" CREATED="1760615463531" MODIFIED="1760615468872">
<node TEXT="program v&#xe9;ge" ID="ID_1501636225" CREATED="1760615465396" MODIFIED="1760615467718"/>
</node>
</node>
</node>
<node TEXT="runner:" ID="ID_1184326779" CREATED="1760615395563" MODIFIED="1760615563367">
<node TEXT="A compiler &#xe1;ltal gener&#xe1;lt k&#xf3;dot futtatja." ID="ID_1303736076" CREATED="1760615542956" MODIFIED="1760615553912"/>
<node TEXT="Data Stack" ID="ID_652414824" CREATED="1760615398099" MODIFIED="1760615407814">
<node TEXT="a param&#xe9;terek, &#xe9;s az eredm&#xe9;ny(ek) t&#xe1;rol&#xe1;s&#xe1;ra" ID="ID_541752601" CREATED="1760615436587" MODIFIED="1760615453656"/>
</node>
<node TEXT="IP" ID="ID_1101814875" CREATED="1760615408019" MODIFIED="1760615409464">
<node TEXT="Az aktu&#xe1;lis utas&#xed;t&#xe1;sra mutat" ID="ID_167519738" CREATED="1760615454619" MODIFIED="1760615461640"/>
</node>
</node>
<node TEXT="pl.:" ID="ID_454859519" CREATED="1760614848243" MODIFIED="1760614850791">
<node STYLE_REF="code" ID="ID_1430249723" CREATED="1760614953574" MODIFIED="1760617596270" LINK="tc_0.svg">
<icon BUILTIN="video"/>
<richcontent TYPE="NODE">

<html>
  <head>
    
  </head>
  <body>
    <p>
      &quot;42&quot;
    </p>
  </body>
</html>
</richcontent>
</node>
<node STYLE_REF="code" ID="ID_1174730099" CREATED="1760614953576" MODIFIED="1760616395255" LINK="tc_2.svg">
<icon BUILTIN="video"/>
<richcontent TYPE="NODE">

<html>
  <head>
    
  </head>
  <body>
    <p>
      &quot;42+24&quot;
    </p>
  </body>
</html>
</richcontent>
</node>
<node STYLE_REF="code" ID="ID_1200750210" CREATED="1760614953574" MODIFIED="1760616395254" LINK="tc_0.svg">
<icon BUILTIN="video"/>
<richcontent TYPE="NODE">

<html>
  <head>
    
  </head>
  <body>
    <p>
      &quot;pi + 2&quot;
    </p>
  </body>
</html>
</richcontent>
</node>
<node STYLE_REF="code" ID="ID_1144923988" CREATED="1760614953576" MODIFIED="1760616395256" LINK="tc_3.svg">
<icon BUILTIN="video"/>
<richcontent TYPE="NODE">

<html>
  <head>
    
  </head>
  <body>
    <p>
      &quot;42+24*33&quot;
    </p>
  </body>
</html>
</richcontent>
</node>
<node STYLE_REF="code" ID="ID_1141374290" CREATED="1760614953577" MODIFIED="1760616395256" LINK="tc_7.svg">
<icon BUILTIN="video"/>
<richcontent TYPE="NODE">

<html>
  <head>
    
  </head>
  <body>
    <p>
      &quot;1 + 2 * 3&quot;
    </p>
  </body>
</html>
</richcontent>
</node>
<node STYLE_REF="code" ID="ID_370554033" CREATED="1760614953578" MODIFIED="1760616395256" LINK="tc_8.svg">
<icon BUILTIN="video"/>
<richcontent TYPE="NODE">

<html>
  <head>
    
  </head>
  <body>
    <p>
      &quot;(1 + 2) * 3&quot;
    </p>
  </body>
</html>
</richcontent>
</node>
<node STYLE_REF="code" ID="ID_847806280" CREATED="1760614953577" MODIFIED="1760616395256" LINK="tc_4.svg">
<icon BUILTIN="video"/>
<richcontent TYPE="NODE">

<html>
  <head>
    
  </head>
  <body>
    <p>
      &quot;((10. * 2) + ((50.0 / 3.1415) - 2.))&quot;
    </p>
  </body>
</html>
</richcontent>
</node>
<node STYLE_REF="code" ID="ID_1220273783" CREATED="1760614953577" MODIFIED="1760616395256" LINK="tc_5.svg">
<icon BUILTIN="video"/>
<richcontent TYPE="NODE">

<html>
  <head>
    
  </head>
  <body>
    <p>
      &quot;((10.*2) + ((60.0 / 3) - 2.))&quot;
    </p>
  </body>
</html>
</richcontent>
</node>
</node>
</node>
</node>
<node TEXT="video timeline" FOLDED="true" ID="ID_1275461413" CREATED="1760623566938" MODIFIED="1760623571814">
<node TEXT="A formula editor &#xe9;s runner k&#xf6;vetelm&#xe9;nyei [0:45]" ID="ID_75323925" CREATED="1760614580597" MODIFIED="1760623930727"/>
<node TEXT="Lehets&#xe9;ges megold&#xe1;sok [3:40]" ID="ID_908406243" CREATED="1760614536516" MODIFIED="1760623959687">
<node TEXT="Be&#xe1;gyazott interpreter [3:50]" ID="ID_1204005487" CREATED="1760615031690" MODIFIED="1760623967974"/>
<node TEXT="Saj&#xe1;t Parser, compiler" ID="ID_1713193499" CREATED="1760615079116" MODIFIED="1760615087110">
<node TEXT="Gener&#xe1;lt Parser [5:00]" ID="ID_483402709" CREATED="1760615104841" MODIFIED="1760624065341"/>
<node TEXT="Hand-Written Parser [6:15]" ID="ID_1168321066" CREATED="1760615130540" MODIFIED="1760624083876"/>
</node>
</node>
<node TEXT="Egy p&#xe9;lda implement&#xe1;ci&#xf3;: gocalc [8:29]" ID="ID_1588470643" CREATED="1760615088291" MODIFIED="1760624146418">
<node TEXT="szintakszis [9:30]" ID="ID_1691361196" CREATED="1760616118945" MODIFIED="1760624192415"/>
<node TEXT="parser [12:33]" ID="ID_1706533694" CREATED="1760615181987" MODIFIED="1760624225018"/>
<node TEXT="interpreter (AST) [21:30]" ID="ID_1993005602" CREATED="1760615490348" MODIFIED="1760624404678"/>
<node TEXT="compiler [23:30]" ID="ID_650980767" CREATED="1760615265731" MODIFIED="1760624465671"/>
<node TEXT="AST gr&#xe1;f + g&#xe9;pi k&#xf3;d p&#xe9;ld&#xe1;k [26:30]" ID="ID_1364036672" CREATED="1760623682145" MODIFIED="1760624601114"/>
<node TEXT="betekint&#xe9;s a gocalc implement&#xe1;ci&#xf3;ba [30:50]" ID="ID_76625587" CREATED="1760624627710" MODIFIED="1760624681233"/>
<node TEXT="benchmark-ok [35:50]" ID="ID_1636551899" CREATED="1760624707308" MODIFIED="1760624719351"/>
</node>
<node TEXT="Q&amp;A [36:00]" ID="ID_971368284" CREATED="1760623669465" MODIFIED="1760624740519"/>
<node TEXT="linkek:" ID="ID_1638222695" CREATED="1760624820946" MODIFIED="1760624823278">
<node TEXT="https://github.com/tombenke/gocalc" ID="ID_555092511" CREATED="1760624859736" MODIFIED="1760624859736" LINK="https://github.com/tombenke/gocalc"/>
<node TEXT="https://github.com/tombenke/parc" ID="ID_211589830" CREATED="1760624847489" MODIFIED="1760624847489" LINK="https://github.com/tombenke/parc"/>
</node>
</node>
</node>
<node TEXT="A parc r&#xe9;szletes ismertet&#xe9;se" FOLDED="true" POSITION="right" ID="ID_370098781" CREATED="1760615568516" MODIFIED="1760618107588">
<node TEXT="Bevezet&#xe9;s" ID="ID_1330058524" CREATED="1760617984285" MODIFIED="1760617987255">
<node TEXT="Mi az a Parser Kombin&#xe1;tor?" ID="ID_349752039" CREATED="1760617963700" MODIFIED="1760618003940"/>
<node TEXT="A parser combin&#xe1;torok helye a parserek palett&#xe1;j&#xe1;n" ID="ID_706846846" CREATED="1760617974709" MODIFIED="1760618007936"/>
<node TEXT="El&#x151;ny&#xf6;k &#xe9;s H&#xe1;tr&#xe1;nyok" ID="ID_930460444" CREATED="1760617974709" MODIFIED="1760618011077">
<node TEXT="El&#x151;ny&#xf6;k (Mikor haszn&#xe1;ld?)" ID="ID_701993505" CREATED="1760617974710" MODIFIED="1760618014445"/>
<node TEXT="H&#xe1;tr&#xe1;nyok (Mikor ne haszn&#xe1;ld?)" ID="ID_1336314243" CREATED="1760617974711" MODIFIED="1760618017980"/>
</node>
</node>
<node ID="ID_577710323" CREATED="1760617974712" MODIFIED="1760618044691"><richcontent TYPE="NODE">

<html>
  <head>
    
  </head>
  <body>
    <p>
      A parc parser combinator csomag
    </p>
  </body>
</html>

</richcontent>
<node TEXT="The Parser Objektumok" ID="ID_975560021" CREATED="1760617974712" MODIFIED="1760618021855"/>
<node TEXT="The Parser State" ID="ID_1373831164" CREATED="1760617974713" MODIFIED="1760618024820"/>
</node>
<node TEXT="Primit&#xed;vek" ID="ID_1950310231" CREATED="1760617974714" MODIFIED="1760618027825"/>
<node TEXT="Felt&#xe9;teles parser-ek" ID="ID_1176344166" CREATED="1760617974715" MODIFIED="1760618030197"/>
<node TEXT="Kombin&#xe1;torok" ID="ID_638510488" CREATED="1760617974715" MODIFIED="1760618032828">
<node TEXT="SequenceOf" ID="ID_1185163903" CREATED="1760617974716" MODIFIED="1760618049325"/>
<node TEXT="Choice" ID="ID_825437640" CREATED="1760617974716" MODIFIED="1760618050772"/>
<node TEXT="Count" ID="ID_572266806" CREATED="1760617974717" MODIFIED="1760618052149"/>
</node>
<node TEXT="Mapping" ID="ID_146919695" CREATED="1760617974718" MODIFIED="1760618035972"/>
<node ID="ID_1418786348" CREATED="1760617974718" MODIFIED="1760618044695"><richcontent TYPE="NODE">

<html>
  <head>
    
  </head>
  <body>
    <p>
      L&#225;ncol&#225;s
    </p>
  </body>
</html>

</richcontent>
</node>
<node ID="ID_1787109661" CREATED="1760617974719" MODIFIED="1760618044696"><richcontent TYPE="NODE">

<html>
  <head>
    
  </head>
  <body>
    <p>
      Hibakezel&#233;s
    </p>
  </body>
</html>

</richcontent>
</node>
<node ID="ID_1771895019" CREATED="1760617974720" MODIFIED="1760618044697"><richcontent TYPE="NODE">

<html>
  <head>
    
  </head>
  <body>
    <p>
      Debugging
    </p>
  </body>
</html>

</richcontent>
</node>
<node ID="ID_409330921" CREATED="1760617974721" MODIFIED="1760618044698"><richcontent TYPE="NODE">

<html>
  <head>
    
  </head>
  <body>
    <p>
      Tesztel&#233;s
    </p>
  </body>
</html>

</richcontent>
</node>
</node>
</node>
</map>

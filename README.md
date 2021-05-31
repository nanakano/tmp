4
- HPLMN（Home. Public Land Mobile Network）の対応するSUPI、DNN、およびS-NSSAIのセッション管理サブスクリプションデータが利用できない場合
　- 選択されたDNN、HPLMNのS-NSSAIが明示的にサブスクライブされていない場合、SMFはセッション管理サブスクリプションデータの代わりにローカル設定を使用することができる。

- 手順3の「リクエストタイプ」が「既存のPDUセッション」または「既存の緊急PDUセッション」を示す場合
  - 既存のPDUセッションのSMコンテキストを更新し、更新した内容をAMFに送信する。
  - Request Typeが「Initial request」であり、Nsmf_PDUSession_CreateSMContext RequestにOld PDU Session IDが含まれている場合
    - SMFはOld PDU Session IDに基づいて解放されるべき既存のPDU Sessionを特定します。
  - SMFは、PDU Sessionに冗長性が必要かどうかを判断し、TS 23.501の5.33.2.1項に記載されているようにRSNを決定します。

7a.
- PDUセッションにダイナミックPCCが使用される場合、SMFはTS 23.501の6.3.7.1項に記載されているようにPCFの選択を行う。
- Request Typeが "Existing PDU Session "または "Existing Emergency PDU Session "を示している場合
  - SMFはPDU Sessionのために既に選択されたPCFを使用する。

7b.
- SMFは、4.16.4項で定義されているSMポリシーアソシエーション確立手順を実行して、PCFとの間でSMポリシーアソシエーションを確立し、PDUセッションのためのデフォルトPCCルールを取得することができる。
- SMFで利用可能であれば、GPSI 3GPP Release 16 96 3GPP TS 23.502 V16.8.0 (2021-03)を含めるものとする。
- ステップ3の要求タイプが「既存のPDUセッション」を示す場合、SMFは、4.16.5.1項に定義されているように、SMFが開始したSMポリシーアソシエーション変更手順によって満たされたポリシー制御要求トリガー条件（複数可）に関する情報を提供してもよい。
- PCFは、5.2.5.4項（およびTS 23.503[20]）で定義されたポリシー情報をSMFに提供することができる。
- PCFは、Emergency DNNに基づいて、TS 23.503 [20]に記載されているように、PCCルールのARPをEmergencyサービス用に予約された値に設定する。
- 注5: ステップ7の目的は、UPFを選択する前にPCCルールを受信することである。UPF選択の入力としてPCCルールが必要ない場合、ステップ7はステップ8の後に実行することができる。

9
- SMFは、4.16.5.1項で定義されたSMF開始のSMポリシーアソシエーション変更手順を実行して、満たされたポリシー制御要求トリガー条件（複数可）に関する情報を提供することができる。
- Request Typeが "initial request "で、dynamic PCCが展開されており、PDU Session TypeがIPv4 or IPv6 or IPv4v6の場合、SMFは、割り当てられたUEのIPアドレス／プレフィックス（es）をPCFに通知します（Policy Control Request Trigger条件が満たされた場合）。
- PCFが配備されている場合、PSデータオフ・ポリシー制御要求トリガが設定されていれば、SMFはさらにPSデータオフ・ステータスをPCFに報告しなければなりません。
- 3GPP PSデータオフに関するSMFおよびPCFの追加動作は、TS 23.503 [20]で定義されています。
- 注7: ステップ7の前にIPアドレス／プレフィックスが割り当てられている場合（例：UDM／UDRに登録されたスタティックIPアドレス／プレフィックス）、またはステップ8の後にステップ7が実行される場合、IPアドレス／プレフィックスはステップ7でPCFに提供することができ、本ステップでのIPアドレス／プレフィックスの通知はスキップすることができる。
- PCFは、更新されたポリシーをSMFに提供してもよい。
- PCFは、5.2.5.4項（およびTS 23.503[20]）で定義されたポリシー情報をSMFに提供することができる。

10a.
- SMFはUPFにN4 Session Establishment/Modification Requestを送信し、このPDU SessionのためにUPFにインストールされるパケット検出、実施、報告のルールを提供します。
- TS 23.501 [2]の5.8.2項に記載されているように、SMFがUPFからIPアドレスの割り当てを要求するように設定されている場合、SMFはUPFに対してIPアドレス／プレフィックスの割り当てを実行するように指示し、UPFが割り当てを実行するために必要な情報を含みます。
- このPDUセッションで選択的なユーザプレーンの非活性化が必要な場合、SMFは非活性化タイマを決定し、UPFに提供します。
- SMF は、Trace Requirements を受信した場合には、Trace Requirements を UPF に提供します。
- TS 23.501 [2]で規定されているように、SMFがPDUセッションに対してReliable Data Serviceを有効にしている場合は、このステップでRDS Configuration情報がUPFに提供されます。
- 必要に応じて、SMFはPDUセッションのSmall Data Rate ControlパラメータをUPFに提供する。
- AMFから受信した場合、SMFはUPFにスモールデータレートコントロールステータスを提供します。
- サービングPLMNがこのPDUセッションに対してサービングPLMNレートコントロール（TS 23.501 [2]の5.31.14.2項参照）を実施しようとする場合、SMFはダウンリンクコントロールプレーンデータパケットのレートを制限するためのサービングPLMNレートコントロールパラメータをUPFに提供するものとします。
- イーサネットタイプのPDUセッションの場合、SMFは（例えば、ある要求されたDNN／S-NSSAIの場合）、ポート番号の提供をUPFに要求する表示を含むことができる。
- TS 23.501 [2]の5.33.1.2項に記載されているように、SMFがPDUセッションの1つ以上のQoSフローに対して冗長伝送を行うことを決定した場合、SMFはUPFから2つのCNトンネル情報を要求する。
- また、SMFはUPFに対して、アップリンク方向のQoSフローについて重複するパケットを排除するように指示する。
- SMFはUPFに対して、TS 23.501 [2]の5.33.2.2項に記載されているPDUセッションの冗長トンネルとして1つのCNトンネル情報が使用されることを示す。
- SMFが、TS 23.501 [2]の5.33.1.2項に記載されているように、冗長伝送のためにPSA UPFとNG-RANの間に2つのI-UPFを挿入することを決定した場合、SMFは、対応するCNトンネル情報を要求し、I-UPFおよびPSA UPFにそれぞれ提供する。
- また、SMFは、アップリンク方向のQoSフローの重複パケットを排除するようPSA UPFに指示する。
- SMFは、TS 23.501 [2]の5.33.2.2項に記載されているPDUセッションの冗長トンネルとして1つのCN Tunnel Infoが使用されることをPSA UPFに示す。
- 注8: 2つのGTP-Uトンネルから受信したパケットに基づいて、RAN／UPF上で消去および並べ替えを行う方法は、RAN／UPFの実装に任されています。2 つの GTP-U トンネルは、同じ RAN ノードおよび UPF で終端されます。
- このPDUセッションに対して制御プレーンCIoT 5GS最適化が有効であり、SMFがステップ8でこのPDUセッションのアンカーとしてNEFを選択した場合、SMFは4.25.2節に記載されているようにSMF-NEF接続確立手順を実行します。

10b.
- UPFは、N4 Session Establishment／Modification Responseを送信することで確認します。
- SMFがステップ10aでIPアドレス／プレフィックスの割り当てをUPFが行うことを示した場合、この応答には要求されたIPアドレス／プレフィックスが含まれます。
- このステップでは、要求されたCNトンネル情報がSMFに提供されます。
- ステップ10aにおいて、SMFがUPFに対してQoSフローのパケット重複排除を指示した場合、UPFによって2つのCN Tunnel Infoが割り当てられ、SMFに提供される。
- SMFが、TS 23.501 [2]の5.33.1.2項に記載されているように、冗長伝送のためにPSA UPFとNG-RANの間に2つのI-UPFを挿入することを決定した場合、2つのI-UPFとUPF（PSA）のCN Tunnel InfoがUPFによって割り当てられ、SMFに提供されます。
- UPFは、TS 23.501 [2]の5.33.2.2項に記載されているように、1つのCNトンネル情報がPDUセッションの冗長トンネルとして使用されることをSMFに示す。SMFがポート番号を提供するようUPFに要求した場合、UPFはTS 23.501 [2]に従ってDS-TTポートとブリッジIDを応答に含めます。複数のUPFがPDUセッションのために選択されている場合、SMFはこのステップでPDUセッションの各UPFとのN4セッション確立／変更手順を開始する。注9: PCFがUE IPアドレス変更ポリシー・コントロール・トリガー（TS 23.503 [20]の6.1.3.5項で規定）に加入している場合、SMFはUPFが割り当てたIPアドレス／プレフィックスをPCFに通知します。これは、図 4.3.2.2.1-1 には示されていません。

16c.
- ステップ3のリクエストタイプが「緊急リクエスト」でも「既存の緊急PDUセッション」でもなく、SMFがまだこのPDUセッションに登録していない場合、SMFは所定のPDUセッションのためにNudm_UECM_Registration（SUPI、DNN、S-NSSAI、PDUセッションID、SMF Identity、Serving PLMN ID、[NID]）を使用して、UDMに登録します。その結果、UDMは以下の情報を保存します。
- SUPI、SMFのアイデンティティおよび関連するDNN、S-NSSAI、PDUセッションID、サービングネットワーク（PLMN ID、[NID]、TS 23.501 [2]の5.18節を参照）。
- UDMはさらに、Nudr_DM_Update (SUPI, Subscription Data, UE context in SMF data)により、この情報をUDRに格納することができます。
- ステップ3で受信したRequest Typeが「Emergency Request」を示している場合。認証されたノンローミングUEの場合、事業者の設定（例えば、事業者が緊急通話のために固定SMFを使用するかどうかに関連するなど）に基づいて、SMFは、緊急サービスに適用される所定のPDUセッションのために、3GPPリリース16 101 3GPP TS 23.502 V16.8.0 (2021-03) Nudm_UECM_Registration（SUPI、PDUセッションID、SMFアイデンティティ、緊急サービスの表示）を使用して、UDMに登録することができる。
- その結果、UDMは、緊急サービスに適用可能なPDU Sessionを保存する。認証されていないUEまたはローミングUEの場合、SMFは所定のPDU SessionのためにUDMに登録してはならない。

19.
SMF to UE：PDU Session Type IPv6 または IPv4v6 の場合、SMF は IPv6 Router Advertisement を生成して UE に送信します。このPDUセッションに対して制御プレーン・CIoT 5GS最適化が有効である場合、SMFは制御プレーン・CIoT 5GS最適化手順におけるモバイル・ターミネーテッド・データ・トランスポートを使用してUEに送信するために、AMFを介してIPv6ルータ広告を送信します（4.24.2項参照）、そうでない場合、SMFはN4およびUPFを介してIPv6ルータ広告を送信します。

20.
UEがポート管理情報コンテナの転送のサポートを示している場合、SMFはPCFに5GSブリッジ情報が利用可能であることを通知します。また、SMFは、DS-TTイーサネットポートのポート番号、DS-TTイーサネットポートのMACアドレス、5GSブリッジID、ポート管理情報コンテナ、UE-DSTT Residence Time（UEから提供されたもの）を含みます。AFは、5GSブリッジIDで示される5GSブリッジにサービスを提供するすべてのNW-TTイーサネットポートのUE-DS-TT Residence Timeを用いて、DS-TTイーサネットポートとNW-TTイーサネットポートで構成される各ポートペアのブリッジ遅延を計算する。SMFは、管理可能なNW-TTイーサネットポートが検出されたことをPCFに通知してもよい。SMFがUPFからポート管理情報コンテナを受信した場合、SMFはTS 23.501 [2]の5.28.3.2項に記載されているように、PCFにポート管理情報コンテナを提供する。

21.
ステップ4の後、PDUセッションの確立が失敗した場合、SMFは以下を実行する。SMFは、SMFがこの（SUPI, DNN, HPLMNのS-NSSAI）に対するUEのPDUセッションを扱わなくなった場合、Nudm_SDM_Unsubscribe（SUPI, セッションマネジメントサブスクリプションデータ, DNN, HPLMNのS-NSSAI）を使用して、対応する（SUPI, DNN, HPLMNのS-NSSAI）に対するセッションマネジメントサブスクリプションデータの修正をアンサブスクライブします。UDMは、Nudr_DM_Unsubscribe (SUPI, Subscription Data, Session Management Subscription data, SNSSAI of the HPLMN, DNN)により、UDRからの修正通知の配信を停止することができます。

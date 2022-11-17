// nolint:all
package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/goccy/go-json"
)

/*
curl $'https://forms.office.com/formapi/api/06f1b89f-07e8-464f-b408-ec1b45703f31/users/906fae39-280f-400a-adfd-21da07a30638/forms(\'n7jxBugHT0a0COwbRXA_MTmub5APKApArf0h2gejBjhUN080SVc4NkVNMzc4RlBKVkdaSVY0RFQ4OC4u\')/responses' \
  -H 'authority: forms.office.com' \
  -H '__requestverificationtoken: 5tgleaVY-H2zbiWrVJl4sd2cjDfdA3Q2HKyQn6dfe5HpZEmheJc4FHmxduDXgNB0h_bw1-EEA2WD6rj7rzTiZybfwSMe_23ju37JPVbLGxb-m1WdFFEicRs7B9-sgXtRpjUjnt1m4PV6sRzTsGVWkQ2' \
  -H 'accept: application/json' \
  -H 'accept-language: en-US,en;q=0.9,vi;q=0.8' \
  -H 'authorization;' \
  -H 'content-type: application/json' \
  -H 'cookie: AADNonce=e247096a-d6de-4598-9299-19b458aab401.637977311960960624; __RequestVerificationToken=gNoIfeD4WB2H-h1Pc2M1W1NOcQFy6FQG6x3MKgIf9vXFujFNKXtvnroVGhXmHsLXd4Rhsy66ZSfPKi9IED1gcHp9PMoH5FgSUkYMn7aiUUI1; MUID=35AC9E04561A6F9121268C17572A6E79; MSFPC=GUID=897b2c5651b4473c990f19229aa671b2&HASH=897b&LV=202209&V=4&LU=1662134251761; RpsAuthNonce=888dc75e-2937-4705-908a-029886647269; RpsAuthNonce=888dc75e-2937-4705-908a-029886647269; FormsWebSessionId=6d066a0c-c378-4a5a-94dc-3d009b2014d4; usenewauthrollout=True; AADAuth.forms=eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsIng1dCI6IjJaUXBKM1VwYmpBWVhZR2FYRUpsOGxWMFRPSSIsImtpZCI6IjJaUXBKM1VwYmpBWVhZR2FYRUpsOGxWMFRPSSJ9.eyJhdWQiOiJjOWE1NTlkMi03YWFiLTRmMTMtYTZlZC1lN2U5YzUyYWVjODciLCJpc3MiOiJodHRwczovL3N0cy53aW5kb3dzLm5ldC8wNmYxYjg5Zi0wN2U4LTQ2NGYtYjQwOC1lYzFiNDU3MDNmMzEvIiwiaWF0IjoxNjY3ODM5MjkxLCJuYmYiOjE2Njc4MzkyOTEsImV4cCI6MTY2Nzg0MzE5MSwiYWlvIjoiQVRRQXkvOFRBQUFBMjk2bGdqelp6RDhRMFpReHZVdVgxbW9ZOE5ObVNWbVI4Nmo3ZWVTQkYwT0hsekxjVjdvT2NwOEhFc2xaZWVEUCIsImFtciI6WyJwd2QiXSwiY19oYXNoIjoiRTQyY1ctNlpIQTZwVVAycHNHWDRQQSIsImlwYWRkciI6IjU4LjE4Ni45MC4xOSIsIm5hbWUiOiJQSEFOIFBIVSBUSEFOSCAyMDE3NjExMyIsIm5vbmNlIjoiNjM4MDM0MzYzOTExNjM2NTE2LlpHTmxNbUl3WXpFdFltRmlZUzAwWldSbExXRXdOelF0WTJReFkySTNOREE0WkRrMVptTTBZamt5TVRBdFptVTJPQzAwWlRGbUxXRmlaVE10TURnd1pXWTBPV1k1TTJaaSIsIm9pZCI6IjkwNmZhZTM5LTI4MGYtNDAwYS1hZGZkLTIxZGEwN2EzMDYzOCIsIm9ucHJlbV9zaWQiOiJTLTEtNS0yMS0yNzQ2MjUxMDA3LTEzMjQ1OTUyMDYtNzgxNjU0MzUxLTQ4OTA0IiwicHVpZCI6IjEwMDNCRkZEQTlGRUIwNTkiLCJyaCI6IjAuQVhJQW43anhCdWdIVDBhMENPd2JSWEFfTWRKWnBjbXJlaE5QcHUzbjZjVXE3SWR5QUNnLiIsInN1YiI6Ik5nWG4zOEplT2NOQ0hMMm9GbUN4SWxpa3BBTzdNX094UUJoZm1rWGZmeU0iLCJ0aWQiOiIwNmYxYjg5Zi0wN2U4LTQ2NGYtYjQwOC1lYzFiNDU3MDNmMzEiLCJ1bmlxdWVfbmFtZSI6IlRIQU5ILlBQMTc2MTEzQHNpcy5odXN0LmVkdS52biIsInVwbiI6IlRIQU5ILlBQMTc2MTEzQHNpcy5odXN0LmVkdS52biIsInV0aSI6IkpBS1FlRE1GekVHcHRTVW5pSGVMQUEiLCJ2ZXIiOiIxLjAiLCJ3aWRzIjpbImI3OWZiZjRkLTNlZjktNDY4OS04MTQzLTc2YjE5NGU4NTUwOSJdfQ.tKDoJ9nUWkK8h9XlA7AkU-lTbMY1ObXkP9nXAd2rdi-foV0GkF2dqZ04A4enCmT0sSoQxELtZqc3XQ0B2mwI9JQBvFMI0NW9yR6Y0gkh8YvDAys-FKV_ZDezK3KBNO_ohGCvvIoYSvOyzUSbbeeCU-FFHA7GlhvC3NNeCty5n8RuU-BIOcw-Qs70LvJn5-N3HSM6S6GQLH8GuO4STrYneTEjX_HBjl4lsqZUn_MAXczeR-0p0mnE-ChlSjy9P8IygJynh8OVgvG8zxj1kBEGA6oeGLoNANcOezMt9liwmtce-1JrREHA5Omh-fByIE9yiFvnOS63__eSquRgj3NEyQ; AADAuthCode.forms=0.AXIAn7jxBugHT0a0COwbRXA_MdJZpcmrehNPpu3n6cUq7IdyACg.AgABAAIAAAD--DLA3VO7QrddgJg7WevrAgDs_wQA9P_PGznnxlQUuM220SvZ22YGdftN2m3QnLLYBZODVHlG4MKyOUH-FmJbON-L4Fe1llUNteIkLrDiWwFOnYLRR34kf0wVGY5BMENAmUrU_Dgp3f0T9KdslvomL2VivIjjjDxWnvv71GUBxR41XCqNgLdfKYi_Y4oPA3WNkpBQw5k282BAk7468AB_kDDzBEl7ZbM1N6eCPCABzxCUKuQH0zSO-NO2WesgReRZMvd635e-obCGzW7OqOY1wre5UlJIUrN4krOYJChKsj7ZYao28oyODPP3e4irI5STotgNs0TYGkQNvKKSp1ibTXGXb-NsI4eOH_71KuVm3Y1-JxUYq0g8d0LFxg-0o6I8SFoc3DKOWeEFWvnhfE6tHWpVLvDqp_xiQMhiWwwOVZNJTJhIHVNC0KFNwxIFJHAwlY5O-hT-R_DC3wQpDS_BBv2ed8onMeLBXYdencUvDk_4OzCsRsDS7pKtgumnhC7OY5QDr7GAyZoT1jKjEa9mlshgMAwgxBZdHLoIFe0z6w2QRkpGZp0LxqHuj2kK59h4s3_tOpie-qbzDsOFsA86wn07GRcsd7fNuW_6QqHC2QFEMPqJaFdQ0-pZ5q5_v0TYaWJpaumj_DkOVPd6t5BqFvtaXzHKXiwG0l5CP5uAXW6VrEIwa6psq7cH46SBl2BFmONER_slMYuee3Y0eUCxp1ld23lt325PIolphhsv3r76W7f8nur3x6Y7zy1192oh5UlNFgvY5PxcMqeJUPUNWbvc_gwdI1A8Ih5bO6XqXOKpSS6Cx0ncPVkPYXeQGil_9zabeccTWlKymTQkFEku8TYPqpzv_fFcJHw8M_SXvx7HamCf4PNZdwxs29_RgtnnZuWoDUBO4wvIyIi_qUZKMp0; OIDCAuth.forms=AWjMFfOGot30ee5hg5UppNhrLxmZhhimt5dhrIPPGS2591Gn3l1GsJidNx2_ylulcp2TR707WokqK425W-_XA7a5yiPkCallX4WlbzD72xozYKJ7zycsq75pxEmh2tDy0nbYYNTXJmJ_UibpWoqYkz2ZsAZjZXOyDM7vQu16sgMT0wunqU6qytpGPCN_qbe5KX50BiRO72I9VcD8WFPKrrjI7cg0QtdNMOBvbU0A5yj0pB6pKwVKZ7PXNt_VLXuQH1HX1EZjEnlSK1eqSGnr7gnhzbcRNV69k8h5PdzM7cUCyLZQbXzQUIqnV65_Prv14hayEv46IKBWxJ29lhJJwD8GGDjHOU1J_6vfSeu-ScuqrF0cOLGrg30rwx4z75V7KsDi_7qZqtamVO4kxy0VNkQGg6A6aLnsYd_iW3XZw5k7T1p3EOUbAFCD4lNwgYAGJeuvI2dvmrdEk4z-O9s00SBkccrubrh_MmUkRrt9vsHgoVDHxxpEdOjEbhNun01BAAaOkx3xGY7_pp90zSiruQdL2ucvM5M9A602AY2tSsYPZqfQLBqR_NMKOh-nwXdgGP5IdyFh4JO4WuiFsJFcNwBTCzAmEOx9EeOPdZmfI-x8wUX3pGh-3lQZ1ii5_BL7K7vbIpGH7FKJyF81spRqwIPr01b0aqu7MwXi9eXFQSUBFOu6neKUyj0xkl_w8_x_oZl9jH7pRlBSjSfCCgI_xGirYEl_db9j7Z81Q1J6FbtFS3hemhFBShBXzrcol1Oadu0ZxN8NR8W-T5_v2_UhceEo90gaMlYsQox4_0fGU0Bnv6pQbg6JCjp31mbxRcvDr7X_mzRS32l-1V7Ax8Zxz0bbnC3o7jVqQvACwJ8CAK2RnqMmgy1FvOhHJFL2f3wstf7vtuKe0Ru2d39vgcDeMbzavhmpClfzhLduQBGostp6ttXlIoYUBWgeLJt02WdNpcAaCJ9hHqipleUolSpk_LtdQ1U_Hp1GehLExVokYHcEbES6xTcq-TQACA2GHEs306gECD8hlqcAxDKw2bepAz3pFqrZJo4l9uHw0YxlEJDEf4LROBc26FNMjJiHSuRsJUyDZOgi_5ai9LLTWkzB6mrm-Qtd1HLN9FxInHFd_UCr97EtSu7GJK3dxj0mQQnnNdFuiuUA_8mZ5M8MmbYq6zodFlxsPvzpR3nYInix_fzBHMk4vsIAkiyG71hKzWQAKBCOnDo_jEl18-a1q-q4XmVNr8dFaFkgxdk8SA2IixrKdUFNFTE3tm13bhgHSLpP7meMOon-M-PFWtwQtWX5E3eiEdEwI4DhbpMxne8Qbl4q6pIooNiCAfZ5a4gCXvV38tKYVXL1sjku4pz2NWWAfAsEwKsGovXKu3QMc0R8oRNwbooIVmySI4RkDyfroek3_Kgvv_aX2fa8r0ZXCwnqkQbVHOh5wAHK-Lg4U6Br6EqR9HsAto_pvSBPNs2IzibFb82Hp2AYb0-ZfsLss08j5i221h6g_9Qz1oBxcYbcmYJGe07CWBF8P8Az5CieECHFGV9jGXjiBkxtow7O1BbXUqfZJKkAWvzoDNEK6juMZjEggD3ksvtCK330ydqveG7aJLkVyL2pmoYivoquxxvuHN7v73Qw3dKTIcZ47oAaycYNJBj2JF17Sj0H6y4xo3Ek-w; MicrosoftApplicationsTelemetryDeviceId=3987c6c0-d668-4f36-a875-afe1e3d4a640; ai_session=kKpKpNGr/G07LdL3WQetjF|1667835523852|1667841468643' \
  -H 'dnt: 1' \
  -H 'odata-maxverion: 4.0' \
  -H 'odata-version: 4.0' \
  -H 'origin: https://forms.office.com' \
  -H 'referer: https://forms.office.com/Pages/ResponsePage.aspx?id=n7jxBugHT0a0COwbRXA_MTmub5APKApArf0h2gejBjhUN080SVc4NkVNMzc4RlBKVkdaSVY0RFQ4OC4u' \
  -H 'sec-ch-ua: "Google Chrome";v="107", "Chromium";v="107", "Not=A?Brand";v="24"' \
  -H 'sec-ch-ua-mobile: ?0' \
  -H 'sec-ch-ua-platform: "Linux"' \
  -H 'sec-fetch-dest: empty' \
  -H 'sec-fetch-mode: cors' \
  -H 'sec-fetch-site: same-origin' \
  -H 'user-agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36' \
  -H 'x-correlationid: 77129adf-d59f-46db-b0be-fa711a9a8cae' \
  -H 'x-ms-form-muid: 35AC9E04561A6F9121268C17572A6E79' \
  -H 'x-ms-form-request-ring: business' \
  -H 'x-ms-form-request-source: ms-formweb' \
  -H 'x-usersessionid: 508731bd-709a-45b5-81b9-25b1de2e008c' \
  --data-raw '{"startDate":"2022-11-07T17:17:48.350Z","submitDate":"2022-11-07T17:18:33.410Z","answers":"[{\"questionId\":\"rbf9a69dacb244700a9de97a7e4912735\",\"answer1\":\"t\"},{\"questionId\":\"r10492d18f1574977a9eb1d6af40625d0\",\"answer1\":\"Có\"},{\"questionId\":\"raaccbea3b3e449b3bc42cdd356c1f98d\",\"answer1\":\"[\\\"Tự nhiên\\\"]\"},{\"questionId\":\"r34373a1a64084a1db540399d473d9a17\",\"answer1\":\"Đã dùng\"},{\"questionId\":\"rc807d0a7d8fe4296a530c6dd1a17815f\",\"answer1\":\"Có, mình rất muốn thử\"},{\"questionId\":\"rd6a0c0b218b143e392397870ff2ee411\",\"answer1\":\"c\"},{\"questionId\":\"rbdc2c061351045f9927bf93916fa53d6\",\"answer1\":\"Đang trong một mối quan hệ\"},{\"questionId\":\"rc13060d56e494305955a823cdfbf2287\",\"answer1\":\"Tốt nghiệp THPT\"},{\"questionId\":\"rd5416c02dd4c48cd8fdda73e7d8f7f22\",\"answer1\":\"Nam\"},{\"questionId\":\"rffd89e998912452bbeab03083a82475f\",\"answer1\":\"21-25 tuổi\"},{\"questionId\":\"r44179e41db714d99b2107d9abb687e57\",\"answer1\":\"Dưới 3 triệu VNĐ\"},{\"questionId\":\"r04f8a63450614a439783eee1d64fcd8e\",\"answer1\":{\"id\":1,\"key\":\"1af6d609-de0e-4a44-8760-bc6b87f48a10\"}},{\"questionId\":\"r1acd5486eaba4640812695314c8ddb50\",\"answer1\":{\"id\":2,\"key\":\"1d70d21c-934e-47b5-9d52-e593942de6e3\"}},{\"questionId\":\"r282e11efca4746699b77b2a97877337e\",\"answer1\":{\"id\":3,\"key\":\"86897558-7fb7-40c9-9add-107b230b6473\"}},{\"questionId\":\"r4caefb532498457eae9bdc147533935b\",\"answer1\":{\"id\":4,\"key\":\"3b52502e-1070-4412-af70-744f544cc8f5\"}},{\"questionId\":\"r54dbbad30f684d0fbae9c944b07701c5\",\"answer1\":{\"id\":5,\"key\":\"7f68b283-0cd7-49f3-a827-fb4b982f23c2\"}},{\"questionId\":\"r9056b52a98d64550b4efca6ef6a850e0\",\"answer1\":{\"id\":6,\"key\":\"e51337b2-28d9-4e39-bfcc-eff66163b9e7\"}},{\"questionId\":\"r9ddef0a0dbaf40cfba7fb639f820d377\",\"answer1\":{\"id\":7,\"key\":\"f3925523-16a6-4697-b707-7c40df45a670\"}},{\"questionId\":\"rc20fc8e50fd04464af9367e88df79e85\",\"answer1\":{\"id\":1,\"key\":\"1af6d609-de0e-4a44-8760-bc6b87f48a10\"}},{\"questionId\":\"rc2efd2b723f543cdb80e604ec5caa597\",\"answer1\":{\"id\":2,\"key\":\"1d70d21c-934e-47b5-9d52-e593942de6e3\"}},{\"questionId\":\"rece97f3a671447f4b795293b533e76b4\",\"answer1\":{\"id\":3,\"key\":\"86897558-7fb7-40c9-9add-107b230b6473\"}},{\"questionId\":\"rde79c7395b4e46eb84605b658ac698f6\",\"answer1\":{\"id\":1,\"key\":\"a0f3757a-2390-43a4-8411-0e0086188714\"}},{\"questionId\":\"r0dc878569da646d781fe630bf8071f85\",\"answer1\":{\"id\":2,\"key\":\"871bb7d8-b45c-48e2-a1a6-fa984f347a06\"}},{\"questionId\":\"rd352b222501f4c0391ddf89a630e05eb\",\"answer1\":{\"id\":3,\"key\":\"f3323ccd-3911-40f8-a1e7-5cfd5ce1fca2\"}},{\"questionId\":\"ra98e685c3af44904b5389c64a16e4913\",\"answer1\":{\"id\":4,\"key\":\"c87a56c7-f200-4331-8d97-40818eb60747\"}},{\"questionId\":\"r968bebf2a04247a6bfa910a16d9a4aa8\",\"answer1\":{\"id\":5,\"key\":\"441201e0-0427-4e54-bb59-5c7676256132\"}},{\"questionId\":\"re1718eb4321d4e76baecd4e894a21e22\",\"answer1\":{\"id\":6,\"key\":\"2009cff6-e7c3-4cae-9569-caf57d9417d8\"}},{\"questionId\":\"rf5a143b617fa4b2cb4fe222b92a3cd3d\",\"answer1\":{\"id\":7,\"key\":\"66670ac7-c3e6-45c9-8ae2-45a1a8dc727a\"}},{\"questionId\":\"r20db2bf8cc6a4027bdda06a667bfddb7\",\"answer1\":{\"id\":1,\"key\":\"a0f3757a-2390-43a4-8411-0e0086188714\"}},{\"questionId\":\"r6090b752758a44f08ea8edc5cd83089a\",\"answer1\":{\"id\":1,\"key\":\"8c7e2bce-da6b-42d9-bd8a-618780af863d\"}},{\"questionId\":\"r89ca726b97504cd79434726227b11c01\",\"answer1\":{\"id\":2,\"key\":\"14a42472-bfe5-4309-bbd2-5553de86d868\"}},{\"questionId\":\"r979bd382cb9e43a18ee27852b1a3d6db\",\"answer1\":{\"id\":3,\"key\":\"522bdb6d-b154-477a-9626-9be7aa401436\"}},{\"questionId\":\"ra883a11e27224ecaa8d6ec4981aaf533\",\"answer1\":{\"id\":4,\"key\":\"4e91e9a7-bf0a-4626-8c96-0ac2d7fe14ba\"}},{\"questionId\":\"rbbe81ab5579f43e499d31f5ccad24b39\",\"answer1\":{\"id\":5,\"key\":\"c61c0522-3444-4acd-ab69-d6378906ad23\"}},{\"questionId\":\"rd096d60cddad4b4da5fed238d689769c\",\"answer1\":{\"id\":6,\"key\":\"14949038-49d5-4d4d-8b9e-40ceecedefe5\"}},{\"questionId\":\"rebfc2d78bf40457386fa5e7d0778565d\",\"answer1\":{\"id\":7,\"key\":\"fa47408f-f7ed-45f3-8963-7f2efc1b559a\"}}]"}' \
  --compressed
*/

func rand1To7() string {
	rand.Seed(time.Now().UnixMilli())

	return fmt.Sprintf("%d", rand.Intn(7)+1)
}

type Req struct {
	StartDate           string `json:"startDate"`
	SubmitDate          string `json:"submitDate"`
	Answers             string `json:"answers"`
	EmailReceiptConsent bool   `json:"emailReceiptConsent"`
}

type AnsTxt struct {
	QuestionID string `json:"questionId"`
	Answer1    string `json:"answer1"`
}

type AnsChoice struct {
	QuestionId string        `json:"questionId"`
	Answer1    AnsChoiceData `json:"answer1"`
}

type AnsMulChoice struct {
	QuestionID string `json:"questionId"`
	Answer1    string `json:"answer1"`
}

type AnsChoiceData struct {
	Id  int    `json:"id"`
	Key string `json:"key"`
}

func main() {
	t := time.NewTicker(time.Second * 5)
	for i := 0; i < 140; i++ {
		time.Local = time.UTC
		reqBody := new(Req)
		// 2022-11-04T17:09:13.522Z
		reqBody.StartDate = time.Now().Format("2006-01-02T15:04:05.999Z07:00")
		reqBody.SubmitDate = time.Now().Add((time.Duration(rand.Intn(60)) + 60) * time.Second).Format("2006-01-02T15:04:05.999Z07:00")

		var ans []interface{}

		rand.Seed(time.Now().UnixMilli())
		for i, q := range quesTxt {
			if i == 2 {
				ansMulChoiceData := AnsMulChoice{
					QuestionID: quesMul.Id,
				}
				count := rand.Intn(len(quesMul.Ans))
				var strAns []string
				for k := 0; k < count; k++ {
					id := rand.Intn(len(quesMul.Ans))
					shouldAdd := true
					for j := range strAns {
						if quesMul.Ans[id] == strAns[j] {
							shouldAdd = false
							break
						}
					}
					if shouldAdd {
						strAns = append(strAns, quesMul.Ans[id])
					}
				}
				if count == 0 {
					noAns := []string{"Ko", "ko", "không nhớ", "không rõ", "không để ý"}
					strAns = append(strAns, noAns[rand.Intn(len(noAns))])
				}
				unmsStrAns, _ := json.Marshal(strAns)
				ansMulChoiceData.Answer1 = string(unmsStrAns)
				ans = append(ans, ansMulChoiceData)

				fmt.Printf("%+v", ansMulChoiceData)
			}
			ans = append(
				ans,
				AnsTxt{
					QuestionID: q.Id,
					Answer1:    q.Ans[rand.Intn(len(q.Ans))],
				},
			)
		}

		rand.Seed(time.Now().UnixMilli())
		for _, qG := range qChoice {
			idx := rand.Intn(len(qG.Answers))
			counter := rand.Intn(2)
			sign := 1
			for _, qC := range qG.Questions {
				a := qG.Answers[idx]
				ans = append(
					ans,
					AnsChoice{
						QuestionId: qC,
						Answer1: AnsChoiceData{
							Id:  a.Id,
							Key: a.Key,
						},
					},
				)
				counter++
				if counter%2 == 0 {
					sign = -sign
					idx += 1 * sign
					if idx < 0 {
						idx = 0
						sign = -sign
					}
					if idx >= len(qG.Answers)-1 {
						idx = len(qG.Answers) - 1
						sign = -sign
					}
				}
			}
		}

		ansData, _ := json.Marshal(ans)

		reqBody.Answers = string(ansData)

		data, _ := json.Marshal(reqBody)

		fmt.Printf("\n****\n%s\n****\n", string(data))

		req, err := http.NewRequest("POST", "https://forms.office.com/formapi/api/06f1b89f-07e8-464f-b408-ec1b45703f31/users/906fae39-280f-400a-adfd-21da07a30638/forms('n7jxBugHT0a0COwbRXA_MTmub5APKApArf0h2gejBjhUN080SVc4NkVNMzc4RlBKVkdaSVY0RFQ4OC4u')/responses", bytes.NewReader(data))
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("authority", "forms.office.com")
		req.Header.Set("__requestverificationtoken", "VjR9ZFoVW5oJky3qH7PSNQhgsL5IWajgeqglc7vA6UPw7wNb_J2wvF2-Ion6M4a0VVVfg1U2GwB4BM8KWTqQdio2TmbonPaV31IJ0iHL7rwTw5AlAtMWX25Nt0DTyrV9tGopoQapLkNckpHE8Vau2w2")
		req.Header.Set("accept", "application/json")
		req.Header.Set("accept-language", "en-US,en;q=0.9,vi;q=0.8")
		req.Header.Set("authorization", "")
		req.Header.Set("content-type", "application/json")
		req.Header.Set("cookie", "RpsAuthNonce=157c7263-ff43-4ba3-83ef-720fb8fd56ae; RpsAuthNonce=157c7263-ff43-4ba3-83ef-720fb8fd56ae; __RequestVerificationToken=kHogBsXxKOODWlQgcyi-etc5-_u-Pl5iHGs8qlNBqXguh0GTRRmA-MQ8iY74ZOHFqJ9V8HeHWdDJbKZ-pml6ehMF20_b_dxw7DZHekg3e481; MUID=1CFD32BACFD46D3F258F20EDCEB26CEF; MSFPC=GUID=81cb6ab657844aeb99b0a632484d9911&HASH=81cb&LV=202211&V=4&LU=1668485257073; FormsWebSessionId=c2aa4330-a47a-4803-b43e-168bc3e18e28; usenewauthrollout=True; AADAuth.forms=eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsIng1dCI6IjJaUXBKM1VwYmpBWVhZR2FYRUpsOGxWMFRPSSIsImtpZCI6IjJaUXBKM1VwYmpBWVhZR2FYRUpsOGxWMFRPSSJ9.eyJhdWQiOiJjOWE1NTlkMi03YWFiLTRmMTMtYTZlZC1lN2U5YzUyYWVjODciLCJpc3MiOiJodHRwczovL3N0cy53aW5kb3dzLm5ldC8wNmYxYjg5Zi0wN2U4LTQ2NGYtYjQwOC1lYzFiNDU3MDNmMzEvIiwiaWF0IjoxNjY4NjIyMTQzLCJuYmYiOjE2Njg2MjIxNDMsImV4cCI6MTY2ODYyNjA0MywiYWlvIjoiQVRRQXkvOFRBQUFBeFhURGUvSndjdWoybWVMdjd1TUJnRHdWbEU1Rkx1SFgyc3hRZGRkQXN5MVRqTXJ6c1QvRVNESU5MV2Q4ajVqNSIsImFtciI6WyJwd2QiXSwiY19oYXNoIjoiMVcydy1zQnJOX0w0SVpiWmxXR2YwQSIsImNjIjoiQ21DWkVhOEtIMlQ1Mk1GdG40ZGtTQVFJWUw0S1JWdHZBa2xnSFhyV1Z5aVYwb3lWaXUxN1RLZTlHNmVPQ2NmUzMyQXpsTnl1NFFQVUVEUzYzdmo1K1RYMkh1MlhZRFFXMmREL2d4NUJoZS9scm0zVDV2RUFaMllqcW9ubE5xM1pUNndTRDNOcGN5NW9kWE4wTG1Wa2RTNTJiaG9TQ2hBckE3RlFic3VmUktTNEc1Q2RRZGFzSWhJS0VBWE5qNEJpRjNOTXFoMisrR1huT1FBb0FUSUNRVk00QVVJSkNRQldqcFk3WGRsSSIsImluX2NvcnAiOiJ0cnVlIiwiaXBhZGRyIjoiMS41NS42LjYxIiwibmFtZSI6IlBIQU4gUEhVIFRIQU5IIDIwMTc2MTEzIiwibm9uY2UiOiI2MzgwNDIxOTIzNDIxMjc1ODguWWpsaU5ERmhZall0TURNeVppMDBZelppTFRrMFl6WXRORE5oWVRSa1ltSmhabU0wWkdKbE1EQXpORGt0TVRZeE9DMDBOVE5oTFRrMU5tRXRNREZoTnpjd01UTmlOMlZpIiwib2lkIjoiOTA2ZmFlMzktMjgwZi00MDBhLWFkZmQtMjFkYTA3YTMwNjM4Iiwib25wcmVtX3NpZCI6IlMtMS01LTIxLTI3NDYyNTEwMDctMTMyNDU5NTIwNi03ODE2NTQzNTEtNDg5MDQiLCJwdWlkIjoiMTAwM0JGRkRBOUZFQjA1OSIsInJoIjoiMC5BWElBbjdqeEJ1Z0hUMGEwQ093YlJYQV9NZEpacGNtcmVoTlBwdTNuNmNVcTdJZHlBQ2cuIiwic3ViIjoiTmdYbjM4SmVPY05DSEwyb0ZtQ3hJbGlrcEFPN01fT3hRQmhmbWtYZmZ5TSIsInRpZCI6IjA2ZjFiODlmLTA3ZTgtNDY0Zi1iNDA4LWVjMWI0NTcwM2YzMSIsInVuaXF1ZV9uYW1lIjoiVEhBTkguUFAxNzYxMTNAc2lzLmh1c3QuZWR1LnZuIiwidXBuIjoiVEhBTkguUFAxNzYxMTNAc2lzLmh1c3QuZWR1LnZuIiwidXRpIjoiQmMyUGdHSVhjMHlxSGI3NFplYzVBQSIsInZlciI6IjEuMCIsIndpZHMiOlsiYjc5ZmJmNGQtM2VmOS00Njg5LTgxNDMtNzZiMTk0ZTg1NTA5Il19.YUckLkLwsqvUrjmI7FbeI9VnXgUFBOYK_M0Z21D67HSaJEeZKWlB82fo1N5sa4lfxK3_eX6wWMsgTnOlh9uRRnXijrAExpwyMUI08tNRxhMKKkQrG2JD8g7h9Grf4R-Lp-KAQzZtT-qwZlx7Lqysdw9lzCxkID2DwV7EtKfbvS0HCZFCD5cGZEqcvJoDqckF9s_5MjAqPXpzWYPyZH79XI9iz2Pch7IwsZdPFkCGtG7Wkl4Iz-QpAk3yCA-_LV9PPydkdUeATpuCreTLJFdASamMpw1EOpu_PMRMj4gGhAZAPYbm3UEN3G61axGs31xB1dFs7rPnX6yfLac4XrHE2Q; AADAuthCode.forms=0.AXIAn7jxBugHT0a0COwbRXA_MdJZpcmrehNPpu3n6cUq7IcBAAA.AgABAAIAAAD--DLA3VO7QrddgJg7WevrAgDs_wQA9P9UWkHW7jNCIv5mxBS-nZ0vx7G3RrpDgd82irt0LJVcM8HGPwTacJyuL8tLGPiyrZb4KVXS88Xmm396wFpkH99Q8w3gEJxWVLsNLl8rTQIAJwluLovyEHSr-OqBOxR3n35uN8jNPuhF3IOMwZczNjazRgMVAL14VAOR9aCBzwwf4LILIaotAk2fMzA7sB_3vO9hwxDCHtRqfA_iaq2ZOHecqMSVwdbqSKYTQk5eBeU6-p_0K-VcNffD7I2OVM0yEBo3SiaVisaSA8nDPYujqI_1H1CnStwEbO7RxDQWA9kG6GFbd-YSaUbkkB6j_OXCBVzo_xQNuKTMHzZtKrLFiE-gK4Xn8dVBbDCl8N_JVL6zh14iCo_lKBa2HRVBFS628csVnBM_DSyaytvAFD3kbBe78Rt_Um-yuYX8Qu9VaVPgu5GKRQXjmSJCqtZSTmj9UwqoJH9hC6IddCFE3yXRjALNv-aBOvYdnVRZwP7nV4ncs6wJ3jXON_ce-oWZpoJD75OVMy4V_MWiG8qRxm5Upmw0wN23DKI2_r3vbwAeWxrFuqosxeTozIYxGTeOYFV0Xov1j-iiH8SxdxXXa4CX_Y3pIicZp_D7CpYguL6P0KbnZ4ZxuGK3FAxaDpSD1GuxxchZedecv1x-ZYCp8QCbVD9Tzltt5xg_2NZ5Y4y-68U4IH1tZNPLi7Bvoo8Pn_ae4D6I76DJM2A1airIfa5c_Hvqiz5uJkfNCK5Dns1mInltfJWFO5XTwxQwjC-tCHMn29WSZK5_gSsTwsiUhTseE9x-nRm0s95tayshYbqQjn2p7cj0Q6To9bxztJ3LVt4eFu-DShZI_5A7lzSKbZTM4Ci-LcluWDNuWs4TDt0cqSSuG3MTO5In; OIDCAuth.forms=Af1FiTjGIxh1q0V72IHCZUiS_JAFqtdAKQ6P_qLLtZQ7KtFqEey-4RRuyvyKc_pDW5d32Q13pkP7SoueZmROhMVSqokhPXo7DD8o35repFMj3zK6zwCWanW7CuAKAwjehgX5JgqYws1AM43iVQ_Qj96SXEhkPa_gh2OSRtdB2asGSgkEeAXzUhLVtC2BqRIm6-O2dh9f8S11QSeTUzcnutl5QE_BB0RlbdxbekL8WCd3DJCnhOozMByMStiCnMVZmzazIUOrU9NtOeBSqBr80ETXyzKCPLzWZ-hw7tgbGfpI-hw6TBl-O1Sr7FiDavSlvSyzOSwhxekuiwutHyPekBF0Da5QeLy_mCyZffQwTzYG5LSyZAjIO_olwhioPT9fxQ_E8nzweg3c8ujEOlIC8Wb1628vpmEs7Af4AvvLejSq2-766GS3i2S_dzsfnesnwvUPJorLfhyNDIiOWcWfim6jsakVpgewLwXYfwCUqp-PZsp5xkGlUT-F9FldSta6o9TN9YGHyvSZn6kPw4vq3dU9PfLfW4j9z9v2Ks3M-xYOE-jCTwfJXK4KDMKV0HhwBBdp-K45i2wt6wcp4L3Uimiy2_DkTBfudUHBRTSVsAbibI0Cig3IkbDW-5IUKL1zPgMSLE1xWyIbrDeot2vyBi8RLtq8fLuKAIQQfBEd_Qy3GHekhEPgFA0NtCTCI2O3NCrG_OyaAH3wuZjcPUC8h8eUTFHGEvXnJnpHbYEXUprX62I-EDTFXGPW_T3FOr7ruwg088cl7FtFIGk-h8MJOo_OfFWIRYFigukieqP72IyUijvRsV3aa6Mg5B68XDp-oo7vVdU0cWe5bZ6JEEWJjpRBrZ8piRKnXVKZ_i1RZo0BlI0-d5n6THE8LpLDQM86YCbGLB-6_2pZeUy_3JxoYVI-gYxJMtKAo2AD2iGlW4xlGjEm4caswmd-C3rc7r9Hq9I_RwcTI_K-mBE2PvKeFo64OZhlpwHNCVJQrx6J6VEFlYrSSMuNedW7SP-o_3NgtWCZubpDW5nkzffyBll60k6jYGzyuQ2-O3ezOLaem8VIDFzVAGM60yuO0BB8b5uFkmsG_J--jdMoHksidwi2A92WQILnHmXJwJ7nsUDj1B7BPoZ03oJ2vOu0V0jVzyQDI9h7UfqC9CLzh-ytDZl3ED6NWCIkogd6REkhuH8haPlKm8mKWltbeAD5owSaDwfiBpVcibzX7-OjAIXf8TumXCVaPjEn_F3vEfHP0vJY_0HdD4_VEsU47ZkfVvGQSgWM7CBz_3JjV3wVY_qQc8bLMaSZAp57bO-wTDqDFBG7Pfvtw1HGyopX_u8Nz3v7Mw6xQh3JFZpdzFQFgrTzBEmAnW1Jk2l6xqz4UqKhJTynNG4i18Z4IdEqAewdwmgmKZqSm7fWEtqALYj1lguSxCJfx01-tr8i8BjHnPfDEHVt56P7yVUZOSXRh2kfriWfe_wcleRAwUtix2T6W6P-hpvmVhMcY67TFBWaY5rdKj7knRquZ8miL4si7-XOCCxccR1STGJq98lRw96M9eLZ2O0ay8MCa6dfMzk2sDRa5SHIRQXIjYzMM_h24IKUT_CBQfGOrbPvoqTR_e4t5CebS4eFPceZlPOx9M-yyspIi1D7rsozvS2uzNY59TSZ2TdpSOn5wPPThQTnZSZlLDAqlM0QHrQZU3R9XW3YztnOfK47KHyZE_otuxqdXS7PVKoQq96165N6jclf5alFFRw-wIr8kTqM3e-mfzdRpY_PqZNBJoaX3ODrIUP3SxbaNuC4qluRfcCJyNrmaE9AQOQM8zJx9GKbl3JUqCf1cd7NVHK3MMa4FLCTL5x6fJLICIA8FO2N3xBqNde3nvfnB5rs5BjduCNu7dR1rCf6OYB0b1XRtkIF1UzzjC2ep8Hu4cDv5sikKSguclYgT7dl1YtwsSG1PTzhsvBCidKOWu2ZJY4MZW1Y; MicrosoftApplicationsTelemetryDeviceId=5be2ee84-3204-45cd-b9b1-8344c5f3a944; ai_session=VIvfrYt3AAtVinwsDx6lvs|1668622418048|1668623086120")
		req.Header.Set("dnt", "1")
		req.Header.Set("odata-maxverion", "4.0")
		req.Header.Set("odata-version", "4.0")
		req.Header.Set("origin", "https://forms.office.com")
		req.Header.Set("referer", "https://forms.office.com/Pages/ResponsePage.aspx?id=n7jxBugHT0a0COwbRXA_MTmub5APKApArf0h2gejBjhUN080SVc4NkVNMzc4RlBKVkdaSVY0RFQ4OC4u")
		req.Header.Set("sec-ch-ua", `"Google Chrome";v="107", "Chromium";v="107", "Not=A?Brand";v="24"`)
		req.Header.Set("sec-ch-ua-mobile", "?0")
		req.Header.Set("sec-ch-ua-platform", `"Linux"`)
		req.Header.Set("sec-fetch-dest", "empty")
		req.Header.Set("sec-fetch-mode", "cors")
		req.Header.Set("sec-fetch-site", "same-origin")
		req.Header.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")
		req.Header.Set("x-correlationid", "2ae6eafa-a4fd-44ff-b1fc-e7b76f143fcf")
		req.Header.Set("x-ms-form-muid", "1CFD32BACFD46D3F258F20EDCEB26CEF")
		req.Header.Set("x-ms-form-request-ring", "business")
		req.Header.Set("x-ms-form-request-source", "ms-formweb")
		req.Header.Set("x-usersessionid", "b8d3a94f-27aa-4fb9-8557-4fea0f6a39c3")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		bodyResp, _ := io.ReadAll(resp.Body)

		fmt.Printf("%+v\n %s", resp.Status, string(bodyResp))
		<-t.C
	}
}

/*
// Generated by curl-to-Go: https://mholt.github.io/curl-to-go

// curl $'https://forms.office.com/formapi/api/24ee7655-0535-4f97-9cc2-bc21a13e2674/users/7eaa579a-9911-4dee-8c56-0a2ef89e7f1f/forms(\'VXbuJDUFl0-cwrwhoT4mdJpXqn4Rme5NjFYKLviefx9UQVFHVUlBUEVSV0dIRFlFTDdHVlNXMVRHUS4u\')/responses' \
//   -H 'authority: forms.office.com' \
//   -H '__requestverificationtoken: KfLakinJfnwZq06YIWv7bFm4aJnmhQYcMba3GXqRftRS9a-NIDARYX6Tt1zPe6wLkVM3rqrbgnYjqBNVS_AnDgyh4SBJ9ESDOald2uJz8QQ1' \
//   -H 'accept: application/json' \
//   -H 'accept-language: en-US,en;q=0.9,vi;q=0.8' \
//   -H 'authorization;' \
//   -H 'content-type: application/json' \
//   -H 'cookie: AADNonce=e247096a-d6de-4598-9299-19b458aab401.637977311960960624; __RequestVerificationToken=gNoIfeD4WB2H-h1Pc2M1W1NOcQFy6FQG6x3MKgIf9vXFujFNKXtvnroVGhXmHsLXd4Rhsy66ZSfPKi9IED1gcHp9PMoH5FgSUkYMn7aiUUI1; MUID=35AC9E04561A6F9121268C17572A6E79; MSFPC=GUID=897b2c5651b4473c990f19229aa671b2&HASH=897b&LV=202209&V=4&LU=1662134251761; RpsAuthNonce=888dc75e-2937-4705-908a-029886647269; RpsAuthNonce=888dc75e-2937-4705-908a-029886647269; MicrosoftApplicationsTelemetryDeviceId=9f2e3245-38ca-4fc1-8c41-c5a881833fae; ai_session=3tVGXTUdWCj4nPb2f1Gm9h|1667580883979|1667585246042' \
//   -H 'dnt: 1' \
//   -H 'odata-maxverion: 4.0' \
//   -H 'odata-version: 4.0' \
//   -H 'origin: https://forms.office.com' \
//   -H 'referer: https://forms.office.com/pages/responsepage.aspx?id=VXbuJDUFl0-cwrwhoT4mdJpXqn4Rme5NjFYKLviefx9UQVFHVUlBUEVSV0dIRFlFTDdHVlNXMVRHUS4u' \
//   -H 'sec-ch-ua: "Google Chrome";v="107", "Chromium";v="107", "Not=A?Brand";v="24"' \
//   -H 'sec-ch-ua-mobile: ?0' \
//   -H 'sec-ch-ua-platform: "Linux"' \
//   -H 'sec-fetch-dest: empty' \
//   -H 'sec-fetch-mode: cors' \
//   -H 'sec-fetch-site: same-origin' \
//   -H 'user-agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36' \
//   -H 'x-correlationid: 6ea193a9-cfaa-4531-ba79-336f80b56557' \
//   -H 'x-ms-form-muid: 35AC9E04561A6F9121268C17572A6E79' \
//   -H 'x-ms-form-request-ring: business' \
//   -H 'x-ms-form-request-source: ms-formweb' \
//   -H 'x-usersessionid: e1f47244-5cbf-4117-8ea3-11be32ad1a9a' \
//   --data-raw '{"startDate":"2022-11-04T17:16:32.290Z","submitDate":"2022-11-04T18:07:40.146Z","answers":"[{\"questionId\":\"r358ca0f05ae448b782abec9ba41be461\",\"answer1\":\"Có\"},{\"questionId\":\"r40a7e02204ea401c84abde74455abd17\",\"answer1\":\"2\"},{\"questionId\":\"rae14fcda382045439e42c1412952c3a5\",\"answer1\":\"test@gmail.com\"},{\"questionId\":\"r345a2d420a2c4170a23dd3296ca7081d\",\"answer1\":\"18-24 \"},{\"questionId\":\"r7297f19b771c4ebe842ec43954491279\",\"answer1\":\"Phi nhị nguyên giới\"},{\"questionId\":\"rde6677ac47a640c79007e7bf76029d00\",\"answer1\":\"Cử nhân\"},{\"questionId\":\"rcbfedb985f23489097cdaa62e53c3766\",\"answer1\":\"Đang sống thử\"},{\"questionId\":\"r09d7ebc6531d4339b59b43ef8c376a91\",\"answer1\":\"Công nghệ\"},{\"questionId\":\"rd51bcc572ba54f2f930591136bf3254a\",\"answer1\":\"1.000.000đ - dưới 4.000.000đ\"},{\"questionId\":\"ra096976dfe744b8586166cc4b8638372\",\"answer1\":null},{\"questionId\":\"rb906468d6d5c493993e7c909c5c404e2\",\"answer1\":{\"id\":4,\"key\":\"ca86cfd0-5724-4b2e-885a-0b2a9e5165ab\"}},{\"questionId\":\"r9c0ce440bec342609c023c9f28d2dc0f\",\"answer1\":{\"id\":4,\"key\":\"ca86cfd0-5724-4b2e-885a-0b2a9e5165ab\"}},{\"questionId\":\"r7e1f495b4bf548b999c34478050e1e45\",\"answer1\":{\"id\":4,\"key\":\"ca86cfd0-5724-4b2e-885a-0b2a9e5165ab\"}},{\"questionId\":\"r8ff21380b92d467d9a4dd4f1ed59bddf\",\"answer1\":{\"id\":4,\"key\":\"ca86cfd0-5724-4b2e-885a-0b2a9e5165ab\"}},{\"questionId\":\"r40a95e2904804766b543da0ab9a3ef30\",\"answer1\":{\"id\":4,\"key\":\"ca86cfd0-5724-4b2e-885a-0b2a9e5165ab\"}},{\"questionId\":\"r09ed9f7550e54b198426dacd72d83d42\",\"answer1\":{\"id\":4,\"key\":\"ca86cfd0-5724-4b2e-885a-0b2a9e5165ab\"}},{\"questionId\":\"rb4fd1e63ea054333b9370a6053ba0058\",\"answer1\":{\"id\":4,\"key\":\"ca86cfd0-5724-4b2e-885a-0b2a9e5165ab\"}},{\"questionId\":\"r88ca2040ca484689a750a840a6021d3a\",\"answer1\":{\"id\":4,\"key\":\"ca86cfd0-5724-4b2e-885a-0b2a9e5165ab\"}},{\"questionId\":\"r2fad0b4bf008408aa033ad02d9e13ee2\",\"answer1\":{\"id\":4,\"key\":\"ca86cfd0-5724-4b2e-885a-0b2a9e5165ab\"}},{\"questionId\":\"reac54be4b6cc4293ade67aee6a70c785\",\"answer1\":{\"id\":4,\"key\":\"ca86cfd0-5724-4b2e-885a-0b2a9e5165ab\"}},{\"questionId\":\"r6055ef45dbe8492cbe19f67d3474a696\",\"answer1\":{\"id\":4,\"key\":\"ca86cfd0-5724-4b2e-885a-0b2a9e5165ab\"}},{\"questionId\":\"rce5d2fd6841b4c8a8152c1efd149db25\",\"answer1\":{\"id\":1,\"key\":\"cad8e26d-b8c6-40cc-b950-cf5a9e966d69\"}},{\"questionId\":\"r997122a017cf4c9ebd1aa337ff903f0f\",\"answer1\":{\"id\":3,\"key\":\"803b010d-bf17-4f98-a602-3ab627a2cb84\"}},{\"questionId\":\"r8e9364c75b8e419fa4a4324dbd887b9b\",\"answer1\":{\"id\":4,\"key\":\"f3a3797f-4b7f-4a70-acd0-555132039239\"}},{\"questionId\":\"r5956259036554904b9f8bf7adf391d50\",\"answer1\":{\"id\":5,\"key\":\"029e39cd-fe1d-4ec2-bb38-c5cca114fd4c\"}},{\"questionId\":\"ra52c083cd2964c19857d4fde5c9bd586\",\"answer1\":{\"id\":6,\"key\":\"a42184c7-4e69-442b-9745-42b18f81296b\"}},{\"questionId\":\"ra6153b5d21bf4b5b8a9c54ec96a1ad63\",\"answer1\":{\"id\":7,\"key\":\"3be74e44-4c5a-43bf-811d-c0a3f0d66364\"}},{\"questionId\":\"r83f55c9279c848f0965ec2d89a277951\",\"answer1\":{\"id\":8,\"key\":\"5220be51-9db2-43c5-b272-ffb14d240a15\"}},{\"questionId\":\"r929c7424dfd14018bc188f5aec092280\",\"answer1\":{\"id\":1,\"key\":\"cad8e26d-b8c6-40cc-b950-cf5a9e966d69\"}},{\"questionId\":\"rb0874f69324a44d0ac11ed67e68a0a11\",\"answer1\":{\"id\":1,\"key\":\"c9cab424-4118-4f52-b45c-d97dfd8ae80f\"}},{\"questionId\":\"r1d5afe61e9c94adb8ba27336dd9bb5b8\",\"answer1\":{\"id\":2,\"key\":\"0c904836-bfd9-468a-ade9-e5d3f034dae3\"}},{\"questionId\":\"r169a11e02ea546f69c36c1a5317eafed\",\"answer1\":{\"id\":3,\"key\":\"1a130efe-37e9-45c2-8ceb-6d00969a9286\"}},{\"questionId\":\"rc915909ee07647afb1dc0749f5108837\",\"answer1\":{\"id\":4,\"key\":\"a1557f49-5f38-47ee-adf2-2310657e0a3b\"}},{\"questionId\":\"r1b8e52083bdf404a9a4e51360f184837\",\"answer1\":{\"id\":6,\"key\":\"0744ed42-c3fc-458b-8ceb-cba100b52cb8\"}},{\"questionId\":\"rab4f505929614cb68ea9438dbc41c9f0\",\"answer1\":{\"id\":7,\"key\":\"e643beb0-e0fb-4b18-afe6-39eeed076bcb\"}},{\"questionId\":\"rd2a76bfef1a1456baed3d6875ed8ec90\",\"answer1\":{\"id\":8,\"key\":\"b2097a15-2436-4adc-a662-9ca543f4fedb\"}},{\"questionId\":\"r9302cc913d334977ac000bb62146fd05\",\"answer1\":{\"id\":1,\"key\":\"3f7600dc-27e2-4bc3-8b99-5d8dd502c77d\"}},{\"questionId\":\"r3e0d7da13f4e431ca79994dccaf2b703\",\"answer1\":{\"id\":2,\"key\":\"8cc9e633-ae7e-4528-9018-dc85a171745a\"}},{\"questionId\":\"rfebd13a157064719bc092f614fc3f002\",\"answer1\":{\"id\":3,\"key\":\"37fb862a-cfea-4656-ad1d-f2111563a043\"}},{\"questionId\":\"r204d4119244c4a8e8537d8054801a3d9\",\"answer1\":{\"id\":4,\"key\":\"516c32c0-da1d-471d-bccc-04b4cdc43dad\"}},{\"questionId\":\"r14ce9d334e8843878156c93ba6225f24\",\"answer1\":{\"id\":5,\"key\":\"50332f36-ef3d-40d6-b873-6355d01188a9\"}},{\"questionId\":\"r179924712d314511891f33123f0b520e\",\"answer1\":{\"id\":6,\"key\":\"74905a95-971e-415a-bf2a-19e9f7206bc0\"}},{\"questionId\":\"r6e27fafe69434f06a2c9680d61fd1ecc\",\"answer1\":{\"id\":7,\"key\":\"80f44d45-1283-4b1b-9472-e3fe516c1fb2\"}}]","emailReceiptConsent":false}' \
//   --compressed

body := strings.NewReader("{\"startDate\":\"2022-11-04T17:16:32.290Z\",\"submitDate\":\"2022-11-04T18:07:40.146Z\",\"answers\":\"[{\"questionId\":\"r358ca0f05ae448b782abec9ba41be461\",\"answer1\":\"Có\"},{\"questionId\":\"r40a7e02204ea401c84abde74455abd17\",\"answer1\":\"2\"},{\"questionId\":\"rae14fcda382045439e42c1412952c3a5\",\"answer1\":\"test@gmail.com\"},{\"questionId\":\"r345a2d420a2c4170a23dd3296ca7081d\",\"answer1\":\"18-24 \"},{\"questionId\":\"r7297f19b771c4ebe842ec43954491279\",\"answer1\":\"Phi nhị nguyên giới\"},{\"questionId\":\"rde6677ac47a640c79007e7bf76029d00\",\"answer1\":\"Cử nhân\"},{\"questionId\":\"rcbfedb985f23489097cdaa62e53c3766\",\"answer1\":\"Đang sống thử\"},{\"questionId\":\"r09d7ebc6531d4339b59b43ef8c376a91\",\"answer1\":\"Công nghệ\"},{\"questionId\":\"rd51bcc572ba54f2f930591136bf3254a\",\"answer1\":\"1.000.000đ - dưới 4.000.000đ\"},{\"questionId\":\"ra096976dfe744b8586166cc4b8638372\",\"answer1\":null},{\"questionId\":\"rb906468d6d5c493993e7c909c5c404e2\",\"answer1\":{\"id\":4,\"key\":\"ca86cfd0-5724-4b2e-885a-0b2a9e5165ab\"}},{\"questionId\":\"r9c0ce440bec342609c023c9f28d2dc0f\",\"answer1\":{\"id\":4,\"key\":\"ca86cfd0-5724-4b2e-885a-0b2a9e5165ab\"}},{\"questionId\":\"r7e1f495b4bf548b999c34478050e1e45\",\"answer1\":{\"id\":4,\"key\":\"ca86cfd0-5724-4b2e-885a-0b2a9e5165ab\"}},{\"questionId\":\"r8ff21380b92d467d9a4dd4f1ed59bddf\",\"answer1\":{\"id\":4,\"key\":\"ca86cfd0-5724-4b2e-885a-0b2a9e5165ab\"}},{\"questionId\":\"r40a95e2904804766b543da0ab9a3ef30\",\"answer1\":{\"id\":4,\"key\":\"ca86cfd0-5724-4b2e-885a-0b2a9e5165ab\"}},{\"questionId\":\"r09ed9f7550e54b198426dacd72d83d42\",\"answer1\":{\"id\":4,\"key\":\"ca86cfd0-5724-4b2e-885a-0b2a9e5165ab\"}},{\"questionId\":\"rb4fd1e63ea054333b9370a6053ba0058\",\"answer1\":{\"id\":4,\"key\":\"ca86cfd0-5724-4b2e-885a-0b2a9e5165ab\"}},{\"questionId\":\"r88ca2040ca484689a750a840a6021d3a\",\"answer1\":{\"id\":4,\"key\":\"ca86cfd0-5724-4b2e-885a-0b2a9e5165ab\"}},{\"questionId\":\"r2fad0b4bf008408aa033ad02d9e13ee2\",\"answer1\":{\"id\":4,\"key\":\"ca86cfd0-5724-4b2e-885a-0b2a9e5165ab\"}},{\"questionId\":\"reac54be4b6cc4293ade67aee6a70c785\",\"answer1\":{\"id\":4,\"key\":\"ca86cfd0-5724-4b2e-885a-0b2a9e5165ab\"}},{\"questionId\":\"r6055ef45dbe8492cbe19f67d3474a696\",\"answer1\":{\"id\":4,\"key\":\"ca86cfd0-5724-4b2e-885a-0b2a9e5165ab\"}},{\"questionId\":\"rce5d2fd6841b4c8a8152c1efd149db25\",\"answer1\":{\"id\":1,\"key\":\"cad8e26d-b8c6-40cc-b950-cf5a9e966d69\"}},{\"questionId\":\"r997122a017cf4c9ebd1aa337ff903f0f\",\"answer1\":{\"id\":3,\"key\":\"803b010d-bf17-4f98-a602-3ab627a2cb84\"}},{\"questionId\":\"r8e9364c75b8e419fa4a4324dbd887b9b\",\"answer1\":{\"id\":4,\"key\":\"f3a3797f-4b7f-4a70-acd0-555132039239\"}},{\"questionId\":\"r5956259036554904b9f8bf7adf391d50\",\"answer1\":{\"id\":5,\"key\":\"029e39cd-fe1d-4ec2-bb38-c5cca114fd4c\"}},{\"questionId\":\"ra52c083cd2964c19857d4fde5c9bd586\",\"answer1\":{\"id\":6,\"key\":\"a42184c7-4e69-442b-9745-42b18f81296b\"}},{\"questionId\":\"ra6153b5d21bf4b5b8a9c54ec96a1ad63\",\"answer1\":{\"id\":7,\"key\":\"3be74e44-4c5a-43bf-811d-c0a3f0d66364\"}},{\"questionId\":\"r83f55c9279c848f0965ec2d89a277951\",\"answer1\":{\"id\":8,\"key\":\"5220be51-9db2-43c5-b272-ffb14d240a15\"}},{\"questionId\":\"r929c7424dfd14018bc188f5aec092280\",\"answer1\":{\"id\":1,\"key\":\"cad8e26d-b8c6-40cc-b950-cf5a9e966d69\"}},{\"questionId\":\"rb0874f69324a44d0ac11ed67e68a0a11\",\"answer1\":{\"id\":1,\"key\":\"c9cab424-4118-4f52-b45c-d97dfd8ae80f\"}},{\"questionId\":\"r1d5afe61e9c94adb8ba27336dd9bb5b8\",\"answer1\":{\"id\":2,\"key\":\"0c904836-bfd9-468a-ade9-e5d3f034dae3\"}},{\"questionId\":\"r169a11e02ea546f69c36c1a5317eafed\",\"answer1\":{\"id\":3,\"key\":\"1a130efe-37e9-45c2-8ceb-6d00969a9286\"}},{\"questionId\":\"rc915909ee07647afb1dc0749f5108837\",\"answer1\":{\"id\":4,\"key\":\"a1557f49-5f38-47ee-adf2-2310657e0a3b\"}},{\"questionId\":\"r1b8e52083bdf404a9a4e51360f184837\",\"answer1\":{\"id\":6,\"key\":\"0744ed42-c3fc-458b-8ceb-cba100b52cb8\"}},{\"questionId\":\"rab4f505929614cb68ea9438dbc41c9f0\",\"answer1\":{\"id\":7,\"key\":\"e643beb0-e0fb-4b18-afe6-39eeed076bcb\"}},{\"questionId\":\"rd2a76bfef1a1456baed3d6875ed8ec90\",\"answer1\":{\"id\":8,\"key\":\"b2097a15-2436-4adc-a662-9ca543f4fedb\"}},{\"questionId\":\"r9302cc913d334977ac000bb62146fd05\",\"answer1\":{\"id\":1,\"key\":\"3f7600dc-27e2-4bc3-8b99-5d8dd502c77d\"}},{\"questionId\":\"r3e0d7da13f4e431ca79994dccaf2b703\",\"answer1\":{\"id\":2,\"key\":\"8cc9e633-ae7e-4528-9018-dc85a171745a\"}},{\"questionId\":\"rfebd13a157064719bc092f614fc3f002\",\"answer1\":{\"id\":3,\"key\":\"37fb862a-cfea-4656-ad1d-f2111563a043\"}},{\"questionId\":\"r204d4119244c4a8e8537d8054801a3d9\",\"answer1\":{\"id\":4,\"key\":\"516c32c0-da1d-471d-bccc-04b4cdc43dad\"}},{\"questionId\":\"r14ce9d334e8843878156c93ba6225f24\",\"answer1\":{\"id\":5,\"key\":\"50332f36-ef3d-40d6-b873-6355d01188a9\"}},{\"questionId\":\"r179924712d314511891f33123f0b520e\",\"answer1\":{\"id\":6,\"key\":\"74905a95-971e-415a-bf2a-19e9f7206bc0\"}},{\"questionId\":\"r6e27fafe69434f06a2c9680d61fd1ecc\",\"answer1\":{\"id\":7,\"key\":\"80f44d45-1283-4b1b-9472-e3fe516c1fb2\"}}]\",\"emailReceiptConsent\":false}")
req, err := http.NewRequest("POST", "https://forms.office.com/formapi/api/24ee7655-0535-4f97-9cc2-bc21a13e2674/users/7eaa579a-9911-4dee-8c56-0a2ef89e7f1f/forms(\\'VXbuJDUFl0-cwrwhoT4mdJpXqn4Rme5NjFYKLviefx9UQVFHVUlBUEVSV0dIRFlFTDdHVlNXMVRHUS4u\\')/responses", body)
if err != nil {
	// handle err
}
req.Header.Set("Authority", "forms.office.com")
req.Header.Set("__requestverificationtoken", "KfLakinJfnwZq06YIWv7bFm4aJnmhQYcMba3GXqRftRS9a-NIDARYX6Tt1zPe6wLkVM3rqrbgnYjqBNVS_AnDgyh4SBJ9ESDOald2uJz8QQ1")
req.Header.Set("Accept", "application/json")
req.Header.Set("Accept-Language", "en-US,en;q=0.9,vi;q=0.8")
req.Header.Set("Content-Type", "application/json")
req.Header.Set("Cookie", "AADNonce=e247096a-d6de-4598-9299-19b458aab401.637977311960960624; __RequestVerificationToken=gNoIfeD4WB2H-h1Pc2M1W1NOcQFy6FQG6x3MKgIf9vXFujFNKXtvnroVGhXmHsLXd4Rhsy66ZSfPKi9IED1gcHp9PMoH5FgSUkYMn7aiUUI1; MUID=35AC9E04561A6F9121268C17572A6E79; MSFPC=GUID=897b2c5651b4473c990f19229aa671b2&HASH=897b&LV=202209&V=4&LU=1662134251761; RpsAuthNonce=888dc75e-2937-4705-908a-029886647269; RpsAuthNonce=888dc75e-2937-4705-908a-029886647269; MicrosoftApplicationsTelemetryDeviceId=9f2e3245-38ca-4fc1-8c41-c5a881833fae; ai_session=3tVGXTUdWCj4nPb2f1Gm9h|1667580883979|1667585246042")
req.Header.Set("Dnt", "1")
req.Header.Set("Odata-Maxverion", "4.0")
req.Header.Set("Odata-Version", "4.0")
req.Header.Set("Origin", "https://forms.office.com")
req.Header.Set("Referer", "https://forms.office.com/pages/responsepage.aspx?id=VXbuJDUFl0-cwrwhoT4mdJpXqn4Rme5NjFYKLviefx9UQVFHVUlBUEVSV0dIRFlFTDdHVlNXMVRHUS4u")
req.Header.Set("Sec-Ch-Ua", "\"Google Chrome\";v=\"107\", \"Chromium\";v=\"107\", \"Not=A?Brand\";v=\"24\"")
req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
req.Header.Set("Sec-Ch-Ua-Platform", "\"Linux\"")
req.Header.Set("Sec-Fetch-Dest", "empty")
req.Header.Set("Sec-Fetch-Mode", "cors")
req.Header.Set("Sec-Fetch-Site", "same-origin")
req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")
req.Header.Set("X-Correlationid", "6ea193a9-cfaa-4531-ba79-336f80b56557")
req.Header.Set("X-Ms-Form-Muid", "35AC9E04561A6F9121268C17572A6E79")
req.Header.Set("X-Ms-Form-Request-Ring", "business")
req.Header.Set("X-Ms-Form-Request-Source", "ms-formweb")
req.Header.Set("X-Usersessionid", "e1f47244-5cbf-4117-8ea3-11be32ad1a9a")

resp, err := http.DefaultClient.Do(req)
if err != nil {
	// handle err
}
defer resp.Body.Close()
*/

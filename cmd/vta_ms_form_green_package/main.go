// nolint:all
package main

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"

	"github.com/goccy/go-json"
)

/*
curl $'https://forms.office.com/formapi/api/24ee7655-0535-4f97-9cc2-bc21a13e2674/users/7eaa579a-9911-4dee-8c56-0a2ef89e7f1f/forms(\'VXbuJDUFl0-cwrwhoT4mdJpXqn4Rme5NjFYKLviefx9UQVFHVUlBUEVSV0dIRFlFTDdHVlNXMVRHUS4u\')/responses' \
  -H 'authority: forms.office.com' \
  -H '__requestverificationtoken: KfLakinJfnwZq06YIWv7bFm4aJnmhQYcMba3GXqRftRS9a-NIDARYX6Tt1zPe6wLkVM3rqrbgnYjqBNVS_AnDgyh4SBJ9ESDOald2uJz8QQ1' \
  -H 'accept: application/json' \
  -H 'accept-language: en-US,en;q=0.9,vi;q=0.8' \
  -H 'authorization;' \
  -H 'content-type: application/json' \
  -H 'cookie: AADNonce=e247096a-d6de-4598-9299-19b458aab401.637977311960960624; __RequestVerificationToken=gNoIfeD4WB2H-h1Pc2M1W1NOcQFy6FQG6x3MKgIf9vXFujFNKXtvnroVGhXmHsLXd4Rhsy66ZSfPKi9IED1gcHp9PMoH5FgSUkYMn7aiUUI1; MUID=35AC9E04561A6F9121268C17572A6E79; MSFPC=GUID=897b2c5651b4473c990f19229aa671b2&HASH=897b&LV=202209&V=4&LU=1662134251761; RpsAuthNonce=888dc75e-2937-4705-908a-029886647269; RpsAuthNonce=888dc75e-2937-4705-908a-029886647269; MicrosoftApplicationsTelemetryDeviceId=9f2e3245-38ca-4fc1-8c41-c5a881833fae; ai_session=3tVGXTUdWCj4nPb2f1Gm9h|1667580883979|1667581829859' \
  -H 'dnt: 1' \
  -H 'odata-maxverion: 4.0' \
  -H 'odata-version: 4.0' \
  -H 'origin: https://forms.office.com' \
  -H 'referer: https://forms.office.com/pages/responsepage.aspx?id=VXbuJDUFl0-cwrwhoT4mdJpXqn4Rme5NjFYKLviefx9UQVFHVUlBUEVSV0dIRFlFTDdHVlNXMVRHUS4u' \
  -H 'sec-ch-ua: "Google Chrome";v="107", "Chromium";v="107", "Not=A?Brand";v="24"' \
  -H 'sec-ch-ua-mobile: ?0' \
  -H 'sec-ch-ua-platform: "Linux"' \
  -H 'sec-fetch-dest: empty' \
  -H 'sec-fetch-mode: cors' \
  -H 'sec-fetch-site: same-origin' \
  -H 'user-agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36' \
  -H 'x-correlationid: 2beef1aa-3c23-446c-a953-137f6569816b' \
  -H 'x-ms-form-muid: 35AC9E04561A6F9121268C17572A6E79' \
  -H 'x-ms-form-request-ring: business' \
  -H 'x-ms-form-request-source: ms-formweb' \
  -H 'x-usersessionid: e1f47244-5cbf-4117-8ea3-11be32ad1a9a' \
  --data-raw '{"startDate":"2022-11-04T17:09:13.522Z","submitDate":"2022-11-04T17:10:31.778Z","answers":"[{\"questionId\":\"r358ca0f05ae448b782abec9ba41be461\",\"answer1\":\"Có\"},{\"questionId\":\"r40a7e02204ea401c84abde74455abd17\",\"answer1\":\"cocoon\"},{\"questionId\":\"rae14fcda382045439e42c1412952c3a5\",\"answer1\":\"test@gmail.com\"},{\"questionId\":\"r345a2d420a2c4170a23dd3296ca7081d\",\"answer1\":\"65+\"},{\"questionId\":\"r7297f19b771c4ebe842ec43954491279\",\"answer1\":\"Phi nhị nguyên giới\"},{\"questionId\":\"rde6677ac47a640c79007e7bf76029d00\",\"answer1\":\"Sinh viên\"},{\"questionId\":\"rcbfedb985f23489097cdaa62e53c3766\",\"answer1\":\"Đã ly hôn\"},{\"questionId\":\"r09d7ebc6531d4339b59b43ef8c376a91\",\"answer1\":\"Công nghệ\"},{\"questionId\":\"rd51bcc572ba54f2f930591136bf3254a\",\"answer1\":\">= 6.000.000đ\"},
  {\"questionId\":\"ra096976dfe744b8586166cc4b8638372\",\"answer1\":\"thanks\"},{\"questionId\":\"rb906468d6d5c493993e7c909c5c404e2\",\"answer1\":{\"id\":1,"}},
  {\"questionId\":\"r9c0ce440bec342609c023c9f28d2dc0f\",\"answer1\":{\"id\":2,\"key\":\"d810d0d7-eca5-4583-987b-3d4c1da72996\"}},
  {\"questionId\":\"r7e1f495b4bf548b999c34478050e1e45\",\"answer1\":{\"id\":3,\"key\":\"97707f9e-ce3d-4328-be58-08b5e23294d3\"}},
  {\"questionId\":\"r8ff21380b92d467d9a4dd4f1ed59bddf\",\"answer1\":{\"id\":4,\"key\":\"ca86cfd0-5724-4b2e-885a-0b2a9e5165ab\"}},
  {\"questionId\":\"r40a95e2904804766b543da0ab9a3ef30\",\"answer1\":{\"id\":5,\"key\":\"b8bab2ef-b67d-496d-af8c-6075b9642548\"}},
  {\"questionId\":\"r09ed9f7550e54b198426dacd72d83d42\",\"answer1\":{\"id\":6,\"key\":\"b8e5c5b5-afad-4e00-8168-23d43ebf38d4\"}},
  {\"questionId\":\"rb4fd1e63ea054333b9370a6053ba0058\",\"answer1\":{\"id\":7,\"key\":\"4a2270ba-65ef-4518-b618-1b62a42f748c\"}},
  {\"questionId\":\"r88ca2040ca484689a750a840a6021d3a\",\"answer1\":{\"id\":1,\"key\":\"65ac90a1-1931-4a46-b6c9-2f2acc16fc8a\"}},
  {\"questionId\":\"r2fad0b4bf008408aa033ad02d9e13ee2\",\"answer1\":{\"id\":2,\"key\":\"d810d0d7-eca5-4583-987b-3d4c1da72996\"}},
  {\"questionId\":\"reac54be4b6cc4293ade67aee6a70c785\",\"answer1\":{\"id\":3,\"key\":\"97707f9e-ce3d-4328-be58-08b5e23294d3\"}},
  {\"questionId\":\"r6055ef45dbe8492cbe19f67d3474a696\",\"answer1\":{\"id\":4,\"key\":\"ca86cfd0-5724-4b2e-885a-0b2a9e5165ab\"}},

  {\"questionId\":\"rce5d2fd6841b4c8a8152c1efd149db25\",\"answer1\":{\"id\":1,\"key\":\"cad8e26d-b8c6-40cc-b950-cf5a9e966d69\"}},
  {\"questionId\":\"r997122a017cf4c9ebd1aa337ff903f0f\",\"answer1\":{\"id\":3,\"key\":\"803b010d-bf17-4f98-a602-3ab627a2cb84\"}},
  {\"questionId\":\"r8e9364c75b8e419fa4a4324dbd887b9b\",\"answer1\":{\"id\":4,\"key\":\"f3a3797f-4b7f-4a70-acd0-555132039239\"}},
  {\"questionId\":\"r5956259036554904b9f8bf7adf391d50\",\"answer1\":{\"id\":5,\"key\":\"029e39cd-fe1d-4ec2-bb38-c5cca114fd4c\"}},
  {\"questionId\":\"ra52c083cd2964c19857d4fde5c9bd586\",\"answer1\":{\"id\":6,\"key\":\"a42184c7-4e69-442b-9745-42b18f81296b\"}},
  {\"questionId\":\"ra6153b5d21bf4b5b8a9c54ec96a1ad63\",\"answer1\":{\"id\":7,\"key\":\"3be74e44-4c5a-43bf-811d-c0a3f0d66364\"}},
  {\"questionId\":\"r83f55c9279c848f0965ec2d89a277951\",\"answer1\":{\"id\":8,\"key\":\"5220be51-9db2-43c5-b272-ffb14d240a15\"}},
  {\"questionId\":\"r929c7424dfd14018bc188f5aec092280\",\"answer1\":{\"id\":1,\"key\":\"cad8e26d-b8c6-40cc-b950-cf5a9e966d69\"}},

  {\"questionId\":\"rb0874f69324a44d0ac11ed67e68a0a11\",\"answer1\":{\"id\":1,\"key\":\"c9cab424-4118-4f52-b45c-d97dfd8ae80f\"}},
  {\"questionId\":\"r1d5afe61e9c94adb8ba27336dd9bb5b8\",\"answer1\":{\"id\":2,\"key\":\"0c904836-bfd9-468a-ade9-e5d3f034dae3\"}},
  {\"questionId\":\"r169a11e02ea546f69c36c1a5317eafed\",\"answer1\":{\"id\":3,\"key\":\"1a130efe-37e9-45c2-8ceb-6d00969a9286\"}},
  {\"questionId\":\"rc915909ee07647afb1dc0749f5108837\",\"answer1\":{\"id\":4,\"key\":\"a1557f49-5f38-47ee-adf2-2310657e0a3b\"}},
  {\"questionId\":\"r1b8e52083bdf404a9a4e51360f184837\",\"answer1\":{\"id\":6,\"key\":\"0744ed42-c3fc-458b-8ceb-cba100b52cb8\"}},
  {\"questionId\":\"rab4f505929614cb68ea9438dbc41c9f0\",\"answer1\":{\"id\":7,\"key\":\"e643beb0-e0fb-4b18-afe6-39eeed076bcb\"}},
  {\"questionId\":\"rd2a76bfef1a1456baed3d6875ed8ec90\",\"answer1\":{\"id\":8,\"key\":\"b2097a15-2436-4adc-a662-9ca543f4fedb\"}},

  {\"questionId\":\"r9302cc913d334977ac000bb62146fd05\",\"answer1\":{\"id\":1,\"key\":\"3f7600dc-27e2-4bc3-8b99-5d8dd502c77d\"}},
  {\"questionId\":\"r3e0d7da13f4e431ca79994dccaf2b703\",\"answer1\":{\"id\":2,\"key\":\"8cc9e633-ae7e-4528-9018-dc85a171745a\"}},
  {\"questionId\":\"rfebd13a157064719bc092f614fc3f002\",\"answer1\":{\"id\":3,\"key\":\"37fb862a-cfea-4656-ad1d-f2111563a043\"}},
  {\"questionId\":\"r204d4119244c4a8e8537d8054801a3d9\",\"answer1\":{\"id\":4,\"key\":\"516c32c0-da1d-471d-bccc-04b4cdc43dad\"}},
  {\"questionId\":\"r14ce9d334e8843878156c93ba6225f24\",\"answer1\":{\"id\":5,\"key\":\"50332f36-ef3d-40d6-b873-6355d01188a9\"}},
  {\"questionId\":\"r179924712d314511891f33123f0b520e\",\"answer1\":{\"id\":6,\"key\":\"74905a95-971e-415a-bf2a-19e9f7206bc0\"}},
  {\"questionId\":\"r6e27fafe69434f06a2c9680d61fd1ecc\",\"answer1\":{\"id\":7,\"key\":\"80f44d45-1283-4b1b-9472-e3fe516c1fb2\"}}]","emailReceiptConsent":false}' \
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

type AnsChoiceData struct {
	Id  int    `json:"id"`
	Key string `json:"key"`
}

func main() {
	t := time.NewTicker(time.Second * 5)
	for i := 0; i < 20; i++ {
		<-t.C
		time.Local = time.UTC
		reqBody := new(Req)
		// 2022-11-04T17:09:13.522Z
		reqBody.StartDate = time.Now().Format("2006-01-02T15:04:05.999Z07:00")
		reqBody.SubmitDate = time.Now().Add(time.Minute).Format("2006-01-02T15:04:05.999Z07:00")

		var ans []interface{}

		for _, q := range quesTxt {
			rand.Seed(time.Now().UnixMilli())
			ans = append(
				ans,
				AnsTxt{
					QuestionID: q.Id,
					Answer1:    q.Ans[rand.Intn(len(q.Ans))],
				},
			)
		}

		for _, qG := range qChoice {
			for _, qC := range qG.Questions {
				rand.Seed(time.Now().UnixMilli())
				a := qG.Answers[rand.Intn(len(qG.Answers))]
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
			}
		}

		ansData, _ := json.Marshal(ans)

		reqBody.Answers = string(ansData)

		data, _ := json.Marshal(reqBody)

		fmt.Println(string(data))

		req, err := http.NewRequest(http.MethodPost, "https://forms.office.com/formapi/api/24ee7655-0535-4f97-9cc2-bc21a13e2674/users/7eaa579a-9911-4dee-8c56-0a2ef89e7f1f/forms('VXbuJDUFl0-cwrwhoT4mdJpXqn4Rme5NjFYKLviefx9UQVFHVUlBUEVSV0dIRFlFTDdHVlNXMVRHUS4u')/responses", bytes.NewReader(data))
		if err != nil {
			panic(err)
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
			panic(err)
		}
		defer resp.Body.Close()

		bodyResp, _ := io.ReadAll(resp.Body)

		fmt.Printf("%+v\n %s", resp.Status, string(bodyResp))
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

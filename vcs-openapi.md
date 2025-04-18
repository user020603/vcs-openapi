# Tổng quan OpenAPI Specification (OAS)

## 1. Giới thiệu 

OpenAPI Specification (OAS), trước đây được gọi là Swagger Specification, là một định dạng mô tả API (Giao diện lập trình ứng dụng) chuẩn hóa, độc lập về ngôn ngữ lập trình cho các dịch vụ web RESTful. Nó cho phép cả con người và máy tính khám phá, hiểu và tương tác với các khả năng của một dịch vụ từ xa mà không cần truy cập vào mã nguồn, tài liệu bổ sung hay phân tích lưu lượng mạng.

Mục tiêu chính của OAS là định nghĩa cấu trúc của một API một cách rõ ràng và nhất quán, giúp đơn giản hóa quá trình thiết kế, xây dựng, tài liệu hóa và sử dụng API.

> Note: (*) = optional

## 2. Các khái niệm cốt lõi

Một document OpenAPI (thường là file YAML hoặc JSON) mô tả API bằng cách sử dụng các thành phần chính sau: 

- `openapi`: Chuỗi ký tự xác định phiên bản của OpenAPI Specification đang được sử dụng(ví dụ: `3.0.3`, `3.1.0`).

- `info`: Cung cấp siêu dữ liệu về API bao gồm: 
    - `title`: Tên của API.
    - `version`: Phiên bản của API (khác với phiên bản của OAS).
    - `description` (*): Mô tả chi tiết về API. 
    - `contact` (*): Thông tin liên hệ của người/tổ chức phát triển API. 
    - `license` (*): Thông tin giấy phép của API. 

- `servers` (optional): Một mảng các đối tượng Server, xác định các URL máy chủ cơ sở mà API có thể được truy cập (ví dụ: máy chủ phát triển, sản phẩm). 

- `paths`: Phần quan trọng nhất, định nghĩa các endpoints có sẵn trong API và các phương thức HTTP có thể được thực hiện.
    - **Path Item Object**: Đại điện cho một đường dẫn cụ thể (ví dụ: `/users`, `/users/{userId}`).
    - **Operation Object**: Mô tả một phương thức HTTP duy nhất trên một đường dẫn (ví dụ: `GET /users`). Nó bao gồm: 
        - `tags` (*): Dùng để nhóm các phương thức trong các công cụ như Swagger UI. 
        - `summary` (*): Mô tả ngắn gọn về phương thức. 
        - `description` (*): Mô tả chi tiết hơn. 
        - `operationId` (*): Một định danh duy nhất cho phương thức.
        - `parameters` (*): Danh sách các tham số được chấp nhận (path, query, header, cookie).
        - `requestBody` (*): Mô tả nội dung request (dùng cho POST, PUT, PATCH).
        - `response`: Định nghĩa các phản hồi được trả về, được xác định bằng mã trạng thái HTTP (ví dụ: `200`, `404`, `500`). Mỗi phản hồi mô tả cấu trúc dữ liệu trả về. 
        - `security` (*): Xác định các yêu cầu bảo mật cho phương thức.

- `components`: Một đối tượng chứa các định nghĩa có thể tái sử dụng trong toàn bộ document OpenAPI, giúp tránh lặp lại và dễ quản lý hơn. Các thành phần chính: 
    - `schemas`: Định nghĩa cấu trúc dữ liệu (data models) sử dụng trong `requestBody` và `responses`. Thường sử dụng JSON Schema.
    - `parameters`: Định nghĩa các tham số có thể tái sử dụng. 
    - `requestBodies`: Định nghĩa các nội dung request có thể tái sử dụng.
    - `responses`: Định nghĩa các phản hồi có thể tái sử dụng. 
    - `securitySchemes`: Định nghĩa các cơ chế bảo mật được API hỗ trợ (ví dụ: API Key, OAuth2, HTTP Basic).
    - `headers`: Định nghĩa các header phản hồi có thể tái sử dụng. 
    - `examples`: Định nghĩa các ví dụ dữ liệu có thể tái sử dụng.
    - `callbacks`: Mô tả các hoạt động không đồng bộ, out-of-band mà API có thể kích hoạt.

- `security` (*): Khai báo các yêu cầu bảo mật áp dụng cho toàn bộ API.

- `tags` (*): Danh sách các thẻ được dùng để sắp xếp logic trong document.

- `externalDocs` (*): Liên kết đến các các document bên ngoài bổ sung.

## 3. Lợi ích sử dụng OpenAPI.

- **Thiết kế API (Design-First)**: Cho phép định nghĩa và thảo luận về API trước khi code, đảm bảo đồng thuận và thiết kế tốt hơn. 

- **Auto documenting**: Các công cụ như Swagger UI có thể tự động tạo ra API tương tác, đẹp mắt và luôn cập nhật từ tệp định nghĩa OAS.

- **Code Generation**: Các công cụ như Swagger Codegen có thể tự động tạo mã khung máy chủ (server stubs) và SDK máy khách (client SDKs) bằng nhiều ngôn ngữ lập trình, giảm thiểu công việc thủ công và lỗi.

- **Kiểm thử API**: Định nghĩa OAS có thể được sử dụng để tạo các trường hợp kiểm thử tự động, giảm thiểu công việc thủ công và lỗi. 

- **Tiêu chuẩn hóa**: Cung cấp một ngôn ngữ chung để mô tả API RESTful, giúp cải thiện sự hợp tác và tích hợp giữa các nhóm và hệ thống khác nhau. 

- **Khám phá và tích hợp**: Giúp người dùng và các ứng dụng khác dễ dàng tiếp cận cách sử dụng API. 

## 4. Cấu trúc mẫu (YAML)

```yaml 
openapi: 3.0.3
info:
  title: Simple Pet Store API
  version: 1.0.0
  description: A simple API to manage pets in a store.
servers:
  - url: https://api.example.com/v1
    description: Production server
paths:
  /pets:
    get:
      summary: List all pets
      operationId: listPets
      tags:
        - pets
      parameters:
        - name: limit
          in: query
          description: How many items to return at one time (max 100)
          required: false
          schema:
            type: integer
            format: int32
      responses:
        '200':
          description: A list of pets.
          content:
            application/json:
              schema:
                type: array
                items:
                  $ref: '#/components/schemas/Pet' # Tham chiếu đến schema tái sử dụng
        default:
          description: unexpected error
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Error'
    post:
      summary: Create a pet
      operationId: createPet
      tags:
        - pets
      requestBody:
        description: Pet object to add to the store
        required: true
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/PetInput'
      responses:
        '201':
          description: Null response
        default:
          description: unexpected error
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Error'
  /pets/{petId}:
    get:
      summary: Info for a specific pet
      operationId: showPetById
      tags:
        - pets
      parameters:
        - name: petId
          in: path
          required: true
          description: The id of the pet to retrieve
          schema:
            type: string
      responses:
        '200':
          description: Expected response to a valid request
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Pet'
        default:
          description: unexpected error
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Error'
components:
  schemas:
    Pet:
      type: object
      required:
        - id
        - name
      properties:
        id:
          type: integer
          format: int64
        name:
          type: string
        tag:
          type: string
    PetInput:
      type: object
      required:
        - name
      properties:
        name:
          type: string
        tag:
          type: string
    Error:
      type: object
      required:
        - code
        - message
      properties:
        code:
          type: integer
          format: int32
        message:
          type: string
```

## 5. Các công cụ phổ biến

*   **Swagger UI:** Tạo tài liệu API tương tác từ định nghĩa OAS.
*   **Swagger Editor:** Trình soạn thảo trực tuyến hoặc cục bộ để viết và xác thực định nghĩa OAS.
*   **Swagger Codegen / OpenAPI Generator:** Tạo mã client/server từ định nghĩa OAS. (OpenAPI Generator là một fork cộng đồng từ Swagger Codegen).
*   **Stoplight Studio:** Một công cụ thiết kế và tài liệu hóa API mạnh mẽ hỗ trợ OAS.
*   **Postman:** Có thể nhập định nghĩa OAS để tạo collection và thực hiện yêu cầu API.
*   **Thư viện xác thực:** Nhiều thư viện cho các ngôn ngữ khác nhau để xác thực yêu cầu/phản hồi dựa trên OAS.

## 6. References 

*   **Trang chủ OpenAPI Initiative:** [https://www.openapis.org/](https://www.openapis.org/)
*   **Đặc tả OpenAPI (GitHub):** [https://github.com/OAI/OpenAPI-Specification](https://github.com/OAI/OpenAPI-Specification)
*   **Tài liệu học OpenAPI:** [https://learn.openapis.org/](https://learn.openapis.org/)
*   **Swagger Documentation:** [https://swagger.io/docs/](https://swagger.io/docs/)
*   **JSON Schema:** [https://json-schema.org/](https://json-schema.org/) (Quan trọng để hiểu `components/schemas`)


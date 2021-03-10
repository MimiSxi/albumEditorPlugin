README

# 相册

### graphql example:

```
mutation {
  albumEditor {
    projs {
      action {
        create(cover: "1", name: "name", userId: 1, pages: "[{\"renderRes\":\"11\",\"status\":1,\"direction\":1,\"pType\":1,\"canvasJson\":\"canvasJson1\"}, {\"renderRes\":\"22\",\"status\":1,\"pType\":2,\"canvasJson\":\"canvasJson2\"}]") {
          id
        }
      }
    }
  }
}

mutation {
  albumEditor {
    proj(id: 1) {
      id
      name
      tempUsed{
        id
      }
      pages {
        totalCount
        edges {
          proJId
          status
          direction
          font
          canvasJson
        }
      }
    }
  }
}


```

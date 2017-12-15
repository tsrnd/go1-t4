package http_test

func TestGetByID(t *testing.T) {
  var mockBook models.Book
  err := faker.FakeData(&mockBook)
  assert.NoError(t, err)

  mockUCase := new(mocks.BookUsecase)
  mockID := int(mockBook.ID)
  mockUCase.On("GetByID", int64(mockID)).Return(&mockBook, nil)
  e := echo.New()
  req, err := http.NewRequest(echo.GET, "/book/" +
              strconv.Itoa(int(mockID)), strings.NewReader(""))
  assert.NoError(t, err)

  rec := httptest.NewRecorder()
  c := e.NewContext(req, rec)
  c.SetPath("book/:id")
  c.SetParamNames("id")
  c.SetParamValues(strconv.Itoa(mockID))

  handler:= bookHttp.BookHandler{
            AUsecase: mockUCase,
            Helper: httpHelper.HttpHelper{}
  }
  handler.GetByID(c)

  assert.Equal(t, http.StatusOK, rec.Code)
  mockUCase.AssertCalled(t, "GetByID", int64(mockID))
}

/*
 * Copyright (c) 2018 Miguel Ángel Ortuño.
 * See the LICENSE file for more information.
 */

package memstorage

import (
	"testing"

	"github.com/ortuman/jackal/xml"
	"github.com/stretchr/testify/require"
)

func TestMockStorageInsertPrivateXML(t *testing.T) {
	private := xml.NewElementNamespace("exodus", "exodus:ns")

	s := New()
	s.ActivateMockedError()
	err := s.InsertOrUpdatePrivateXML([]xml.XElement{private}, "exodus:ns", "ortuman")
	require.Equal(t, ErrMockedError, err)
	s.DeactivateMockedError()
	err = s.InsertOrUpdatePrivateXML([]xml.XElement{private}, "exodus:ns", "ortuman")
	require.Nil(t, err)
}

func TestMockStorageFetchPrivateXML(t *testing.T) {
	private := xml.NewElementNamespace("exodus", "exodus:ns")

	s := New()
	s.InsertOrUpdatePrivateXML([]xml.XElement{private}, "exodus:ns", "ortuman")

	s.ActivateMockedError()
	_, err := s.FetchPrivateXML("exodus:ns", "ortuman")
	require.Equal(t, ErrMockedError, err)
	s.DeactivateMockedError()
	elems, _ := s.FetchPrivateXML("exodus:ns", "ortuman")
	require.Equal(t, 1, len(elems))
}

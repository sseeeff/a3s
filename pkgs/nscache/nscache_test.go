package nscache

import (
	"context"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
	"go.aporeto.io/a3s/pkgs/notification"
	"go.aporeto.io/bahamut"
)

func TestNewCache(t *testing.T) {

	Convey("Given I create a new cache", t, func() {

		pubsub := bahamut.NewLocalPubSubClient()
		cache := New(pubsub, 12, true)

		Convey("Then it should be correct", func() {
			So(cache, ShouldNotBeNil)
			So(cache.pubsub, ShouldEqual, pubsub)
			So(cache.cache, ShouldNotBeNil)
			So(cache.invalidateOnAuth, ShouldBeTrue)
		})
	})
}

func TestCacheBehavior(t *testing.T) {

	Convey("Given I create a new cache with some keys", t, func() {

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		pubsub := bahamut.NewLocalPubSubClient()
		_ = pubsub.Connect(ctx)

		cache := New(pubsub, 12, true)
		cache.Start(ctx)

		cache.Set("/hello", "", "hello0", time.Minute)
		cache.Set("/hello/world", "", "hello1", time.Minute)
		cache.Set("/hello/world/cool", "user1", "hello2", time.Minute)
		cache.Set("/hello/world/cool", "user2", "hello3", time.Minute)

		So(cache.Get("/hello", "").Value(), ShouldEqual, "hello0")
		So(cache.Get("/hello/world", "").Value(), ShouldEqual, "hello1")
		So(cache.Get("/hello/world/cool", "user1").Value(), ShouldEqual, "hello2")
		So(cache.Get("/hello/world/cool", "user2").Value(), ShouldEqual, "hello3")

		Convey("When I receive a notification for the deletion of / the cache should be emptied", func() {

			pub := bahamut.NewPublication("notifications.changes.namespace")
			_ = pub.Encode(notification.Message{Data: "/"})
			_ = pubsub.Publish(pub)

			time.Sleep(300 * time.Millisecond)

			So(cache.Get("/hello", ""), ShouldBeNil)
			So(cache.Get("/hello/world", ""), ShouldBeNil)
			So(cache.Get("/hello/world/cool", "user1"), ShouldBeNil)
			So(cache.Get("/hello/world/cool", "user2"), ShouldBeNil)
		})

		Convey("When I receive a notification for the deletion of /hello/world it should remove the corresponding branch", func() {

			pub := bahamut.NewPublication("notifications.changes.namespace")
			_ = pub.Encode(notification.Message{Data: "/hello/world"})
			_ = pubsub.Publish(pub)

			time.Sleep(300 * time.Millisecond)

			So(cache.Get("/hello", "").Value(), ShouldEqual, "hello0")
			So(cache.Get("/hello/world", ""), ShouldBeNil)
			So(cache.Get("/hello/world/cool", "user1"), ShouldBeNil)
			So(cache.Get("/hello/world/cool", "user2"), ShouldBeNil)
		})

		Convey("When I receive a notification for a policy in /hello/world it should remove the corresponding branch", func() {

			pub := bahamut.NewPublication("notifications.changes.authorization")
			_ = pub.Encode(notification.Message{Data: "/hello/world"})
			_ = pubsub.Publish(pub)

			time.Sleep(300 * time.Millisecond)

			So(cache.Get("/hello", "").Value(), ShouldEqual, "hello0")
			So(cache.Get("/hello/world", ""), ShouldBeNil)
			So(cache.Get("/hello/world/cool", "user1"), ShouldBeNil)
			So(cache.Get("/hello/world/cool", "user2"), ShouldBeNil)
		})

		Convey("When I receive a notification for a policy in /hello/world it should nothing when ignored", func() {

			cancel()
			cache.invalidateOnAuth = false

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			cache.Start(ctx)

			pub := bahamut.NewPublication("notifications.changes.authorization")
			_ = pub.Encode(notification.Message{Data: "/hello/world"})
			_ = pubsub.Publish(pub)

			time.Sleep(300 * time.Millisecond)

			So(cache.Get("/hello", "").Value(), ShouldEqual, "hello0")
			So(cache.Get("/hello/world", "").Value(), ShouldEqual, "hello1")
			So(cache.Get("/hello/world/cool", "user1").Value(), ShouldEqual, "hello2")
			So(cache.Get("/hello/world/cool", "user2").Value(), ShouldEqual, "hello3")
		})
	})
}

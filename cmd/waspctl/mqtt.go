package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/vx-labs/mqtt-protocol/packet"
	"github.com/vx-labs/wasp/wasp/api"
	"go.uber.org/zap"
)

func Sessions(ctx context.Context, config *viper.Viper) *cobra.Command {
	sessions := &cobra.Command{
		Use: "sessions",
	}
	sessions.AddCommand(&cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Run: func(cmd *cobra.Command, _ []string) {
			conn, l := mustDial(ctx, cmd, config)
			ctx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
			out, err := api.NewMQTTClient(conn).ListSessionMetadatas(ctx, &api.ListSessionMetadatasRequest{})
			if err != nil {
				l.Fatal("failed to list connected sessions", zap.Error(err))
			}
			cancel()
			table := getTable([]string{"ID", "Client ID", "Peer", "Mount Point", "Connected Since"}, cmd.OutOrStdout())
			for _, member := range out.GetSessionMetadatasList() {
				table.Append([]string{
					member.GetSessionID(),
					member.GetClientID(),
					fmt.Sprintf("%x", member.GetPeer()),
					member.GetMountPoint(),
					time.Since(time.Unix(member.GetConnectedAt(), 0)).String(),
				})
			}
			table.Render()
		},
	})
	return sessions
}
func Messages(ctx context.Context, config *viper.Viper) *cobra.Command {
	messages := &cobra.Command{
		Use: "messages",
	}
	distributeMessage := &cobra.Command{
		Use: "distribute",
		Run: func(cmd *cobra.Command, _ []string) {
			conn, l := mustDial(ctx, cmd, config)
			var payload []byte
			if p := config.GetString("payload"); p != "" {
				payload = []byte(p)
			}
			ctx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
			_, err := api.NewMQTTClient(conn).DistributeMessage(ctx, &api.DistributeMessageRequest{
				Message: &packet.Publish{
					Header: &packet.Header{
						Dup:    config.GetBool("dup"),
						Qos:    config.GetInt32("qos"),
						Retain: config.GetBool("retain"),
					},
					Topic:   []byte(config.GetString("topic")),
					Payload: payload,
				},
			})
			cancel()
			if err != nil {
				l.Fatal("failed to distribute message", zap.Error(err))
			}
		},
	}
	distributeMessage.Flags().Bool("dup", false, "Mark the message as duplicate.")
	distributeMessage.Flags().BoolP("retain", "r", false, "Mark the message as retained.")
	distributeMessage.Flags().Int32P("qos", "q", int32(0), "Set the Message QoS.")
	distributeMessage.Flags().StringP("topic", "t", "", "Set the Message topic.")
	distributeMessage.Flags().StringP("payload", "p", "", "Set the Message payload.")
	distributeMessage.MarkFlagRequired("topic")
	messages.AddCommand(distributeMessage)

	scheduleMessage := &cobra.Command{
		Use: "schedule",
		Run: func(cmd *cobra.Command, _ []string) {
			conn, l := mustDial(ctx, cmd, config)
			var payload []byte
			if p := config.GetString("payload"); p != "" {
				payload = []byte(p)
			}
			ctx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
			_, err := api.NewMQTTClient(conn).ScheduleMessage(ctx, &api.ScheduleMessageRequest{
				Message: &packet.Publish{
					Header: &packet.Header{
						Dup:    config.GetBool("dup"),
						Qos:    config.GetInt32("qos"),
						Retain: config.GetBool("retain"),
					},
					Topic:   []byte(config.GetString("topic")),
					Payload: payload,
				},
			})
			cancel()
			if err != nil {
				l.Fatal("failed to schedule message", zap.Error(err))
			}
		},
	}
	scheduleMessage.Flags().Bool("dup", false, "Mark the message as duplicate.")
	scheduleMessage.Flags().BoolP("retain", "r", false, "Mark the message as retained.")
	scheduleMessage.Flags().Int32P("qos", "q", int32(0), "Set the Message QoS.")
	scheduleMessage.Flags().StringP("topic", "t", "", "Set the Message topic.")
	scheduleMessage.Flags().StringP("payload", "p", "", "Set the Message payload.")
	scheduleMessage.MarkFlagRequired("topic")
	messages.AddCommand(scheduleMessage)
	return messages
}

func Subscriptions(ctx context.Context, config *viper.Viper) *cobra.Command {
	subscriptions := &cobra.Command{
		Use: "subscriptions",
	}
	listSubscriptions := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Run: func(cmd *cobra.Command, _ []string) {
			conn, l := mustDial(ctx, cmd, config)
			ctx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
			out, err := api.NewMQTTClient(conn).ListSubscriptions(ctx, &api.ListSubscriptionsRequest{})
			cancel()
			if err != nil {
				l.Fatal("failed to list subscriptions", zap.Error(err))
			}
			table := getTable([]string{"Client ID", "Peer", "Pattern", "QoS"}, cmd.OutOrStdout())
			for _, member := range out.GetSubscriptions() {
				table.Append([]string{
					member.GetSessionID(),
					fmt.Sprintf("%x", member.GetPeer()),
					string(member.GetPattern()),
					fmt.Sprintf("%x", member.GetQoS())})
			}
			table.Render()
		},
	}
	createSubscription := &cobra.Command{
		Use:     "create",
		Aliases: []string{"new"},
		Run: func(cmd *cobra.Command, _ []string) {
			conn, l := mustDial(ctx, cmd, config)
			ctx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
			peerIDStr := config.GetString("peer")
			peerID, err := strconv.ParseUint(peerIDStr, 10, 64)
			if err != nil {
				peerID, err = strconv.ParseUint(peerIDStr, 16, 64)
				if err != nil {
					l.Fatal("failed to parse peer-id", zap.Error(err))
				}
			}
			_, err = api.NewMQTTClient(conn).CreateSubscription(ctx, &api.CreateSubscriptionRequest{
				SessionID: config.GetString("session-id"),
				Pattern:   []byte(config.GetString("pattern")),
				Peer:      peerID,
				QoS:       config.GetInt32("qos"),
			})
			cancel()
			if err != nil {
				l.Fatal("failed to create subscription", zap.Error(err))
			}
			fmt.Printf("subscribed %s:%s to %s (qos %v)\n",
				config.GetString("peer"), config.GetString("session-id"), config.GetString("pattern"), config.GetInt32("qos"))
		},
	}
	createSubscription.Flags().Int32P("qos", "q", int32(0), "Set the Subscription's  QoS.")
	createSubscription.Flags().StringP("session-id", "i", "", "Set the Subscription's session id.")
	createSubscription.Flags().StringP("pattern", "t", "", "Set the Subscription's pattern.")
	createSubscription.Flags().StringP("peer", "p", "", "Set the Subscription's Peer.")

	createSubscription.MarkFlagRequired("pattern")
	createSubscription.MarkFlagRequired("peer")
	createSubscription.MarkFlagRequired("session-id")

	deleteSubscription := &cobra.Command{
		Use:     "delete",
		Aliases: []string{"rm"},
		Run: func(cmd *cobra.Command, _ []string) {
			conn, l := mustDial(ctx, cmd, config)
			ctx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
			peerIDStr := config.GetString("peer")
			peerID, err := strconv.ParseUint(peerIDStr, 10, 64)
			if err != nil {
				peerID, err = strconv.ParseUint(peerIDStr, 16, 64)
				if err != nil {
					l.Fatal("failed to parse peer-id", zap.Error(err))
				}
			}
			_, err = api.NewMQTTClient(conn).DeleteSubscription(ctx, &api.DeleteSubscriptionRequest{
				SessionID: config.GetString("session-id"),
				Pattern:   []byte(config.GetString("pattern")),
				Peer:      peerID,
			})
			cancel()
			if err != nil {
				l.Fatal("failed to create subscription", zap.Error(err))
			}
		},
	}
	deleteSubscription.Flags().StringP("session-id", "i", "", "Set the Subscription's session id.")
	deleteSubscription.Flags().StringP("pattern", "t", "", "Set the Subscription's pattern.")
	deleteSubscription.Flags().StringP("peer", "p", "", "Set the Subscription's peer id.")

	deleteSubscription.MarkFlagRequired("pattern")
	deleteSubscription.MarkFlagRequired("peer")
	deleteSubscription.MarkFlagRequired("session-id")

	subscriptions.AddCommand(listSubscriptions)
	subscriptions.AddCommand(createSubscription)
	subscriptions.AddCommand(deleteSubscription)
	return subscriptions
}

var qosString = map[int32]string{
	0: "at-most-once",
	1: "at-least-once",
	2: "exactly-once",
}

func Topics(ctx context.Context, config *viper.Viper) *cobra.Command {
	topics := &cobra.Command{
		Use: "topics",
	}
	topics.AddCommand(&cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Run: func(cmd *cobra.Command, patterns []string) {
			conn, l := mustDial(ctx, cmd, config)
			if len(patterns) == 0 {
				patterns = []string{"#"}
			}
			for _, pattern := range patterns {
				ctx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
				out, err := api.NewMQTTClient(conn).ListRetainedMessages(ctx, &api.ListRetainedMessagesRequest{
					Pattern: []byte(pattern),
				})
				cancel()
				if err != nil {
					l.Fatal("failed to list topics", zap.Error(err))
				}
				table := getTable([]string{"Topic", "Payload", "QoS", "Age"}, cmd.OutOrStdout())
				for _, member := range out.GetRetainedMessages() {
					table.Append([]string{
						string(member.Publish.GetTopic()),
						string(member.Publish.GetPayload()),
						qosString[member.Publish.Header.GetQos()],
						humanize.Time(time.Unix(0, member.LastAdded)),
					})
				}
				table.Render()
			}
		},
	})
	topics.AddCommand(&cobra.Command{
		Use:     "delete",
		Aliases: []string{"rm"},
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, topics []string) {
			conn, l := mustDial(ctx, cmd, config)
			for _, topic := range topics {
				ctx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
				_, err := api.NewMQTTClient(conn).DeleteRetainedMessage(ctx, &api.DeleteRetainedMessageRequest{
					Topic: []byte(topic),
				})
				cancel()
				if err != nil {
					l.Fatal("failed to delete topic", zap.Error(err))
				}
				fmt.Println(string(topic))

			}
		},
	})
	return topics
}

func Cluster(ctx context.Context, config *viper.Viper) *cobra.Command {
	cluster := &cobra.Command{
		Use: "cluster",
	}
	cluster.AddCommand(&cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Run: func(cmd *cobra.Command, _ []string) {
			conn, l := mustDial(ctx, cmd, config)
			ctx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
			out, err := api.NewMQTTClient(conn).ListClusterMembers(ctx, &api.ListClusterMembersRequest{})
			cancel()
			if err != nil {
				l.Fatal("failed to list cluster members", zap.Error(err))
			}
			table := getTable([]string{"ID", "RPC Address", "Health"}, cmd.OutOrStdout())
			for _, member := range out.GetClusterMembers() {
				table.Append([]string{
					fmt.Sprintf("%x", member.ID),
					member.Address,
					member.HealthState,
				})
				table.Render()
			}
		},
	})
	return cluster
}

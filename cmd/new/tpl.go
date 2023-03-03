package new

// TplProjectStructure 项目结构
const TplProjectStructure = `
├── LICENSE
├── Makefile
├── README.md
├── cmd
│   ├── generator
│   │   └── entgo
│   │       └── main.go
│   └── omgind
│       ├── common
│       │   └── ent_client.go
│       ├── enter
│       │   └── root.go
│       ├── main.go
│       ├── migrate
│       │   ├── district.go
│       │   ├── dump.go
│       │   ├── load.go
│       │   └── migrate.go
│       ├── version
│       │   └── version.go
│       └── web
│           └── web.go
├── configs
│   ├── config.dev.toml
│   ├── config.docker.toml
│   ├── config.prod.toml
│   ├── config.toml -> config.dev.toml
│   ├── menu.yaml
│   └── model.conf
├── data
│   ├── config
│   │   ├── mysql
│   │   │   ├── conf.d
│   │   │   ├── my.cnf
│   │   │   └── mysql-files
│   │   └── redis
│   │       └── redis.conf
│   ├── influxdb
│   │   ├── engine
│   │   │   └── data
│   │   └── influxd.bolt
│   ├── jwt_auth.db
│   ├── mysql
│   │   └── data
│   │       ├── #ib_16384_0.dblwr
│   │       ├── #ib_16384_1.dblwr
│   │       ├── #innodb_temp
│   │       │   ├── temp_1.ibt
│   │       │   ├── temp_10.ibt
│   │       │   ├── temp_2.ibt
│   │       │   ├── temp_3.ibt
│   │       │   ├── temp_4.ibt
│   │       │   ├── temp_5.ibt
│   │       │   ├── temp_6.ibt
│   │       │   ├── temp_7.ibt
│   │       │   ├── temp_8.ibt
│   │       │   └── temp_9.ibt
│   │       ├── auto.cnf
│   │       ├── binlog.000001
│   │       ├── binlog.000002
│   │       ├── binlog.000003
│   │       ├── binlog.000004
│   │       ├── binlog.index
│   │       ├── ca-key.pem
│   │       ├── ca.pem
│   │       ├── client-cert.pem
│   │       ├── client-key.pem
│   │       ├── ib_buffer_pool
│   │       ├── ib_logfile0
│   │       ├── ib_logfile1
│   │       ├── ibdata1
│   │       ├── ibtmp1
│   │       ├── mysql
│   │       │   ├── general_log.CSM
│   │       │   ├── general_log.CSV
│   │       │   ├── general_log_213.sdi
│   │       │   ├── slow_log.CSM
│   │       │   ├── slow_log.CSV
│   │       │   └── slow_log_214.sdi
│   │       ├── mysql.ibd
│   │       ├── omgind
│   │       ├── performance_schema
│   │       │   ├── accounts_145.sdi
│   │       │   ├── binary_log_trans_189.sdi
│   │       │   ├── cond_instances_82.sdi
│   │       │   ├── data_lock_waits_161.sdi
│   │       │   ├── data_locks_160.sdi
│   │       │   ├── error_log_83.sdi
│   │       │   ├── events_errors_su_139.sdi
│   │       │   ├── events_errors_su_140.sdi
│   │       │   ├── events_errors_su_141.sdi
│   │       │   ├── events_errors_su_142.sdi
│   │       │   ├── events_errors_su_143.sdi
│   │       │   ├── events_stages_cu_111.sdi
│   │       │   ├── events_stages_hi_112.sdi
│   │       │   ├── events_stages_hi_113.sdi
│   │       │   ├── events_stages_su_114.sdi
│   │       │   ├── events_stages_su_115.sdi
│   │       │   ├── events_stages_su_116.sdi
│   │       │   ├── events_stages_su_117.sdi
│   │       │   ├── events_stages_su_118.sdi
│   │       │   ├── events_statement_119.sdi
│   │       │   ├── events_statement_120.sdi
│   │       │   ├── events_statement_121.sdi
│   │       │   ├── events_statement_122.sdi
│   │       │   ├── events_statement_123.sdi
│   │       │   ├── events_statement_124.sdi
│   │       │   ├── events_statement_125.sdi
│   │       │   ├── events_statement_126.sdi
│   │       │   ├── events_statement_127.sdi
│   │       │   ├── events_statement_128.sdi
│   │       │   ├── events_statement_129.sdi
│   │       │   ├── events_statement_130.sdi
│   │       │   ├── events_transacti_131.sdi
│   │       │   ├── events_transacti_132.sdi
│   │       │   ├── events_transacti_133.sdi
│   │       │   ├── events_transacti_134.sdi
│   │       │   ├── events_transacti_135.sdi
│   │       │   ├── events_transacti_136.sdi
│   │       │   ├── events_transacti_137.sdi
│   │       │   ├── events_transacti_138.sdi
│   │       │   ├── events_waits_cur_84.sdi
│   │       │   ├── events_waits_his_85.sdi
│   │       │   ├── events_waits_his_86.sdi
│   │       │   ├── events_waits_sum_87.sdi
│   │       │   ├── events_waits_sum_88.sdi
│   │       │   ├── events_waits_sum_89.sdi
│   │       │   ├── events_waits_sum_90.sdi
│   │       │   ├── events_waits_sum_91.sdi
│   │       │   ├── events_waits_sum_92.sdi
│   │       │   ├── file_instances_93.sdi
│   │       │   ├── file_summary_by__94.sdi
│   │       │   ├── file_summary_by__95.sdi
│   │       │   ├── global_status_181.sdi
│   │       │   ├── global_variables_184.sdi
│   │       │   ├── host_cache_96.sdi
│   │       │   ├── hosts_146.sdi
│   │       │   ├── keyring_componen_191.sdi
│   │       │   ├── keyring_keys_152.sdi
│   │       │   ├── log_status_174.sdi
│   │       │   ├── memory_summary_b_154.sdi
│   │       │   ├── memory_summary_b_155.sdi
│   │       │   ├── memory_summary_b_156.sdi
│   │       │   ├── memory_summary_b_157.sdi
│   │       │   ├── memory_summary_g_153.sdi
│   │       │   ├── metadata_locks_159.sdi
│   │       │   ├── mutex_instances_97.sdi
│   │       │   ├── objects_summary__98.sdi
│   │       │   ├── performance_time_99.sdi
│   │       │   ├── persisted_variab_187.sdi
│   │       │   ├── prepared_stateme_175.sdi
│   │       │   ├── processlist_100.sdi
│   │       │   ├── replication_appl_165.sdi
│   │       │   ├── replication_appl_166.sdi
│   │       │   ├── replication_appl_167.sdi
│   │       │   ├── replication_appl_168.sdi
│   │       │   ├── replication_appl_170.sdi
│   │       │   ├── replication_appl_171.sdi
│   │       │   ├── replication_asyn_172.sdi
│   │       │   ├── replication_asyn_173.sdi
│   │       │   ├── replication_conn_162.sdi
│   │       │   ├── replication_conn_164.sdi
│   │       │   ├── replication_grou_163.sdi
│   │       │   ├── replication_grou_169.sdi
│   │       │   ├── rwlock_instances_101.sdi
│   │       │   ├── session_account__151.sdi
│   │       │   ├── session_connect__150.sdi
│   │       │   ├── session_status_182.sdi
│   │       │   ├── session_variable_185.sdi
│   │       │   ├── setup_actors_102.sdi
│   │       │   ├── setup_consumers_103.sdi
│   │       │   ├── setup_instrument_104.sdi
│   │       │   ├── setup_objects_105.sdi
│   │       │   ├── setup_threads_106.sdi
│   │       │   ├── socket_instances_147.sdi
│   │       │   ├── socket_summary_b_148.sdi
│   │       │   ├── socket_summary_b_149.sdi
│   │       │   ├── status_by_accoun_177.sdi
│   │       │   ├── status_by_host_178.sdi
│   │       │   ├── status_by_thread_179.sdi
│   │       │   ├── status_by_user_180.sdi
│   │       │   ├── table_handles_158.sdi
│   │       │   ├── table_io_waits_s_107.sdi
│   │       │   ├── table_io_waits_s_108.sdi
│   │       │   ├── table_lock_waits_109.sdi
│   │       │   ├── threads_110.sdi
│   │       │   ├── tls_channel_stat_190.sdi
│   │       │   ├── user_defined_fun_188.sdi
│   │       │   ├── user_variables_b_176.sdi
│   │       │   ├── users_144.sdi
│   │       │   ├── variables_by_thr_183.sdi
│   │       │   └── variables_info_186.sdi
│   │       ├── private_key.pem
│   │       ├── public_key.pem
│   │       ├── server-cert.pem
│   │       ├── server-key.pem
│   │       ├── sys
│   │       │   └── sys_config.ibd
│   │       ├── undo_001
│   │       └── undo_002
│   ├── omgind.db
│   ├── rabbitmq
│   │   └── mnesia
│   │       ├── rabbit@0e915e524d2e
│   │       │   ├── DECISION_TAB.LOG
│   │       │   ├── LATEST.LOG
│   │       │   ├── cluster_nodes.config
│   │       │   ├── msg_stores
│   │       │   │   └── vhosts
│   │       │   │       └── 628WB79CIFDYO9LJI6DKMI09L
│   │       │   │           ├── msg_store_persistent
│   │       │   │           │   ├── 0.rdq
│   │       │   │           │   ├── clean.dot
│   │       │   │           │   ├── file_summary.ets
│   │       │   │           │   └── msg_store_index.ets
│   │       │   │           ├── msg_store_transient
│   │       │   │           │   ├── 0.rdq
│   │       │   │           │   ├── clean.dot
│   │       │   │           │   ├── file_summary.ets
│   │       │   │           │   └── msg_store_index.ets
│   │       │   │           └── recovery.dets
│   │       │   ├── nodes_running_at_shutdown
│   │       │   ├── quorum
│   │       │   │   └── rabbit@0e915e524d2e
│   │       │   │       ├── 00000001.wal
│   │       │   │       ├── meta.dets
│   │       │   │       └── names.dets
│   │       │   ├── rabbit_durable_exchange.DCD
│   │       │   ├── rabbit_durable_exchange.DCL
│   │       │   ├── rabbit_durable_queue.DCD
│   │       │   ├── rabbit_durable_queue.DCL
│   │       │   ├── rabbit_durable_route.DCD
│   │       │   ├── rabbit_runtime_parameters.DCD
│   │       │   ├── rabbit_runtime_parameters.DCL
│   │       │   ├── rabbit_serial
│   │       │   ├── rabbit_topic_permission.DCD
│   │       │   ├── rabbit_user.DCD
│   │       │   ├── rabbit_user_permission.DCD
│   │       │   ├── rabbit_user_permission.DCL
│   │       │   ├── rabbit_vhost.DCD
│   │       │   ├── schema.DAT
│   │       │   └── schema_version
│   │       ├── rabbit@0e915e524d2e-feature_flags
│   │       ├── rabbit@0e915e524d2e-plugins-expand
│   │       │   ├── cowboy-2.8.0
│   │       │   │   └── ebin
│   │       │   │       ├── cowboy.app
│   │       │   │       ├── cowboy.beam
│   │       │   │       ├── cowboy_app.beam
│   │       │   │       ├── cowboy_bstr.beam
│   │       │   │       ├── cowboy_children.beam
│   │       │   │       ├── cowboy_clear.beam
│   │       │   │       ├── cowboy_clock.beam
│   │       │   │       ├── cowboy_compress_h.beam
│   │       │   │       ├── cowboy_constraints.beam
│   │       │   │       ├── cowboy_handler.beam
│   │       │   │       ├── cowboy_http.beam
│   │       │   │       ├── cowboy_http2.beam
│   │       │   │       ├── cowboy_loop.beam
│   │       │   │       ├── cowboy_metrics_h.beam
│   │       │   │       ├── cowboy_middleware.beam
│   │       │   │       ├── cowboy_req.beam
│   │       │   │       ├── cowboy_rest.beam
│   │       │   │       ├── cowboy_router.beam
│   │       │   │       ├── cowboy_static.beam
│   │       │   │       ├── cowboy_stream.beam
│   │       │   │       ├── cowboy_stream_h.beam
│   │       │   │       ├── cowboy_sub_protocol.beam
│   │       │   │       ├── cowboy_sup.beam
│   │       │   │       ├── cowboy_tls.beam
│   │       │   │       ├── cowboy_tracer_h.beam
│   │       │   │       └── cowboy_websocket.beam
│   │       │   ├── cowlib-2.9.1
│   │       │   │   ├── ebin
│   │       │   │   │   ├── cow_base64url.beam
│   │       │   │   │   ├── cow_cookie.beam
│   │       │   │   │   ├── cow_date.beam
│   │       │   │   │   ├── cow_hpack.beam
│   │       │   │   │   ├── cow_http.beam
│   │       │   │   │   ├── cow_http2.beam
│   │       │   │   │   ├── cow_http2_machine.beam
│   │       │   │   │   ├── cow_http_hd.beam
│   │       │   │   │   ├── cow_http_struct_hd.beam
│   │       │   │   │   ├── cow_http_te.beam
│   │       │   │   │   ├── cow_iolists.beam
│   │       │   │   │   ├── cow_link.beam
│   │       │   │   │   ├── cow_mimetypes.beam
│   │       │   │   │   ├── cow_multipart.beam
│   │       │   │   │   ├── cow_qs.beam
│   │       │   │   │   ├── cow_spdy.beam
│   │       │   │   │   ├── cow_sse.beam
│   │       │   │   │   ├── cow_uri.beam
│   │       │   │   │   ├── cow_uri_template.beam
│   │       │   │   │   ├── cow_ws.beam
│   │       │   │   │   └── cowlib.app
│   │       │   │   └── include
│   │       │   │       ├── cow_inline.hrl
│   │       │   │       └── cow_parse.hrl
│   │       │   ├── prometheus-4.6.0
│   │       │   │   ├── ebin
│   │       │   │   │   ├── prometheus.app
│   │       │   │   │   ├── prometheus.beam
│   │       │   │   │   ├── prometheus_boolean.beam
│   │       │   │   │   ├── prometheus_buckets.beam
│   │       │   │   │   ├── prometheus_collector.beam
│   │       │   │   │   ├── prometheus_counter.beam
│   │       │   │   │   ├── prometheus_format.beam
│   │       │   │   │   ├── prometheus_gauge.beam
│   │       │   │   │   ├── prometheus_histogram.beam
│   │       │   │   │   ├── prometheus_http.beam
│   │       │   │   │   ├── prometheus_instrumenter.beam
│   │       │   │   │   ├── prometheus_metric.beam
│   │       │   │   │   ├── prometheus_metric_spec.beam
│   │       │   │   │   ├── prometheus_misc.beam
│   │       │   │   │   ├── prometheus_mnesia.beam
│   │       │   │   │   ├── prometheus_mnesia_collector.beam
│   │       │   │   │   ├── prometheus_model.beam
│   │       │   │   │   ├── prometheus_model_helpers.beam
│   │       │   │   │   ├── prometheus_protobuf_format.beam
│   │       │   │   │   ├── prometheus_registry.beam
│   │       │   │   │   ├── prometheus_summary.beam
│   │       │   │   │   ├── prometheus_sup.beam
│   │       │   │   │   ├── prometheus_test_instrumenter.beam
│   │       │   │   │   ├── prometheus_text_format.beam
│   │       │   │   │   ├── prometheus_time.beam
│   │       │   │   │   ├── prometheus_vm_dist_collector.beam
│   │       │   │   │   ├── prometheus_vm_memory_collector.beam
│   │       │   │   │   ├── prometheus_vm_msacc_collector.beam
│   │       │   │   │   ├── prometheus_vm_statistics_collector.beam
│   │       │   │   │   └── prometheus_vm_system_info_collector.beam
│   │       │   │   └── include
│   │       │   │       ├── prometheus.hrl
│   │       │   │       └── prometheus_model.hrl
│   │       │   ├── rabbitmq_management_agent-3.8.16
│   │       │   │   ├── ebin
│   │       │   │   │   ├── Elixir.RabbitMQ.CLI.Ctl.Commands.ResetStatsDbCommand.beam
│   │       │   │   │   ├── exometer_slide.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_app.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_config.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_sup.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_sup_sup.beam
│   │       │   │   │   ├── rabbit_mgmt_data.beam
│   │       │   │   │   ├── rabbit_mgmt_data_compat.beam
│   │       │   │   │   ├── rabbit_mgmt_db_handler.beam
│   │       │   │   │   ├── rabbit_mgmt_external_stats.beam
│   │       │   │   │   ├── rabbit_mgmt_ff.beam
│   │       │   │   │   ├── rabbit_mgmt_format.beam
│   │       │   │   │   ├── rabbit_mgmt_gc.beam
│   │       │   │   │   ├── rabbit_mgmt_metrics_collector.beam
│   │       │   │   │   ├── rabbit_mgmt_metrics_gc.beam
│   │       │   │   │   ├── rabbit_mgmt_storage.beam
│   │       │   │   │   └── rabbitmq_management_agent.app
│   │       │   │   ├── include
│   │       │   │   │   ├── rabbit_mgmt_agent.hrl
│   │       │   │   │   ├── rabbit_mgmt_metrics.hrl
│   │       │   │   │   └── rabbit_mgmt_records.hrl
│   │       │   │   └── priv
│   │       │   │       └── schema
│   │       │   │           └── rabbitmq_management_agent.schema
│   │       │   ├── rabbitmq_prometheus-3.8.16
│   │       │   │   ├── ebin
│   │       │   │   │   ├── prometheus_rabbitmq_core_metrics_collector.beam
│   │       │   │   │   ├── rabbit_prometheus_app.beam
│   │       │   │   │   ├── rabbit_prometheus_dispatcher.beam
│   │       │   │   │   ├── rabbit_prometheus_handler.beam
│   │       │   │   │   └── rabbitmq_prometheus.app
│   │       │   │   └── priv
│   │       │   │       └── schema
│   │       │   │           └── rabbitmq_prometheus.schema
│   │       │   └── rabbitmq_web_dispatch-3.8.16
│   │       │       └── ebin
│   │       │           ├── rabbit_cowboy_middleware.beam
│   │       │           ├── rabbit_cowboy_redirect.beam
│   │       │           ├── rabbit_cowboy_stream_h.beam
│   │       │           ├── rabbit_web_dispatch.beam
│   │       │           ├── rabbit_web_dispatch_app.beam
│   │       │           ├── rabbit_web_dispatch_listing_handler.beam
│   │       │           ├── rabbit_web_dispatch_registry.beam
│   │       │           ├── rabbit_web_dispatch_sup.beam
│   │       │           ├── rabbit_web_dispatch_util.beam
│   │       │           ├── rabbitmq_web_dispatch.app
│   │       │           ├── webmachine_log.beam
│   │       │           └── webmachine_log_handler.beam
│   │       ├── rabbit@1be3aacf8bef
│   │       │   ├── DECISION_TAB.LOG
│   │       │   ├── LATEST.LOG
│   │       │   ├── cluster_nodes.config
│   │       │   ├── msg_stores
│   │       │   │   └── vhosts
│   │       │   │       └── 628WB79CIFDYO9LJI6DKMI09L
│   │       │   │           ├── msg_store_persistent
│   │       │   │           │   ├── 0.rdq
│   │       │   │           │   ├── clean.dot
│   │       │   │           │   ├── file_summary.ets
│   │       │   │           │   └── msg_store_index.ets
│   │       │   │           ├── msg_store_transient
│   │       │   │           │   ├── 0.rdq
│   │       │   │           │   ├── clean.dot
│   │       │   │           │   ├── file_summary.ets
│   │       │   │           │   └── msg_store_index.ets
│   │       │   │           └── recovery.dets
│   │       │   ├── nodes_running_at_shutdown
│   │       │   ├── quorum
│   │       │   │   └── rabbit@1be3aacf8bef
│   │       │   │       ├── 00000001.wal
│   │       │   │       ├── meta.dets
│   │       │   │       └── names.dets
│   │       │   ├── rabbit_durable_exchange.DCD
│   │       │   ├── rabbit_durable_queue.DCD
│   │       │   ├── rabbit_durable_queue.DCL
│   │       │   ├── rabbit_durable_route.DCD
│   │       │   ├── rabbit_runtime_parameters.DCD
│   │       │   ├── rabbit_serial
│   │       │   ├── rabbit_topic_permission.DCD
│   │       │   ├── rabbit_user.DCD
│   │       │   ├── rabbit_user.DCL
│   │       │   ├── rabbit_user_permission.DCD
│   │       │   ├── rabbit_vhost.DCD
│   │       │   ├── rabbit_vhost.DCL
│   │       │   ├── schema.DAT
│   │       │   └── schema_version
│   │       ├── rabbit@1be3aacf8bef-feature_flags
│   │       ├── rabbit@1be3aacf8bef-plugins-expand
│   │       │   ├── cowboy-2.8.0
│   │       │   │   └── ebin
│   │       │   │       ├── cowboy.app
│   │       │   │       ├── cowboy.beam
│   │       │   │       ├── cowboy_app.beam
│   │       │   │       ├── cowboy_bstr.beam
│   │       │   │       ├── cowboy_children.beam
│   │       │   │       ├── cowboy_clear.beam
│   │       │   │       ├── cowboy_clock.beam
│   │       │   │       ├── cowboy_compress_h.beam
│   │       │   │       ├── cowboy_constraints.beam
│   │       │   │       ├── cowboy_handler.beam
│   │       │   │       ├── cowboy_http.beam
│   │       │   │       ├── cowboy_http2.beam
│   │       │   │       ├── cowboy_loop.beam
│   │       │   │       ├── cowboy_metrics_h.beam
│   │       │   │       ├── cowboy_middleware.beam
│   │       │   │       ├── cowboy_req.beam
│   │       │   │       ├── cowboy_rest.beam
│   │       │   │       ├── cowboy_router.beam
│   │       │   │       ├── cowboy_static.beam
│   │       │   │       ├── cowboy_stream.beam
│   │       │   │       ├── cowboy_stream_h.beam
│   │       │   │       ├── cowboy_sub_protocol.beam
│   │       │   │       ├── cowboy_sup.beam
│   │       │   │       ├── cowboy_tls.beam
│   │       │   │       ├── cowboy_tracer_h.beam
│   │       │   │       └── cowboy_websocket.beam
│   │       │   ├── cowlib-2.9.1
│   │       │   │   ├── ebin
│   │       │   │   │   ├── cow_base64url.beam
│   │       │   │   │   ├── cow_cookie.beam
│   │       │   │   │   ├── cow_date.beam
│   │       │   │   │   ├── cow_hpack.beam
│   │       │   │   │   ├── cow_http.beam
│   │       │   │   │   ├── cow_http2.beam
│   │       │   │   │   ├── cow_http2_machine.beam
│   │       │   │   │   ├── cow_http_hd.beam
│   │       │   │   │   ├── cow_http_struct_hd.beam
│   │       │   │   │   ├── cow_http_te.beam
│   │       │   │   │   ├── cow_iolists.beam
│   │       │   │   │   ├── cow_link.beam
│   │       │   │   │   ├── cow_mimetypes.beam
│   │       │   │   │   ├── cow_multipart.beam
│   │       │   │   │   ├── cow_qs.beam
│   │       │   │   │   ├── cow_spdy.beam
│   │       │   │   │   ├── cow_sse.beam
│   │       │   │   │   ├── cow_uri.beam
│   │       │   │   │   ├── cow_uri_template.beam
│   │       │   │   │   ├── cow_ws.beam
│   │       │   │   │   └── cowlib.app
│   │       │   │   └── include
│   │       │   │       ├── cow_inline.hrl
│   │       │   │       └── cow_parse.hrl
│   │       │   ├── prometheus-4.6.0
│   │       │   │   ├── ebin
│   │       │   │   │   ├── prometheus.app
│   │       │   │   │   ├── prometheus.beam
│   │       │   │   │   ├── prometheus_boolean.beam
│   │       │   │   │   ├── prometheus_buckets.beam
│   │       │   │   │   ├── prometheus_collector.beam
│   │       │   │   │   ├── prometheus_counter.beam
│   │       │   │   │   ├── prometheus_format.beam
│   │       │   │   │   ├── prometheus_gauge.beam
│   │       │   │   │   ├── prometheus_histogram.beam
│   │       │   │   │   ├── prometheus_http.beam
│   │       │   │   │   ├── prometheus_instrumenter.beam
│   │       │   │   │   ├── prometheus_metric.beam
│   │       │   │   │   ├── prometheus_metric_spec.beam
│   │       │   │   │   ├── prometheus_misc.beam
│   │       │   │   │   ├── prometheus_mnesia.beam
│   │       │   │   │   ├── prometheus_mnesia_collector.beam
│   │       │   │   │   ├── prometheus_model.beam
│   │       │   │   │   ├── prometheus_model_helpers.beam
│   │       │   │   │   ├── prometheus_protobuf_format.beam
│   │       │   │   │   ├── prometheus_registry.beam
│   │       │   │   │   ├── prometheus_summary.beam
│   │       │   │   │   ├── prometheus_sup.beam
│   │       │   │   │   ├── prometheus_test_instrumenter.beam
│   │       │   │   │   ├── prometheus_text_format.beam
│   │       │   │   │   ├── prometheus_time.beam
│   │       │   │   │   ├── prometheus_vm_dist_collector.beam
│   │       │   │   │   ├── prometheus_vm_memory_collector.beam
│   │       │   │   │   ├── prometheus_vm_msacc_collector.beam
│   │       │   │   │   ├── prometheus_vm_statistics_collector.beam
│   │       │   │   │   └── prometheus_vm_system_info_collector.beam
│   │       │   │   └── include
│   │       │   │       ├── prometheus.hrl
│   │       │   │       └── prometheus_model.hrl
│   │       │   ├── rabbitmq_management_agent-3.8.16
│   │       │   │   ├── ebin
│   │       │   │   │   ├── Elixir.RabbitMQ.CLI.Ctl.Commands.ResetStatsDbCommand.beam
│   │       │   │   │   ├── exometer_slide.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_app.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_config.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_sup.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_sup_sup.beam
│   │       │   │   │   ├── rabbit_mgmt_data.beam
│   │       │   │   │   ├── rabbit_mgmt_data_compat.beam
│   │       │   │   │   ├── rabbit_mgmt_db_handler.beam
│   │       │   │   │   ├── rabbit_mgmt_external_stats.beam
│   │       │   │   │   ├── rabbit_mgmt_ff.beam
│   │       │   │   │   ├── rabbit_mgmt_format.beam
│   │       │   │   │   ├── rabbit_mgmt_gc.beam
│   │       │   │   │   ├── rabbit_mgmt_metrics_collector.beam
│   │       │   │   │   ├── rabbit_mgmt_metrics_gc.beam
│   │       │   │   │   ├── rabbit_mgmt_storage.beam
│   │       │   │   │   └── rabbitmq_management_agent.app
│   │       │   │   ├── include
│   │       │   │   │   ├── rabbit_mgmt_agent.hrl
│   │       │   │   │   ├── rabbit_mgmt_metrics.hrl
│   │       │   │   │   └── rabbit_mgmt_records.hrl
│   │       │   │   └── priv
│   │       │   │       └── schema
│   │       │   │           └── rabbitmq_management_agent.schema
│   │       │   ├── rabbitmq_prometheus-3.8.16
│   │       │   │   ├── ebin
│   │       │   │   │   ├── prometheus_rabbitmq_core_metrics_collector.beam
│   │       │   │   │   ├── rabbit_prometheus_app.beam
│   │       │   │   │   ├── rabbit_prometheus_dispatcher.beam
│   │       │   │   │   ├── rabbit_prometheus_handler.beam
│   │       │   │   │   └── rabbitmq_prometheus.app
│   │       │   │   └── priv
│   │       │   │       └── schema
│   │       │   │           └── rabbitmq_prometheus.schema
│   │       │   └── rabbitmq_web_dispatch-3.8.16
│   │       │       └── ebin
│   │       │           ├── rabbit_cowboy_middleware.beam
│   │       │           ├── rabbit_cowboy_redirect.beam
│   │       │           ├── rabbit_cowboy_stream_h.beam
│   │       │           ├── rabbit_web_dispatch.beam
│   │       │           ├── rabbit_web_dispatch_app.beam
│   │       │           ├── rabbit_web_dispatch_listing_handler.beam
│   │       │           ├── rabbit_web_dispatch_registry.beam
│   │       │           ├── rabbit_web_dispatch_sup.beam
│   │       │           ├── rabbit_web_dispatch_util.beam
│   │       │           ├── rabbitmq_web_dispatch.app
│   │       │           ├── webmachine_log.beam
│   │       │           └── webmachine_log_handler.beam
│   │       ├── rabbit@2493aba9a177
│   │       │   ├── DECISION_TAB.LOG
│   │       │   ├── LATEST.LOG
│   │       │   ├── cluster_nodes.config
│   │       │   ├── msg_stores
│   │       │   │   └── vhosts
│   │       │   │       └── 628WB79CIFDYO9LJI6DKMI09L
│   │       │   │           ├── msg_store_persistent
│   │       │   │           │   ├── 0.rdq
│   │       │   │           │   ├── clean.dot
│   │       │   │           │   ├── file_summary.ets
│   │       │   │           │   └── msg_store_index.ets
│   │       │   │           ├── msg_store_transient
│   │       │   │           │   ├── 0.rdq
│   │       │   │           │   ├── clean.dot
│   │       │   │           │   ├── file_summary.ets
│   │       │   │           │   └── msg_store_index.ets
│   │       │   │           └── recovery.dets
│   │       │   ├── nodes_running_at_shutdown
│   │       │   ├── quorum
│   │       │   │   └── rabbit@2493aba9a177
│   │       │   │       ├── 00000001.wal
│   │       │   │       ├── meta.dets
│   │       │   │       └── names.dets
│   │       │   ├── rabbit_durable_exchange.DCD
│   │       │   ├── rabbit_durable_queue.DCD
│   │       │   ├── rabbit_durable_queue.DCL
│   │       │   ├── rabbit_durable_route.DCD
│   │       │   ├── rabbit_runtime_parameters.DCD
│   │       │   ├── rabbit_serial
│   │       │   ├── rabbit_topic_permission.DCD
│   │       │   ├── rabbit_user.DCD
│   │       │   ├── rabbit_user.DCL
│   │       │   ├── rabbit_user_permission.DCD
│   │       │   ├── rabbit_vhost.DCD
│   │       │   ├── rabbit_vhost.DCL
│   │       │   ├── schema.DAT
│   │       │   └── schema_version
│   │       ├── rabbit@2493aba9a177-feature_flags
│   │       ├── rabbit@2493aba9a177-plugins-expand
│   │       │   ├── cowboy-2.8.0
│   │       │   │   └── ebin
│   │       │   │       ├── cowboy.app
│   │       │   │       ├── cowboy.beam
│   │       │   │       ├── cowboy_app.beam
│   │       │   │       ├── cowboy_bstr.beam
│   │       │   │       ├── cowboy_children.beam
│   │       │   │       ├── cowboy_clear.beam
│   │       │   │       ├── cowboy_clock.beam
│   │       │   │       ├── cowboy_compress_h.beam
│   │       │   │       ├── cowboy_constraints.beam
│   │       │   │       ├── cowboy_handler.beam
│   │       │   │       ├── cowboy_http.beam
│   │       │   │       ├── cowboy_http2.beam
│   │       │   │       ├── cowboy_loop.beam
│   │       │   │       ├── cowboy_metrics_h.beam
│   │       │   │       ├── cowboy_middleware.beam
│   │       │   │       ├── cowboy_req.beam
│   │       │   │       ├── cowboy_rest.beam
│   │       │   │       ├── cowboy_router.beam
│   │       │   │       ├── cowboy_static.beam
│   │       │   │       ├── cowboy_stream.beam
│   │       │   │       ├── cowboy_stream_h.beam
│   │       │   │       ├── cowboy_sub_protocol.beam
│   │       │   │       ├── cowboy_sup.beam
│   │       │   │       ├── cowboy_tls.beam
│   │       │   │       ├── cowboy_tracer_h.beam
│   │       │   │       └── cowboy_websocket.beam
│   │       │   ├── cowlib-2.9.1
│   │       │   │   ├── ebin
│   │       │   │   │   ├── cow_base64url.beam
│   │       │   │   │   ├── cow_cookie.beam
│   │       │   │   │   ├── cow_date.beam
│   │       │   │   │   ├── cow_hpack.beam
│   │       │   │   │   ├── cow_http.beam
│   │       │   │   │   ├── cow_http2.beam
│   │       │   │   │   ├── cow_http2_machine.beam
│   │       │   │   │   ├── cow_http_hd.beam
│   │       │   │   │   ├── cow_http_struct_hd.beam
│   │       │   │   │   ├── cow_http_te.beam
│   │       │   │   │   ├── cow_iolists.beam
│   │       │   │   │   ├── cow_link.beam
│   │       │   │   │   ├── cow_mimetypes.beam
│   │       │   │   │   ├── cow_multipart.beam
│   │       │   │   │   ├── cow_qs.beam
│   │       │   │   │   ├── cow_spdy.beam
│   │       │   │   │   ├── cow_sse.beam
│   │       │   │   │   ├── cow_uri.beam
│   │       │   │   │   ├── cow_uri_template.beam
│   │       │   │   │   ├── cow_ws.beam
│   │       │   │   │   └── cowlib.app
│   │       │   │   └── include
│   │       │   │       ├── cow_inline.hrl
│   │       │   │       └── cow_parse.hrl
│   │       │   ├── prometheus-4.6.0
│   │       │   │   ├── ebin
│   │       │   │   │   ├── prometheus.app
│   │       │   │   │   ├── prometheus.beam
│   │       │   │   │   ├── prometheus_boolean.beam
│   │       │   │   │   ├── prometheus_buckets.beam
│   │       │   │   │   ├── prometheus_collector.beam
│   │       │   │   │   ├── prometheus_counter.beam
│   │       │   │   │   ├── prometheus_format.beam
│   │       │   │   │   ├── prometheus_gauge.beam
│   │       │   │   │   ├── prometheus_histogram.beam
│   │       │   │   │   ├── prometheus_http.beam
│   │       │   │   │   ├── prometheus_instrumenter.beam
│   │       │   │   │   ├── prometheus_metric.beam
│   │       │   │   │   ├── prometheus_metric_spec.beam
│   │       │   │   │   ├── prometheus_misc.beam
│   │       │   │   │   ├── prometheus_mnesia.beam
│   │       │   │   │   ├── prometheus_mnesia_collector.beam
│   │       │   │   │   ├── prometheus_model.beam
│   │       │   │   │   ├── prometheus_model_helpers.beam
│   │       │   │   │   ├── prometheus_protobuf_format.beam
│   │       │   │   │   ├── prometheus_registry.beam
│   │       │   │   │   ├── prometheus_summary.beam
│   │       │   │   │   ├── prometheus_sup.beam
│   │       │   │   │   ├── prometheus_test_instrumenter.beam
│   │       │   │   │   ├── prometheus_text_format.beam
│   │       │   │   │   ├── prometheus_time.beam
│   │       │   │   │   ├── prometheus_vm_dist_collector.beam
│   │       │   │   │   ├── prometheus_vm_memory_collector.beam
│   │       │   │   │   ├── prometheus_vm_msacc_collector.beam
│   │       │   │   │   ├── prometheus_vm_statistics_collector.beam
│   │       │   │   │   └── prometheus_vm_system_info_collector.beam
│   │       │   │   └── include
│   │       │   │       ├── prometheus.hrl
│   │       │   │       └── prometheus_model.hrl
│   │       │   ├── rabbitmq_management_agent-3.8.16
│   │       │   │   ├── ebin
│   │       │   │   │   ├── Elixir.RabbitMQ.CLI.Ctl.Commands.ResetStatsDbCommand.beam
│   │       │   │   │   ├── exometer_slide.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_app.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_config.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_sup.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_sup_sup.beam
│   │       │   │   │   ├── rabbit_mgmt_data.beam
│   │       │   │   │   ├── rabbit_mgmt_data_compat.beam
│   │       │   │   │   ├── rabbit_mgmt_db_handler.beam
│   │       │   │   │   ├── rabbit_mgmt_external_stats.beam
│   │       │   │   │   ├── rabbit_mgmt_ff.beam
│   │       │   │   │   ├── rabbit_mgmt_format.beam
│   │       │   │   │   ├── rabbit_mgmt_gc.beam
│   │       │   │   │   ├── rabbit_mgmt_metrics_collector.beam
│   │       │   │   │   ├── rabbit_mgmt_metrics_gc.beam
│   │       │   │   │   ├── rabbit_mgmt_storage.beam
│   │       │   │   │   └── rabbitmq_management_agent.app
│   │       │   │   ├── include
│   │       │   │   │   ├── rabbit_mgmt_agent.hrl
│   │       │   │   │   ├── rabbit_mgmt_metrics.hrl
│   │       │   │   │   └── rabbit_mgmt_records.hrl
│   │       │   │   └── priv
│   │       │   │       └── schema
│   │       │   │           └── rabbitmq_management_agent.schema
│   │       │   ├── rabbitmq_prometheus-3.8.16
│   │       │   │   ├── ebin
│   │       │   │   │   ├── prometheus_rabbitmq_core_metrics_collector.beam
│   │       │   │   │   ├── rabbit_prometheus_app.beam
│   │       │   │   │   ├── rabbit_prometheus_dispatcher.beam
│   │       │   │   │   ├── rabbit_prometheus_handler.beam
│   │       │   │   │   └── rabbitmq_prometheus.app
│   │       │   │   └── priv
│   │       │   │       └── schema
│   │       │   │           └── rabbitmq_prometheus.schema
│   │       │   └── rabbitmq_web_dispatch-3.8.16
│   │       │       └── ebin
│   │       │           ├── rabbit_cowboy_middleware.beam
│   │       │           ├── rabbit_cowboy_redirect.beam
│   │       │           ├── rabbit_cowboy_stream_h.beam
│   │       │           ├── rabbit_web_dispatch.beam
│   │       │           ├── rabbit_web_dispatch_app.beam
│   │       │           ├── rabbit_web_dispatch_listing_handler.beam
│   │       │           ├── rabbit_web_dispatch_registry.beam
│   │       │           ├── rabbit_web_dispatch_sup.beam
│   │       │           ├── rabbit_web_dispatch_util.beam
│   │       │           ├── rabbitmq_web_dispatch.app
│   │       │           ├── webmachine_log.beam
│   │       │           └── webmachine_log_handler.beam
│   │       ├── rabbit@2768dda1aa30
│   │       │   ├── DECISION_TAB.LOG
│   │       │   ├── LATEST.LOG
│   │       │   ├── cluster_nodes.config
│   │       │   ├── msg_stores
│   │       │   │   └── vhosts
│   │       │   │       └── 628WB79CIFDYO9LJI6DKMI09L
│   │       │   │           ├── msg_store_persistent
│   │       │   │           │   ├── 0.rdq
│   │       │   │           │   ├── clean.dot
│   │       │   │           │   ├── file_summary.ets
│   │       │   │           │   └── msg_store_index.ets
│   │       │   │           ├── msg_store_transient
│   │       │   │           │   ├── 0.rdq
│   │       │   │           │   ├── clean.dot
│   │       │   │           │   ├── file_summary.ets
│   │       │   │           │   └── msg_store_index.ets
│   │       │   │           └── recovery.dets
│   │       │   ├── nodes_running_at_shutdown
│   │       │   ├── quorum
│   │       │   │   └── rabbit@2768dda1aa30
│   │       │   │       ├── 00000002.wal
│   │       │   │       ├── meta.dets
│   │       │   │       └── names.dets
│   │       │   ├── rabbit_durable_exchange.DCD
│   │       │   ├── rabbit_durable_queue.DCD
│   │       │   ├── rabbit_durable_route.DCD
│   │       │   ├── rabbit_runtime_parameters.DCD
│   │       │   ├── rabbit_serial
│   │       │   ├── rabbit_topic_permission.DCD
│   │       │   ├── rabbit_user.DCD
│   │       │   ├── rabbit_user_permission.DCD
│   │       │   ├── rabbit_vhost.DCD
│   │       │   ├── schema.DAT
│   │       │   └── schema_version
│   │       ├── rabbit@2768dda1aa30-feature_flags
│   │       ├── rabbit@2768dda1aa30-plugins-expand
│   │       │   ├── cowboy-2.8.0
│   │       │   │   └── ebin
│   │       │   │       ├── cowboy.app
│   │       │   │       ├── cowboy.beam
│   │       │   │       ├── cowboy_app.beam
│   │       │   │       ├── cowboy_bstr.beam
│   │       │   │       ├── cowboy_children.beam
│   │       │   │       ├── cowboy_clear.beam
│   │       │   │       ├── cowboy_clock.beam
│   │       │   │       ├── cowboy_compress_h.beam
│   │       │   │       ├── cowboy_constraints.beam
│   │       │   │       ├── cowboy_handler.beam
│   │       │   │       ├── cowboy_http.beam
│   │       │   │       ├── cowboy_http2.beam
│   │       │   │       ├── cowboy_loop.beam
│   │       │   │       ├── cowboy_metrics_h.beam
│   │       │   │       ├── cowboy_middleware.beam
│   │       │   │       ├── cowboy_req.beam
│   │       │   │       ├── cowboy_rest.beam
│   │       │   │       ├── cowboy_router.beam
│   │       │   │       ├── cowboy_static.beam
│   │       │   │       ├── cowboy_stream.beam
│   │       │   │       ├── cowboy_stream_h.beam
│   │       │   │       ├── cowboy_sub_protocol.beam
│   │       │   │       ├── cowboy_sup.beam
│   │       │   │       ├── cowboy_tls.beam
│   │       │   │       ├── cowboy_tracer_h.beam
│   │       │   │       └── cowboy_websocket.beam
│   │       │   ├── cowlib-2.9.1
│   │       │   │   ├── ebin
│   │       │   │   │   ├── cow_base64url.beam
│   │       │   │   │   ├── cow_cookie.beam
│   │       │   │   │   ├── cow_date.beam
│   │       │   │   │   ├── cow_hpack.beam
│   │       │   │   │   ├── cow_http.beam
│   │       │   │   │   ├── cow_http2.beam
│   │       │   │   │   ├── cow_http2_machine.beam
│   │       │   │   │   ├── cow_http_hd.beam
│   │       │   │   │   ├── cow_http_struct_hd.beam
│   │       │   │   │   ├── cow_http_te.beam
│   │       │   │   │   ├── cow_iolists.beam
│   │       │   │   │   ├── cow_link.beam
│   │       │   │   │   ├── cow_mimetypes.beam
│   │       │   │   │   ├── cow_multipart.beam
│   │       │   │   │   ├── cow_qs.beam
│   │       │   │   │   ├── cow_spdy.beam
│   │       │   │   │   ├── cow_sse.beam
│   │       │   │   │   ├── cow_uri.beam
│   │       │   │   │   ├── cow_uri_template.beam
│   │       │   │   │   ├── cow_ws.beam
│   │       │   │   │   └── cowlib.app
│   │       │   │   └── include
│   │       │   │       ├── cow_inline.hrl
│   │       │   │       └── cow_parse.hrl
│   │       │   ├── prometheus-4.6.0
│   │       │   │   ├── ebin
│   │       │   │   │   ├── prometheus.app
│   │       │   │   │   ├── prometheus.beam
│   │       │   │   │   ├── prometheus_boolean.beam
│   │       │   │   │   ├── prometheus_buckets.beam
│   │       │   │   │   ├── prometheus_collector.beam
│   │       │   │   │   ├── prometheus_counter.beam
│   │       │   │   │   ├── prometheus_format.beam
│   │       │   │   │   ├── prometheus_gauge.beam
│   │       │   │   │   ├── prometheus_histogram.beam
│   │       │   │   │   ├── prometheus_http.beam
│   │       │   │   │   ├── prometheus_instrumenter.beam
│   │       │   │   │   ├── prometheus_metric.beam
│   │       │   │   │   ├── prometheus_metric_spec.beam
│   │       │   │   │   ├── prometheus_misc.beam
│   │       │   │   │   ├── prometheus_mnesia.beam
│   │       │   │   │   ├── prometheus_mnesia_collector.beam
│   │       │   │   │   ├── prometheus_model.beam
│   │       │   │   │   ├── prometheus_model_helpers.beam
│   │       │   │   │   ├── prometheus_protobuf_format.beam
│   │       │   │   │   ├── prometheus_registry.beam
│   │       │   │   │   ├── prometheus_summary.beam
│   │       │   │   │   ├── prometheus_sup.beam
│   │       │   │   │   ├── prometheus_test_instrumenter.beam
│   │       │   │   │   ├── prometheus_text_format.beam
│   │       │   │   │   ├── prometheus_time.beam
│   │       │   │   │   ├── prometheus_vm_dist_collector.beam
│   │       │   │   │   ├── prometheus_vm_memory_collector.beam
│   │       │   │   │   ├── prometheus_vm_msacc_collector.beam
│   │       │   │   │   ├── prometheus_vm_statistics_collector.beam
│   │       │   │   │   └── prometheus_vm_system_info_collector.beam
│   │       │   │   └── include
│   │       │   │       ├── prometheus.hrl
│   │       │   │       └── prometheus_model.hrl
│   │       │   ├── rabbitmq_management_agent-3.8.16
│   │       │   │   ├── ebin
│   │       │   │   │   ├── Elixir.RabbitMQ.CLI.Ctl.Commands.ResetStatsDbCommand.beam
│   │       │   │   │   ├── exometer_slide.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_app.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_config.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_sup.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_sup_sup.beam
│   │       │   │   │   ├── rabbit_mgmt_data.beam
│   │       │   │   │   ├── rabbit_mgmt_data_compat.beam
│   │       │   │   │   ├── rabbit_mgmt_db_handler.beam
│   │       │   │   │   ├── rabbit_mgmt_external_stats.beam
│   │       │   │   │   ├── rabbit_mgmt_ff.beam
│   │       │   │   │   ├── rabbit_mgmt_format.beam
│   │       │   │   │   ├── rabbit_mgmt_gc.beam
│   │       │   │   │   ├── rabbit_mgmt_metrics_collector.beam
│   │       │   │   │   ├── rabbit_mgmt_metrics_gc.beam
│   │       │   │   │   ├── rabbit_mgmt_storage.beam
│   │       │   │   │   └── rabbitmq_management_agent.app
│   │       │   │   ├── include
│   │       │   │   │   ├── rabbit_mgmt_agent.hrl
│   │       │   │   │   ├── rabbit_mgmt_metrics.hrl
│   │       │   │   │   └── rabbit_mgmt_records.hrl
│   │       │   │   └── priv
│   │       │   │       └── schema
│   │       │   │           └── rabbitmq_management_agent.schema
│   │       │   ├── rabbitmq_prometheus-3.8.16
│   │       │   │   ├── ebin
│   │       │   │   │   ├── prometheus_rabbitmq_core_metrics_collector.beam
│   │       │   │   │   ├── rabbit_prometheus_app.beam
│   │       │   │   │   ├── rabbit_prometheus_dispatcher.beam
│   │       │   │   │   ├── rabbit_prometheus_handler.beam
│   │       │   │   │   └── rabbitmq_prometheus.app
│   │       │   │   └── priv
│   │       │   │       └── schema
│   │       │   │           └── rabbitmq_prometheus.schema
│   │       │   └── rabbitmq_web_dispatch-3.8.16
│   │       │       └── ebin
│   │       │           ├── rabbit_cowboy_middleware.beam
│   │       │           ├── rabbit_cowboy_redirect.beam
│   │       │           ├── rabbit_cowboy_stream_h.beam
│   │       │           ├── rabbit_web_dispatch.beam
│   │       │           ├── rabbit_web_dispatch_app.beam
│   │       │           ├── rabbit_web_dispatch_listing_handler.beam
│   │       │           ├── rabbit_web_dispatch_registry.beam
│   │       │           ├── rabbit_web_dispatch_sup.beam
│   │       │           ├── rabbit_web_dispatch_util.beam
│   │       │           ├── rabbitmq_web_dispatch.app
│   │       │           ├── webmachine_log.beam
│   │       │           └── webmachine_log_handler.beam
│   │       ├── rabbit@2c4a131d3bd5
│   │       │   ├── DECISION_TAB.LOG
│   │       │   ├── LATEST.LOG
│   │       │   ├── cluster_nodes.config
│   │       │   ├── msg_stores
│   │       │   │   └── vhosts
│   │       │   │       └── 628WB79CIFDYO9LJI6DKMI09L
│   │       │   │           ├── msg_store_persistent
│   │       │   │           │   ├── 0.rdq
│   │       │   │           │   ├── clean.dot
│   │       │   │           │   ├── file_summary.ets
│   │       │   │           │   └── msg_store_index.ets
│   │       │   │           ├── msg_store_transient
│   │       │   │           │   ├── 0.rdq
│   │       │   │           │   ├── clean.dot
│   │       │   │           │   ├── file_summary.ets
│   │       │   │           │   └── msg_store_index.ets
│   │       │   │           └── recovery.dets
│   │       │   ├── nodes_running_at_shutdown
│   │       │   ├── quorum
│   │       │   │   └── rabbit@2c4a131d3bd5
│   │       │   │       ├── 00000001.wal
│   │       │   │       ├── meta.dets
│   │       │   │       └── names.dets
│   │       │   ├── rabbit_durable_exchange.DCD
│   │       │   ├── rabbit_durable_exchange.DCL
│   │       │   ├── rabbit_durable_queue.DCD
│   │       │   ├── rabbit_durable_queue.DCL
│   │       │   ├── rabbit_durable_route.DCD
│   │       │   ├── rabbit_runtime_parameters.DCD
│   │       │   ├── rabbit_runtime_parameters.DCL
│   │       │   ├── rabbit_serial
│   │       │   ├── rabbit_topic_permission.DCD
│   │       │   ├── rabbit_user.DCD
│   │       │   ├── rabbit_user_permission.DCD
│   │       │   ├── rabbit_user_permission.DCL
│   │       │   ├── rabbit_vhost.DCD
│   │       │   ├── schema.DAT
│   │       │   └── schema_version
│   │       ├── rabbit@2c4a131d3bd5-feature_flags
│   │       ├── rabbit@2c4a131d3bd5-plugins-expand
│   │       │   ├── cowboy-2.8.0
│   │       │   │   └── ebin
│   │       │   │       ├── cowboy.app
│   │       │   │       ├── cowboy.beam
│   │       │   │       ├── cowboy_app.beam
│   │       │   │       ├── cowboy_bstr.beam
│   │       │   │       ├── cowboy_children.beam
│   │       │   │       ├── cowboy_clear.beam
│   │       │   │       ├── cowboy_clock.beam
│   │       │   │       ├── cowboy_compress_h.beam
│   │       │   │       ├── cowboy_constraints.beam
│   │       │   │       ├── cowboy_handler.beam
│   │       │   │       ├── cowboy_http.beam
│   │       │   │       ├── cowboy_http2.beam
│   │       │   │       ├── cowboy_loop.beam
│   │       │   │       ├── cowboy_metrics_h.beam
│   │       │   │       ├── cowboy_middleware.beam
│   │       │   │       ├── cowboy_req.beam
│   │       │   │       ├── cowboy_rest.beam
│   │       │   │       ├── cowboy_router.beam
│   │       │   │       ├── cowboy_static.beam
│   │       │   │       ├── cowboy_stream.beam
│   │       │   │       ├── cowboy_stream_h.beam
│   │       │   │       ├── cowboy_sub_protocol.beam
│   │       │   │       ├── cowboy_sup.beam
│   │       │   │       ├── cowboy_tls.beam
│   │       │   │       ├── cowboy_tracer_h.beam
│   │       │   │       └── cowboy_websocket.beam
│   │       │   ├── cowlib-2.9.1
│   │       │   │   ├── ebin
│   │       │   │   │   ├── cow_base64url.beam
│   │       │   │   │   ├── cow_cookie.beam
│   │       │   │   │   ├── cow_date.beam
│   │       │   │   │   ├── cow_hpack.beam
│   │       │   │   │   ├── cow_http.beam
│   │       │   │   │   ├── cow_http2.beam
│   │       │   │   │   ├── cow_http2_machine.beam
│   │       │   │   │   ├── cow_http_hd.beam
│   │       │   │   │   ├── cow_http_struct_hd.beam
│   │       │   │   │   ├── cow_http_te.beam
│   │       │   │   │   ├── cow_iolists.beam
│   │       │   │   │   ├── cow_link.beam
│   │       │   │   │   ├── cow_mimetypes.beam
│   │       │   │   │   ├── cow_multipart.beam
│   │       │   │   │   ├── cow_qs.beam
│   │       │   │   │   ├── cow_spdy.beam
│   │       │   │   │   ├── cow_sse.beam
│   │       │   │   │   ├── cow_uri.beam
│   │       │   │   │   ├── cow_uri_template.beam
│   │       │   │   │   ├── cow_ws.beam
│   │       │   │   │   └── cowlib.app
│   │       │   │   └── include
│   │       │   │       ├── cow_inline.hrl
│   │       │   │       └── cow_parse.hrl
│   │       │   ├── prometheus-4.6.0
│   │       │   │   ├── ebin
│   │       │   │   │   ├── prometheus.app
│   │       │   │   │   ├── prometheus.beam
│   │       │   │   │   ├── prometheus_boolean.beam
│   │       │   │   │   ├── prometheus_buckets.beam
│   │       │   │   │   ├── prometheus_collector.beam
│   │       │   │   │   ├── prometheus_counter.beam
│   │       │   │   │   ├── prometheus_format.beam
│   │       │   │   │   ├── prometheus_gauge.beam
│   │       │   │   │   ├── prometheus_histogram.beam
│   │       │   │   │   ├── prometheus_http.beam
│   │       │   │   │   ├── prometheus_instrumenter.beam
│   │       │   │   │   ├── prometheus_metric.beam
│   │       │   │   │   ├── prometheus_metric_spec.beam
│   │       │   │   │   ├── prometheus_misc.beam
│   │       │   │   │   ├── prometheus_mnesia.beam
│   │       │   │   │   ├── prometheus_mnesia_collector.beam
│   │       │   │   │   ├── prometheus_model.beam
│   │       │   │   │   ├── prometheus_model_helpers.beam
│   │       │   │   │   ├── prometheus_protobuf_format.beam
│   │       │   │   │   ├── prometheus_registry.beam
│   │       │   │   │   ├── prometheus_summary.beam
│   │       │   │   │   ├── prometheus_sup.beam
│   │       │   │   │   ├── prometheus_test_instrumenter.beam
│   │       │   │   │   ├── prometheus_text_format.beam
│   │       │   │   │   ├── prometheus_time.beam
│   │       │   │   │   ├── prometheus_vm_dist_collector.beam
│   │       │   │   │   ├── prometheus_vm_memory_collector.beam
│   │       │   │   │   ├── prometheus_vm_msacc_collector.beam
│   │       │   │   │   ├── prometheus_vm_statistics_collector.beam
│   │       │   │   │   └── prometheus_vm_system_info_collector.beam
│   │       │   │   └── include
│   │       │   │       ├── prometheus.hrl
│   │       │   │       └── prometheus_model.hrl
│   │       │   ├── rabbitmq_management_agent-3.8.16
│   │       │   │   ├── ebin
│   │       │   │   │   ├── Elixir.RabbitMQ.CLI.Ctl.Commands.ResetStatsDbCommand.beam
│   │       │   │   │   ├── exometer_slide.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_app.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_config.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_sup.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_sup_sup.beam
│   │       │   │   │   ├── rabbit_mgmt_data.beam
│   │       │   │   │   ├── rabbit_mgmt_data_compat.beam
│   │       │   │   │   ├── rabbit_mgmt_db_handler.beam
│   │       │   │   │   ├── rabbit_mgmt_external_stats.beam
│   │       │   │   │   ├── rabbit_mgmt_ff.beam
│   │       │   │   │   ├── rabbit_mgmt_format.beam
│   │       │   │   │   ├── rabbit_mgmt_gc.beam
│   │       │   │   │   ├── rabbit_mgmt_metrics_collector.beam
│   │       │   │   │   ├── rabbit_mgmt_metrics_gc.beam
│   │       │   │   │   ├── rabbit_mgmt_storage.beam
│   │       │   │   │   └── rabbitmq_management_agent.app
│   │       │   │   ├── include
│   │       │   │   │   ├── rabbit_mgmt_agent.hrl
│   │       │   │   │   ├── rabbit_mgmt_metrics.hrl
│   │       │   │   │   └── rabbit_mgmt_records.hrl
│   │       │   │   └── priv
│   │       │   │       └── schema
│   │       │   │           └── rabbitmq_management_agent.schema
│   │       │   ├── rabbitmq_prometheus-3.8.16
│   │       │   │   ├── ebin
│   │       │   │   │   ├── prometheus_rabbitmq_core_metrics_collector.beam
│   │       │   │   │   ├── rabbit_prometheus_app.beam
│   │       │   │   │   ├── rabbit_prometheus_dispatcher.beam
│   │       │   │   │   ├── rabbit_prometheus_handler.beam
│   │       │   │   │   └── rabbitmq_prometheus.app
│   │       │   │   └── priv
│   │       │   │       └── schema
│   │       │   │           └── rabbitmq_prometheus.schema
│   │       │   └── rabbitmq_web_dispatch-3.8.16
│   │       │       └── ebin
│   │       │           ├── rabbit_cowboy_middleware.beam
│   │       │           ├── rabbit_cowboy_redirect.beam
│   │       │           ├── rabbit_cowboy_stream_h.beam
│   │       │           ├── rabbit_web_dispatch.beam
│   │       │           ├── rabbit_web_dispatch_app.beam
│   │       │           ├── rabbit_web_dispatch_listing_handler.beam
│   │       │           ├── rabbit_web_dispatch_registry.beam
│   │       │           ├── rabbit_web_dispatch_sup.beam
│   │       │           ├── rabbit_web_dispatch_util.beam
│   │       │           ├── rabbitmq_web_dispatch.app
│   │       │           ├── webmachine_log.beam
│   │       │           └── webmachine_log_handler.beam
│   │       ├── rabbit@6587544dab80
│   │       │   ├── DECISION_TAB.LOG
│   │       │   ├── LATEST.LOG
│   │       │   ├── cluster_nodes.config
│   │       │   ├── msg_stores
│   │       │   │   └── vhosts
│   │       │   │       └── 628WB79CIFDYO9LJI6DKMI09L
│   │       │   │           ├── msg_store_persistent
│   │       │   │           │   ├── 0.rdq
│   │       │   │           │   ├── clean.dot
│   │       │   │           │   ├── file_summary.ets
│   │       │   │           │   └── msg_store_index.ets
│   │       │   │           ├── msg_store_transient
│   │       │   │           │   ├── 0.rdq
│   │       │   │           │   ├── clean.dot
│   │       │   │           │   ├── file_summary.ets
│   │       │   │           │   └── msg_store_index.ets
│   │       │   │           └── recovery.dets
│   │       │   ├── nodes_running_at_shutdown
│   │       │   ├── quorum
│   │       │   │   └── rabbit@6587544dab80
│   │       │   │       ├── 00000001.wal
│   │       │   │       ├── meta.dets
│   │       │   │       └── names.dets
│   │       │   ├── rabbit_durable_exchange.DCD
│   │       │   ├── rabbit_durable_exchange.DCL
│   │       │   ├── rabbit_durable_queue.DCD
│   │       │   ├── rabbit_durable_queue.DCL
│   │       │   ├── rabbit_durable_route.DCD
│   │       │   ├── rabbit_runtime_parameters.DCD
│   │       │   ├── rabbit_runtime_parameters.DCL
│   │       │   ├── rabbit_serial
│   │       │   ├── rabbit_topic_permission.DCD
│   │       │   ├── rabbit_user.DCD
│   │       │   ├── rabbit_user_permission.DCD
│   │       │   ├── rabbit_user_permission.DCL
│   │       │   ├── rabbit_vhost.DCD
│   │       │   ├── schema.DAT
│   │       │   └── schema_version
│   │       ├── rabbit@6587544dab80-feature_flags
│   │       ├── rabbit@6587544dab80-plugins-expand
│   │       │   ├── cowboy-2.8.0
│   │       │   │   └── ebin
│   │       │   │       ├── cowboy.app
│   │       │   │       ├── cowboy.beam
│   │       │   │       ├── cowboy_app.beam
│   │       │   │       ├── cowboy_bstr.beam
│   │       │   │       ├── cowboy_children.beam
│   │       │   │       ├── cowboy_clear.beam
│   │       │   │       ├── cowboy_clock.beam
│   │       │   │       ├── cowboy_compress_h.beam
│   │       │   │       ├── cowboy_constraints.beam
│   │       │   │       ├── cowboy_handler.beam
│   │       │   │       ├── cowboy_http.beam
│   │       │   │       ├── cowboy_http2.beam
│   │       │   │       ├── cowboy_loop.beam
│   │       │   │       ├── cowboy_metrics_h.beam
│   │       │   │       ├── cowboy_middleware.beam
│   │       │   │       ├── cowboy_req.beam
│   │       │   │       ├── cowboy_rest.beam
│   │       │   │       ├── cowboy_router.beam
│   │       │   │       ├── cowboy_static.beam
│   │       │   │       ├── cowboy_stream.beam
│   │       │   │       ├── cowboy_stream_h.beam
│   │       │   │       ├── cowboy_sub_protocol.beam
│   │       │   │       ├── cowboy_sup.beam
│   │       │   │       ├── cowboy_tls.beam
│   │       │   │       ├── cowboy_tracer_h.beam
│   │       │   │       └── cowboy_websocket.beam
│   │       │   ├── cowlib-2.9.1
│   │       │   │   ├── ebin
│   │       │   │   │   ├── cow_base64url.beam
│   │       │   │   │   ├── cow_cookie.beam
│   │       │   │   │   ├── cow_date.beam
│   │       │   │   │   ├── cow_hpack.beam
│   │       │   │   │   ├── cow_http.beam
│   │       │   │   │   ├── cow_http2.beam
│   │       │   │   │   ├── cow_http2_machine.beam
│   │       │   │   │   ├── cow_http_hd.beam
│   │       │   │   │   ├── cow_http_struct_hd.beam
│   │       │   │   │   ├── cow_http_te.beam
│   │       │   │   │   ├── cow_iolists.beam
│   │       │   │   │   ├── cow_link.beam
│   │       │   │   │   ├── cow_mimetypes.beam
│   │       │   │   │   ├── cow_multipart.beam
│   │       │   │   │   ├── cow_qs.beam
│   │       │   │   │   ├── cow_spdy.beam
│   │       │   │   │   ├── cow_sse.beam
│   │       │   │   │   ├── cow_uri.beam
│   │       │   │   │   ├── cow_uri_template.beam
│   │       │   │   │   ├── cow_ws.beam
│   │       │   │   │   └── cowlib.app
│   │       │   │   └── include
│   │       │   │       ├── cow_inline.hrl
│   │       │   │       └── cow_parse.hrl
│   │       │   ├── prometheus-4.6.0
│   │       │   │   ├── ebin
│   │       │   │   │   ├── prometheus.app
│   │       │   │   │   ├── prometheus.beam
│   │       │   │   │   ├── prometheus_boolean.beam
│   │       │   │   │   ├── prometheus_buckets.beam
│   │       │   │   │   ├── prometheus_collector.beam
│   │       │   │   │   ├── prometheus_counter.beam
│   │       │   │   │   ├── prometheus_format.beam
│   │       │   │   │   ├── prometheus_gauge.beam
│   │       │   │   │   ├── prometheus_histogram.beam
│   │       │   │   │   ├── prometheus_http.beam
│   │       │   │   │   ├── prometheus_instrumenter.beam
│   │       │   │   │   ├── prometheus_metric.beam
│   │       │   │   │   ├── prometheus_metric_spec.beam
│   │       │   │   │   ├── prometheus_misc.beam
│   │       │   │   │   ├── prometheus_mnesia.beam
│   │       │   │   │   ├── prometheus_mnesia_collector.beam
│   │       │   │   │   ├── prometheus_model.beam
│   │       │   │   │   ├── prometheus_model_helpers.beam
│   │       │   │   │   ├── prometheus_protobuf_format.beam
│   │       │   │   │   ├── prometheus_registry.beam
│   │       │   │   │   ├── prometheus_summary.beam
│   │       │   │   │   ├── prometheus_sup.beam
│   │       │   │   │   ├── prometheus_test_instrumenter.beam
│   │       │   │   │   ├── prometheus_text_format.beam
│   │       │   │   │   ├── prometheus_time.beam
│   │       │   │   │   ├── prometheus_vm_dist_collector.beam
│   │       │   │   │   ├── prometheus_vm_memory_collector.beam
│   │       │   │   │   ├── prometheus_vm_msacc_collector.beam
│   │       │   │   │   ├── prometheus_vm_statistics_collector.beam
│   │       │   │   │   └── prometheus_vm_system_info_collector.beam
│   │       │   │   └── include
│   │       │   │       ├── prometheus.hrl
│   │       │   │       └── prometheus_model.hrl
│   │       │   ├── rabbitmq_management_agent-3.8.16
│   │       │   │   ├── ebin
│   │       │   │   │   ├── Elixir.RabbitMQ.CLI.Ctl.Commands.ResetStatsDbCommand.beam
│   │       │   │   │   ├── exometer_slide.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_app.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_config.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_sup.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_sup_sup.beam
│   │       │   │   │   ├── rabbit_mgmt_data.beam
│   │       │   │   │   ├── rabbit_mgmt_data_compat.beam
│   │       │   │   │   ├── rabbit_mgmt_db_handler.beam
│   │       │   │   │   ├── rabbit_mgmt_external_stats.beam
│   │       │   │   │   ├── rabbit_mgmt_ff.beam
│   │       │   │   │   ├── rabbit_mgmt_format.beam
│   │       │   │   │   ├── rabbit_mgmt_gc.beam
│   │       │   │   │   ├── rabbit_mgmt_metrics_collector.beam
│   │       │   │   │   ├── rabbit_mgmt_metrics_gc.beam
│   │       │   │   │   ├── rabbit_mgmt_storage.beam
│   │       │   │   │   └── rabbitmq_management_agent.app
│   │       │   │   ├── include
│   │       │   │   │   ├── rabbit_mgmt_agent.hrl
│   │       │   │   │   ├── rabbit_mgmt_metrics.hrl
│   │       │   │   │   └── rabbit_mgmt_records.hrl
│   │       │   │   └── priv
│   │       │   │       └── schema
│   │       │   │           └── rabbitmq_management_agent.schema
│   │       │   ├── rabbitmq_prometheus-3.8.16
│   │       │   │   ├── ebin
│   │       │   │   │   ├── prometheus_rabbitmq_core_metrics_collector.beam
│   │       │   │   │   ├── rabbit_prometheus_app.beam
│   │       │   │   │   ├── rabbit_prometheus_dispatcher.beam
│   │       │   │   │   ├── rabbit_prometheus_handler.beam
│   │       │   │   │   └── rabbitmq_prometheus.app
│   │       │   │   └── priv
│   │       │   │       └── schema
│   │       │   │           └── rabbitmq_prometheus.schema
│   │       │   └── rabbitmq_web_dispatch-3.8.16
│   │       │       └── ebin
│   │       │           ├── rabbit_cowboy_middleware.beam
│   │       │           ├── rabbit_cowboy_redirect.beam
│   │       │           ├── rabbit_cowboy_stream_h.beam
│   │       │           ├── rabbit_web_dispatch.beam
│   │       │           ├── rabbit_web_dispatch_app.beam
│   │       │           ├── rabbit_web_dispatch_listing_handler.beam
│   │       │           ├── rabbit_web_dispatch_registry.beam
│   │       │           ├── rabbit_web_dispatch_sup.beam
│   │       │           ├── rabbit_web_dispatch_util.beam
│   │       │           ├── rabbitmq_web_dispatch.app
│   │       │           ├── webmachine_log.beam
│   │       │           └── webmachine_log_handler.beam
│   │       ├── rabbit@6a4a614ba284
│   │       │   ├── DECISION_TAB.LOG
│   │       │   ├── LATEST.LOG
│   │       │   ├── cluster_nodes.config
│   │       │   ├── msg_stores
│   │       │   │   └── vhosts
│   │       │   │       └── 628WB79CIFDYO9LJI6DKMI09L
│   │       │   │           ├── msg_store_persistent
│   │       │   │           │   ├── 0.rdq
│   │       │   │           │   ├── clean.dot
│   │       │   │           │   ├── file_summary.ets
│   │       │   │           │   └── msg_store_index.ets
│   │       │   │           ├── msg_store_transient
│   │       │   │           │   ├── 0.rdq
│   │       │   │           │   ├── clean.dot
│   │       │   │           │   ├── file_summary.ets
│   │       │   │           │   └── msg_store_index.ets
│   │       │   │           └── recovery.dets
│   │       │   ├── nodes_running_at_shutdown
│   │       │   ├── quorum
│   │       │   │   └── rabbit@6a4a614ba284
│   │       │   │       ├── 00000001.wal
│   │       │   │       ├── meta.dets
│   │       │   │       └── names.dets
│   │       │   ├── rabbit_durable_exchange.DCD
│   │       │   ├── rabbit_durable_exchange.DCL
│   │       │   ├── rabbit_durable_queue.DCD
│   │       │   ├── rabbit_durable_queue.DCL
│   │       │   ├── rabbit_durable_route.DCD
│   │       │   ├── rabbit_runtime_parameters.DCD
│   │       │   ├── rabbit_runtime_parameters.DCL
│   │       │   ├── rabbit_serial
│   │       │   ├── rabbit_topic_permission.DCD
│   │       │   ├── rabbit_user.DCD
│   │       │   ├── rabbit_user_permission.DCD
│   │       │   ├── rabbit_user_permission.DCL
│   │       │   ├── rabbit_vhost.DCD
│   │       │   ├── schema.DAT
│   │       │   └── schema_version
│   │       ├── rabbit@6a4a614ba284-feature_flags
│   │       ├── rabbit@6a4a614ba284-plugins-expand
│   │       │   ├── cowboy-2.8.0
│   │       │   │   └── ebin
│   │       │   │       ├── cowboy.app
│   │       │   │       ├── cowboy.beam
│   │       │   │       ├── cowboy_app.beam
│   │       │   │       ├── cowboy_bstr.beam
│   │       │   │       ├── cowboy_children.beam
│   │       │   │       ├── cowboy_clear.beam
│   │       │   │       ├── cowboy_clock.beam
│   │       │   │       ├── cowboy_compress_h.beam
│   │       │   │       ├── cowboy_constraints.beam
│   │       │   │       ├── cowboy_handler.beam
│   │       │   │       ├── cowboy_http.beam
│   │       │   │       ├── cowboy_http2.beam
│   │       │   │       ├── cowboy_loop.beam
│   │       │   │       ├── cowboy_metrics_h.beam
│   │       │   │       ├── cowboy_middleware.beam
│   │       │   │       ├── cowboy_req.beam
│   │       │   │       ├── cowboy_rest.beam
│   │       │   │       ├── cowboy_router.beam
│   │       │   │       ├── cowboy_static.beam
│   │       │   │       ├── cowboy_stream.beam
│   │       │   │       ├── cowboy_stream_h.beam
│   │       │   │       ├── cowboy_sub_protocol.beam
│   │       │   │       ├── cowboy_sup.beam
│   │       │   │       ├── cowboy_tls.beam
│   │       │   │       ├── cowboy_tracer_h.beam
│   │       │   │       └── cowboy_websocket.beam
│   │       │   ├── cowlib-2.9.1
│   │       │   │   ├── ebin
│   │       │   │   │   ├── cow_base64url.beam
│   │       │   │   │   ├── cow_cookie.beam
│   │       │   │   │   ├── cow_date.beam
│   │       │   │   │   ├── cow_hpack.beam
│   │       │   │   │   ├── cow_http.beam
│   │       │   │   │   ├── cow_http2.beam
│   │       │   │   │   ├── cow_http2_machine.beam
│   │       │   │   │   ├── cow_http_hd.beam
│   │       │   │   │   ├── cow_http_struct_hd.beam
│   │       │   │   │   ├── cow_http_te.beam
│   │       │   │   │   ├── cow_iolists.beam
│   │       │   │   │   ├── cow_link.beam
│   │       │   │   │   ├── cow_mimetypes.beam
│   │       │   │   │   ├── cow_multipart.beam
│   │       │   │   │   ├── cow_qs.beam
│   │       │   │   │   ├── cow_spdy.beam
│   │       │   │   │   ├── cow_sse.beam
│   │       │   │   │   ├── cow_uri.beam
│   │       │   │   │   ├── cow_uri_template.beam
│   │       │   │   │   ├── cow_ws.beam
│   │       │   │   │   └── cowlib.app
│   │       │   │   └── include
│   │       │   │       ├── cow_inline.hrl
│   │       │   │       └── cow_parse.hrl
│   │       │   ├── prometheus-4.6.0
│   │       │   │   ├── ebin
│   │       │   │   │   ├── prometheus.app
│   │       │   │   │   ├── prometheus.beam
│   │       │   │   │   ├── prometheus_boolean.beam
│   │       │   │   │   ├── prometheus_buckets.beam
│   │       │   │   │   ├── prometheus_collector.beam
│   │       │   │   │   ├── prometheus_counter.beam
│   │       │   │   │   ├── prometheus_format.beam
│   │       │   │   │   ├── prometheus_gauge.beam
│   │       │   │   │   ├── prometheus_histogram.beam
│   │       │   │   │   ├── prometheus_http.beam
│   │       │   │   │   ├── prometheus_instrumenter.beam
│   │       │   │   │   ├── prometheus_metric.beam
│   │       │   │   │   ├── prometheus_metric_spec.beam
│   │       │   │   │   ├── prometheus_misc.beam
│   │       │   │   │   ├── prometheus_mnesia.beam
│   │       │   │   │   ├── prometheus_mnesia_collector.beam
│   │       │   │   │   ├── prometheus_model.beam
│   │       │   │   │   ├── prometheus_model_helpers.beam
│   │       │   │   │   ├── prometheus_protobuf_format.beam
│   │       │   │   │   ├── prometheus_registry.beam
│   │       │   │   │   ├── prometheus_summary.beam
│   │       │   │   │   ├── prometheus_sup.beam
│   │       │   │   │   ├── prometheus_test_instrumenter.beam
│   │       │   │   │   ├── prometheus_text_format.beam
│   │       │   │   │   ├── prometheus_time.beam
│   │       │   │   │   ├── prometheus_vm_dist_collector.beam
│   │       │   │   │   ├── prometheus_vm_memory_collector.beam
│   │       │   │   │   ├── prometheus_vm_msacc_collector.beam
│   │       │   │   │   ├── prometheus_vm_statistics_collector.beam
│   │       │   │   │   └── prometheus_vm_system_info_collector.beam
│   │       │   │   └── include
│   │       │   │       ├── prometheus.hrl
│   │       │   │       └── prometheus_model.hrl
│   │       │   ├── rabbitmq_management_agent-3.8.16
│   │       │   │   ├── ebin
│   │       │   │   │   ├── Elixir.RabbitMQ.CLI.Ctl.Commands.ResetStatsDbCommand.beam
│   │       │   │   │   ├── exometer_slide.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_app.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_config.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_sup.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_sup_sup.beam
│   │       │   │   │   ├── rabbit_mgmt_data.beam
│   │       │   │   │   ├── rabbit_mgmt_data_compat.beam
│   │       │   │   │   ├── rabbit_mgmt_db_handler.beam
│   │       │   │   │   ├── rabbit_mgmt_external_stats.beam
│   │       │   │   │   ├── rabbit_mgmt_ff.beam
│   │       │   │   │   ├── rabbit_mgmt_format.beam
│   │       │   │   │   ├── rabbit_mgmt_gc.beam
│   │       │   │   │   ├── rabbit_mgmt_metrics_collector.beam
│   │       │   │   │   ├── rabbit_mgmt_metrics_gc.beam
│   │       │   │   │   ├── rabbit_mgmt_storage.beam
│   │       │   │   │   └── rabbitmq_management_agent.app
│   │       │   │   ├── include
│   │       │   │   │   ├── rabbit_mgmt_agent.hrl
│   │       │   │   │   ├── rabbit_mgmt_metrics.hrl
│   │       │   │   │   └── rabbit_mgmt_records.hrl
│   │       │   │   └── priv
│   │       │   │       └── schema
│   │       │   │           └── rabbitmq_management_agent.schema
│   │       │   ├── rabbitmq_prometheus-3.8.16
│   │       │   │   ├── ebin
│   │       │   │   │   ├── prometheus_rabbitmq_core_metrics_collector.beam
│   │       │   │   │   ├── rabbit_prometheus_app.beam
│   │       │   │   │   ├── rabbit_prometheus_dispatcher.beam
│   │       │   │   │   ├── rabbit_prometheus_handler.beam
│   │       │   │   │   └── rabbitmq_prometheus.app
│   │       │   │   └── priv
│   │       │   │       └── schema
│   │       │   │           └── rabbitmq_prometheus.schema
│   │       │   └── rabbitmq_web_dispatch-3.8.16
│   │       │       └── ebin
│   │       │           ├── rabbit_cowboy_middleware.beam
│   │       │           ├── rabbit_cowboy_redirect.beam
│   │       │           ├── rabbit_cowboy_stream_h.beam
│   │       │           ├── rabbit_web_dispatch.beam
│   │       │           ├── rabbit_web_dispatch_app.beam
│   │       │           ├── rabbit_web_dispatch_listing_handler.beam
│   │       │           ├── rabbit_web_dispatch_registry.beam
│   │       │           ├── rabbit_web_dispatch_sup.beam
│   │       │           ├── rabbit_web_dispatch_util.beam
│   │       │           ├── rabbitmq_web_dispatch.app
│   │       │           ├── webmachine_log.beam
│   │       │           └── webmachine_log_handler.beam
│   │       ├── rabbit@6aa85d289c17
│   │       │   ├── DECISION_TAB.LOG
│   │       │   ├── LATEST.LOG
│   │       │   ├── cluster_nodes.config
│   │       │   ├── msg_stores
│   │       │   │   └── vhosts
│   │       │   │       └── 628WB79CIFDYO9LJI6DKMI09L
│   │       │   │           ├── msg_store_persistent
│   │       │   │           │   ├── 0.rdq
│   │       │   │           │   ├── clean.dot
│   │       │   │           │   ├── file_summary.ets
│   │       │   │           │   └── msg_store_index.ets
│   │       │   │           ├── msg_store_transient
│   │       │   │           │   ├── 0.rdq
│   │       │   │           │   ├── clean.dot
│   │       │   │           │   ├── file_summary.ets
│   │       │   │           │   └── msg_store_index.ets
│   │       │   │           └── recovery.dets
│   │       │   ├── nodes_running_at_shutdown
│   │       │   ├── quorum
│   │       │   │   └── rabbit@6aa85d289c17
│   │       │   │       ├── 00000001.wal
│   │       │   │       ├── meta.dets
│   │       │   │       └── names.dets
│   │       │   ├── rabbit_durable_exchange.DCD
│   │       │   ├── rabbit_durable_exchange.DCL
│   │       │   ├── rabbit_durable_queue.DCD
│   │       │   ├── rabbit_durable_queue.DCL
│   │       │   ├── rabbit_durable_route.DCD
│   │       │   ├── rabbit_runtime_parameters.DCD
│   │       │   ├── rabbit_runtime_parameters.DCL
│   │       │   ├── rabbit_serial
│   │       │   ├── rabbit_topic_permission.DCD
│   │       │   ├── rabbit_user.DCD
│   │       │   ├── rabbit_user_permission.DCD
│   │       │   ├── rabbit_user_permission.DCL
│   │       │   ├── rabbit_vhost.DCD
│   │       │   ├── schema.DAT
│   │       │   └── schema_version
│   │       ├── rabbit@6aa85d289c17-feature_flags
│   │       ├── rabbit@6aa85d289c17-plugins-expand
│   │       │   ├── cowboy-2.8.0
│   │       │   │   └── ebin
│   │       │   │       ├── cowboy.app
│   │       │   │       ├── cowboy.beam
│   │       │   │       ├── cowboy_app.beam
│   │       │   │       ├── cowboy_bstr.beam
│   │       │   │       ├── cowboy_children.beam
│   │       │   │       ├── cowboy_clear.beam
│   │       │   │       ├── cowboy_clock.beam
│   │       │   │       ├── cowboy_compress_h.beam
│   │       │   │       ├── cowboy_constraints.beam
│   │       │   │       ├── cowboy_handler.beam
│   │       │   │       ├── cowboy_http.beam
│   │       │   │       ├── cowboy_http2.beam
│   │       │   │       ├── cowboy_loop.beam
│   │       │   │       ├── cowboy_metrics_h.beam
│   │       │   │       ├── cowboy_middleware.beam
│   │       │   │       ├── cowboy_req.beam
│   │       │   │       ├── cowboy_rest.beam
│   │       │   │       ├── cowboy_router.beam
│   │       │   │       ├── cowboy_static.beam
│   │       │   │       ├── cowboy_stream.beam
│   │       │   │       ├── cowboy_stream_h.beam
│   │       │   │       ├── cowboy_sub_protocol.beam
│   │       │   │       ├── cowboy_sup.beam
│   │       │   │       ├── cowboy_tls.beam
│   │       │   │       ├── cowboy_tracer_h.beam
│   │       │   │       └── cowboy_websocket.beam
│   │       │   ├── cowlib-2.9.1
│   │       │   │   ├── ebin
│   │       │   │   │   ├── cow_base64url.beam
│   │       │   │   │   ├── cow_cookie.beam
│   │       │   │   │   ├── cow_date.beam
│   │       │   │   │   ├── cow_hpack.beam
│   │       │   │   │   ├── cow_http.beam
│   │       │   │   │   ├── cow_http2.beam
│   │       │   │   │   ├── cow_http2_machine.beam
│   │       │   │   │   ├── cow_http_hd.beam
│   │       │   │   │   ├── cow_http_struct_hd.beam
│   │       │   │   │   ├── cow_http_te.beam
│   │       │   │   │   ├── cow_iolists.beam
│   │       │   │   │   ├── cow_link.beam
│   │       │   │   │   ├── cow_mimetypes.beam
│   │       │   │   │   ├── cow_multipart.beam
│   │       │   │   │   ├── cow_qs.beam
│   │       │   │   │   ├── cow_spdy.beam
│   │       │   │   │   ├── cow_sse.beam
│   │       │   │   │   ├── cow_uri.beam
│   │       │   │   │   ├── cow_uri_template.beam
│   │       │   │   │   ├── cow_ws.beam
│   │       │   │   │   └── cowlib.app
│   │       │   │   └── include
│   │       │   │       ├── cow_inline.hrl
│   │       │   │       └── cow_parse.hrl
│   │       │   ├── prometheus-4.6.0
│   │       │   │   ├── ebin
│   │       │   │   │   ├── prometheus.app
│   │       │   │   │   ├── prometheus.beam
│   │       │   │   │   ├── prometheus_boolean.beam
│   │       │   │   │   ├── prometheus_buckets.beam
│   │       │   │   │   ├── prometheus_collector.beam
│   │       │   │   │   ├── prometheus_counter.beam
│   │       │   │   │   ├── prometheus_format.beam
│   │       │   │   │   ├── prometheus_gauge.beam
│   │       │   │   │   ├── prometheus_histogram.beam
│   │       │   │   │   ├── prometheus_http.beam
│   │       │   │   │   ├── prometheus_instrumenter.beam
│   │       │   │   │   ├── prometheus_metric.beam
│   │       │   │   │   ├── prometheus_metric_spec.beam
│   │       │   │   │   ├── prometheus_misc.beam
│   │       │   │   │   ├── prometheus_mnesia.beam
│   │       │   │   │   ├── prometheus_mnesia_collector.beam
│   │       │   │   │   ├── prometheus_model.beam
│   │       │   │   │   ├── prometheus_model_helpers.beam
│   │       │   │   │   ├── prometheus_protobuf_format.beam
│   │       │   │   │   ├── prometheus_registry.beam
│   │       │   │   │   ├── prometheus_summary.beam
│   │       │   │   │   ├── prometheus_sup.beam
│   │       │   │   │   ├── prometheus_test_instrumenter.beam
│   │       │   │   │   ├── prometheus_text_format.beam
│   │       │   │   │   ├── prometheus_time.beam
│   │       │   │   │   ├── prometheus_vm_dist_collector.beam
│   │       │   │   │   ├── prometheus_vm_memory_collector.beam
│   │       │   │   │   ├── prometheus_vm_msacc_collector.beam
│   │       │   │   │   ├── prometheus_vm_statistics_collector.beam
│   │       │   │   │   └── prometheus_vm_system_info_collector.beam
│   │       │   │   └── include
│   │       │   │       ├── prometheus.hrl
│   │       │   │       └── prometheus_model.hrl
│   │       │   ├── rabbitmq_management_agent-3.8.16
│   │       │   │   ├── ebin
│   │       │   │   │   ├── Elixir.RabbitMQ.CLI.Ctl.Commands.ResetStatsDbCommand.beam
│   │       │   │   │   ├── exometer_slide.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_app.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_config.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_sup.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_sup_sup.beam
│   │       │   │   │   ├── rabbit_mgmt_data.beam
│   │       │   │   │   ├── rabbit_mgmt_data_compat.beam
│   │       │   │   │   ├── rabbit_mgmt_db_handler.beam
│   │       │   │   │   ├── rabbit_mgmt_external_stats.beam
│   │       │   │   │   ├── rabbit_mgmt_ff.beam
│   │       │   │   │   ├── rabbit_mgmt_format.beam
│   │       │   │   │   ├── rabbit_mgmt_gc.beam
│   │       │   │   │   ├── rabbit_mgmt_metrics_collector.beam
│   │       │   │   │   ├── rabbit_mgmt_metrics_gc.beam
│   │       │   │   │   ├── rabbit_mgmt_storage.beam
│   │       │   │   │   └── rabbitmq_management_agent.app
│   │       │   │   ├── include
│   │       │   │   │   ├── rabbit_mgmt_agent.hrl
│   │       │   │   │   ├── rabbit_mgmt_metrics.hrl
│   │       │   │   │   └── rabbit_mgmt_records.hrl
│   │       │   │   └── priv
│   │       │   │       └── schema
│   │       │   │           └── rabbitmq_management_agent.schema
│   │       │   ├── rabbitmq_prometheus-3.8.16
│   │       │   │   ├── ebin
│   │       │   │   │   ├── prometheus_rabbitmq_core_metrics_collector.beam
│   │       │   │   │   ├── rabbit_prometheus_app.beam
│   │       │   │   │   ├── rabbit_prometheus_dispatcher.beam
│   │       │   │   │   ├── rabbit_prometheus_handler.beam
│   │       │   │   │   └── rabbitmq_prometheus.app
│   │       │   │   └── priv
│   │       │   │       └── schema
│   │       │   │           └── rabbitmq_prometheus.schema
│   │       │   └── rabbitmq_web_dispatch-3.8.16
│   │       │       └── ebin
│   │       │           ├── rabbit_cowboy_middleware.beam
│   │       │           ├── rabbit_cowboy_redirect.beam
│   │       │           ├── rabbit_cowboy_stream_h.beam
│   │       │           ├── rabbit_web_dispatch.beam
│   │       │           ├── rabbit_web_dispatch_app.beam
│   │       │           ├── rabbit_web_dispatch_listing_handler.beam
│   │       │           ├── rabbit_web_dispatch_registry.beam
│   │       │           ├── rabbit_web_dispatch_sup.beam
│   │       │           ├── rabbit_web_dispatch_util.beam
│   │       │           ├── rabbitmq_web_dispatch.app
│   │       │           ├── webmachine_log.beam
│   │       │           └── webmachine_log_handler.beam
│   │       ├── rabbit@6c57fd82a7fc
│   │       │   ├── DECISION_TAB.LOG
│   │       │   ├── LATEST.LOG
│   │       │   ├── cluster_nodes.config
│   │       │   ├── msg_stores
│   │       │   │   └── vhosts
│   │       │   │       └── 628WB79CIFDYO9LJI6DKMI09L
│   │       │   │           ├── msg_store_persistent
│   │       │   │           │   ├── 0.rdq
│   │       │   │           │   ├── clean.dot
│   │       │   │           │   ├── file_summary.ets
│   │       │   │           │   └── msg_store_index.ets
│   │       │   │           ├── msg_store_transient
│   │       │   │           │   ├── 0.rdq
│   │       │   │           │   ├── clean.dot
│   │       │   │           │   ├── file_summary.ets
│   │       │   │           │   └── msg_store_index.ets
│   │       │   │           └── recovery.dets
│   │       │   ├── nodes_running_at_shutdown
│   │       │   ├── quorum
│   │       │   │   └── rabbit@6c57fd82a7fc
│   │       │   │       ├── 00000001.wal
│   │       │   │       ├── meta.dets
│   │       │   │       └── names.dets
│   │       │   ├── rabbit_durable_exchange.DCD
│   │       │   ├── rabbit_durable_exchange.DCL
│   │       │   ├── rabbit_durable_queue.DCD
│   │       │   ├── rabbit_durable_queue.DCL
│   │       │   ├── rabbit_durable_route.DCD
│   │       │   ├── rabbit_runtime_parameters.DCD
│   │       │   ├── rabbit_runtime_parameters.DCL
│   │       │   ├── rabbit_serial
│   │       │   ├── rabbit_topic_permission.DCD
│   │       │   ├── rabbit_user.DCD
│   │       │   ├── rabbit_user_permission.DCD
│   │       │   ├── rabbit_user_permission.DCL
│   │       │   ├── rabbit_vhost.DCD
│   │       │   ├── schema.DAT
│   │       │   └── schema_version
│   │       ├── rabbit@6c57fd82a7fc-feature_flags
│   │       ├── rabbit@6c57fd82a7fc-plugins-expand
│   │       │   ├── cowboy-2.8.0
│   │       │   │   └── ebin
│   │       │   │       ├── cowboy.app
│   │       │   │       ├── cowboy.beam
│   │       │   │       ├── cowboy_app.beam
│   │       │   │       ├── cowboy_bstr.beam
│   │       │   │       ├── cowboy_children.beam
│   │       │   │       ├── cowboy_clear.beam
│   │       │   │       ├── cowboy_clock.beam
│   │       │   │       ├── cowboy_compress_h.beam
│   │       │   │       ├── cowboy_constraints.beam
│   │       │   │       ├── cowboy_handler.beam
│   │       │   │       ├── cowboy_http.beam
│   │       │   │       ├── cowboy_http2.beam
│   │       │   │       ├── cowboy_loop.beam
│   │       │   │       ├── cowboy_metrics_h.beam
│   │       │   │       ├── cowboy_middleware.beam
│   │       │   │       ├── cowboy_req.beam
│   │       │   │       ├── cowboy_rest.beam
│   │       │   │       ├── cowboy_router.beam
│   │       │   │       ├── cowboy_static.beam
│   │       │   │       ├── cowboy_stream.beam
│   │       │   │       ├── cowboy_stream_h.beam
│   │       │   │       ├── cowboy_sub_protocol.beam
│   │       │   │       ├── cowboy_sup.beam
│   │       │   │       ├── cowboy_tls.beam
│   │       │   │       ├── cowboy_tracer_h.beam
│   │       │   │       └── cowboy_websocket.beam
│   │       │   ├── cowlib-2.9.1
│   │       │   │   ├── ebin
│   │       │   │   │   ├── cow_base64url.beam
│   │       │   │   │   ├── cow_cookie.beam
│   │       │   │   │   ├── cow_date.beam
│   │       │   │   │   ├── cow_hpack.beam
│   │       │   │   │   ├── cow_http.beam
│   │       │   │   │   ├── cow_http2.beam
│   │       │   │   │   ├── cow_http2_machine.beam
│   │       │   │   │   ├── cow_http_hd.beam
│   │       │   │   │   ├── cow_http_struct_hd.beam
│   │       │   │   │   ├── cow_http_te.beam
│   │       │   │   │   ├── cow_iolists.beam
│   │       │   │   │   ├── cow_link.beam
│   │       │   │   │   ├── cow_mimetypes.beam
│   │       │   │   │   ├── cow_multipart.beam
│   │       │   │   │   ├── cow_qs.beam
│   │       │   │   │   ├── cow_spdy.beam
│   │       │   │   │   ├── cow_sse.beam
│   │       │   │   │   ├── cow_uri.beam
│   │       │   │   │   ├── cow_uri_template.beam
│   │       │   │   │   ├── cow_ws.beam
│   │       │   │   │   └── cowlib.app
│   │       │   │   └── include
│   │       │   │       ├── cow_inline.hrl
│   │       │   │       └── cow_parse.hrl
│   │       │   ├── prometheus-4.6.0
│   │       │   │   ├── ebin
│   │       │   │   │   ├── prometheus.app
│   │       │   │   │   ├── prometheus.beam
│   │       │   │   │   ├── prometheus_boolean.beam
│   │       │   │   │   ├── prometheus_buckets.beam
│   │       │   │   │   ├── prometheus_collector.beam
│   │       │   │   │   ├── prometheus_counter.beam
│   │       │   │   │   ├── prometheus_format.beam
│   │       │   │   │   ├── prometheus_gauge.beam
│   │       │   │   │   ├── prometheus_histogram.beam
│   │       │   │   │   ├── prometheus_http.beam
│   │       │   │   │   ├── prometheus_instrumenter.beam
│   │       │   │   │   ├── prometheus_metric.beam
│   │       │   │   │   ├── prometheus_metric_spec.beam
│   │       │   │   │   ├── prometheus_misc.beam
│   │       │   │   │   ├── prometheus_mnesia.beam
│   │       │   │   │   ├── prometheus_mnesia_collector.beam
│   │       │   │   │   ├── prometheus_model.beam
│   │       │   │   │   ├── prometheus_model_helpers.beam
│   │       │   │   │   ├── prometheus_protobuf_format.beam
│   │       │   │   │   ├── prometheus_registry.beam
│   │       │   │   │   ├── prometheus_summary.beam
│   │       │   │   │   ├── prometheus_sup.beam
│   │       │   │   │   ├── prometheus_test_instrumenter.beam
│   │       │   │   │   ├── prometheus_text_format.beam
│   │       │   │   │   ├── prometheus_time.beam
│   │       │   │   │   ├── prometheus_vm_dist_collector.beam
│   │       │   │   │   ├── prometheus_vm_memory_collector.beam
│   │       │   │   │   ├── prometheus_vm_msacc_collector.beam
│   │       │   │   │   ├── prometheus_vm_statistics_collector.beam
│   │       │   │   │   └── prometheus_vm_system_info_collector.beam
│   │       │   │   └── include
│   │       │   │       ├── prometheus.hrl
│   │       │   │       └── prometheus_model.hrl
│   │       │   ├── rabbitmq_management_agent-3.8.16
│   │       │   │   ├── ebin
│   │       │   │   │   ├── Elixir.RabbitMQ.CLI.Ctl.Commands.ResetStatsDbCommand.beam
│   │       │   │   │   ├── exometer_slide.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_app.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_config.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_sup.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_sup_sup.beam
│   │       │   │   │   ├── rabbit_mgmt_data.beam
│   │       │   │   │   ├── rabbit_mgmt_data_compat.beam
│   │       │   │   │   ├── rabbit_mgmt_db_handler.beam
│   │       │   │   │   ├── rabbit_mgmt_external_stats.beam
│   │       │   │   │   ├── rabbit_mgmt_ff.beam
│   │       │   │   │   ├── rabbit_mgmt_format.beam
│   │       │   │   │   ├── rabbit_mgmt_gc.beam
│   │       │   │   │   ├── rabbit_mgmt_metrics_collector.beam
│   │       │   │   │   ├── rabbit_mgmt_metrics_gc.beam
│   │       │   │   │   ├── rabbit_mgmt_storage.beam
│   │       │   │   │   └── rabbitmq_management_agent.app
│   │       │   │   ├── include
│   │       │   │   │   ├── rabbit_mgmt_agent.hrl
│   │       │   │   │   ├── rabbit_mgmt_metrics.hrl
│   │       │   │   │   └── rabbit_mgmt_records.hrl
│   │       │   │   └── priv
│   │       │   │       └── schema
│   │       │   │           └── rabbitmq_management_agent.schema
│   │       │   ├── rabbitmq_prometheus-3.8.16
│   │       │   │   ├── ebin
│   │       │   │   │   ├── prometheus_rabbitmq_core_metrics_collector.beam
│   │       │   │   │   ├── rabbit_prometheus_app.beam
│   │       │   │   │   ├── rabbit_prometheus_dispatcher.beam
│   │       │   │   │   ├── rabbit_prometheus_handler.beam
│   │       │   │   │   └── rabbitmq_prometheus.app
│   │       │   │   └── priv
│   │       │   │       └── schema
│   │       │   │           └── rabbitmq_prometheus.schema
│   │       │   └── rabbitmq_web_dispatch-3.8.16
│   │       │       └── ebin
│   │       │           ├── rabbit_cowboy_middleware.beam
│   │       │           ├── rabbit_cowboy_redirect.beam
│   │       │           ├── rabbit_cowboy_stream_h.beam
│   │       │           ├── rabbit_web_dispatch.beam
│   │       │           ├── rabbit_web_dispatch_app.beam
│   │       │           ├── rabbit_web_dispatch_listing_handler.beam
│   │       │           ├── rabbit_web_dispatch_registry.beam
│   │       │           ├── rabbit_web_dispatch_sup.beam
│   │       │           ├── rabbit_web_dispatch_util.beam
│   │       │           ├── rabbitmq_web_dispatch.app
│   │       │           ├── webmachine_log.beam
│   │       │           └── webmachine_log_handler.beam
│   │       ├── rabbit@7424799d16a3
│   │       │   ├── DECISION_TAB.LOG
│   │       │   ├── LATEST.LOG
│   │       │   ├── cluster_nodes.config
│   │       │   ├── msg_stores
│   │       │   │   └── vhosts
│   │       │   │       └── 628WB79CIFDYO9LJI6DKMI09L
│   │       │   │           ├── msg_store_persistent
│   │       │   │           │   ├── 0.rdq
│   │       │   │           │   ├── clean.dot
│   │       │   │           │   ├── file_summary.ets
│   │       │   │           │   └── msg_store_index.ets
│   │       │   │           ├── msg_store_transient
│   │       │   │           │   ├── 0.rdq
│   │       │   │           │   ├── clean.dot
│   │       │   │           │   ├── file_summary.ets
│   │       │   │           │   └── msg_store_index.ets
│   │       │   │           └── recovery.dets
│   │       │   ├── nodes_running_at_shutdown
│   │       │   ├── quorum
│   │       │   │   └── rabbit@7424799d16a3
│   │       │   │       ├── 00000001.wal
│   │       │   │       ├── meta.dets
│   │       │   │       └── names.dets
│   │       │   ├── rabbit_durable_exchange.DCD
│   │       │   ├── rabbit_durable_queue.DCD
│   │       │   ├── rabbit_durable_queue.DCL
│   │       │   ├── rabbit_durable_route.DCD
│   │       │   ├── rabbit_runtime_parameters.DCD
│   │       │   ├── rabbit_serial
│   │       │   ├── rabbit_topic_permission.DCD
│   │       │   ├── rabbit_user.DCD
│   │       │   ├── rabbit_user.DCL
│   │       │   ├── rabbit_user_permission.DCD
│   │       │   ├── rabbit_vhost.DCD
│   │       │   ├── rabbit_vhost.DCL
│   │       │   ├── schema.DAT
│   │       │   └── schema_version
│   │       ├── rabbit@7424799d16a3-feature_flags
│   │       ├── rabbit@7424799d16a3-plugins-expand
│   │       │   ├── cowboy-2.8.0
│   │       │   │   └── ebin
│   │       │   │       ├── cowboy.app
│   │       │   │       ├── cowboy.beam
│   │       │   │       ├── cowboy_app.beam
│   │       │   │       ├── cowboy_bstr.beam
│   │       │   │       ├── cowboy_children.beam
│   │       │   │       ├── cowboy_clear.beam
│   │       │   │       ├── cowboy_clock.beam
│   │       │   │       ├── cowboy_compress_h.beam
│   │       │   │       ├── cowboy_constraints.beam
│   │       │   │       ├── cowboy_handler.beam
│   │       │   │       ├── cowboy_http.beam
│   │       │   │       ├── cowboy_http2.beam
│   │       │   │       ├── cowboy_loop.beam
│   │       │   │       ├── cowboy_metrics_h.beam
│   │       │   │       ├── cowboy_middleware.beam
│   │       │   │       ├── cowboy_req.beam
│   │       │   │       ├── cowboy_rest.beam
│   │       │   │       ├── cowboy_router.beam
│   │       │   │       ├── cowboy_static.beam
│   │       │   │       ├── cowboy_stream.beam
│   │       │   │       ├── cowboy_stream_h.beam
│   │       │   │       ├── cowboy_sub_protocol.beam
│   │       │   │       ├── cowboy_sup.beam
│   │       │   │       ├── cowboy_tls.beam
│   │       │   │       ├── cowboy_tracer_h.beam
│   │       │   │       └── cowboy_websocket.beam
│   │       │   ├── cowlib-2.9.1
│   │       │   │   ├── ebin
│   │       │   │   │   ├── cow_base64url.beam
│   │       │   │   │   ├── cow_cookie.beam
│   │       │   │   │   ├── cow_date.beam
│   │       │   │   │   ├── cow_hpack.beam
│   │       │   │   │   ├── cow_http.beam
│   │       │   │   │   ├── cow_http2.beam
│   │       │   │   │   ├── cow_http2_machine.beam
│   │       │   │   │   ├── cow_http_hd.beam
│   │       │   │   │   ├── cow_http_struct_hd.beam
│   │       │   │   │   ├── cow_http_te.beam
│   │       │   │   │   ├── cow_iolists.beam
│   │       │   │   │   ├── cow_link.beam
│   │       │   │   │   ├── cow_mimetypes.beam
│   │       │   │   │   ├── cow_multipart.beam
│   │       │   │   │   ├── cow_qs.beam
│   │       │   │   │   ├── cow_spdy.beam
│   │       │   │   │   ├── cow_sse.beam
│   │       │   │   │   ├── cow_uri.beam
│   │       │   │   │   ├── cow_uri_template.beam
│   │       │   │   │   ├── cow_ws.beam
│   │       │   │   │   └── cowlib.app
│   │       │   │   └── include
│   │       │   │       ├── cow_inline.hrl
│   │       │   │       └── cow_parse.hrl
│   │       │   ├── prometheus-4.6.0
│   │       │   │   ├── ebin
│   │       │   │   │   ├── prometheus.app
│   │       │   │   │   ├── prometheus.beam
│   │       │   │   │   ├── prometheus_boolean.beam
│   │       │   │   │   ├── prometheus_buckets.beam
│   │       │   │   │   ├── prometheus_collector.beam
│   │       │   │   │   ├── prometheus_counter.beam
│   │       │   │   │   ├── prometheus_format.beam
│   │       │   │   │   ├── prometheus_gauge.beam
│   │       │   │   │   ├── prometheus_histogram.beam
│   │       │   │   │   ├── prometheus_http.beam
│   │       │   │   │   ├── prometheus_instrumenter.beam
│   │       │   │   │   ├── prometheus_metric.beam
│   │       │   │   │   ├── prometheus_metric_spec.beam
│   │       │   │   │   ├── prometheus_misc.beam
│   │       │   │   │   ├── prometheus_mnesia.beam
│   │       │   │   │   ├── prometheus_mnesia_collector.beam
│   │       │   │   │   ├── prometheus_model.beam
│   │       │   │   │   ├── prometheus_model_helpers.beam
│   │       │   │   │   ├── prometheus_protobuf_format.beam
│   │       │   │   │   ├── prometheus_registry.beam
│   │       │   │   │   ├── prometheus_summary.beam
│   │       │   │   │   ├── prometheus_sup.beam
│   │       │   │   │   ├── prometheus_test_instrumenter.beam
│   │       │   │   │   ├── prometheus_text_format.beam
│   │       │   │   │   ├── prometheus_time.beam
│   │       │   │   │   ├── prometheus_vm_dist_collector.beam
│   │       │   │   │   ├── prometheus_vm_memory_collector.beam
│   │       │   │   │   ├── prometheus_vm_msacc_collector.beam
│   │       │   │   │   ├── prometheus_vm_statistics_collector.beam
│   │       │   │   │   └── prometheus_vm_system_info_collector.beam
│   │       │   │   └── include
│   │       │   │       ├── prometheus.hrl
│   │       │   │       └── prometheus_model.hrl
│   │       │   ├── rabbitmq_management_agent-3.8.16
│   │       │   │   ├── ebin
│   │       │   │   │   ├── Elixir.RabbitMQ.CLI.Ctl.Commands.ResetStatsDbCommand.beam
│   │       │   │   │   ├── exometer_slide.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_app.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_config.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_sup.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_sup_sup.beam
│   │       │   │   │   ├── rabbit_mgmt_data.beam
│   │       │   │   │   ├── rabbit_mgmt_data_compat.beam
│   │       │   │   │   ├── rabbit_mgmt_db_handler.beam
│   │       │   │   │   ├── rabbit_mgmt_external_stats.beam
│   │       │   │   │   ├── rabbit_mgmt_ff.beam
│   │       │   │   │   ├── rabbit_mgmt_format.beam
│   │       │   │   │   ├── rabbit_mgmt_gc.beam
│   │       │   │   │   ├── rabbit_mgmt_metrics_collector.beam
│   │       │   │   │   ├── rabbit_mgmt_metrics_gc.beam
│   │       │   │   │   ├── rabbit_mgmt_storage.beam
│   │       │   │   │   └── rabbitmq_management_agent.app
│   │       │   │   ├── include
│   │       │   │   │   ├── rabbit_mgmt_agent.hrl
│   │       │   │   │   ├── rabbit_mgmt_metrics.hrl
│   │       │   │   │   └── rabbit_mgmt_records.hrl
│   │       │   │   └── priv
│   │       │   │       └── schema
│   │       │   │           └── rabbitmq_management_agent.schema
│   │       │   ├── rabbitmq_prometheus-3.8.16
│   │       │   │   ├── ebin
│   │       │   │   │   ├── prometheus_rabbitmq_core_metrics_collector.beam
│   │       │   │   │   ├── rabbit_prometheus_app.beam
│   │       │   │   │   ├── rabbit_prometheus_dispatcher.beam
│   │       │   │   │   ├── rabbit_prometheus_handler.beam
│   │       │   │   │   └── rabbitmq_prometheus.app
│   │       │   │   └── priv
│   │       │   │       └── schema
│   │       │   │           └── rabbitmq_prometheus.schema
│   │       │   └── rabbitmq_web_dispatch-3.8.16
│   │       │       └── ebin
│   │       │           ├── rabbit_cowboy_middleware.beam
│   │       │           ├── rabbit_cowboy_redirect.beam
│   │       │           ├── rabbit_cowboy_stream_h.beam
│   │       │           ├── rabbit_web_dispatch.beam
│   │       │           ├── rabbit_web_dispatch_app.beam
│   │       │           ├── rabbit_web_dispatch_listing_handler.beam
│   │       │           ├── rabbit_web_dispatch_registry.beam
│   │       │           ├── rabbit_web_dispatch_sup.beam
│   │       │           ├── rabbit_web_dispatch_util.beam
│   │       │           ├── rabbitmq_web_dispatch.app
│   │       │           ├── webmachine_log.beam
│   │       │           └── webmachine_log_handler.beam
│   │       ├── rabbit@74d0f5ae8d6f
│   │       │   ├── DECISION_TAB.LOG
│   │       │   ├── LATEST.LOG
│   │       │   ├── cluster_nodes.config
│   │       │   ├── msg_stores
│   │       │   │   └── vhosts
│   │       │   │       └── 628WB79CIFDYO9LJI6DKMI09L
│   │       │   │           ├── msg_store_persistent
│   │       │   │           │   ├── 0.rdq
│   │       │   │           │   ├── clean.dot
│   │       │   │           │   ├── file_summary.ets
│   │       │   │           │   └── msg_store_index.ets
│   │       │   │           ├── msg_store_transient
│   │       │   │           │   ├── 0.rdq
│   │       │   │           │   ├── clean.dot
│   │       │   │           │   ├── file_summary.ets
│   │       │   │           │   └── msg_store_index.ets
│   │       │   │           └── recovery.dets
│   │       │   ├── nodes_running_at_shutdown
│   │       │   ├── quorum
│   │       │   │   └── rabbit@74d0f5ae8d6f
│   │       │   │       ├── 00000001.wal
│   │       │   │       ├── meta.dets
│   │       │   │       └── names.dets
│   │       │   ├── rabbit_durable_exchange.DCD
│   │       │   ├── rabbit_durable_queue.DCD
│   │       │   ├── rabbit_durable_queue.DCL
│   │       │   ├── rabbit_durable_route.DCD
│   │       │   ├── rabbit_runtime_parameters.DCD
│   │       │   ├── rabbit_serial
│   │       │   ├── rabbit_topic_permission.DCD
│   │       │   ├── rabbit_user.DCD
│   │       │   ├── rabbit_user.DCL
│   │       │   ├── rabbit_user_permission.DCD
│   │       │   ├── rabbit_vhost.DCD
│   │       │   ├── rabbit_vhost.DCL
│   │       │   ├── schema.DAT
│   │       │   └── schema_version
│   │       ├── rabbit@74d0f5ae8d6f-feature_flags
│   │       ├── rabbit@74d0f5ae8d6f-plugins-expand
│   │       │   ├── cowboy-2.8.0
│   │       │   │   └── ebin
│   │       │   │       ├── cowboy.app
│   │       │   │       ├── cowboy.beam
│   │       │   │       ├── cowboy_app.beam
│   │       │   │       ├── cowboy_bstr.beam
│   │       │   │       ├── cowboy_children.beam
│   │       │   │       ├── cowboy_clear.beam
│   │       │   │       ├── cowboy_clock.beam
│   │       │   │       ├── cowboy_compress_h.beam
│   │       │   │       ├── cowboy_constraints.beam
│   │       │   │       ├── cowboy_handler.beam
│   │       │   │       ├── cowboy_http.beam
│   │       │   │       ├── cowboy_http2.beam
│   │       │   │       ├── cowboy_loop.beam
│   │       │   │       ├── cowboy_metrics_h.beam
│   │       │   │       ├── cowboy_middleware.beam
│   │       │   │       ├── cowboy_req.beam
│   │       │   │       ├── cowboy_rest.beam
│   │       │   │       ├── cowboy_router.beam
│   │       │   │       ├── cowboy_static.beam
│   │       │   │       ├── cowboy_stream.beam
│   │       │   │       ├── cowboy_stream_h.beam
│   │       │   │       ├── cowboy_sub_protocol.beam
│   │       │   │       ├── cowboy_sup.beam
│   │       │   │       ├── cowboy_tls.beam
│   │       │   │       ├── cowboy_tracer_h.beam
│   │       │   │       └── cowboy_websocket.beam
│   │       │   ├── cowlib-2.9.1
│   │       │   │   ├── ebin
│   │       │   │   │   ├── cow_base64url.beam
│   │       │   │   │   ├── cow_cookie.beam
│   │       │   │   │   ├── cow_date.beam
│   │       │   │   │   ├── cow_hpack.beam
│   │       │   │   │   ├── cow_http.beam
│   │       │   │   │   ├── cow_http2.beam
│   │       │   │   │   ├── cow_http2_machine.beam
│   │       │   │   │   ├── cow_http_hd.beam
│   │       │   │   │   ├── cow_http_struct_hd.beam
│   │       │   │   │   ├── cow_http_te.beam
│   │       │   │   │   ├── cow_iolists.beam
│   │       │   │   │   ├── cow_link.beam
│   │       │   │   │   ├── cow_mimetypes.beam
│   │       │   │   │   ├── cow_multipart.beam
│   │       │   │   │   ├── cow_qs.beam
│   │       │   │   │   ├── cow_spdy.beam
│   │       │   │   │   ├── cow_sse.beam
│   │       │   │   │   ├── cow_uri.beam
│   │       │   │   │   ├── cow_uri_template.beam
│   │       │   │   │   ├── cow_ws.beam
│   │       │   │   │   └── cowlib.app
│   │       │   │   └── include
│   │       │   │       ├── cow_inline.hrl
│   │       │   │       └── cow_parse.hrl
│   │       │   ├── prometheus-4.6.0
│   │       │   │   ├── ebin
│   │       │   │   │   ├── prometheus.app
│   │       │   │   │   ├── prometheus.beam
│   │       │   │   │   ├── prometheus_boolean.beam
│   │       │   │   │   ├── prometheus_buckets.beam
│   │       │   │   │   ├── prometheus_collector.beam
│   │       │   │   │   ├── prometheus_counter.beam
│   │       │   │   │   ├── prometheus_format.beam
│   │       │   │   │   ├── prometheus_gauge.beam
│   │       │   │   │   ├── prometheus_histogram.beam
│   │       │   │   │   ├── prometheus_http.beam
│   │       │   │   │   ├── prometheus_instrumenter.beam
│   │       │   │   │   ├── prometheus_metric.beam
│   │       │   │   │   ├── prometheus_metric_spec.beam
│   │       │   │   │   ├── prometheus_misc.beam
│   │       │   │   │   ├── prometheus_mnesia.beam
│   │       │   │   │   ├── prometheus_mnesia_collector.beam
│   │       │   │   │   ├── prometheus_model.beam
│   │       │   │   │   ├── prometheus_model_helpers.beam
│   │       │   │   │   ├── prometheus_protobuf_format.beam
│   │       │   │   │   ├── prometheus_registry.beam
│   │       │   │   │   ├── prometheus_summary.beam
│   │       │   │   │   ├── prometheus_sup.beam
│   │       │   │   │   ├── prometheus_test_instrumenter.beam
│   │       │   │   │   ├── prometheus_text_format.beam
│   │       │   │   │   ├── prometheus_time.beam
│   │       │   │   │   ├── prometheus_vm_dist_collector.beam
│   │       │   │   │   ├── prometheus_vm_memory_collector.beam
│   │       │   │   │   ├── prometheus_vm_msacc_collector.beam
│   │       │   │   │   ├── prometheus_vm_statistics_collector.beam
│   │       │   │   │   └── prometheus_vm_system_info_collector.beam
│   │       │   │   └── include
│   │       │   │       ├── prometheus.hrl
│   │       │   │       └── prometheus_model.hrl
│   │       │   ├── rabbitmq_management_agent-3.8.16
│   │       │   │   ├── ebin
│   │       │   │   │   ├── Elixir.RabbitMQ.CLI.Ctl.Commands.ResetStatsDbCommand.beam
│   │       │   │   │   ├── exometer_slide.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_app.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_config.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_sup.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_sup_sup.beam
│   │       │   │   │   ├── rabbit_mgmt_data.beam
│   │       │   │   │   ├── rabbit_mgmt_data_compat.beam
│   │       │   │   │   ├── rabbit_mgmt_db_handler.beam
│   │       │   │   │   ├── rabbit_mgmt_external_stats.beam
│   │       │   │   │   ├── rabbit_mgmt_ff.beam
│   │       │   │   │   ├── rabbit_mgmt_format.beam
│   │       │   │   │   ├── rabbit_mgmt_gc.beam
│   │       │   │   │   ├── rabbit_mgmt_metrics_collector.beam
│   │       │   │   │   ├── rabbit_mgmt_metrics_gc.beam
│   │       │   │   │   ├── rabbit_mgmt_storage.beam
│   │       │   │   │   └── rabbitmq_management_agent.app
│   │       │   │   ├── include
│   │       │   │   │   ├── rabbit_mgmt_agent.hrl
│   │       │   │   │   ├── rabbit_mgmt_metrics.hrl
│   │       │   │   │   └── rabbit_mgmt_records.hrl
│   │       │   │   └── priv
│   │       │   │       └── schema
│   │       │   │           └── rabbitmq_management_agent.schema
│   │       │   ├── rabbitmq_prometheus-3.8.16
│   │       │   │   ├── ebin
│   │       │   │   │   ├── prometheus_rabbitmq_core_metrics_collector.beam
│   │       │   │   │   ├── rabbit_prometheus_app.beam
│   │       │   │   │   ├── rabbit_prometheus_dispatcher.beam
│   │       │   │   │   ├── rabbit_prometheus_handler.beam
│   │       │   │   │   └── rabbitmq_prometheus.app
│   │       │   │   └── priv
│   │       │   │       └── schema
│   │       │   │           └── rabbitmq_prometheus.schema
│   │       │   └── rabbitmq_web_dispatch-3.8.16
│   │       │       └── ebin
│   │       │           ├── rabbit_cowboy_middleware.beam
│   │       │           ├── rabbit_cowboy_redirect.beam
│   │       │           ├── rabbit_cowboy_stream_h.beam
│   │       │           ├── rabbit_web_dispatch.beam
│   │       │           ├── rabbit_web_dispatch_app.beam
│   │       │           ├── rabbit_web_dispatch_listing_handler.beam
│   │       │           ├── rabbit_web_dispatch_registry.beam
│   │       │           ├── rabbit_web_dispatch_sup.beam
│   │       │           ├── rabbit_web_dispatch_util.beam
│   │       │           ├── rabbitmq_web_dispatch.app
│   │       │           ├── webmachine_log.beam
│   │       │           └── webmachine_log_handler.beam
│   │       ├── rabbit@74dbdeccd9a7
│   │       │   ├── DECISION_TAB.LOG
│   │       │   ├── LATEST.LOG
│   │       │   ├── cluster_nodes.config
│   │       │   ├── msg_stores
│   │       │   │   └── vhosts
│   │       │   │       └── 628WB79CIFDYO9LJI6DKMI09L
│   │       │   │           ├── msg_store_persistent
│   │       │   │           │   ├── 0.rdq
│   │       │   │           │   ├── clean.dot
│   │       │   │           │   ├── file_summary.ets
│   │       │   │           │   └── msg_store_index.ets
│   │       │   │           ├── msg_store_transient
│   │       │   │           │   ├── 0.rdq
│   │       │   │           │   ├── clean.dot
│   │       │   │           │   ├── file_summary.ets
│   │       │   │           │   └── msg_store_index.ets
│   │       │   │           └── recovery.dets
│   │       │   ├── nodes_running_at_shutdown
│   │       │   ├── quorum
│   │       │   │   └── rabbit@74dbdeccd9a7
│   │       │   │       ├── 00000002.wal
│   │       │   │       ├── meta.dets
│   │       │   │       └── names.dets
│   │       │   ├── rabbit_durable_exchange.DCD
│   │       │   ├── rabbit_durable_queue.DCD
│   │       │   ├── rabbit_durable_route.DCD
│   │       │   ├── rabbit_runtime_parameters.DCD
│   │       │   ├── rabbit_serial
│   │       │   ├── rabbit_topic_permission.DCD
│   │       │   ├── rabbit_user.DCD
│   │       │   ├── rabbit_user_permission.DCD
│   │       │   ├── rabbit_vhost.DCD
│   │       │   ├── schema.DAT
│   │       │   └── schema_version
│   │       ├── rabbit@74dbdeccd9a7-feature_flags
│   │       ├── rabbit@74dbdeccd9a7-plugins-expand
│   │       │   ├── cowboy-2.8.0
│   │       │   │   └── ebin
│   │       │   │       ├── cowboy.app
│   │       │   │       ├── cowboy.beam
│   │       │   │       ├── cowboy_app.beam
│   │       │   │       ├── cowboy_bstr.beam
│   │       │   │       ├── cowboy_children.beam
│   │       │   │       ├── cowboy_clear.beam
│   │       │   │       ├── cowboy_clock.beam
│   │       │   │       ├── cowboy_compress_h.beam
│   │       │   │       ├── cowboy_constraints.beam
│   │       │   │       ├── cowboy_handler.beam
│   │       │   │       ├── cowboy_http.beam
│   │       │   │       ├── cowboy_http2.beam
│   │       │   │       ├── cowboy_loop.beam
│   │       │   │       ├── cowboy_metrics_h.beam
│   │       │   │       ├── cowboy_middleware.beam
│   │       │   │       ├── cowboy_req.beam
│   │       │   │       ├── cowboy_rest.beam
│   │       │   │       ├── cowboy_router.beam
│   │       │   │       ├── cowboy_static.beam
│   │       │   │       ├── cowboy_stream.beam
│   │       │   │       ├── cowboy_stream_h.beam
│   │       │   │       ├── cowboy_sub_protocol.beam
│   │       │   │       ├── cowboy_sup.beam
│   │       │   │       ├── cowboy_tls.beam
│   │       │   │       ├── cowboy_tracer_h.beam
│   │       │   │       └── cowboy_websocket.beam
│   │       │   ├── cowlib-2.9.1
│   │       │   │   ├── ebin
│   │       │   │   │   ├── cow_base64url.beam
│   │       │   │   │   ├── cow_cookie.beam
│   │       │   │   │   ├── cow_date.beam
│   │       │   │   │   ├── cow_hpack.beam
│   │       │   │   │   ├── cow_http.beam
│   │       │   │   │   ├── cow_http2.beam
│   │       │   │   │   ├── cow_http2_machine.beam
│   │       │   │   │   ├── cow_http_hd.beam
│   │       │   │   │   ├── cow_http_struct_hd.beam
│   │       │   │   │   ├── cow_http_te.beam
│   │       │   │   │   ├── cow_iolists.beam
│   │       │   │   │   ├── cow_link.beam
│   │       │   │   │   ├── cow_mimetypes.beam
│   │       │   │   │   ├── cow_multipart.beam
│   │       │   │   │   ├── cow_qs.beam
│   │       │   │   │   ├── cow_spdy.beam
│   │       │   │   │   ├── cow_sse.beam
│   │       │   │   │   ├── cow_uri.beam
│   │       │   │   │   ├── cow_uri_template.beam
│   │       │   │   │   ├── cow_ws.beam
│   │       │   │   │   └── cowlib.app
│   │       │   │   └── include
│   │       │   │       ├── cow_inline.hrl
│   │       │   │       └── cow_parse.hrl
│   │       │   ├── prometheus-4.6.0
│   │       │   │   ├── ebin
│   │       │   │   │   ├── prometheus.app
│   │       │   │   │   ├── prometheus.beam
│   │       │   │   │   ├── prometheus_boolean.beam
│   │       │   │   │   ├── prometheus_buckets.beam
│   │       │   │   │   ├── prometheus_collector.beam
│   │       │   │   │   ├── prometheus_counter.beam
│   │       │   │   │   ├── prometheus_format.beam
│   │       │   │   │   ├── prometheus_gauge.beam
│   │       │   │   │   ├── prometheus_histogram.beam
│   │       │   │   │   ├── prometheus_http.beam
│   │       │   │   │   ├── prometheus_instrumenter.beam
│   │       │   │   │   ├── prometheus_metric.beam
│   │       │   │   │   ├── prometheus_metric_spec.beam
│   │       │   │   │   ├── prometheus_misc.beam
│   │       │   │   │   ├── prometheus_mnesia.beam
│   │       │   │   │   ├── prometheus_mnesia_collector.beam
│   │       │   │   │   ├── prometheus_model.beam
│   │       │   │   │   ├── prometheus_model_helpers.beam
│   │       │   │   │   ├── prometheus_protobuf_format.beam
│   │       │   │   │   ├── prometheus_registry.beam
│   │       │   │   │   ├── prometheus_summary.beam
│   │       │   │   │   ├── prometheus_sup.beam
│   │       │   │   │   ├── prometheus_test_instrumenter.beam
│   │       │   │   │   ├── prometheus_text_format.beam
│   │       │   │   │   ├── prometheus_time.beam
│   │       │   │   │   ├── prometheus_vm_dist_collector.beam
│   │       │   │   │   ├── prometheus_vm_memory_collector.beam
│   │       │   │   │   ├── prometheus_vm_msacc_collector.beam
│   │       │   │   │   ├── prometheus_vm_statistics_collector.beam
│   │       │   │   │   └── prometheus_vm_system_info_collector.beam
│   │       │   │   └── include
│   │       │   │       ├── prometheus.hrl
│   │       │   │       └── prometheus_model.hrl
│   │       │   ├── rabbitmq_management_agent-3.8.16
│   │       │   │   ├── ebin
│   │       │   │   │   ├── Elixir.RabbitMQ.CLI.Ctl.Commands.ResetStatsDbCommand.beam
│   │       │   │   │   ├── exometer_slide.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_app.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_config.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_sup.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_sup_sup.beam
│   │       │   │   │   ├── rabbit_mgmt_data.beam
│   │       │   │   │   ├── rabbit_mgmt_data_compat.beam
│   │       │   │   │   ├── rabbit_mgmt_db_handler.beam
│   │       │   │   │   ├── rabbit_mgmt_external_stats.beam
│   │       │   │   │   ├── rabbit_mgmt_ff.beam
│   │       │   │   │   ├── rabbit_mgmt_format.beam
│   │       │   │   │   ├── rabbit_mgmt_gc.beam
│   │       │   │   │   ├── rabbit_mgmt_metrics_collector.beam
│   │       │   │   │   ├── rabbit_mgmt_metrics_gc.beam
│   │       │   │   │   ├── rabbit_mgmt_storage.beam
│   │       │   │   │   └── rabbitmq_management_agent.app
│   │       │   │   ├── include
│   │       │   │   │   ├── rabbit_mgmt_agent.hrl
│   │       │   │   │   ├── rabbit_mgmt_metrics.hrl
│   │       │   │   │   └── rabbit_mgmt_records.hrl
│   │       │   │   └── priv
│   │       │   │       └── schema
│   │       │   │           └── rabbitmq_management_agent.schema
│   │       │   ├── rabbitmq_prometheus-3.8.16
│   │       │   │   ├── ebin
│   │       │   │   │   ├── prometheus_rabbitmq_core_metrics_collector.beam
│   │       │   │   │   ├── rabbit_prometheus_app.beam
│   │       │   │   │   ├── rabbit_prometheus_dispatcher.beam
│   │       │   │   │   ├── rabbit_prometheus_handler.beam
│   │       │   │   │   └── rabbitmq_prometheus.app
│   │       │   │   └── priv
│   │       │   │       └── schema
│   │       │   │           └── rabbitmq_prometheus.schema
│   │       │   └── rabbitmq_web_dispatch-3.8.16
│   │       │       └── ebin
│   │       │           ├── rabbit_cowboy_middleware.beam
│   │       │           ├── rabbit_cowboy_redirect.beam
│   │       │           ├── rabbit_cowboy_stream_h.beam
│   │       │           ├── rabbit_web_dispatch.beam
│   │       │           ├── rabbit_web_dispatch_app.beam
│   │       │           ├── rabbit_web_dispatch_listing_handler.beam
│   │       │           ├── rabbit_web_dispatch_registry.beam
│   │       │           ├── rabbit_web_dispatch_sup.beam
│   │       │           ├── rabbit_web_dispatch_util.beam
│   │       │           ├── rabbitmq_web_dispatch.app
│   │       │           ├── webmachine_log.beam
│   │       │           └── webmachine_log_handler.beam
│   │       ├── rabbit@a838b7530c3c
│   │       │   ├── DECISION_TAB.LOG
│   │       │   ├── LATEST.LOG
│   │       │   ├── cluster_nodes.config
│   │       │   ├── msg_stores
│   │       │   │   └── vhosts
│   │       │   │       └── 628WB79CIFDYO9LJI6DKMI09L
│   │       │   │           ├── msg_store_persistent
│   │       │   │           │   ├── 0.rdq
│   │       │   │           │   ├── clean.dot
│   │       │   │           │   ├── file_summary.ets
│   │       │   │           │   └── msg_store_index.ets
│   │       │   │           ├── msg_store_transient
│   │       │   │           │   ├── 0.rdq
│   │       │   │           │   ├── clean.dot
│   │       │   │           │   ├── file_summary.ets
│   │       │   │           │   └── msg_store_index.ets
│   │       │   │           └── recovery.dets
│   │       │   ├── nodes_running_at_shutdown
│   │       │   ├── quorum
│   │       │   │   └── rabbit@a838b7530c3c
│   │       │   │       ├── 00000004.wal
│   │       │   │       ├── meta.dets
│   │       │   │       └── names.dets
│   │       │   ├── rabbit_durable_exchange.DCD
│   │       │   ├── rabbit_durable_queue.DCD
│   │       │   ├── rabbit_durable_route.DCD
│   │       │   ├── rabbit_runtime_parameters.DCD
│   │       │   ├── rabbit_serial
│   │       │   ├── rabbit_topic_permission.DCD
│   │       │   ├── rabbit_user.DCD
│   │       │   ├── rabbit_user_permission.DCD
│   │       │   ├── rabbit_vhost.DCD
│   │       │   ├── schema.DAT
│   │       │   └── schema_version
│   │       ├── rabbit@a838b7530c3c-feature_flags
│   │       ├── rabbit@a838b7530c3c-plugins-expand
│   │       │   ├── cowboy-2.8.0
│   │       │   │   └── ebin
│   │       │   │       ├── cowboy.app
│   │       │   │       ├── cowboy.beam
│   │       │   │       ├── cowboy_app.beam
│   │       │   │       ├── cowboy_bstr.beam
│   │       │   │       ├── cowboy_children.beam
│   │       │   │       ├── cowboy_clear.beam
│   │       │   │       ├── cowboy_clock.beam
│   │       │   │       ├── cowboy_compress_h.beam
│   │       │   │       ├── cowboy_constraints.beam
│   │       │   │       ├── cowboy_handler.beam
│   │       │   │       ├── cowboy_http.beam
│   │       │   │       ├── cowboy_http2.beam
│   │       │   │       ├── cowboy_loop.beam
│   │       │   │       ├── cowboy_metrics_h.beam
│   │       │   │       ├── cowboy_middleware.beam
│   │       │   │       ├── cowboy_req.beam
│   │       │   │       ├── cowboy_rest.beam
│   │       │   │       ├── cowboy_router.beam
│   │       │   │       ├── cowboy_static.beam
│   │       │   │       ├── cowboy_stream.beam
│   │       │   │       ├── cowboy_stream_h.beam
│   │       │   │       ├── cowboy_sub_protocol.beam
│   │       │   │       ├── cowboy_sup.beam
│   │       │   │       ├── cowboy_tls.beam
│   │       │   │       ├── cowboy_tracer_h.beam
│   │       │   │       └── cowboy_websocket.beam
│   │       │   ├── cowlib-2.9.1
│   │       │   │   ├── ebin
│   │       │   │   │   ├── cow_base64url.beam
│   │       │   │   │   ├── cow_cookie.beam
│   │       │   │   │   ├── cow_date.beam
│   │       │   │   │   ├── cow_hpack.beam
│   │       │   │   │   ├── cow_http.beam
│   │       │   │   │   ├── cow_http2.beam
│   │       │   │   │   ├── cow_http2_machine.beam
│   │       │   │   │   ├── cow_http_hd.beam
│   │       │   │   │   ├── cow_http_struct_hd.beam
│   │       │   │   │   ├── cow_http_te.beam
│   │       │   │   │   ├── cow_iolists.beam
│   │       │   │   │   ├── cow_link.beam
│   │       │   │   │   ├── cow_mimetypes.beam
│   │       │   │   │   ├── cow_multipart.beam
│   │       │   │   │   ├── cow_qs.beam
│   │       │   │   │   ├── cow_spdy.beam
│   │       │   │   │   ├── cow_sse.beam
│   │       │   │   │   ├── cow_uri.beam
│   │       │   │   │   ├── cow_uri_template.beam
│   │       │   │   │   ├── cow_ws.beam
│   │       │   │   │   └── cowlib.app
│   │       │   │   └── include
│   │       │   │       ├── cow_inline.hrl
│   │       │   │       └── cow_parse.hrl
│   │       │   ├── prometheus-4.6.0
│   │       │   │   ├── ebin
│   │       │   │   │   ├── prometheus.app
│   │       │   │   │   ├── prometheus.beam
│   │       │   │   │   ├── prometheus_boolean.beam
│   │       │   │   │   ├── prometheus_buckets.beam
│   │       │   │   │   ├── prometheus_collector.beam
│   │       │   │   │   ├── prometheus_counter.beam
│   │       │   │   │   ├── prometheus_format.beam
│   │       │   │   │   ├── prometheus_gauge.beam
│   │       │   │   │   ├── prometheus_histogram.beam
│   │       │   │   │   ├── prometheus_http.beam
│   │       │   │   │   ├── prometheus_instrumenter.beam
│   │       │   │   │   ├── prometheus_metric.beam
│   │       │   │   │   ├── prometheus_metric_spec.beam
│   │       │   │   │   ├── prometheus_misc.beam
│   │       │   │   │   ├── prometheus_mnesia.beam
│   │       │   │   │   ├── prometheus_mnesia_collector.beam
│   │       │   │   │   ├── prometheus_model.beam
│   │       │   │   │   ├── prometheus_model_helpers.beam
│   │       │   │   │   ├── prometheus_protobuf_format.beam
│   │       │   │   │   ├── prometheus_registry.beam
│   │       │   │   │   ├── prometheus_summary.beam
│   │       │   │   │   ├── prometheus_sup.beam
│   │       │   │   │   ├── prometheus_test_instrumenter.beam
│   │       │   │   │   ├── prometheus_text_format.beam
│   │       │   │   │   ├── prometheus_time.beam
│   │       │   │   │   ├── prometheus_vm_dist_collector.beam
│   │       │   │   │   ├── prometheus_vm_memory_collector.beam
│   │       │   │   │   ├── prometheus_vm_msacc_collector.beam
│   │       │   │   │   ├── prometheus_vm_statistics_collector.beam
│   │       │   │   │   └── prometheus_vm_system_info_collector.beam
│   │       │   │   └── include
│   │       │   │       ├── prometheus.hrl
│   │       │   │       └── prometheus_model.hrl
│   │       │   ├── rabbitmq_management_agent-3.8.16
│   │       │   │   ├── ebin
│   │       │   │   │   ├── Elixir.RabbitMQ.CLI.Ctl.Commands.ResetStatsDbCommand.beam
│   │       │   │   │   ├── exometer_slide.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_app.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_config.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_sup.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_sup_sup.beam
│   │       │   │   │   ├── rabbit_mgmt_data.beam
│   │       │   │   │   ├── rabbit_mgmt_data_compat.beam
│   │       │   │   │   ├── rabbit_mgmt_db_handler.beam
│   │       │   │   │   ├── rabbit_mgmt_external_stats.beam
│   │       │   │   │   ├── rabbit_mgmt_ff.beam
│   │       │   │   │   ├── rabbit_mgmt_format.beam
│   │       │   │   │   ├── rabbit_mgmt_gc.beam
│   │       │   │   │   ├── rabbit_mgmt_metrics_collector.beam
│   │       │   │   │   ├── rabbit_mgmt_metrics_gc.beam
│   │       │   │   │   ├── rabbit_mgmt_storage.beam
│   │       │   │   │   └── rabbitmq_management_agent.app
│   │       │   │   ├── include
│   │       │   │   │   ├── rabbit_mgmt_agent.hrl
│   │       │   │   │   ├── rabbit_mgmt_metrics.hrl
│   │       │   │   │   └── rabbit_mgmt_records.hrl
│   │       │   │   └── priv
│   │       │   │       └── schema
│   │       │   │           └── rabbitmq_management_agent.schema
│   │       │   ├── rabbitmq_prometheus-3.8.16
│   │       │   │   ├── ebin
│   │       │   │   │   ├── prometheus_rabbitmq_core_metrics_collector.beam
│   │       │   │   │   ├── rabbit_prometheus_app.beam
│   │       │   │   │   ├── rabbit_prometheus_dispatcher.beam
│   │       │   │   │   ├── rabbit_prometheus_handler.beam
│   │       │   │   │   └── rabbitmq_prometheus.app
│   │       │   │   └── priv
│   │       │   │       └── schema
│   │       │   │           └── rabbitmq_prometheus.schema
│   │       │   └── rabbitmq_web_dispatch-3.8.16
│   │       │       └── ebin
│   │       │           ├── rabbit_cowboy_middleware.beam
│   │       │           ├── rabbit_cowboy_redirect.beam
│   │       │           ├── rabbit_cowboy_stream_h.beam
│   │       │           ├── rabbit_web_dispatch.beam
│   │       │           ├── rabbit_web_dispatch_app.beam
│   │       │           ├── rabbit_web_dispatch_listing_handler.beam
│   │       │           ├── rabbit_web_dispatch_registry.beam
│   │       │           ├── rabbit_web_dispatch_sup.beam
│   │       │           ├── rabbit_web_dispatch_util.beam
│   │       │           ├── rabbitmq_web_dispatch.app
│   │       │           ├── webmachine_log.beam
│   │       │           └── webmachine_log_handler.beam
│   │       ├── rabbit@c609f6389bbf
│   │       │   ├── DECISION_TAB.LOG
│   │       │   ├── LATEST.LOG
│   │       │   ├── cluster_nodes.config
│   │       │   ├── msg_stores
│   │       │   │   └── vhosts
│   │       │   │       └── 628WB79CIFDYO9LJI6DKMI09L
│   │       │   │           ├── msg_store_persistent
│   │       │   │           │   ├── 0.rdq
│   │       │   │           │   ├── clean.dot
│   │       │   │           │   ├── file_summary.ets
│   │       │   │           │   └── msg_store_index.ets
│   │       │   │           ├── msg_store_transient
│   │       │   │           │   ├── 0.rdq
│   │       │   │           │   ├── clean.dot
│   │       │   │           │   ├── file_summary.ets
│   │       │   │           │   └── msg_store_index.ets
│   │       │   │           └── recovery.dets
│   │       │   ├── nodes_running_at_shutdown
│   │       │   ├── quorum
│   │       │   │   └── rabbit@c609f6389bbf
│   │       │   │       ├── 00000001.wal
│   │       │   │       ├── meta.dets
│   │       │   │       └── names.dets
│   │       │   ├── rabbit_durable_exchange.DCD
│   │       │   ├── rabbit_durable_exchange.DCL
│   │       │   ├── rabbit_durable_queue.DCD
│   │       │   ├── rabbit_durable_queue.DCL
│   │       │   ├── rabbit_durable_route.DCD
│   │       │   ├── rabbit_runtime_parameters.DCD
│   │       │   ├── rabbit_runtime_parameters.DCL
│   │       │   ├── rabbit_serial
│   │       │   ├── rabbit_topic_permission.DCD
│   │       │   ├── rabbit_user.DCD
│   │       │   ├── rabbit_user_permission.DCD
│   │       │   ├── rabbit_user_permission.DCL
│   │       │   ├── rabbit_vhost.DCD
│   │       │   ├── schema.DAT
│   │       │   └── schema_version
│   │       ├── rabbit@c609f6389bbf-feature_flags
│   │       ├── rabbit@c609f6389bbf-plugins-expand
│   │       │   ├── cowboy-2.8.0
│   │       │   │   └── ebin
│   │       │   │       ├── cowboy.app
│   │       │   │       ├── cowboy.beam
│   │       │   │       ├── cowboy_app.beam
│   │       │   │       ├── cowboy_bstr.beam
│   │       │   │       ├── cowboy_children.beam
│   │       │   │       ├── cowboy_clear.beam
│   │       │   │       ├── cowboy_clock.beam
│   │       │   │       ├── cowboy_compress_h.beam
│   │       │   │       ├── cowboy_constraints.beam
│   │       │   │       ├── cowboy_handler.beam
│   │       │   │       ├── cowboy_http.beam
│   │       │   │       ├── cowboy_http2.beam
│   │       │   │       ├── cowboy_loop.beam
│   │       │   │       ├── cowboy_metrics_h.beam
│   │       │   │       ├── cowboy_middleware.beam
│   │       │   │       ├── cowboy_req.beam
│   │       │   │       ├── cowboy_rest.beam
│   │       │   │       ├── cowboy_router.beam
│   │       │   │       ├── cowboy_static.beam
│   │       │   │       ├── cowboy_stream.beam
│   │       │   │       ├── cowboy_stream_h.beam
│   │       │   │       ├── cowboy_sub_protocol.beam
│   │       │   │       ├── cowboy_sup.beam
│   │       │   │       ├── cowboy_tls.beam
│   │       │   │       ├── cowboy_tracer_h.beam
│   │       │   │       └── cowboy_websocket.beam
│   │       │   ├── cowlib-2.9.1
│   │       │   │   ├── ebin
│   │       │   │   │   ├── cow_base64url.beam
│   │       │   │   │   ├── cow_cookie.beam
│   │       │   │   │   ├── cow_date.beam
│   │       │   │   │   ├── cow_hpack.beam
│   │       │   │   │   ├── cow_http.beam
│   │       │   │   │   ├── cow_http2.beam
│   │       │   │   │   ├── cow_http2_machine.beam
│   │       │   │   │   ├── cow_http_hd.beam
│   │       │   │   │   ├── cow_http_struct_hd.beam
│   │       │   │   │   ├── cow_http_te.beam
│   │       │   │   │   ├── cow_iolists.beam
│   │       │   │   │   ├── cow_link.beam
│   │       │   │   │   ├── cow_mimetypes.beam
│   │       │   │   │   ├── cow_multipart.beam
│   │       │   │   │   ├── cow_qs.beam
│   │       │   │   │   ├── cow_spdy.beam
│   │       │   │   │   ├── cow_sse.beam
│   │       │   │   │   ├── cow_uri.beam
│   │       │   │   │   ├── cow_uri_template.beam
│   │       │   │   │   ├── cow_ws.beam
│   │       │   │   │   └── cowlib.app
│   │       │   │   └── include
│   │       │   │       ├── cow_inline.hrl
│   │       │   │       └── cow_parse.hrl
│   │       │   ├── prometheus-4.6.0
│   │       │   │   ├── ebin
│   │       │   │   │   ├── prometheus.app
│   │       │   │   │   ├── prometheus.beam
│   │       │   │   │   ├── prometheus_boolean.beam
│   │       │   │   │   ├── prometheus_buckets.beam
│   │       │   │   │   ├── prometheus_collector.beam
│   │       │   │   │   ├── prometheus_counter.beam
│   │       │   │   │   ├── prometheus_format.beam
│   │       │   │   │   ├── prometheus_gauge.beam
│   │       │   │   │   ├── prometheus_histogram.beam
│   │       │   │   │   ├── prometheus_http.beam
│   │       │   │   │   ├── prometheus_instrumenter.beam
│   │       │   │   │   ├── prometheus_metric.beam
│   │       │   │   │   ├── prometheus_metric_spec.beam
│   │       │   │   │   ├── prometheus_misc.beam
│   │       │   │   │   ├── prometheus_mnesia.beam
│   │       │   │   │   ├── prometheus_mnesia_collector.beam
│   │       │   │   │   ├── prometheus_model.beam
│   │       │   │   │   ├── prometheus_model_helpers.beam
│   │       │   │   │   ├── prometheus_protobuf_format.beam
│   │       │   │   │   ├── prometheus_registry.beam
│   │       │   │   │   ├── prometheus_summary.beam
│   │       │   │   │   ├── prometheus_sup.beam
│   │       │   │   │   ├── prometheus_test_instrumenter.beam
│   │       │   │   │   ├── prometheus_text_format.beam
│   │       │   │   │   ├── prometheus_time.beam
│   │       │   │   │   ├── prometheus_vm_dist_collector.beam
│   │       │   │   │   ├── prometheus_vm_memory_collector.beam
│   │       │   │   │   ├── prometheus_vm_msacc_collector.beam
│   │       │   │   │   ├── prometheus_vm_statistics_collector.beam
│   │       │   │   │   └── prometheus_vm_system_info_collector.beam
│   │       │   │   └── include
│   │       │   │       ├── prometheus.hrl
│   │       │   │       └── prometheus_model.hrl
│   │       │   ├── rabbitmq_management_agent-3.8.16
│   │       │   │   ├── ebin
│   │       │   │   │   ├── Elixir.RabbitMQ.CLI.Ctl.Commands.ResetStatsDbCommand.beam
│   │       │   │   │   ├── exometer_slide.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_app.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_config.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_sup.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_sup_sup.beam
│   │       │   │   │   ├── rabbit_mgmt_data.beam
│   │       │   │   │   ├── rabbit_mgmt_data_compat.beam
│   │       │   │   │   ├── rabbit_mgmt_db_handler.beam
│   │       │   │   │   ├── rabbit_mgmt_external_stats.beam
│   │       │   │   │   ├── rabbit_mgmt_ff.beam
│   │       │   │   │   ├── rabbit_mgmt_format.beam
│   │       │   │   │   ├── rabbit_mgmt_gc.beam
│   │       │   │   │   ├── rabbit_mgmt_metrics_collector.beam
│   │       │   │   │   ├── rabbit_mgmt_metrics_gc.beam
│   │       │   │   │   ├── rabbit_mgmt_storage.beam
│   │       │   │   │   └── rabbitmq_management_agent.app
│   │       │   │   ├── include
│   │       │   │   │   ├── rabbit_mgmt_agent.hrl
│   │       │   │   │   ├── rabbit_mgmt_metrics.hrl
│   │       │   │   │   └── rabbit_mgmt_records.hrl
│   │       │   │   └── priv
│   │       │   │       └── schema
│   │       │   │           └── rabbitmq_management_agent.schema
│   │       │   ├── rabbitmq_prometheus-3.8.16
│   │       │   │   ├── ebin
│   │       │   │   │   ├── prometheus_rabbitmq_core_metrics_collector.beam
│   │       │   │   │   ├── rabbit_prometheus_app.beam
│   │       │   │   │   ├── rabbit_prometheus_dispatcher.beam
│   │       │   │   │   ├── rabbit_prometheus_handler.beam
│   │       │   │   │   └── rabbitmq_prometheus.app
│   │       │   │   └── priv
│   │       │   │       └── schema
│   │       │   │           └── rabbitmq_prometheus.schema
│   │       │   └── rabbitmq_web_dispatch-3.8.16
│   │       │       └── ebin
│   │       │           ├── rabbit_cowboy_middleware.beam
│   │       │           ├── rabbit_cowboy_redirect.beam
│   │       │           ├── rabbit_cowboy_stream_h.beam
│   │       │           ├── rabbit_web_dispatch.beam
│   │       │           ├── rabbit_web_dispatch_app.beam
│   │       │           ├── rabbit_web_dispatch_listing_handler.beam
│   │       │           ├── rabbit_web_dispatch_registry.beam
│   │       │           ├── rabbit_web_dispatch_sup.beam
│   │       │           ├── rabbit_web_dispatch_util.beam
│   │       │           ├── rabbitmq_web_dispatch.app
│   │       │           ├── webmachine_log.beam
│   │       │           └── webmachine_log_handler.beam
│   │       ├── rabbit@e126dc313b3c
│   │       │   ├── DECISION_TAB.LOG
│   │       │   ├── LATEST.LOG
│   │       │   ├── cluster_nodes.config
│   │       │   ├── msg_stores
│   │       │   │   └── vhosts
│   │       │   │       └── 628WB79CIFDYO9LJI6DKMI09L
│   │       │   │           ├── msg_store_persistent
│   │       │   │           │   ├── 0.rdq
│   │       │   │           │   ├── clean.dot
│   │       │   │           │   ├── file_summary.ets
│   │       │   │           │   └── msg_store_index.ets
│   │       │   │           ├── msg_store_transient
│   │       │   │           │   ├── 0.rdq
│   │       │   │           │   ├── clean.dot
│   │       │   │           │   ├── file_summary.ets
│   │       │   │           │   └── msg_store_index.ets
│   │       │   │           └── recovery.dets
│   │       │   ├── nodes_running_at_shutdown
│   │       │   ├── quorum
│   │       │   │   └── rabbit@e126dc313b3c
│   │       │   │       ├── 00000001.wal
│   │       │   │       ├── meta.dets
│   │       │   │       └── names.dets
│   │       │   ├── rabbit_durable_exchange.DCD
│   │       │   ├── rabbit_durable_queue.DCD
│   │       │   ├── rabbit_durable_queue.DCL
│   │       │   ├── rabbit_durable_route.DCD
│   │       │   ├── rabbit_runtime_parameters.DCD
│   │       │   ├── rabbit_serial
│   │       │   ├── rabbit_topic_permission.DCD
│   │       │   ├── rabbit_user.DCD
│   │       │   ├── rabbit_user.DCL
│   │       │   ├── rabbit_user_permission.DCD
│   │       │   ├── rabbit_vhost.DCD
│   │       │   ├── rabbit_vhost.DCL
│   │       │   ├── schema.DAT
│   │       │   └── schema_version
│   │       ├── rabbit@e126dc313b3c-feature_flags
│   │       ├── rabbit@e126dc313b3c-plugins-expand
│   │       │   ├── cowboy-2.8.0
│   │       │   │   └── ebin
│   │       │   │       ├── cowboy.app
│   │       │   │       ├── cowboy.beam
│   │       │   │       ├── cowboy_app.beam
│   │       │   │       ├── cowboy_bstr.beam
│   │       │   │       ├── cowboy_children.beam
│   │       │   │       ├── cowboy_clear.beam
│   │       │   │       ├── cowboy_clock.beam
│   │       │   │       ├── cowboy_compress_h.beam
│   │       │   │       ├── cowboy_constraints.beam
│   │       │   │       ├── cowboy_handler.beam
│   │       │   │       ├── cowboy_http.beam
│   │       │   │       ├── cowboy_http2.beam
│   │       │   │       ├── cowboy_loop.beam
│   │       │   │       ├── cowboy_metrics_h.beam
│   │       │   │       ├── cowboy_middleware.beam
│   │       │   │       ├── cowboy_req.beam
│   │       │   │       ├── cowboy_rest.beam
│   │       │   │       ├── cowboy_router.beam
│   │       │   │       ├── cowboy_static.beam
│   │       │   │       ├── cowboy_stream.beam
│   │       │   │       ├── cowboy_stream_h.beam
│   │       │   │       ├── cowboy_sub_protocol.beam
│   │       │   │       ├── cowboy_sup.beam
│   │       │   │       ├── cowboy_tls.beam
│   │       │   │       ├── cowboy_tracer_h.beam
│   │       │   │       └── cowboy_websocket.beam
│   │       │   ├── cowlib-2.9.1
│   │       │   │   ├── ebin
│   │       │   │   │   ├── cow_base64url.beam
│   │       │   │   │   ├── cow_cookie.beam
│   │       │   │   │   ├── cow_date.beam
│   │       │   │   │   ├── cow_hpack.beam
│   │       │   │   │   ├── cow_http.beam
│   │       │   │   │   ├── cow_http2.beam
│   │       │   │   │   ├── cow_http2_machine.beam
│   │       │   │   │   ├── cow_http_hd.beam
│   │       │   │   │   ├── cow_http_struct_hd.beam
│   │       │   │   │   ├── cow_http_te.beam
│   │       │   │   │   ├── cow_iolists.beam
│   │       │   │   │   ├── cow_link.beam
│   │       │   │   │   ├── cow_mimetypes.beam
│   │       │   │   │   ├── cow_multipart.beam
│   │       │   │   │   ├── cow_qs.beam
│   │       │   │   │   ├── cow_spdy.beam
│   │       │   │   │   ├── cow_sse.beam
│   │       │   │   │   ├── cow_uri.beam
│   │       │   │   │   ├── cow_uri_template.beam
│   │       │   │   │   ├── cow_ws.beam
│   │       │   │   │   └── cowlib.app
│   │       │   │   └── include
│   │       │   │       ├── cow_inline.hrl
│   │       │   │       └── cow_parse.hrl
│   │       │   ├── prometheus-4.6.0
│   │       │   │   ├── ebin
│   │       │   │   │   ├── prometheus.app
│   │       │   │   │   ├── prometheus.beam
│   │       │   │   │   ├── prometheus_boolean.beam
│   │       │   │   │   ├── prometheus_buckets.beam
│   │       │   │   │   ├── prometheus_collector.beam
│   │       │   │   │   ├── prometheus_counter.beam
│   │       │   │   │   ├── prometheus_format.beam
│   │       │   │   │   ├── prometheus_gauge.beam
│   │       │   │   │   ├── prometheus_histogram.beam
│   │       │   │   │   ├── prometheus_http.beam
│   │       │   │   │   ├── prometheus_instrumenter.beam
│   │       │   │   │   ├── prometheus_metric.beam
│   │       │   │   │   ├── prometheus_metric_spec.beam
│   │       │   │   │   ├── prometheus_misc.beam
│   │       │   │   │   ├── prometheus_mnesia.beam
│   │       │   │   │   ├── prometheus_mnesia_collector.beam
│   │       │   │   │   ├── prometheus_model.beam
│   │       │   │   │   ├── prometheus_model_helpers.beam
│   │       │   │   │   ├── prometheus_protobuf_format.beam
│   │       │   │   │   ├── prometheus_registry.beam
│   │       │   │   │   ├── prometheus_summary.beam
│   │       │   │   │   ├── prometheus_sup.beam
│   │       │   │   │   ├── prometheus_test_instrumenter.beam
│   │       │   │   │   ├── prometheus_text_format.beam
│   │       │   │   │   ├── prometheus_time.beam
│   │       │   │   │   ├── prometheus_vm_dist_collector.beam
│   │       │   │   │   ├── prometheus_vm_memory_collector.beam
│   │       │   │   │   ├── prometheus_vm_msacc_collector.beam
│   │       │   │   │   ├── prometheus_vm_statistics_collector.beam
│   │       │   │   │   └── prometheus_vm_system_info_collector.beam
│   │       │   │   └── include
│   │       │   │       ├── prometheus.hrl
│   │       │   │       └── prometheus_model.hrl
│   │       │   ├── rabbitmq_management_agent-3.8.16
│   │       │   │   ├── ebin
│   │       │   │   │   ├── Elixir.RabbitMQ.CLI.Ctl.Commands.ResetStatsDbCommand.beam
│   │       │   │   │   ├── exometer_slide.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_app.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_config.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_sup.beam
│   │       │   │   │   ├── rabbit_mgmt_agent_sup_sup.beam
│   │       │   │   │   ├── rabbit_mgmt_data.beam
│   │       │   │   │   ├── rabbit_mgmt_data_compat.beam
│   │       │   │   │   ├── rabbit_mgmt_db_handler.beam
│   │       │   │   │   ├── rabbit_mgmt_external_stats.beam
│   │       │   │   │   ├── rabbit_mgmt_ff.beam
│   │       │   │   │   ├── rabbit_mgmt_format.beam
│   │       │   │   │   ├── rabbit_mgmt_gc.beam
│   │       │   │   │   ├── rabbit_mgmt_metrics_collector.beam
│   │       │   │   │   ├── rabbit_mgmt_metrics_gc.beam
│   │       │   │   │   ├── rabbit_mgmt_storage.beam
│   │       │   │   │   └── rabbitmq_management_agent.app
│   │       │   │   ├── include
│   │       │   │   │   ├── rabbit_mgmt_agent.hrl
│   │       │   │   │   ├── rabbit_mgmt_metrics.hrl
│   │       │   │   │   └── rabbit_mgmt_records.hrl
│   │       │   │   └── priv
│   │       │   │       └── schema
│   │       │   │           └── rabbitmq_management_agent.schema
│   │       │   ├── rabbitmq_prometheus-3.8.16
│   │       │   │   ├── ebin
│   │       │   │   │   ├── prometheus_rabbitmq_core_metrics_collector.beam
│   │       │   │   │   ├── rabbit_prometheus_app.beam
│   │       │   │   │   ├── rabbit_prometheus_dispatcher.beam
│   │       │   │   │   ├── rabbit_prometheus_handler.beam
│   │       │   │   │   └── rabbitmq_prometheus.app
│   │       │   │   └── priv
│   │       │   │       └── schema
│   │       │   │           └── rabbitmq_prometheus.schema
│   │       │   └── rabbitmq_web_dispatch-3.8.16
│   │       │       └── ebin
│   │       │           ├── rabbit_cowboy_middleware.beam
│   │       │           ├── rabbit_cowboy_redirect.beam
│   │       │           ├── rabbit_cowboy_stream_h.beam
│   │       │           ├── rabbit_web_dispatch.beam
│   │       │           ├── rabbit_web_dispatch_app.beam
│   │       │           ├── rabbit_web_dispatch_listing_handler.beam
│   │       │           ├── rabbit_web_dispatch_registry.beam
│   │       │           ├── rabbit_web_dispatch_sup.beam
│   │       │           ├── rabbit_web_dispatch_util.beam
│   │       │           ├── rabbitmq_web_dispatch.app
│   │       │           ├── webmachine_log.beam
│   │       │           └── webmachine_log_handler.beam
│   │       ├── rabbit@ffa8cb643571
│   │       │   ├── DECISION_TAB.LOG
│   │       │   ├── LATEST.LOG
│   │       │   ├── cluster_nodes.config
│   │       │   ├── msg_stores
│   │       │   │   └── vhosts
│   │       │   │       └── 628WB79CIFDYO9LJI6DKMI09L
│   │       │   │           ├── msg_store_persistent
│   │       │   │           │   ├── 0.rdq
│   │       │   │           │   ├── clean.dot
│   │       │   │           │   ├── file_summary.ets
│   │       │   │           │   └── msg_store_index.ets
│   │       │   │           ├── msg_store_transient
│   │       │   │           │   ├── 0.rdq
│   │       │   │           │   ├── clean.dot
│   │       │   │           │   ├── file_summary.ets
│   │       │   │           │   └── msg_store_index.ets
│   │       │   │           └── recovery.dets
│   │       │   ├── nodes_running_at_shutdown
│   │       │   ├── quorum
│   │       │   │   └── rabbit@ffa8cb643571
│   │       │   │       ├── 00000001.wal
│   │       │   │       ├── meta.dets
│   │       │   │       └── names.dets
│   │       │   ├── rabbit_durable_exchange.DCD
│   │       │   ├── rabbit_durable_exchange.DCL
│   │       │   ├── rabbit_durable_queue.DCD
│   │       │   ├── rabbit_durable_queue.DCL
│   │       │   ├── rabbit_durable_route.DCD
│   │       │   ├── rabbit_runtime_parameters.DCD
│   │       │   ├── rabbit_runtime_parameters.DCL
│   │       │   ├── rabbit_serial
│   │       │   ├── rabbit_topic_permission.DCD
│   │       │   ├── rabbit_user.DCD
│   │       │   ├── rabbit_user_permission.DCD
│   │       │   ├── rabbit_user_permission.DCL
│   │       │   ├── rabbit_vhost.DCD
│   │       │   ├── schema.DAT
│   │       │   └── schema_version
│   │       ├── rabbit@ffa8cb643571-feature_flags
│   │       └── rabbit@ffa8cb643571-plugins-expand
│   │           ├── cowboy-2.8.0
│   │           │   └── ebin
│   │           │       ├── cowboy.app
│   │           │       ├── cowboy.beam
│   │           │       ├── cowboy_app.beam
│   │           │       ├── cowboy_bstr.beam
│   │           │       ├── cowboy_children.beam
│   │           │       ├── cowboy_clear.beam
│   │           │       ├── cowboy_clock.beam
│   │           │       ├── cowboy_compress_h.beam
│   │           │       ├── cowboy_constraints.beam
│   │           │       ├── cowboy_handler.beam
│   │           │       ├── cowboy_http.beam
│   │           │       ├── cowboy_http2.beam
│   │           │       ├── cowboy_loop.beam
│   │           │       ├── cowboy_metrics_h.beam
│   │           │       ├── cowboy_middleware.beam
│   │           │       ├── cowboy_req.beam
│   │           │       ├── cowboy_rest.beam
│   │           │       ├── cowboy_router.beam
│   │           │       ├── cowboy_static.beam
│   │           │       ├── cowboy_stream.beam
│   │           │       ├── cowboy_stream_h.beam
│   │           │       ├── cowboy_sub_protocol.beam
│   │           │       ├── cowboy_sup.beam
│   │           │       ├── cowboy_tls.beam
│   │           │       ├── cowboy_tracer_h.beam
│   │           │       └── cowboy_websocket.beam
│   │           ├── cowlib-2.9.1
│   │           │   ├── ebin
│   │           │   │   ├── cow_base64url.beam
│   │           │   │   ├── cow_cookie.beam
│   │           │   │   ├── cow_date.beam
│   │           │   │   ├── cow_hpack.beam
│   │           │   │   ├── cow_http.beam
│   │           │   │   ├── cow_http2.beam
│   │           │   │   ├── cow_http2_machine.beam
│   │           │   │   ├── cow_http_hd.beam
│   │           │   │   ├── cow_http_struct_hd.beam
│   │           │   │   ├── cow_http_te.beam
│   │           │   │   ├── cow_iolists.beam
│   │           │   │   ├── cow_link.beam
│   │           │   │   ├── cow_mimetypes.beam
│   │           │   │   ├── cow_multipart.beam
│   │           │   │   ├── cow_qs.beam
│   │           │   │   ├── cow_spdy.beam
│   │           │   │   ├── cow_sse.beam
│   │           │   │   ├── cow_uri.beam
│   │           │   │   ├── cow_uri_template.beam
│   │           │   │   ├── cow_ws.beam
│   │           │   │   └── cowlib.app
│   │           │   └── include
│   │           │       ├── cow_inline.hrl
│   │           │       └── cow_parse.hrl
│   │           ├── prometheus-4.6.0
│   │           │   ├── ebin
│   │           │   │   ├── prometheus.app
│   │           │   │   ├── prometheus.beam
│   │           │   │   ├── prometheus_boolean.beam
│   │           │   │   ├── prometheus_buckets.beam
│   │           │   │   ├── prometheus_collector.beam
│   │           │   │   ├── prometheus_counter.beam
│   │           │   │   ├── prometheus_format.beam
│   │           │   │   ├── prometheus_gauge.beam
│   │           │   │   ├── prometheus_histogram.beam
│   │           │   │   ├── prometheus_http.beam
│   │           │   │   ├── prometheus_instrumenter.beam
│   │           │   │   ├── prometheus_metric.beam
│   │           │   │   ├── prometheus_metric_spec.beam
│   │           │   │   ├── prometheus_misc.beam
│   │           │   │   ├── prometheus_mnesia.beam
│   │           │   │   ├── prometheus_mnesia_collector.beam
│   │           │   │   ├── prometheus_model.beam
│   │           │   │   ├── prometheus_model_helpers.beam
│   │           │   │   ├── prometheus_protobuf_format.beam
│   │           │   │   ├── prometheus_registry.beam
│   │           │   │   ├── prometheus_summary.beam
│   │           │   │   ├── prometheus_sup.beam
│   │           │   │   ├── prometheus_test_instrumenter.beam
│   │           │   │   ├── prometheus_text_format.beam
│   │           │   │   ├── prometheus_time.beam
│   │           │   │   ├── prometheus_vm_dist_collector.beam
│   │           │   │   ├── prometheus_vm_memory_collector.beam
│   │           │   │   ├── prometheus_vm_msacc_collector.beam
│   │           │   │   ├── prometheus_vm_statistics_collector.beam
│   │           │   │   └── prometheus_vm_system_info_collector.beam
│   │           │   └── include
│   │           │       ├── prometheus.hrl
│   │           │       └── prometheus_model.hrl
│   │           ├── rabbitmq_management_agent-3.8.16
│   │           │   ├── ebin
│   │           │   │   ├── Elixir.RabbitMQ.CLI.Ctl.Commands.ResetStatsDbCommand.beam
│   │           │   │   ├── exometer_slide.beam
│   │           │   │   ├── rabbit_mgmt_agent_app.beam
│   │           │   │   ├── rabbit_mgmt_agent_config.beam
│   │           │   │   ├── rabbit_mgmt_agent_sup.beam
│   │           │   │   ├── rabbit_mgmt_agent_sup_sup.beam
│   │           │   │   ├── rabbit_mgmt_data.beam
│   │           │   │   ├── rabbit_mgmt_data_compat.beam
│   │           │   │   ├── rabbit_mgmt_db_handler.beam
│   │           │   │   ├── rabbit_mgmt_external_stats.beam
│   │           │   │   ├── rabbit_mgmt_ff.beam
│   │           │   │   ├── rabbit_mgmt_format.beam
│   │           │   │   ├── rabbit_mgmt_gc.beam
│   │           │   │   ├── rabbit_mgmt_metrics_collector.beam
│   │           │   │   ├── rabbit_mgmt_metrics_gc.beam
│   │           │   │   ├── rabbit_mgmt_storage.beam
│   │           │   │   └── rabbitmq_management_agent.app
│   │           │   ├── include
│   │           │   │   ├── rabbit_mgmt_agent.hrl
│   │           │   │   ├── rabbit_mgmt_metrics.hrl
│   │           │   │   └── rabbit_mgmt_records.hrl
│   │           │   └── priv
│   │           │       └── schema
│   │           │           └── rabbitmq_management_agent.schema
│   │           ├── rabbitmq_prometheus-3.8.16
│   │           │   ├── ebin
│   │           │   │   ├── prometheus_rabbitmq_core_metrics_collector.beam
│   │           │   │   ├── rabbit_prometheus_app.beam
│   │           │   │   ├── rabbit_prometheus_dispatcher.beam
│   │           │   │   ├── rabbit_prometheus_handler.beam
│   │           │   │   └── rabbitmq_prometheus.app
│   │           │   └── priv
│   │           │       └── schema
│   │           │           └── rabbitmq_prometheus.schema
│   │           └── rabbitmq_web_dispatch-3.8.16
│   │               └── ebin
│   │                   ├── rabbit_cowboy_middleware.beam
│   │                   ├── rabbit_cowboy_redirect.beam
│   │                   ├── rabbit_cowboy_stream_h.beam
│   │                   ├── rabbit_web_dispatch.beam
│   │                   ├── rabbit_web_dispatch_app.beam
│   │                   ├── rabbit_web_dispatch_listing_handler.beam
│   │                   ├── rabbit_web_dispatch_registry.beam
│   │                   ├── rabbit_web_dispatch_sup.beam
│   │                   ├── rabbit_web_dispatch_util.beam
│   │                   ├── rabbitmq_web_dispatch.app
│   │                   ├── webmachine_log.beam
│   │                   └── webmachine_log_handler.beam
│   └── redis
│       └── data
│           └── dump.rdb
├── deploy
│   ├── compose.dev
│   │   ├── docker-compose.app.dev.yml
│   │   ├── docker-compose.dev.yml
│   │   ├── docker-compose.influxdb.dev.yml
│   │   ├── docker-compose.mysql.dev.yml
│   │   ├── docker-compose.network.dev.yml
│   │   ├── docker-compose.pg.dev.yml
│   │   ├── docker-compose.rabbitmq.dev.yml
│   │   └── docker-compose.redis.dev.yml
│   ├── config.dev.txt
│   ├── const.dev.sh
│   └── dockerfile
│       ├── dockerfile.dev
│       └── run.sh
├── docs
│   └── data_model.md
├── go.mod
├── go.sum
├── internal
│   ├── api
│   │   └── v2
│   │       ├── demo.api.go
│   │       ├── main.go
│   │       ├── mock
│   │       │   ├── demo.mock.go
│   │       │   ├── mock.go
│   │       │   ├── signin.mock.go
│   │       │   ├── sys_address.mock.go
│   │       │   ├── sys_dict.mock.go
│   │       │   ├── sys_district.mock.go
│   │       │   ├── sys_menu.mock.go
│   │       │   ├── sys_role.mock.go
│   │       │   └── sys_user.mock.go
│   │       ├── signin.api.go
│   │       ├── sys_address.api.go
│   │       ├── sys_dict.api.go
│   │       ├── sys_district.api.go
│   │       ├── sys_menu.api.go
│   │       ├── sys_role.api.go
│   │       └── sys_user.api.go
│   ├── app
│   │   ├── app.go
│   │   ├── auth.go
│   │   ├── casbin.go
│   │   ├── contextx
│   │   │   └── contextx.go
│   │   ├── ginx
│   │   │   └── ginx.go
│   │   ├── influxdb.go
│   │   ├── injector.go
│   │   ├── logger.go
│   │   ├── middleware
│   │   │   ├── healthchek
│   │   │   │   └── mw_healthcheck.go
│   │   │   ├── middleware.go
│   │   │   ├── mw_auth.go
│   │   │   ├── mw_casbin.go
│   │   │   ├── mw_copy_body.go
│   │   │   ├── mw_cors.go
│   │   │   ├── mw_logger.go
│   │   │   ├── mw_rate_limiter.go
│   │   │   ├── mw_recover.go
│   │   │   ├── mw_trace.go
│   │   │   └── mw_www.go
│   │   ├── module
│   │   │   └── adapter
│   │   │       └── casbin.go
│   │   ├── rabbitmq.go
│   │   ├── redis.go
│   │   ├── resolver
│   │   │   └── resolver.go
│   │   ├── schema
│   │   │   ├── demo.sch.go
│   │   │   ├── main.go
│   │   │   ├── signin.sch.go
│   │   │   ├── sys_address.sch.go
│   │   │   ├── sys_dict.sch.go
│   │   │   ├── sys_dict_item.sch.go
│   │   │   ├── sys_district.sch.go
│   │   │   ├── sys_menu.sch.go
│   │   │   ├── sys_menu_action.sch.go
│   │   │   ├── sys_menu_action_resouce.sch.go
│   │   │   ├── sys_role.sch.go
│   │   │   └── sys_user.sch.go
│   │   ├── service
│   │   │   ├── casbin.srv.go
│   │   │   ├── common.srv.go
│   │   │   ├── demo.srv.go
│   │   │   ├── main.go
│   │   │   ├── signin.srv.go
│   │   │   ├── sys_address.srv.go
│   │   │   ├── sys_dict.srv.go
│   │   │   ├── sys_district.srv.go
│   │   │   ├── sys_menu.srv.go
│   │   │   ├── sys_role.srv.go
│   │   │   └── sys_user.srv.go
│   │   ├── swagger
│   │   │   ├── docs.go
│   │   │   ├── swagger.json
│   │   │   └── swagger.yaml
│   │   ├── test
│   │   │   ├── data
│   │   │   │   └── jwt_auth.db
│   │   │   ├── doc.go
│   │   │   ├── t_demo_test.go
│   │   │   ├── t_init_test.go
│   │   │   ├── t_menu_test.go
│   │   │   ├── t_role_test.go
│   │   │   └── t_user_test.go
│   │   ├── vcode.go
│   │   ├── web.go
│   │   ├── wire.go
│   │   └── wire_gen.go
│   ├── gen
│   │   └── ent
│   │       ├── client.go
│   │       ├── debug.go
│   │       ├── ent.go
│   │       ├── enttest
│   │       │   └── enttest.go
│   │       ├── entviz.go
│   │       ├── hook
│   │       │   └── hook.go
│   │       ├── internal
│   │       │   ├── schema.go
│   │       │   └── schemaconfig.go
│   │       ├── migrate
│   │       │   ├── migrate.go
│   │       │   └── schema.go
│   │       ├── mutation.go
│   │       ├── mutation_input.go
│   │       ├── predicate
│   │       │   └── predicate.go
│   │       ├── runtime
│   │       │   └── runtime.go
│   │       ├── runtime.go
│   │       ├── schema-viz.html
│   │       ├── stringer.go
│   │       ├── sysaddress
│   │       │   ├── sysaddress.go
│   │       │   └── where.go
│   │       ├── sysaddress.go
│   │       ├── sysaddress_create.go
│   │       ├── sysaddress_delete.go
│   │       ├── sysaddress_query.go
│   │       ├── sysaddress_update.go
│   │       ├── sysdict
│   │       │   ├── sysdict.go
│   │       │   └── where.go
│   │       ├── sysdict.go
│   │       ├── sysdict_create.go
│   │       ├── sysdict_delete.go
│   │       ├── sysdict_query.go
│   │       ├── sysdict_update.go
│   │       ├── sysdictitem
│   │       │   ├── sysdictitem.go
│   │       │   └── where.go
│   │       ├── sysdictitem.go
│   │       ├── sysdictitem_create.go
│   │       ├── sysdictitem_delete.go
│   │       ├── sysdictitem_query.go
│   │       ├── sysdictitem_update.go
│   │       ├── sysdistrict
│   │       │   ├── sysdistrict.go
│   │       │   └── where.go
│   │       ├── sysdistrict.go
│   │       ├── sysdistrict_create.go
│   │       ├── sysdistrict_delete.go
│   │       ├── sysdistrict_query.go
│   │       ├── sysdistrict_update.go
│   │       ├── sysjwtblock
│   │       │   ├── sysjwtblock.go
│   │       │   └── where.go
│   │       ├── sysjwtblock.go
│   │       ├── sysjwtblock_create.go
│   │       ├── sysjwtblock_delete.go
│   │       ├── sysjwtblock_query.go
│   │       ├── sysjwtblock_update.go
│   │       ├── syslogging
│   │       │   ├── syslogging.go
│   │       │   └── where.go
│   │       ├── syslogging.go
│   │       ├── syslogging_create.go
│   │       ├── syslogging_delete.go
│   │       ├── syslogging_query.go
│   │       ├── syslogging_update.go
│   │       ├── sysmenu
│   │       │   ├── sysmenu.go
│   │       │   └── where.go
│   │       ├── sysmenu.go
│   │       ├── sysmenu_create.go
│   │       ├── sysmenu_delete.go
│   │       ├── sysmenu_query.go
│   │       ├── sysmenu_update.go
│   │       ├── sysmenuaction
│   │       │   ├── sysmenuaction.go
│   │       │   └── where.go
│   │       ├── sysmenuaction.go
│   │       ├── sysmenuaction_create.go
│   │       ├── sysmenuaction_delete.go
│   │       ├── sysmenuaction_query.go
│   │       ├── sysmenuaction_update.go
│   │       ├── sysmenuactionresource
│   │       │   ├── sysmenuactionresource.go
│   │       │   └── where.go
│   │       ├── sysmenuactionresource.go
│   │       ├── sysmenuactionresource_create.go
│   │       ├── sysmenuactionresource_delete.go
│   │       ├── sysmenuactionresource_query.go
│   │       ├── sysmenuactionresource_update.go
│   │       ├── sysrole
│   │       │   ├── sysrole.go
│   │       │   └── where.go
│   │       ├── sysrole.go
│   │       ├── sysrole_create.go
│   │       ├── sysrole_delete.go
│   │       ├── sysrole_query.go
│   │       ├── sysrole_update.go
│   │       ├── sysrolemenu
│   │       │   ├── sysrolemenu.go
│   │       │   └── where.go
│   │       ├── sysrolemenu.go
│   │       ├── sysrolemenu_create.go
│   │       ├── sysrolemenu_delete.go
│   │       ├── sysrolemenu_query.go
│   │       ├── sysrolemenu_update.go
│   │       ├── sysuser
│   │       │   ├── sysuser.go
│   │       │   └── where.go
│   │       ├── sysuser.go
│   │       ├── sysuser_create.go
│   │       ├── sysuser_delete.go
│   │       ├── sysuser_query.go
│   │       ├── sysuser_update.go
│   │       ├── sysuserrole
│   │       │   ├── sysuserrole.go
│   │       │   └── where.go
│   │       ├── sysuserrole.go
│   │       ├── sysuserrole_create.go
│   │       ├── sysuserrole_delete.go
│   │       ├── sysuserrole_query.go
│   │       ├── sysuserrole_update.go
│   │       ├── tx.go
│   │       ├── xxxdemo
│   │       │   ├── where.go
│   │       │   └── xxxdemo.go
│   │       ├── xxxdemo.go
│   │       ├── xxxdemo_create.go
│   │       ├── xxxdemo_delete.go
│   │       ├── xxxdemo_query.go
│   │       └── xxxdemo_update.go
│   ├── router
│   │   ├── r_api.go
│   │   ├── router.go
│   │   ├── v2_demo.go
│   │   ├── v2_sys_address.go
│   │   ├── v2_sys_dict.go
│   │   ├── v2_sys_district.go
│   │   ├── v2_sys_menu.go
│   │   ├── v2_sys_role.go
│   │   └── v2_sys_user.go
│   └── schema
│       ├── client.go
│       ├── entity
│       │   ├── entity.go
│       │   ├── hooks.go
│       │   ├── sys_address.entity.go
│       │   ├── sys_casbin_rule.entity.go
│       │   ├── sys_dict.entity.go
│       │   ├── sys_dict_item.entity.go
│       │   ├── sys_district.entity.go
│       │   ├── sys_jwt_blocklist.go
│       │   ├── sys_logging.go
│       │   ├── sys_menu.entity.go
│       │   ├── sys_menu_action.entity.go
│       │   ├── sys_menu_action_resource.entity.go
│       │   ├── sys_role.entity.go
│       │   ├── sys_role_menu.entity.go
│       │   ├── sys_user.entity.go
│       │   ├── sys_user_role.entity.go
│       │   └── xxx_demo.go
│       ├── mixin
│       │   ├── active.go
│       │   ├── creator.go
│       │   ├── hooks.go
│       │   ├── id.go
│       │   ├── meno.go
│       │   ├── mptt_tree.go
│       │   ├── organ.go
│       │   ├── owner.go
│       │   ├── softdel.go
│       │   ├── sort.go
│       │   ├── time.go
│       │   └── tree.go
│       ├── repo
│       │   ├── common.go
│       │   ├── demo.repo.go
│       │   ├── main.go
│       │   ├── sys_address.repo.go
│       │   ├── sys_dict.repo.go
│       │   ├── sys_dict_item.repo.go
│       │   ├── sys_district.repo.go
│       │   ├── sys_menu.repo.go
│       │   ├── sys_menu_action.repo.go
│       │   ├── sys_menu_action_resource.repo.go
│       │   ├── sys_role.repo.go
│       │   ├── sys_role_menu.repo.go
│       │   ├── sys_user.repo.go
│       │   ├── sys_user_role.repo.go
│       │   └── trans.repo.go
│       └── template
│           ├── debug.tpl
│           ├── mutation_input.tpl
│           └── stringer.tpl
├── pkg
│   ├── auth
│   │   ├── auth.go
│   │   └── jwtauth
│   │       ├── auth.go
│   │       ├── auth_test.go
│   │       ├── store
│   │       │   ├── buntdb
│   │       │   │   ├── buntdb.go
│   │       │   │   └── buntdb_test.go
│   │       │   └── redis
│   │       │       ├── redis.go
│   │       │       └── redis_test.go
│   │       ├── store.go
│   │       └── token.go
│   ├── config
│   │   ├── config.go
│   │   └── option
│   │       ├── config.captcha.go
│   │       ├── config.casbin.go
│   │       ├── config.ent.go
│   │       ├── config.http.go
│   │       ├── config.influxdb.go
│   │       ├── config.log.go
│   │       ├── config.menu.go
│   │       ├── config.moniter.go
│   │       ├── config.mysql.go
│   │       ├── config.postgres.go
│   │       ├── config.rabbit.go
│   │       ├── config.redis.go
│   │       ├── config.root.go
│   │       ├── config.sqlite.go
│   │       └── config.system.go
│   ├── errors
│   │   ├── errors.go
│   │   └── response.go
│   ├── global
│   │   └── global.go
│   ├── helper
│   │   ├── deepcopier
│   │   │   └── deepcopier.go
│   │   ├── hash
│   │   │   └── hash.go
│   │   ├── json
│   │   │   └── json.go
│   │   ├── pulid
│   │   │   └── pulid.go
│   │   ├── str
│   │   │   ├── conv.go
│   │   │   └── random.go
│   │   ├── structure
│   │   │   └── structure.go
│   │   ├── trace
│   │   │   └── trace.go
│   │   ├── uid
│   │   │   ├── id.go
│   │   │   ├── ulid
│   │   │   │   └── ulid.go
│   │   │   └── uuid
│   │   │       └── uuid.go
│   │   └── yaml
│   │       └── yaml.go
│   ├── logger
│   │   ├── hook
│   │   │   ├── entgo
│   │   │   │   └── entgo.go
│   │   │   ├── gorm
│   │   │   │   └── gorm.go
│   │   │   └── hook.go
│   │   └── logger.go
│   └── vcode
│       ├── store
│       │   └── redis.go
│       └── vcode.go
└── scripts
    ├── build-image.sh
    ├── compose.dev
    │   ├── docker-compose.app.yml
    │   ├── docker-compose.dev.yml
    │   ├── docker-compose.external.yml
    │   ├── docker-compose.influxdb.yml
    │   ├── docker-compose.mysql.yml
    │   ├── docker-compose.network.ipv4.yml
    │   ├── docker-compose.network.ipv6.yml
    │   ├── docker-compose.pg.yml
    │   ├── docker-compose.rabbitmq.yml
    │   └── docker-compose.redis.yml
    ├── config.dev.txt
    ├── config_init
    │   ├── mysql
    │   │   └── my.cnf
    │   └── redis
    │       └── redis.conf
    ├── const.dev.sh
    ├── ctrl.dev.sh
    ├── ctrl.help.sh
    ├── g
    │   ├── address.yaml
    │   ├── district.yaml
    │   └── organ.yaml
    ├── simple
    │   ├── docker.golang.sh
    │   ├── docker.influxdb.sh
    │   └── docker.rabbitmq.sh
    ├── sql
    │   ├── district.csv
    │   ├── init_mysql.sql
    │   └── init_postgres.sql
    └── utils.sh
`

load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "anchore_image.go",
        "anchore_image_list.go",
        "anchore_image_tag_summary.go",
        "api_client.go",
        "api_error_response.go",
        "api_response.go",
        "catalog_api.go",
        "configuration.go",
        "content_response.go",
        "default_api.go",
        "feed_group_metadata.go",
        "feed_metadata.go",
        "gate_spec.go",
        "image_analysis_report.go",
        "image_analysis_request.go",
        "image_content.go",
        "image_content_api.go",
        "image_detail.go",
        "image_filter.go",
        "image_ref.go",
        "image_selection_rule.go",
        "image_tags_api.go",
        "images_api.go",
        "mapping_rule.go",
        "policies_api.go",
        "policy.go",
        "policy_api.go",
        "policy_bundle.go",
        "policy_bundle_list.go",
        "policy_bundle_record.go",
        "policy_evaluation.go",
        "policy_evaluation_api.go",
        "policy_rule.go",
        "policy_rule_params.go",
        "prune_candidate.go",
        "prune_candidate_list.go",
        "registries_api.go",
        "registry_configuration.go",
        "registry_configuration_list.go",
        "repository_tag_list.go",
        "service.go",
        "service_list.go",
        "services_api.go",
        "status_response.go",
        "subscription.go",
        "subscription_list.go",
        "subscription_request.go",
        "subscription_update.go",
        "subscriptions_api.go",
        "system_api.go",
        "system_status_response.go",
        "trigger_param_spec.go",
        "trigger_spec.go",
        "vulnerabilities_api.go",
        "vulnerability.go",
        "vulnerability_list.go",
        "whitelist.go",
        "whitelist_item.go",
    ],
    importpath = "github.com/omnicate/loltel/infra/paranoia-dashboard/go-anchore",
    visibility = ["//visibility:public"],
    deps = [
        "//infra/paranoia-dashboard/go-anchore/vendor/golang.org/x/net/context:go_default_library",
        "//infra/paranoia-dashboard/go-anchore/vendor/golang.org/x/oauth2:go_default_library",
    ],
)
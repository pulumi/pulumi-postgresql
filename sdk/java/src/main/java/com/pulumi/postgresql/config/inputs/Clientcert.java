// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.postgresql.config.inputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class Clientcert {
    private String cert;
    private String key;
    private @Nullable Boolean sslinline;

    private Clientcert() {}
    public String cert() {
        return this.cert;
    }
    public String key() {
        return this.key;
    }
    public Optional<Boolean> sslinline() {
        return Optional.ofNullable(this.sslinline);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(Clientcert defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String cert;
        private String key;
        private @Nullable Boolean sslinline;
        public Builder() {}
        public Builder(Clientcert defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.cert = defaults.cert;
    	      this.key = defaults.key;
    	      this.sslinline = defaults.sslinline;
        }

        @CustomType.Setter
        public Builder cert(String cert) {
            this.cert = Objects.requireNonNull(cert);
            return this;
        }
        @CustomType.Setter
        public Builder key(String key) {
            this.key = Objects.requireNonNull(key);
            return this;
        }
        @CustomType.Setter
        public Builder sslinline(@Nullable Boolean sslinline) {
            this.sslinline = sslinline;
            return this;
        }
        public Clientcert build() {
            final var _resultValue = new Clientcert();
            _resultValue.cert = cert;
            _resultValue.key = key;
            _resultValue.sslinline = sslinline;
            return _resultValue;
        }
    }
}

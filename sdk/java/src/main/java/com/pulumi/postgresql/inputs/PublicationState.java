// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.postgresql.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PublicationState extends com.pulumi.resources.ResourceArgs {

    public static final PublicationState Empty = new PublicationState();

    /**
     * Should be ALL TABLES added to the publication. Defaults to &#39;false&#39;
     * 
     */
    @Import(name="allTables")
    private @Nullable Output<Boolean> allTables;

    /**
     * @return Should be ALL TABLES added to the publication. Defaults to &#39;false&#39;
     * 
     */
    public Optional<Output<Boolean>> allTables() {
        return Optional.ofNullable(this.allTables);
    }

    /**
     * Which database to create the publication on. Defaults to provider database.
     * 
     */
    @Import(name="database")
    private @Nullable Output<String> database;

    /**
     * @return Which database to create the publication on. Defaults to provider database.
     * 
     */
    public Optional<Output<String>> database() {
        return Optional.ofNullable(this.database);
    }

    /**
     * Should all subsequent resources of the publication be dropped. Defaults to &#39;false&#39;
     * 
     */
    @Import(name="dropCascade")
    private @Nullable Output<Boolean> dropCascade;

    /**
     * @return Should all subsequent resources of the publication be dropped. Defaults to &#39;false&#39;
     * 
     */
    public Optional<Output<Boolean>> dropCascade() {
        return Optional.ofNullable(this.dropCascade);
    }

    /**
     * The name of the publication.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the publication.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Who owns the publication. Defaults to provider user.
     * 
     */
    @Import(name="owner")
    private @Nullable Output<String> owner;

    /**
     * @return Who owns the publication. Defaults to provider user.
     * 
     */
    public Optional<Output<String>> owner() {
        return Optional.ofNullable(this.owner);
    }

    /**
     * Which &#39;publish&#39; options should be turned on. Default to &#39;insert&#39;,&#39;update&#39;,&#39;delete&#39;
     * 
     */
    @Import(name="publishParams")
    private @Nullable Output<List<String>> publishParams;

    /**
     * @return Which &#39;publish&#39; options should be turned on. Default to &#39;insert&#39;,&#39;update&#39;,&#39;delete&#39;
     * 
     */
    public Optional<Output<List<String>>> publishParams() {
        return Optional.ofNullable(this.publishParams);
    }

    /**
     * Should be option &#39;publish_via_partition_root&#39; be turned on. Default to &#39;false&#39;
     * 
     */
    @Import(name="publishViaPartitionRootParam")
    private @Nullable Output<Boolean> publishViaPartitionRootParam;

    /**
     * @return Should be option &#39;publish_via_partition_root&#39; be turned on. Default to &#39;false&#39;
     * 
     */
    public Optional<Output<Boolean>> publishViaPartitionRootParam() {
        return Optional.ofNullable(this.publishViaPartitionRootParam);
    }

    /**
     * Which tables add to the publication. By defaults no tables added. Format of table is `&lt;schema_name&gt;.&lt;table_name&gt;`. If `&lt;schema_name&gt;` is not specified - default database schema will be used.
     * 
     */
    @Import(name="tables")
    private @Nullable Output<List<String>> tables;

    /**
     * @return Which tables add to the publication. By defaults no tables added. Format of table is `&lt;schema_name&gt;.&lt;table_name&gt;`. If `&lt;schema_name&gt;` is not specified - default database schema will be used.
     * 
     */
    public Optional<Output<List<String>>> tables() {
        return Optional.ofNullable(this.tables);
    }

    private PublicationState() {}

    private PublicationState(PublicationState $) {
        this.allTables = $.allTables;
        this.database = $.database;
        this.dropCascade = $.dropCascade;
        this.name = $.name;
        this.owner = $.owner;
        this.publishParams = $.publishParams;
        this.publishViaPartitionRootParam = $.publishViaPartitionRootParam;
        this.tables = $.tables;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PublicationState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PublicationState $;

        public Builder() {
            $ = new PublicationState();
        }

        public Builder(PublicationState defaults) {
            $ = new PublicationState(Objects.requireNonNull(defaults));
        }

        /**
         * @param allTables Should be ALL TABLES added to the publication. Defaults to &#39;false&#39;
         * 
         * @return builder
         * 
         */
        public Builder allTables(@Nullable Output<Boolean> allTables) {
            $.allTables = allTables;
            return this;
        }

        /**
         * @param allTables Should be ALL TABLES added to the publication. Defaults to &#39;false&#39;
         * 
         * @return builder
         * 
         */
        public Builder allTables(Boolean allTables) {
            return allTables(Output.of(allTables));
        }

        /**
         * @param database Which database to create the publication on. Defaults to provider database.
         * 
         * @return builder
         * 
         */
        public Builder database(@Nullable Output<String> database) {
            $.database = database;
            return this;
        }

        /**
         * @param database Which database to create the publication on. Defaults to provider database.
         * 
         * @return builder
         * 
         */
        public Builder database(String database) {
            return database(Output.of(database));
        }

        /**
         * @param dropCascade Should all subsequent resources of the publication be dropped. Defaults to &#39;false&#39;
         * 
         * @return builder
         * 
         */
        public Builder dropCascade(@Nullable Output<Boolean> dropCascade) {
            $.dropCascade = dropCascade;
            return this;
        }

        /**
         * @param dropCascade Should all subsequent resources of the publication be dropped. Defaults to &#39;false&#39;
         * 
         * @return builder
         * 
         */
        public Builder dropCascade(Boolean dropCascade) {
            return dropCascade(Output.of(dropCascade));
        }

        /**
         * @param name The name of the publication.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the publication.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param owner Who owns the publication. Defaults to provider user.
         * 
         * @return builder
         * 
         */
        public Builder owner(@Nullable Output<String> owner) {
            $.owner = owner;
            return this;
        }

        /**
         * @param owner Who owns the publication. Defaults to provider user.
         * 
         * @return builder
         * 
         */
        public Builder owner(String owner) {
            return owner(Output.of(owner));
        }

        /**
         * @param publishParams Which &#39;publish&#39; options should be turned on. Default to &#39;insert&#39;,&#39;update&#39;,&#39;delete&#39;
         * 
         * @return builder
         * 
         */
        public Builder publishParams(@Nullable Output<List<String>> publishParams) {
            $.publishParams = publishParams;
            return this;
        }

        /**
         * @param publishParams Which &#39;publish&#39; options should be turned on. Default to &#39;insert&#39;,&#39;update&#39;,&#39;delete&#39;
         * 
         * @return builder
         * 
         */
        public Builder publishParams(List<String> publishParams) {
            return publishParams(Output.of(publishParams));
        }

        /**
         * @param publishParams Which &#39;publish&#39; options should be turned on. Default to &#39;insert&#39;,&#39;update&#39;,&#39;delete&#39;
         * 
         * @return builder
         * 
         */
        public Builder publishParams(String... publishParams) {
            return publishParams(List.of(publishParams));
        }

        /**
         * @param publishViaPartitionRootParam Should be option &#39;publish_via_partition_root&#39; be turned on. Default to &#39;false&#39;
         * 
         * @return builder
         * 
         */
        public Builder publishViaPartitionRootParam(@Nullable Output<Boolean> publishViaPartitionRootParam) {
            $.publishViaPartitionRootParam = publishViaPartitionRootParam;
            return this;
        }

        /**
         * @param publishViaPartitionRootParam Should be option &#39;publish_via_partition_root&#39; be turned on. Default to &#39;false&#39;
         * 
         * @return builder
         * 
         */
        public Builder publishViaPartitionRootParam(Boolean publishViaPartitionRootParam) {
            return publishViaPartitionRootParam(Output.of(publishViaPartitionRootParam));
        }

        /**
         * @param tables Which tables add to the publication. By defaults no tables added. Format of table is `&lt;schema_name&gt;.&lt;table_name&gt;`. If `&lt;schema_name&gt;` is not specified - default database schema will be used.
         * 
         * @return builder
         * 
         */
        public Builder tables(@Nullable Output<List<String>> tables) {
            $.tables = tables;
            return this;
        }

        /**
         * @param tables Which tables add to the publication. By defaults no tables added. Format of table is `&lt;schema_name&gt;.&lt;table_name&gt;`. If `&lt;schema_name&gt;` is not specified - default database schema will be used.
         * 
         * @return builder
         * 
         */
        public Builder tables(List<String> tables) {
            return tables(Output.of(tables));
        }

        /**
         * @param tables Which tables add to the publication. By defaults no tables added. Format of table is `&lt;schema_name&gt;.&lt;table_name&gt;`. If `&lt;schema_name&gt;` is not specified - default database schema will be used.
         * 
         * @return builder
         * 
         */
        public Builder tables(String... tables) {
            return tables(List.of(tables));
        }

        public PublicationState build() {
            return $;
        }
    }

}
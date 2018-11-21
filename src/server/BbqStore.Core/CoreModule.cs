using System;
using System.Collections.Generic;
using System.Text;
using Autofac;
using BbqStore.Core.Database;
using BbqStore.Core.Services;
using Marten;

namespace BbqStore.Core
{
    public class CoreModule : Module
    {
        protected override void Load(ContainerBuilder builder)
        {
            base.Load(builder);

            builder.RegisterType<ConfigurationService>().As<IConfigurationService>();

            builder.Register(ds =>
            {
                var configurationService = ds.Resolve<IConfigurationService>();
                return DocumentStore.For(_ =>
                {
                    _.Connection(configurationService.Get("Data:BbqStore:ConnectionString"));
                    _.AutoCreateSchemaObjects = AutoCreate.All;

                    _.InitialData.Add(new BbqStoreInitialData(InitialDataSets.Stores));
                    _.InitialData.Add(new BbqStoreInitialData(InitialDataSets.Products));

                });
            }).As<IDocumentStore>().SingleInstance();

            builder.Register(c => c.Resolve<IDocumentStore>().LightweightSession()).As<IDocumentSession>()
                .InstancePerLifetimeScope();

            builder.RegisterAssemblyTypes(ThisAssembly).Where(t => t.Name.EndsWith("Service"))
                .AsImplementedInterfaces();
        }
    }
}

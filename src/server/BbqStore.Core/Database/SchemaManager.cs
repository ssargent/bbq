using System;
using System.Collections.Generic;
using System.Reflection;
using System.Text;
using DbUp;

namespace BbqStore.Core.Database
{
    public class SchemaManager
    {
        public static bool Upgrade(string connectionString)
        {
            var databseConnectionString = connectionString;

            var upgrader = DeployChanges.To.PostgresqlDatabase(databseConnectionString)
                .WithScriptsEmbeddedInAssembly(Assembly.GetExecutingAssembly())
                .JournalToPostgresqlTable("public", "schema_journal")
                .Build();

            var result = upgrader.PerformUpgrade();

            if (!result.Successful)
            {
                Console.ForegroundColor = ConsoleColor.Red;
                Console.WriteLine(result.Error);
                Console.ResetColor();

#if DEBUG
                Console.ReadLine();
#endif
                return false;
            }

            Console.ForegroundColor = ConsoleColor.Green;
            Console.WriteLine("Success!");
            Console.ResetColor();

            return true;
        }
    }
}


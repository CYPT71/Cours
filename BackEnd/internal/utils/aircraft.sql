-- phpMyAdmin SQL Dump
-- version 5.1.0
-- https://www.phpmyadmin.net/
--
-- Hôte : db
-- Généré le : ven. 19 mars 2021 à 21:11
-- Version du serveur :  8.0.23
-- Version de PHP : 7.4.15

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de données : `aircraft`
--

-- --------------------------------------------------------

--
-- Structure de la table `CabinCrew`
--

CREATE TABLE IF NOT EXISTS `CabinCrew` (
  `id` int NOT NULL,
  `amoung` date NOT NULL,
  `fonction` varchar(666) NOT NULL,
  `staff_id` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Structure de la table `Departus`
--

CREATE TABLE IF NOT EXISTS `Departus` (
  `id` int NOT NULL,
  `date` date NOT NULL,
  `pilote` int NOT NULL,
  `copilote` int NOT NULL,
  `aircrew` varchar(250) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `free_places` int NOT NULL,
  `occupied` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Structure de la table `Device`
--

CREATE TABLE IF NOT EXISTS `Device` (
  `id` int NOT NULL,
  `capacity` int NOT NULL,
  `type` varchar(200) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Structure de la table `Employees`
--

CREATE TABLE IF NOT EXISTS `Employees` (
  `id` int NOT NULL,
  `aircrew` boolean NOT NULL,
  `ground` boolean NOT NULL,
  `social_security` int NOT NULL,
  `name` varchar(250) NOT NULL,
  `first_name` varchar(250) NOT NULL,
  `address` varchar(250) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Structure de la table `Fligth`
--

CREATE TABLE IF NOT EXISTS `Fligth` (
  `id` int NOT NULL,
  `id_departures` int NOT NULL,
  `arrival` date NOT NULL,
  `id_route` int NOT NULL,
  `id_device` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Structure de la table `Passenger`
--

CREATE TABLE IF NOT EXISTS `Passenger` (
  `id` int NOT NULL,
  `name` varchar(666) NOT NULL,
  `first_name` varchar(666) NOT NULL,
  `adress` varchar(666) NOT NULL,
  `profession` varchar(666) NOT NULL,
  `bank` varchar(666) NOT NULL,
  `ticket_id` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Structure de la table `Pilote`
--

CREATE TABLE IF NOT EXISTS `Pilote` (
  `id` int NOT NULL,
  `licence` date NOT NULL,
  `among` time NOT NULL,
  `staff_id` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Structure de la table `Route`
--

CREATE TABLE IF NOT EXISTS `Route` (
  `id` int NOT NULL,
  `origin` varchar(250) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `arrival` varchar(250) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Structure de la table `Tickets`
--

CREATE TABLE IF NOT EXISTS `Tickets` (
  `id` int NOT NULL,
  `expire` date NOT NULL,
  `price` int NOT NULL,
  `departures_id` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Index pour les tables déchargées
--

--
-- Index pour la table `CabinCrew`
--
ALTER TABLE `CabinCrew`
  ADD PRIMARY KEY (`id`),
  ADD KEY `staff_id` (`staff_id`);

--
-- Index pour la table `Departus`
--
ALTER TABLE `Departus`
  ADD PRIMARY KEY (`id`),
  ADD KEY `pilote` (`pilote`,`copilote`),
  ADD KEY `copilote` (`copilote`),
  ADD KEY `pilote_2` (`pilote`,`copilote`);

--
-- Index pour la table `Device`
--
ALTER TABLE `Device`
  ADD PRIMARY KEY (`id`);

--
-- Index pour la table `Employees`
--
ALTER TABLE `Employees`
  ADD PRIMARY KEY (`id`);

--
-- Index pour la table `Fligth`
--
ALTER TABLE `Fligth`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `id_departures` (`id_departures`),
  ADD UNIQUE KEY `id_route` (`id_route`),
  ADD KEY `id_device` (`id_device`),
  ADD KEY `id_device_2` (`id_device`);

--
-- Index pour la table `Passenger`
--
ALTER TABLE `Passenger`
  ADD PRIMARY KEY (`id`),
  ADD KEY `ticket_id` (`ticket_id`);

--
-- Index pour la table `Pilote`
--
ALTER TABLE `Pilote`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `staff_id` (`staff_id`);

--
-- Index pour la table `Route`
--
ALTER TABLE `Route`
  ADD PRIMARY KEY (`id`);

--
-- Index pour la table `Tickets`
--
ALTER TABLE `Tickets`
  ADD PRIMARY KEY (`id`),
  ADD KEY `departures_id` (`departures_id`);

--
-- AUTO_INCREMENT pour les tables déchargées
--

--
-- AUTO_INCREMENT pour la table `CabinCrew`
--
ALTER TABLE `CabinCrew`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT pour la table `Departus`
--
ALTER TABLE `Departus`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT pour la table `Device`
--
ALTER TABLE `Device`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT pour la table `Employees`
--
ALTER TABLE `Employees`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT pour la table `Fligth`
--
ALTER TABLE `Fligth`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT pour la table `Passenger`
--
ALTER TABLE `Passenger`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT pour la table `Pilote`
--
ALTER TABLE `Pilote`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT pour la table `Route`
--
ALTER TABLE `Route`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT pour la table `Tickets`
--
ALTER TABLE `Tickets`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- Contraintes pour les tables déchargées
--

--
-- Contraintes pour la table `CabinCrew`
--
ALTER TABLE `CabinCrew`
  ADD CONSTRAINT `CabinCrew_ibfk_1` FOREIGN KEY (`staff_id`) REFERENCES `Employees` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Contraintes pour la table `Departus`
--
ALTER TABLE `Departus`
  ADD CONSTRAINT `Departus_ibfk_1` FOREIGN KEY (`pilote`) REFERENCES `Pilote` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `Departus_ibfk_2` FOREIGN KEY (`copilote`) REFERENCES `Pilote` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Contraintes pour la table `Fligth`
--
ALTER TABLE `Fligth`
  ADD CONSTRAINT `Fligth_ibfk_1` FOREIGN KEY (`id_departures`) REFERENCES `Departus` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `Fligth_ibfk_2` FOREIGN KEY (`id_route`) REFERENCES `Route` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `Fligth_ibfk_3` FOREIGN KEY (`id_device`) REFERENCES `Device` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Contraintes pour la table `Passenger`
--
ALTER TABLE `Passenger`
  ADD CONSTRAINT `Passenger_ibfk_1` FOREIGN KEY (`ticket_id`) REFERENCES `Tickets` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Contraintes pour la table `Pilote`
--
ALTER TABLE `Pilote`
  ADD CONSTRAINT `Pilote_ibfk_1` FOREIGN KEY (`staff_id`) REFERENCES `Employees` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Contraintes pour la table `Tickets`
--
ALTER TABLE `Tickets`
  ADD CONSTRAINT `Tickets_ibfk_1` FOREIGN KEY (`departures_id`) REFERENCES `Departus` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
